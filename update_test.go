package update

import (
  "bytes"
  "encoding/json"
  "log"

  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
)

var _ = Describe("Test", func() {

  var (
    body []byte
    err  error
  )

  Context("Post", func() {

    BeforeEach(func() {
      test := Version{}
      body, err = json.Marshal(test)
      if err != nil {
        log.Println("Unable to marshal test")
      }
    })

    It("returns a 200 Status Code", func() {
      PostRequest("POST", "/v1/updates/1", HandleNewVersion, bytes.NewReader(body))
      Expect(response.Code).To(Equal(200))
    })
  })

  Context("Get", func() {
    It("returns a 200 Status Code", func() {
      Request("GET", "/v1/updates/1/100", GetVersion)
      Expect(response.Code).To(Equal(200))
      //Expect(response.Body).To(MatchJSON(`[{"Name":"keep things green"}]`))
    })
  })
})
