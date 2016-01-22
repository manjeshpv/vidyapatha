package dbconfig

import (
	mgo "gopkg.in/mgo.v2"
)

type DB struct {
	Session *mgo.Session
}

func (db *DB) DoDial() (s *mgo.Session, err error) {
	return mgo.Dial(DBUrl())
}

func (db *DB) Name() string {
	return "my_awesome_app"
}

func DBUrl() string {
	dburl := "mongodb://qnotify:quezx123@ds063439.mongolab.com:63439/tsc"

	if dburl == "" {
		dburl = "mongodb://qnotify:quezx123@ds063439.mongolab.com:63439/tsc"
	}

	return dburl
}
