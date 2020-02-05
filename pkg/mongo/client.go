package mongo

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

// Client is the interface to be implemented by MongoClient
type Client interface {
	Close()
	FindOne(col string, query interface{}, result interface{}) error
	FindAll(col string, query interface{}, result interface{}) error
	Insert(col string, docs ...interface{}) error
	Update(col string, sel interface{}, update interface{}) error
	DeleteOne(col string, id string) error
	DeleteAll(col string, sel interface{}) error
}

// Config stores the configuration required to setup a new MongoClient
type Config struct {
	Host     string
	Database string
	AuthDB   string
	User     string
	Pass     string
}

// MongoClient is an abstract wrapper for MongoDB operations
type mongoClient struct {
	session *mgo.Session
}

func NewMongoClient(url string) (*mongoClient, error) {
	session, err := mgo.Dial(url)
	if err != nil {
		return nil, err
	}

	return &mongoClient{session}, nil
}

func (c *mongoClient) find(s *mgo.Session, col string, query interface{}) *mgo.Query {
	return s.DB("").C(col).Find(query)
}

func (c *mongoClient) FindOne(col string, query interface{}, result interface{}) error {
	s := c.session.Copy()
	defer s.Close()

	return c.find(s, col, query).One(result)
}

func (c *mongoClient) FindAll(col string, query interface{}, result interface{}) error {
	s := c.session.Copy()
	defer s.Close()

	return c.find(s, col, query).All(result)
}

func (c *mongoClient) Insert(col string, docs ...interface{}) error {
	s := c.session.Copy()
	defer s.Close()

	return s.DB("").C(col).Insert(docs...)
}

func (c *mongoClient) Update(col string, sel interface{}, update interface{}) error {
	s := c.session.Copy()
	defer s.Close()

	return s.DB("").C(col).Update(sel, update)
}

func (c *mongoClient) DeleteOne(col string, id string) error {
	s := c.session.Copy()
	defer s.Close()

	return s.DB("").C(col).RemoveId(bson.ObjectIdHex(id))
}

func (c *mongoClient) DeleteAll(col string, sel interface{}) error {
	s := c.session.Copy()
	defer s.Close()

	_, err := s.DB("").C(col).RemoveAll(sel)
	return err
}

// Close closes the client's MongoDB session
func (c *mongoClient) Close() {
	defer c.session.Close()
}
