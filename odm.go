package goose

import (
	"log"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type docBase struct {
	ID bson.ObjectId `bson:"_id,omitempty" json:"_id"`
}

type schema struct {
	Name       string
	Collection *mgo.Collection
	Definition interface{}
}

type Goose struct {
	Session *mgo.Session
	DB      *mgo.Database
	Schemas map[string]schema
}

var instantiated *Goose = nil

func Init() *Goose {
	if instantiated == nil {
		instantiated = new(Goose)
	}
	instantiated.Schemas = make(map[string]schema)
	return instantiated
}

func (goose *Goose) Connect(connectionString, dbName string) (*mgo.Session, *mgo.Database) {
	session, err := mgo.Dial(connectionString)
	if err != nil {
		panic(err)
	}
	db := session.DB(dbName)
	goose.Session = session
	goose.DB = db
	return session, db
}

func (goose *Goose) Definition(name string) schema {
	_schema := schema{}
	if goose.Schemas[name] == (schema{}) {
		collection := goose.DB.C(name)
		_schema.Name = name
		_schema.Collection = collection
		goose.Schemas[name] = _schema
	}
	return goose.Schemas[name]
}

// Insert a new document
func (sch *schema) Save(doc interface{}) error {
	err := sch.Collection.Insert(doc)
	if err != nil {
		log.Printf("ERROR! %s\n", err)
	}
	return err
}

// Query documents in the collections
func (sch *schema) Index(page, count int, selectData []string, filter []byte) []bson.M {
	if page < 1 {
		page = 1
	}
	skip := count * (page - 1)
	var filterQuery interface{}
	bson.UnmarshalJSON(filter, &filterQuery)
	query := sch.Collection.Find(filterQuery).Sort("_id").Skip(skip).Limit(count)

	if len(selectData) > 1 {
		selectQuery := bson.M{}
		for _, v := range selectData {
			selectQuery[v] = 1
		}
		query = query.Select(selectQuery)
	}

	r := query.Iter()
	var resultSet []bson.M
	var p interface{}
	for {
		flag := r.Next(&p)
		if !flag {
			break
		}
		resultSet = append(resultSet, p.(bson.M))
	}

	return resultSet
}

// retrieve count of documents
func (sch *schema) Count(filter []byte) int {
	var filterQuery interface{}
	bson.UnmarshalJSON(filter, &filterQuery)
	query := sch.Collection.Find(filterQuery).Sort("_id")

	count, _ := query.Count()
	return count
}

// Get document
func (sch *schema) Get(id string, selectData []string) bson.M {
	query := sch.Collection.Find(bson.M{"_id": bson.ObjectIdHex(id)})

	if len(selectData) > 1 {
		selectQuery := bson.M{}
		for _, v := range selectData {
			selectQuery[v] = 1
		}
		query = query.Select(selectQuery)
	}

	var p bson.M
	query.One(&p)
	return p
}

// Update documents
func (sch *schema) Update(id string, _doc interface{}) (interface{}, bson.M) {
	var result bson.M
	doc := _doc.(map[string]interface{})
	delete(doc, "_id")
	change := mgo.Change{
		Update:    bson.M{"$set": doc},
		ReturnNew: true,
	}
	d, err := sch.Collection.Find(bson.M{"_id": bson.ObjectIdHex(id)}).Apply(change, &result)
	if err != nil {
		log.Println(err)
	}
	return d, result
}

// Delete documents
func (sch *schema) Delete(id string) error {
	err := sch.Collection.Remove(bson.M{"_id": bson.ObjectIdHex(id)})
	return err
}

func (sch *schema) SetIndex(keys []string, unique, dropDups, backgroud, sparse bool) {
	index := mgo.Index{
		Key:        keys,
		Unique:     unique,
		DropDups:   dropDups,
		Background: backgroud,
		Sparse:     sparse,
	}
	sch.Collection.EnsureIndex(index)
}
