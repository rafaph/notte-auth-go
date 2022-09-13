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

	for _, handler := range f.handlers {
		mux.HandleFunc(handler.endpoint.path, func(responseWriter http.ResponseWriter, request *http.Request) {
			if request.Method != handler.endpoint.method {
				http.Error(responseWriter, "invalid http method", http.StatusMethodNotAllowed)
				return
			}

			responseWriter.Header().Add("Content-Type", "application/json")
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
