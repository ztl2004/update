package main

import (
  "github.com/codegangsta/martini"
  "github.com/codegangsta/martini-contrib/binding"
  "github.com/codegangsta/martini-contrib/render"
  "log"
  "net/http"
  "strconv"
)

const url = "/v1/updates/"

func main() {
  m := martini.Classic()
  m.Group("/v1/updates", func(r martini.Router) {
    m.Get(url+":app/:version", GetVersion)
    m.Post(url, binding.Json(Version{}), PostVersion)
    m.Put(url+":app/:version", PutVersion)
    m.Delete(url+":app/:version", DelVersion)
  })
  http.ListenAndServe(":3000", m)
}

func DelVersion(params martini.Params, r render.Render) {
  version := new(Version)
  appId, err := strconv.ParseInt(params["app"], 0, 64)
  version.App = appId
  version.Version = params["versioin"]
  affected, err := x.Cols("App", "Version").Delete(version)
  if err != nil {
    r.JSON(400, map[string]interface{}{"Errors": "invalid json"})
  }
  if affected == 0 {
    r.JSON(400, version)
  }
  r.JSON(200, version)
}

func PutVersion(params martini.Params, r render.Render) {
  version := new(Version)
  appId, err := strconv.ParseInt(params["app"], 0, 64)
  version.App = appId
  version.Version = params["versioin"]
  affected, err := x.Cols("App", "Version").Update(version)
  if err != nil {
    r.JSON(400, map[string]interface{}{"Errors": "invalid json"})
  }
  if affected == 0 {
    r.JSON(400, version)
  }
  r.JSON(200, version)
}

func GetVersion(params martini.Params, r render.Render) {
  version := new(Version)
  appId, err := strconv.ParseInt(params["app"], 0, 64)
  version.App = appId
  version.Version = params["version"]
  has, err := x.Get(version)
  log.Println(version, has, params["version"])
  if err != nil {
    r.JSON(400, map[string]interface{}{"Errors": "invalid json"})
  }
  if !has {
    r.JSON(400, version)
    return
  }
  r.JSON(200, version)
}

func PostVersion(version Version, r render.Render) {
  _, err := x.Insert(version)
  if err != nil {
    r.JSON(400, map[string]interface{}{"Errors": "invalid json"})
    return
  }
  //r.JSON(200, map[string]interface{}{"version added": "ok"})
  r.JSON(200, version)
}
