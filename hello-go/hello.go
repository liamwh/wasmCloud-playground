package main

import (
	// http "github.com/liamwh/wasmcloud-playground/examples/golang/actors/http-hello-world/gen"
	http "github.com/liamwh/wasmCloud-playground/hello-go/gen"
)

// Helper type aliases to make code more readable
type (
	HttpRequest          = http.ExportsWasiHttp0_2_0_IncomingHandlerIncomingRequest
	HttpResponseWriter   = http.ExportsWasiHttp0_2_0_IncomingHandlerResponseOutparam
	HttpOutgoingResponse = http.WasiHttp0_2_0_TypesOutgoingResponse
	HttpError            = http.WasiHttp0_2_0_TypesErrorCode
)

type HttpServer struct{}

func init() {
	httpserver := HttpServer{}
	// configuration, err := openfga.NewConfiguration(openfga.Configuration{
	// 	ApiUrl: "http://openfga:8080",
	// 	// ApiUrl:               os.Getenv("FGA_API_URL"), // required, e.g. https://api.fga.example
	// })
	// if err != nil {
	// 	panic(err)
	// }
	// Set the incoming handler struct to HttpServer
	http.SetExportsWasiHttp0_2_0_IncomingHandler(httpserver)
}

func (h HttpServer) Handle(request HttpRequest, responseWriter HttpResponseWriter) {
	// Construct HttpResponse to send back
	headers := http.NewFields()
	httpResponse := http.NewOutgoingResponse(headers)
	httpResponse.SetStatusCode(200)
	httpResponse.Body().Unwrap().Write().Unwrap().BlockingWriteAndFlush([]uint8("Hello from Go!\n")).Unwrap()

	// Send HTTP response
	okResponse := http.Ok[HttpOutgoingResponse, HttpError](httpResponse)
	http.StaticResponseOutparamSet(responseWriter, okResponse)
}

//go:generate wit-bindgen tiny-go wit --out-dir=gen --gofmt
func main() {}
