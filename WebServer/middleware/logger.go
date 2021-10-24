package middleware

import (
	"fmt"
	"net/http"
)

func Logger(handlerFunc http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		request.ParseForm()
		fmt.Println("Logger MiddleWare:")
		fmt.Println("	- Default : ", request.Form)
		fmt.Println("	- Path : ", request.URL.Path)
		handlerFunc(writer, request)
	}
}