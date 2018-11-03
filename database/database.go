package db

import (
	"time"

	"github.com/globalsign/mgo"
)

var DB *mgo.Database

func Connect(user, password, dbname, host string) (*mgo.Database, error) {
	mgoDialInfo := &mgo.DialInfo{
		Addrs:    []string{host},
		Username: user,
		Password: password,
		Database: dbname,
		Timeout:  5 * time.Second,
	}

	session, err := mgo.DialWithInfo(mgoDialInfo)

	if err != nil {
		return nil, err
	}

	DB = session.DB(dbname)

	return DB, nil
}

func GetDB() *mgo.Database {
	return DB
}
