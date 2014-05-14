package handler

import (
  "github.com/arkors/update/model"
  "github.com/codegangsta/martini"
  "github.com/go-xorm/xorm"
  //"github.com/codegangsta/martini-contrib/binding"
  "github.com/codegangsta/martini-contrib/render"
  "log"
  "strconv"
)

func DelVersion(db *xorm.Engine, params martini.Params, r render.Render) {
  version := new(model.Version)
  appId, err := strconv.ParseInt(params["app"], 0, 64)
  version.App = appId
  version.Version = params["versioin"]
  affected, err := db.Cols("App", "Version").Delete(version)
  if err != nil {
    r.JSON(400, map[string]interface{}{"Errors": "invalid json"})
  }
  if affected == 0 {
    r.JSON(400, version)
  }
  r.JSON(200, version)
}

func PutVersion(db *xorm.Engine, params martini.Params, r render.Render) {
  version := new(model.Version)
  appId, err := strconv.ParseInt(params["app"], 0, 64)
  version.App = appId
  version.Version = params["versioin"]
  affected, err := db.Cols("App", "Version").Update(version)
  if err != nil {
    r.JSON(400, map[string]interface{}{"Errors": "invalid json"})
  }
  if affected == 0 {
    r.JSON(400, version)
  }
  r.JSON(200, version)
}

func GetVersion(db *xorm.Engine, params martini.Params, r render.Render) {
  version := new(model.Version)
  appId, err := strconv.ParseInt(params["app"], 0, 64)
  version.App = appId
  version.Version = params["version"]
  has, err := db.Get(version)
  //log.Println(version, has, params["version"])
  if err != nil {
    r.JSON(400, map[string]interface{}{"Errors": "invalid json"})
  }
  if !has {
    r.JSON(400, version)
    return
  }
  r.JSON(200, version)
}

func PostVersion(db *xorm.Engine, params martini.Params, version model.Version, r render.Render) {
  if version.App == 0 && version.Version == "" {

    r.JSON(400, map[string]interface{}{"Errors": "invalid json"})
  }
  _, err := db.Insert(version)
  if err != nil {
    r.JSON(400, map[string]interface{}{"Errors": "db error"})
    return
  }
  //r.JSON(200, map[string]interface{}{"version added": "ok"})
  r.JSON(200, version)
}
