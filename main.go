package main

import (
  "github.com/arkors/update/handler"
  "github.com/arkors/update/model"
  "github.com/codegangsta/martini"
  "github.com/codegangsta/martini-contrib/binding"
  "github.com/go-xorm/xorm"
  "log"
  "net/http"
)

const url = "/v1/updates/"

var db *xorm.Engine

func init() {
  var err error
  db, err = xorm.NewEngine("mysql", "arkors:arkors@/arkors_update?charset=utf8")
  if err != nil {
    log.Fatalf("Fail to create engine: %v\n", err)
  }

  if err = db.Sync(new(model.Version)); err != nil {
    log.Fatalf("Fail to sync database: %v\n", err)
  }
}

func Db() martini.Handler {
  return func(c martini.Context) {
    c.Map(db)
  }
}

func main() {
  m := martini.Classic()
  m.Use(Db())
  m.Group("/v1/updates", func(r martini.Router) {
    m.Get("/:app/:version", handler.GetVersion)
    m.Post("/:app", binding.Json(model.Version{}), handler.PostVersion)
    m.Put("/:app/:version", handler.PutVersion)
    m.Delete("/:app/:version", handler.DelVersion)
  })
  http.ListenAndServe(":3000", m)
}
