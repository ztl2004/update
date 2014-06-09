package main

import (
  "github.com/arkors/update/model"
  _ "github.com/go-sql-driver/mysql"
  "github.com/go-xorm/xorm"
  "log"
)

var x *xorm.Engine

func init() {
  var err error
  x, err = xorm.NewEngine("mysql", "arkors:arkors@/arkors_update?charset=utf8")
  if err != nil {
    log.Fatalf("Fail to create engine: %v\n", err)
  }

  if err = x.Sync(new(model.Version)); err != nil {
    log.Fatalf("Fail to sync database: %v\n", err)
  }
}
