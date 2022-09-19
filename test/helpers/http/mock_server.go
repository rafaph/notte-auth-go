package http

import (
	"github.com/rafaph/notte-auth/lib/validator"
	"log"
	"net/http"
	"net/http/httptest"
)

type MockResponse struct {
	StatusCode int
	Body       interface{}
}

type mockRoute struct {
	method string `validate:"required"`
	path   string `validate:"required"`
}

type mockHandler struct {
	route    mockRoute
	response MockResponse
}

type mockServer struct {
	handlers     []mockHandler
	currentRoute mockRoute
}

func (f *mockServer) When(method string, path string) *mockServer {
	f.currentRoute = mockRoute{method, path}

	return f
}

func (f *mockServer) Return(response MockResponse) {
	if err := validator.Validate(f.currentRoute); err != nil {
		log.Panicln("invalid route, please call `When` first.")
	}

	f.handlers = append(f.handlers, mockHandler{f.currentRoute, response})
	f.currentRoute = mockRoute{}
}

func (f *mockServer) getHandler() http.Handler {
	return http.HandlerFunc(func(responseWriter http.ResponseWriter, request *http.Request) {
		responseWriter.Header().Add("Content-Type", "application/json; charset=utf-8")

		for _, handler := range f.handlers {
			if request.URL.Path == handler.route.path {
				if request.Method != handler.route.method {
					responseWriter.WriteHeader(http.StatusMethodNotAllowed)
					_, _ = responseWriter.Write([]byte(`{"message": "method not allowed"}`))
					return
				}

				responseWriter.WriteHeader(handler.response.StatusCode)
				_, _ = responseWriter.Write([]byte(handler.response.Body.(string)))
				return
			}
		}
		responseWriter.WriteHeader(http.StatusNotFound)
		_, _ = responseWriter.Write([]byte(`{"message": "not found"}`))
	})
}

func (f *mockServer) Run(callback func(baseUrl string)) {
	server := httptest.NewServer(f.getHandler())
	defer server.Close()

	callback(server.URL)
}

func NewMockServer() *mockServer {
	return &mockServer{
		[]mockHandler{},
		mockRoute{},
	}
}
