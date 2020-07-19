package health

import (
	"github.com/danmaina/HttpResponse"
	"github.com/danmaina/logger"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"os"
)

type Health struct {
	State string `json:"state"`
	Pid   int    `json:"pid"`
}

const (
	applicationRunState = "Running"
)

func GetAppStatus(rw http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	logger.DEBUG("Received Request. \nHeaders: ", r.Header, "\nMethod: ", r.Method, "\nContentLength: ", r.ContentLength)
	//Set Content Type to JSON
	rw.Header().Set("content-type", "application/json")

	// Set Application Health
	applicationHealth := Health{
		State: applicationRunState,
		Pid:   os.Getpid(),
	}

	logger.DEBUG("Returning application health as:", applicationHealth)

	errReturningHealth := handlers.Response{
		Status: http.StatusOK,
		Error:  nil,
		Body:   applicationHealth,
	}.ReturnResponse(rw)

	if errReturningHealth != nil {
		logger.ERR("Error While trying to return response: ", errReturningHealth)
	}
}
