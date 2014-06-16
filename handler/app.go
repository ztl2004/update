package handler

import (
  "bytes"
  "github.com/arkors/update/model"
  "github.com/go-martini/martini"
  "github.com/go-xorm/xorm"
  //"github.com/codegangsta/martini-contrib/binding"
  "fmt"
  "github.com/hoisie/redis"
  "github.com/martini-contrib/render"
  "net/http"
  "strconv"
  "strings"
  //"time"
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
      //Indert into redis
      //sql := "select * from version where app=" + params["app"]
      //result := new(model.Version)
      //has, err := db.Sql(sql).Get(result)
      //if err == nil && has {
      var client redis.Client
      var buffer bytes.Buffer
      //idTrans := strconv.FormatInt(result.Id, 10)
      //appTrans := strconv.FormatInt(result.App, 10)
      //buffer.WriteString("Id=" + idTrans + "|")
      //buffer.WriteString("App=" + appTrans + "|")
      //buffer.WriteString("Version=" + result.Version + "|")
      //buffer.WriteString("Name=" + result.Name + "|")
      //buffer.WriteString("Updated=" + String(result.Updated) + "|")
      //buffer.WriteString("Changed=" + result.Changed + "|")
      //buffer.WriteString("Client=" + result.Client + "|")
      //buffer.WriteString("Url=" + result.Url + "|")
      //buffer.WriteString("Compatible=" + result.Compatible)
      //client.Set(appTrans+"@"+result.Version, []byte(buffer.String()))
      buffer.WriteString("App=" + params["app"] + "|")
      buffer.WriteString("Version=" + params["version"] + "|")
      buffer.WriteString("Name=" + version.Name + "|")
      //buffer.WriteString("Updated=" + String(result.Updated) + "|")
      buffer.WriteString("Changed=" + version.Changed + "|")
      buffer.WriteString("Client=" + version.Client + "|")
      buffer.WriteString("Url=" + version.Url + "|")
      buffer.WriteString("Compatible=" + version.Compatible)
      client.Set(params["app"]+"@"+params["version"], []byte(buffer.String()))
      r.JSON(201, version)
      return
      // }
    }
  }
}
func GetVersion(db *xorm.Engine, params martini.Params, r render.Render, res *http.Request) {
  appId, errAppId := strconv.ParseInt(params["app"], 0, 64)
  if errAppId != nil {
    r.JSON(400, map[string]interface{}{"error": "The application's id must be numrical"})
    return
  }
  _, log := res.Header["X-Arkors-Application-Log"]
  _, id := res.Header["X-Arkors-Application-Id"]
  _, Token := res.Header["X-Arkors-Application-Token"]
  _, clientHeader := res.Header["X-Arkors-Application-Client"]
  if log != true || clientHeader != true || id != true || Token != true {
    r.JSON(400, map[string]interface{}{"error": "Invalid request header,it should be include 'X-Arkors-Application-log' and 'X-Arkors-Application-Client'."})
    return
  }
  var client redis.Client
  result, err := client.Get(params["app"] + "@" + params["version"])
  if err == nil && result != nil {
    fmt.Println("EnterUpdateModule")
    fmt.Println(string(result))
    versionStringArray := strings.Split(string(result), "|")
    versionModel := new(model.Version)
    for i, _ := range versionStringArray {
      colum := strings.Split(versionStringArray[i], "=")[0]
      value := strings.Split(versionStringArray[i], "=")[1]
      fmt.Println(colum + "=" + value)
      if colum == "Id" {
        idTrans, _ := strconv.ParseInt(value, 0, 64)
        versionModel.Id = idTrans
      } else if colum == "App" {
        versionModel.App = appId
      } else if colum == "Version" {
        versionModel.Version = value
      } else if colum == "Name" {
        versionModel.Name = value
      } else if colum == "Updated" {
        //versionModel.Updated=value
      } else if colum == "Changed" {
        versionModel.Changed = value
      } else if colum == "Client" {
        versionModel.Client = value
      } else if colum == "Url" {
        versionModel.Url = value
      } else if colum == "Compatible" {
        versionModel.Compatible = value
      }
    }
    r.JSON(200, versionModel)
    return
  }
  /*
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
  */

}

func UpdateApp(db *xorm.Engine, params martini.Params, version model.Version, r render.Render, res *http.Request) {
  appId, errAppId := strconv.ParseInt(params["app"], 0, 64)
  if errAppId != nil {
    r.JSON(400, map[string]interface{}{"error": "The application's id must be numrical"})
    return
  }
  _, log := res.Header["X-Arkors-Application-Log"]
  _, clientHeader := res.Header["X-Arkors-Application-Client"]
  if log != true || clientHeader != true {
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
    var client redis.Client
    var buffer bytes.Buffer
    buffer.WriteString("App=" + params["app"] + "|")
    buffer.WriteString("Version=" + params["version"] + "|")
    buffer.WriteString("Name=" + version.Name + "|")
    //buffer.WriteString("Updated=" + String(result.Updated) + "|")
    buffer.WriteString("Changed=" + version.Changed + "|")
    buffer.WriteString("Client=" + version.Client + "|")
    buffer.WriteString("Url=" + version.Url + "|")
    buffer.WriteString("Compatible=" + version.Compatible)
    client.Set(params["app"]+"@"+params["version"], []byte(buffer.String()))
    val, _ := client.Get(params["app"] + "@" + params["version"])
    key := params["app"] + "@" + params["version"]
    fmt.Println("key=="+key, "val==="+string(val))
    r.JSON(200, version)
    return
  }
}

func DelVersion(db *xorm.Engine, params martini.Params, version model.Version, r render.Render, res *http.Request) {
  _, err := strconv.ParseInt(params["app"], 0, 64)
  if err != nil {
    r.JSON(400, map[string]interface{}{"error": "The application's id must be numrical"})
    return
  }
  _, log := res.Header["X-Arkors-Application-Log"]
  _, clientHeader := res.Header["X-Arkors-Application-Client"]
  if log != true || clientHeader != true {
    r.JSON(400, map[string]interface{}{"error": "Invalid request header,it should be include 'X-Arkors-Application-log' and 'X-Arkors-Application-Client'."})
    return
  }
  sql := "select * from version where app=" + params["app"] + " and version='" + params["version"] + "'"
  result := new(model.Version)
  has, err := db.Sql(sql).Get(result)
  if err != nil {
    r.JSON(400, map[string]interface{}{"error": "Datebase Error"})
    return
  }
  if has {
    //db.In("Id", id).Delete(result)
    _, err = db.Exec("delete from version where app=" + params["app"] + " and version='" + params["version"] + "'")
    if err == nil {
      var client redis.Client
      fmt.Println(params["app"] + "@" + params["version"])
      client.Del(params["app"] + "@" + params["version"])
      r.JSON(200, result)
      return
    }
  } else {
    r.JSON(404, map[string]interface{}{"error": "Application's ID is not exist!"})
    return
  }
}
