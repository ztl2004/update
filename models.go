package main

import (
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"log"
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

func getVersion() (*Version, error) {
	a := &Version{}
	has, err := x.Id("").Get(a)
	if err != nil {
		return nil, err
	} else if !has {
		return nil, errors.New("Account does not exist")
	}
	return a, nil
}