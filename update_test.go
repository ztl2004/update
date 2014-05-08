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

  Context("Just a Test", func() {

    BeforeEach(func() {
      test := Version{}
      body, err = json.Marshal(test)
      if err != nil {
        log.Println("Unable to marshal test")
      }
    })

    It("returns a 200 Status Code", func() {
      PostRequest("POST", "/test", HandleNewVersion, bytes.NewReader(body))
      Expect(response.Code).To(Equal(200))
    })
  })
})
