package main

import (
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"reflect"
	"runtime"
	"time"
)

func GetFunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

func main() {
	http.HandleFunc("/", EchoHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}

func EchoHandler(writer http.ResponseWriter, request *http.Request) {
	logger.WithFields(logrus.Fields{
		"remoteAddr": request.RemoteAddr,
		"method":     request.Method,
		"uri":        request.RequestURI,
		"host":       request.Host,
		"protocol":   request.Proto,
		"function":	GetFunctionName(EchoHandler),
		"user-agent":	request.UserAgent(),
	}).Info("")

	hostname, _ := os.Hostname()
	request.Header.Del("Cookie")
	request.Header.Set("Time", time.Now().String())
	request.Header.Set("Hostname", hostname)
	err := request.Write(writer)
	if err != nil {
		return 
	}


}
