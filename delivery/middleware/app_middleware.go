package middleware

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"go_api/delivery/commonresp"
	"net/http"
)

func DummyMiddleWare(c *gin.Context) {
	fmt.Println("DUMMY")
	c.Next()
}

func TokenAuthMiddleWare(token string) gin.HandlerFunc {
	requiredToken := token
	if requiredToken == "" {
		panic("API_TOKEN not exist")
	}

	return func(c *gin.Context) {
		token := c.Request.Header.Get("api_token")
		if token == "" {
			//c.AbortWithStatusJSON(http.StatusUnauthorized, "Unauthorized Token")
			e := commonresp.NewErrorMessage(http.StatusUnauthorized, "01", "Unauthorized")
			c.Abort()
			c.Error(fmt.Errorf("%s", e.ToJson()))
			return
		}
		if token != requiredToken {
			//c.AbortWithStatusJSON(http.StatusUnauthorized, "Unauthorized Token")
			e := commonresp.NewErrorMessage(http.StatusUnauthorized, "02", "Unauthorized")
			c.Abort()
			c.Error(fmt.Errorf("%s", e.ToJson()))
			return
		}
		c.Next()
	}
}

func ErrorMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		detectedError := c.Errors.Last()
		if detectedError == nil {
			return
		}
		e := detectedError.Error()
		errResp := commonresp.ErrorMessage{}
		err := json.Unmarshal([]byte(e), &errResp)
		if err != nil {
			errResp.HttpCode = http.StatusInternalServerError
			errResp.ErrorDescription = commonresp.ErrorDescription{
				Status:       "Error",
				ResponseCode: "06",
				Description:  "convert json field",
			}
		}
	}
}
