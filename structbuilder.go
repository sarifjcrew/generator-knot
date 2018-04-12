package main

import (
	"gopkg.in/mgo.v2"
)

type structbuilder struct {
	DBName         string
	CollectionName string
}

func (s *structbuilder) init(DBName, CollectionName string) *structbuilder {
	return &structbuilder{DBName, CollectionName}
}

func (s *structbuilder) connectdb() (session *mgo.Session, err error) {
	session, err = mgo.Dial("localhost")
	if err != nil {
		return
	}

	session.SetMode(mgo.Monotonic, true)
	return
}

func (s *structbuilder) getfirstdatadb(dbName string, collectionName string) (datas interface{}, err error) {
	var fatalErr error
	defer func() {
		if fatalErr != nil {
			datas = nil
			err = fatalErr
		}
	}()

	session, err := s.connectdb()
	if err != nil {
		fatalErr = err
		return
	}

	c := session.DB("Generator").C()
}
