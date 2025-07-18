package usecase

import (
	"time"

	"go-web-api/config"
	"go-web-api/constant"
	"go-web-api/pkg/logging"
	"go-web-api/pkg/service_errors"
	dto "go-web-api/usecase/dto"

	"github.com/golang-jwt/jwt"
)

type TokenUsecase struct {
	logger logging.Logger
	cfg    *config.Config
}

type tokenDto struct {
	UserId       int
	FirstName    string
	LastName     string
	Username     string
	MobileNumber string
	Email        string
	Roles        []string
}

func NewTokenUsecase(cfg *config.Config) *TokenUsecase {
	logger := logging.NewLogger(cfg)
	return &TokenUsecase{
		cfg:    cfg,
		logger: logger,
	}
}

func (s *TokenUsecase) GenerateToken(token tokenDto) (*dto.TokenDetail, error) {
	td := &dto.TokenDetail{}
	td.AccessTokenExpireTime = time.Now().Add(s.cfg.JWT.AccessTokenExpireDuration * time.Minute).Unix()
	td.RefreshTokenExpireTime = time.Now().Add(s.cfg.JWT.RefreshTokenExpireDuration * time.Minute).Unix()

	atc := jwt.MapClaims{}

	atc[constant.UserIdKey] = token.UserId
	atc[constant.FirstNameKey] = token.FirstName
	atc[constant.LastNameKey] = token.LastName
	atc[constant.UsernameKey] = token.Username
	atc[constant.EmailKey] = token.Email
	atc[constant.MobileNumberKey] = token.MobileNumber
	atc[constant.RolesKey] = token.Roles
	atc[constant.ExpireTimeKey] = td.AccessTokenExpireTime

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atc)

	var err error
	td.AccessToken, err = at.SignedString([]byte(s.cfg.JWT.Secret))

	if err != nil {
		return nil, err
	}

	rtc := jwt.MapClaims{}

	rtc[constant.UserIdKey] = token.UserId
	rtc[constant.ExpireTimeKey] = td.RefreshTokenExpireTime

	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtc)

	td.RefreshToken, err = rt.SignedString([]byte(s.cfg.JWT.RefreshSecret))

	if err != nil {
		return nil, err
	}

	return td, nil
}

func (s *TokenUsecase) VerifyToken(token string) (*jwt.Token, error) {
	at, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, &service_errors.ServiceError{EndUserMessage: service_errors.UnExpectedError}
		}
		return []byte(s.cfg.JWT.Secret), nil
	})
	if err != nil {
		return nil, err
	}
	if !at.Valid {
		return nil, &service_errors.ServiceError{EndUserMessage: service_errors.TokenExpired}
	}
	claims, ok := at.Claims.(jwt.MapClaims)
	if !ok {
		return nil, &service_errors.ServiceError{EndUserMessage: service_errors.ClaimsNotFound}
	}
	exp, ok := claims["Exp"].(float64)
	if !ok {
		return nil, &service_errors.ServiceError{EndUserMessage: service_errors.TokenRequired}
	}
	if ok {
		expTime := int64(exp)
		timeNow := time.Now().Unix()
		if timeNow > expTime {
			return nil, &service_errors.ServiceError{EndUserMessage: service_errors.TokenExpired}
		}
	}
	return at, nil
}

func (s *TokenUsecase) GetClaims(token string) (claimMap map[string]interface{}, err error) {
	claimMap = map[string]interface{}{}

	verifyToken, err := s.VerifyToken(token)
	if err != nil {
		return nil, err
	}
	claims, ok := verifyToken.Claims.(jwt.MapClaims)
	if ok && verifyToken.Valid {
		for k, v := range claims {
			claimMap[k] = v
		}
		return claimMap, nil
	}
	return nil, &service_errors.ServiceError{EndUserMessage: service_errors.ClaimsNotFound}
}
