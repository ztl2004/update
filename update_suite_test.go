package update

import (
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"

  "testing"

  "github.com/codegangsta/martini"
  "github.com/codegangsta/martini-contrib/binding"
  "github.com/codegangsta/martini-contrib/render"
  "io"
  "net/http"
  "net/http/httptest"
)

var (
  response *httptest.ResponseRecorder
)

func TestUpdate(t *testing.T) {
  RegisterFailHandler(Fail)
  RunSpecs(t, "Update Suite")
}

func Request(method string, route string, handler martini.Handler) {
  m := martini.Classic()
  m.Get(route, handler)
  m.Use(render.Renderer())
  request, _ := http.NewRequest(method, route, nil)
  response = httptest.NewRecorder()
  m.ServeHTTP(response, request)
}

func DeleteRequest(method string, route string, handler martini.Handler) {
  m := martini.Classic()
  m.Get(route, handler)
  m.Use(render.Renderer())
  request, _ := http.NewRequest(method, route, nil)
  response = httptest.NewRecorder()
  m.ServeHTTP(response, request)
}

func PutRequest(method string, route string, handler martini.Handler) {
  m := martini.Classic()
  m.Get(route, handler)
  m.Use(render.Renderer())
  request, _ := http.NewRequest(method, route, nil)
  response = httptest.NewRecorder()
  m.ServeHTTP(response, request)
}

func PostRequest(method string, route string, handler martini.Handler, body io.Reader) {
  m := martini.Classic()
  m.Post(route, binding.Json(Version{}), handler)
  m.Use(render.Renderer())
  request, _ := http.NewRequest(method, route, body)
  response = httptest.NewRecorder()
  m.ServeHTTP(response, request)
}
