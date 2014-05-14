package main

import (
  "bytes"
  "encoding/json"
  "github.com/arkors/update/handler"
  "github.com/arkors/update/model"
  "github.com/codegangsta/martini"
  "github.com/codegangsta/martini-contrib/binding"
  "github.com/codegangsta/martini-contrib/render"
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
  //"log"
  "net/http"
  "net/http/httptest"
)

var _ = Describe("Test", func() {

  var (
    body []byte
    err  error
  )

  m := martini.Classic()
  m.Use(render.Renderer())
  m.Use(Db())
  m.Group("/v1/updates", func(r martini.Router) {
    m.Get("/:app/:version", handler.GetVersion)
    r.Post("/:app", binding.Json(model.Version{}), handler.PostVersion)
    m.Put("/:app/:version", handler.PutVersion)
    m.Delete("/:app/:version", handler.DelVersion)
  })

  Context("Post", func() {

    It("returns a 200 Status Code", func() {
      test := model.Version{App: 2, Version: "123"}
      body, err = json.Marshal(test)
      request := NewRequest("POST", "/v1/updates/2", body)
      response = httptest.NewRecorder()
      m.ServeHTTP(response, request)
      Expect(response.Code).To(Equal(200))
    })
    It("returns a 400 Status Code for invalid json", func() {
      request := NewRequest("POST", "/v1/updates/2",
        []byte("{\"sign\"\"5024442115e7bd738354c1fac662aed5\"}"))
      response = httptest.NewRecorder()
      m.ServeHTTP(response, request)
      Expect(response.Code).To(Equal(400))
    })
  })

  Context("Get", func() {
    It("returns a 200 Status Code", func() {
      request := NewRequest("GET", "/v1/updates/2/123", nil)
      response = httptest.NewRecorder()
      m.ServeHTTP(response, request)
      Expect(response.Code).To(Equal(200))
    })
  })
})

func NewRequest(method string, url string, body []byte) *http.Request {
  request, _ := http.NewRequest(method, url, bytes.NewReader(body))
  request.Header.Set("X-Arkors-Application-Log", "5024442115e7bd738354c1fac662aed5")
  request.Header.Set("X-Arkors-Application-Client", "127.0.0.1,TEST")
  request.Header.Set("Accept", "application/json")
  return request
}
