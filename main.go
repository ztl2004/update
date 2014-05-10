package update

import (
  "github.com/codegangsta/martini"
  "github.com/codegangsta/martini-contrib/binding"
  "github.com/codegangsta/martini-contrib/render"
  "strconv"
)

const url = "/v1/updates/"

func main() {
  m := martini.Classic()
  m.Get(url+":app/:version", GetVersion)
  m.Post(url, binding.Json(Version{}), HandleNewVersion)
  m.Run()
}

func GetVersion(params martini.Params, r render.Render) {
  appId, err := strconv.ParseInt(params["app"], 10, 64)
  version := &Version{App: appId}
  has, err := x.Get(version)
  if err != nil {
    r.JSON(400, map[string]interface{}{"Errors": "invalid json"})
  }
  if !has {
    r.JSON(400, version)
  }
  r.JSON(200, version)
}

func HandleNewVersion(version Version, r render.Render) {
  _, err := x.Insert(version)
  if err != nil {
    r.JSON(400, map[string]interface{}{"Errors": "invalid json"})
    return
  }
  //r.JSON(200, map[string]interface{}{"version added": "ok"})
  r.JSON(200, version)
}
