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
  "log"
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
    m.Get(url+":app/:version", handler.GetVersion)
    r.Post("/:app", binding.Json(model.Version{}), handler.PostVersion)
    m.Put(url+":app/:version", handler.PutVersion)
    m.Delete(url+":app/:version", handler.DelVersion)
  })

  Context("Post", func() {

    BeforeEach(func() {
      test := model.Version{App: 2, Version: "123"}
      log.Println(test)
      body, err = json.Marshal(test)
      if err != nil {
        log.Println("Unable to marshal test")
      }
    })

    It("returns a 200 Status Code", func() {
      request := NewRequest("POST", "/v1/updates/123", body)
      response = httptest.NewRecorder()
      m.ServeHTTP(response, request)
      Expect(response.Code).To(Equal(200))
    })
  })

  Context("Get", func() {
    It("returns a 200 Status Code", func() {
      //Request("GET", "/v1/updates/1/100", GetVersion)
      Expect(response.Code).To(Equal(200))
      //Expect(response.Body).To(MatchJSON(`[{"Name":"keep things green"}]`))
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
