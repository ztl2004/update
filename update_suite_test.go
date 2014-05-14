package main

import (
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
  //"log"
  "net/http/httptest"
  "testing"
)

var (
  response *httptest.ResponseRecorder
)

func TestUpdate(t *testing.T) {
  RegisterFailHandler(Fail)
  RunSpecs(t, "Update Suite")
}
