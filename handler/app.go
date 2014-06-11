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
func GetVersion(db *xorm.Engine, params martini.Params, versionModel model.Version, r render.Render, res *http.Request) {
  /* appId,err := strconv.ParseInt(params["app"],0,64)
  versionId,err2 := strconv.ParseInt(params["verion"],0,64)

  if err !=nil||err2 != nil {
    r.JSON(400,map[string]interface{"error":"The application's id or version's must be numrical"})
  }

  _,log:=res.Header["X-Arkors-Application-Log"]
  _,client=res.Header["X-Arkors-Application-Client"]
  _,token=res.Header["X-Arkors-Application-Token"]

  if log != true || client != true||token !=true {
    r.JSON(400,map[string]interface{"error":"Invalid request header,it should be include 'X-Arkors-Application-log','X-Arkors-Application-Client' and 'X-Arkors-Application-Token'."})
  }*/
}

func UpdateApp(db *xorm.Engine, params martini.Params, r render.Render) {

}

func DelVersion(db *xorm.Engine, params martini.Params, version model.Version, r render.Render, res *http.Request) {
  _, err := strconv.ParseInt(params["id"], 0, 64)
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
  sql := "select * from version where app=" + params["id"]
  fmt.Println(sql)
  results, err := db.Query(sql)
  //fmt.Println(results)
  if results == nil {
    //fmt.Println("testRepeatid")
    r.JSON(400, map[string]interface{}{"error": "Application's id is not exist"})
    return
  } else {
    var result model.Version
    db.Sql(sql).Get(&result)
    r.JSON(201, result)
  }
}
