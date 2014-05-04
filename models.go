// Copyright 2013-2014 Unknown
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

package main

import (
  "errors"
  "log"

  _ "github.com/go-sql-driver/mysql"
  "github.com/go-xorm/xorm"
)

type Version struct {
  Id         int64
  Name       string `xorm:"unique"`
  Update     string
  changed    string
  url        string
  client     string
  compatible string
}

var x *xorm.Engine

func init() {
  var err error
  x, err = xorm.NewEngine("mysql", "root:root@/update?charset=utf8")
  if err != nil {
    log.Fatalf("Fail to create engine: %v\n", err)
  }

  if err = x.Sync(new(Version)); err != nil {
    log.Fatalf("Fail to sync database: %v\n", err)
  }
}

func newVersion(name string) error {
  _, err := x.Insert(&Version{Name: name})
  return err
}

func getVersion(id Id) (*Version, error) {
  a := &Version{}
  has, err := x.Id(id).Get(a)
  if err != nil {
    return nil, err
  } else if !has {
    return nil, errors.New("Account does not exist")
  }
  return a, nil
}
