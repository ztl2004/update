package handler

import (
  "github.com/arkors/update/model"
  "github.com/codegangsta/martini"
  "github.com/go-xorm/xorm"
  //"github.com/codegangsta/martini-contrib/binding"
  "github.com/codegangsta/martini-contrib/render"
  //"log"
  "strconv"
)

func CreateVersion(db *xorm.Engine, params martini.Params,versionModel model.Version,r render.Render,res *http.Request){
  appId,err := strconv.ParseInt(params["app"],0,64)
  if err !=nil {
    r.JSON(400,map[string]interface{"error":"The application's id must be numrical"})
  }

  _,log:=res.Header["X-Arkors-Application-Log"]
  _,client=res.Header["X-Arkors-Application-Client"]
  if log != true || client != true {
    r.JSON(400,map[string]interface{"error":"Invalid request header,it should be include 'X-Arkors-Application-log' and 'X-Arkors-Application-Client'."})
  }

  raws,err:=db.Raws("select * from arkors where appid="+appId)
  if raws.Next() {
    r.JSON(400,map[string]interface{"error":"The application's id already exist"})
  }
  versionModel.App=appId
  affected,err:=db.Insert(versionModel)
  if err!=nil {
   r.JSON(400,map[string]interface{"error":"Database error"})
  } else{
   r.JSON(201,verisonModel)
  }
}

func GetVersion(db *xorm.Engine, params martini.Params,versionModel model.Version,r render.Render,res *http.Request){
  appId,err := strconv.ParseInt(params["app"],0,64)
  versionId,err2 := strconv.ParseInt(params["verion"],0,64)

  if err !=nil||err2 != nil {
    r.JSON(400,map[string]interface{"error":"The application's id or version's must be numrical"})
  }

  _,log:=res.Header["X-Arkors-Application-Log"]
  _,client=res.Header["X-Arkors-Application-Client"]
  _,token=res.Header["X-Arkors-Application-Token"]

  if log != true || client != true||token !=true {
    r.JSON(400,map[string]interface{"error":"Invalid request header,it should be include 'X-Arkors-Application-log','X-Arkors-Application-Client' and 'X-Arkors-Application-Token'."})
  }

}

func UpdateApp(db *xorm.Engine, params martini.Params, r render.Render){

}

func DelVersion(db *xorm.Engine, params martini.Params, r render.Render){

}
