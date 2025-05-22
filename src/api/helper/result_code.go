package helper

type ResultCode int

const (
	Success         ResultCode = 0
	ValidationError ResultCode = 401
	AuthError       ResultCode = 411
	ForbiddenError  ResultCode = 431
	NotFoundError   ResultCode = 441
	LimiterError    ResultCode = 491
	OtpLimiterError ResultCode = 492
	CustomRecovery  ResultCode = 501
	InternalError   ResultCode = 502
)
