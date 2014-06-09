package model

import (
  _ "github.com/go-sql-driver/mysql"
  "time"
)

type Version struct {
  Id         int64
  App        int64
  Version    string    `xorm:"text"`
  Name       string    `xorm:"text"`
  Updated    time.Time `xorm:"updated"`
  Changed    time.Time `xorm:"changed"`
  Url        string    `xorm:"text"`
  Client     string    `xorm:"text"`
  Compatible string    `xorm:"text"`
}
