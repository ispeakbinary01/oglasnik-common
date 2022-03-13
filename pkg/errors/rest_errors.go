package errors

import (
	"time"

	"github.com/labstack/echo/v4"
	useragent "github.com/mssola/user_agent"
)

// RestError is the default error struct
// with extra fields
type RestError struct {
	Status  int
	Code    string
	Message string
}

// RestErrorOutput defines the structure that the client
// will receive when the app encounters an error
type RestErrorOutput struct {
	Status         int       `json:"status"`
	Code           string    `json:"code"`
	Message        string    `json:"message"`
	Timestamp      time.Time `json:"timestamp"`
	Endpoint       string    `json:"endpoint"`
	BrowserName    string    `json:"browser_name"`
	BrowserVersion string    `json:"browser_version"`
	OS             string    `json:"os"`
	Mobile         bool      `json:"mobile"`
	Bot            bool      `json:"bot"`
	IP             string    `json:"ip"`
	Error          string    `json:"error"`
}

var restErrors = map[string]RestError{
	RestUnableToGetOne: {
		Status:  500,
		Code:    RestUnableToGetOne,
		Message: "unable to get one record",
	},
	RestUnableToGetMany: {
		Status:  500,
		Code:    RestUnableToGetMany,
		Message: "unable to get many records",
	},
}

// SendRestError returns an RestErrorOutput structure
func SendRestError(c echo.Context, err string, msg ...string) error {
	ua := useragent.New(c.Request().UserAgent())
	brName, brVersion := ua.Browser()
	var m string
	if len(msg) > 0 {
		m = msg[0]
	}
	eo := RestErrorOutput{
		Status:         restErrors[err].Status,
		Code:           restErrors[err].Code,
		Message:        restErrors[err].Message,
		Timestamp:      time.Now(),
		Endpoint:       c.Path(),
		BrowserName:    brName,
		BrowserVersion: brVersion,
		OS:             ua.OS(),
		Mobile:         ua.Mobile(),
		Bot:            ua.Bot(),
		IP:             c.RealIP(),
		Error:          m,
	}
	c.JSON(restErrors[err].Status, eo)
	return nil
}

const RestUnableToGetOne = "RestUnableToGetOne"

const RestUnableToGetMany = "RestUnableToGetMany"

const RestUnableToRegisterUser = "RestUnableToRegisterUser"
