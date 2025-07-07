package handler

import (
	"bytes"
	"encoding/json"
	"net/http"

	"go-web-api/api/helper"

	"github.com/gin-gonic/gin"
)

type PythonHandler struct {
}

func NewPythonHandler() *PythonHandler {
	return &PythonHandler{}
}

// PythonHandler ...
// @Summary Call Python service
// @Description Call Python service
// @Tags Python
// @Accept json
// @Produce json
// @Success 200 {object} helper.BaseHttpResponse "Success"
// @Failure 400 {object} helper.BaseHttpResponse "Failed"
// @Router /v1/python/hello [get]
func (h *PythonHandler) Hello(c *gin.Context) {
	go func() {
		resp, err := http.Get("http://localhost:5001/hello")
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}
		defer resp.Body.Close()

		var result map[string]interface{}

		if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, helper.GenerateBaseResponse(gin.H{"res": result}, true, helper.Success))
	}()

}

// PythonHandler ...
// @Summary Call Python service for prediction
// @Description Call Python service for prediction
// @Tags Python
// @Accept json
// @Produce json
// @Param        body  body   map[string]interface{}  true  "Request body"
// @Success 200 {object} helper.BaseHttpResponse "Success"
// @Failure 400 {object} helper.BaseHttpResponse "Failed"
// @Router /v1/python/predict [post]
func (h *PythonHandler) Predict(c *gin.Context) {
	var requestBody map[string]interface{}
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	bodyBytes, err := json.Marshal(requestBody)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to marshal request body"})
		return
	}

	resp, err := http.Post("http://localhost:5001/predict", "application/json", bytes.NewBuffer(bodyBytes))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	defer resp.Body.Close()

	var result map[string]interface{}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, result)
}
