package utils

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func Response(c echo.Context, data interface{}, err error) error {
	hostname, _ := os.Hostname()
	if err != nil {
		return c.JSON(http.StatusBadGateway, map[string]interface{}{
			"message":  err.Error(),
			"status":   false,
			"hostname": hostname,
		})
	} else {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"data":     data,
			"message":  "success",
			"status":   true,
			"hostname": hostname,
		})
	}
}
