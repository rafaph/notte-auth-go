package helpers

import (
	"net/http"
	"net/http/httptest"
)

type Context struct {
	StatusCode   int
	ResponseBody string
}

type Endpont struct {
	method string
	path   string
}

type Handler struct {
	endpoint Endpont
	context  Context
}

type FakeServer struct {
	handlers []Handler
	endpoint Endpont
}

func (f *FakeServer) When(method string, path string) *FakeServer {
	f.endpoint = Endpont{method, path}

	return f
}

func (f *FakeServer) Returns(context Context) {
	if f.endpoint.method == "" || f.endpoint.path == "" {
		panic("invalid endpoint, please call `When` first.")
	}

	f.handlers = append(f.handlers, Handler{f.endpoint, context})
	f.endpoint = Endpont{}
}

func (f *FakeServer) getHandler() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(responseWriter http.ResponseWriter, request *http.Request) {
		responseWriter.Header().Add("Content-Type", "application/json; charset=utf-8")
		responseWriter.WriteHeader(http.StatusNotFound)
		_, _ = responseWriter.Write([]byte(`{"message": "not found"}`))
	})

	for _, handler := range f.handlers {
		mux.HandleFunc(handler.endpoint.path, func(responseWriter http.ResponseWriter, request *http.Request) {
			responseWriter.Header().Add("Content-Type", "application/json; charset=utf-8")
			if request.Method != handler.endpoint.method {
				responseWriter.WriteHeader(http.StatusMethodNotAllowed)
				_, _ = responseWriter.Write([]byte(`{"message": "method not allowed"}`))
				return
			}
			responseWriter.WriteHeader(handler.context.StatusCode)
			_, _ = responseWriter.Write([]byte(handler.context.ResponseBody))
		})
	}

	return mux
}

func (f *FakeServer) Run(callback func(baseUrl string)) {
	server := httptest.NewServer(f.getHandler())
	defer server.Close()

	callback(server.URL)
}

func NewFakeServer() *FakeServer {
	return &FakeServer{
		[]Handler{},
		Endpont{},
	}
}
