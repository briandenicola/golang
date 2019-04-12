package m_example

import (
	"log"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

type Person struct {
	Name string
	Phone string
}

type mongoCollection struct {
	collection *mgo.Collection
}

func (m *mongoCollection) Insert(p *Person) error {	
	err := m.collection.Insert(&p)
	return err 
}

func (m *mongoCollection) Query(name string) Person {	
	result := Person{}
	err := m.collection.Find(bson.M{"name": name}).One(&result)
	if err != nil {
		log.Fatal(err)
	}
	return result
}

func CreateCollectionSession(connStr string, database string, collection string) *mongoCollection {
	session, err := mgo.Dial(connStr)
	if err != nil {
		panic(err)
	}
	session.SetMode(mgo.Monotonic, true)
	c := session.DB(database).C(collection)
	
	return &mongoCollection{collection: c}
}