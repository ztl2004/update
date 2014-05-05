package main

import (
  "github.com/codegangsta/martini"
  "github.com/codegangsta/martini-contrib/render"
  "log"
)

const url = "/v1/updates"

func main() {
  m := martini.Classic()
  m.Get(url, func() string {
    return "Martini, shaken not stirred!"
  })
  //m.Post(url, binding.Json(Version), HandleNewVersion)
  m.Run()
}

func HandleNewVersion(version Version, r render.Render) {
  err := newVersion(version.Name)
  if err != nil {
    log.Println(err)
    return
  }
  r.JSON(200, map[string]interface{}{"version added": "ok"})
}
