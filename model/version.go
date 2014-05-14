package model

import (
  _ "github.com/go-sql-driver/mysql"
)

type Version struct {
  Id         int64
  App        int64
  Version    string
  Name       string
  Update     string
  Changed    string
  Url        string
  Client     string
  Compatible string
}
