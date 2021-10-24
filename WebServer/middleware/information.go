package middleware

import (
	"fmt"
	"net/http"
)

func Information(handlerFunc http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		request.ParseForm()
		fmt.Fprintln(writer, "Information MiddleWare:")
		fmt.Fprintln(writer, "	- Default : ", request.Form)
		fmt.Fprintln(writer, "	- Path : ", request.URL.Path)
		handlerFunc(writer, request)
	}
}