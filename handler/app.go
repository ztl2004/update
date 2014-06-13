package handler

import (
  "github.com/arkors/update/model"
  "github.com/go-martini/martini"
  "github.com/go-xorm/xorm"
  //"github.com/codegangsta/martini-contrib/binding"
  "fmt"
  "github.com/martini-contrib/render"
  "net/http"
  "strconv"
)

func CreateVersion(db *xorm.Engine, params martini.Params, version model.Version, r render.Render, res *http.Request) {
  appId, err := strconv.ParseInt(params["app"], 0, 64)
  if err != nil {
    r.JSON(400, map[string]interface{}{"error": "The application's id must be numrical"})
    return
  }
  _, log := res.Header["X-Arkors-Application-Log"]
  _, client := res.Header["X-Arkors-Application-Client"]
  if log != true || client != true {
    r.JSON(400, map[string]interface{}{"error": "Invalid request header,it should be include 'X-Arkors-Application-log' and 'X-Arkors-Application-Client'."})
    return
  }
  if version.Version == "" || version.Name == "" || version.Changed == "" || version.Url == "" || version.Client == "" || version.Compatible == "" {
    r.JSON(400, map[string]interface{}{"error": "Invalid json body "})
    return
  }
  sql := "select * from version where app=" + params["app"]
  fmt.Println(sql)
  results, err := db.Query(sql)
  //fmt.Println(results)
  if results != nil {
    //fmt.Println("testRepeatid")
    r.JSON(400, map[string]interface{}{"error": "The application's id already exist"})
    return
  } else {
    version.App = appId
    _, err2 := db.Insert(version)
    if err2 != nil {
      r.JSON(400, map[string]interface{}{"error": "Database error"})
      return
    } else {
      r.JSON(201, version)
      return
    }
  }
}
func GetVersion(db *xorm.Engine, params martini.Params, r render.Render, res *http.Request) {
  _, errAppId := strconv.ParseInt(params["app"], 0, 64)
  if errAppId != nil {
    r.JSON(400, map[string]interface{}{"error": "The application's id must be numrical"})
    return
  }
  _, log := res.Header["X-Arkors-Application-Log"]
  _, id := res.Header["X-Arkors-Application-Id"]
  _, Token := res.Header["X-Arkors-Application-Token"]
  _, client := res.Header["X-Arkors-Application-Client"]
  if log != true || client != true || id != true || Token != true {
    r.JSON(400, map[string]interface{}{"error": "Invalid request header,it should be include 'X-Arkors-Application-log' and 'X-Arkors-Application-Client'."})
    return
  }
  sql := "select * from version where app=" + params["app"] + " and version='" + params["version"] + "'"
  result := new(model.Version)
  has, err := db.Sql(sql).Get(result)
  if err != nil {
    r.JSON(400, map[string]interface{}{"error": "Database Error"})
    return
  }
  if has {
    r.JSON(200, result)
    return
  } else {
    r.JSON(404, map[string]interface{}{"error": "No version found"})
    return
  }
}

func UpdateApp(db *xorm.Engine, params martini.Params, version model.Version, r render.Render, res *http.Request) {
  appId, errAppId := strconv.ParseInt(params["app"], 0, 64)
  if errAppId != nil {
    r.JSON(400, map[string]interface{}{"error": "The application's id must be numrical"})
    return
  }
  _, log := res.Header["X-Arkors-Application-Log"]
  _, client := res.Header["X-Arkors-Application-Client"]
  if log != true || client != true {
    fmt.Println("Invalid request header,it should be include 'X-Arkors-Application-log' and 'X-Arkors-Application-Client'.")
    r.JSON(400, map[string]interface{}{"error": "Invalid request header,it should be include 'X-Arkors-Application-log' and 'X-Arkors-Application-Client'."})
    return
  }
  if version.Version == "" || version.Name == "" || version.Changed == "" || version.Url == "" || version.Client == "" || version.Compatible == "" {
    r.JSON(400, map[string]interface{}{"error": "Invalid json body "})
    return
  }
  version.App = appId
  version.Version = params["version"]
  fmt.Println("changed=======" + version.Changed)
  has, err := db.In("App", appId).Update(version)
  if err != nil {
    r.JSON(400, map[string]interface{}{"error": "Database Error"})
    return
  }
  if has == 0 {
    r.JSON(404, map[string]interface{}{"error": "Not found any version records"})
    return
  } else {
    r.JSON(200, version)
    return
  }
}

func DelVersion(db *xorm.Engine, params martini.Params, version model.Version, r render.Render, res *http.Request) {
  id, err := strconv.ParseInt(params["id"], 0, 64)
  if err != nil {
    r.JSON(400, map[string]interface{}{"error": "The application's id must be numrical"})
    return
  }
  _, log := res.Header["X-Arkors-Application-Log"]
  _, client := res.Header["X-Arkors-Application-Client"]
  if log != true || client != true {
    r.JSON(400, map[string]interface{}{"error": "Invalid request header,it should be include 'X-Arkors-Application-log' and 'X-Arkors-Application-Client'."})
    return
  }
  sql := "select * from version where id=" + params["id"]
  result := new(model.Version)
  has, err := db.Sql(sql).Get(result)
  if err != nil {
    r.JSON(400, map[string]interface{}{"error": "Datebase Error"})
    return
  }
  if has {
    db.In("Id", id).Delete(result)
    r.JSON(200, result)
    return
  } else {
    r.JSON(404, map[string]interface{}{"error": "Application's ID is not exist!"})
    return
  }
}
