package main

import (
  "bytes"
  "encoding/json"
  "log"

  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"

  "github.com/codegangsta/martini"
  "github.com/codegangsta/martini-contrib/binding"
  "github.com/codegangsta/martini-contrib/render"
)

var _ = Describe("Test", func() {

  var (
    body []byte
    err  error
  )

  Context("Post", func() {

    BeforeEach(func() {
      test := Version{App: 2, Version: "123"}
      body, err = json.Marshal(test)
      if err != nil {
        log.Println("Unable to marshal test")
      }
    })

    It("returns a 200 Status Code", func() {
      m := martini.Classic()
      m.Use(render.Renderer())
      m.Group("/v1/updates", func(r martini.Router) {
        r.Post("/:app/:version", binding.Json(Version{}), PostVersion)
      })
      response = httptest.NewRecorder()
      request, _ := http.NewRequest("POST", "/v1/updates/", bytes.NewReader(body))
      request.Header.Set("X-Arkors-Application-Log", "5024442115e7bd738354c1fac662aed5")
      request.Header.Set("X-Arkors-Application-Client", "127.0.0.1,TEST")
      request.Header.Set("Accept", "application/json")
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
