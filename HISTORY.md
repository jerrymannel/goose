# History

## 0.1
---
* The following features 
    - Connect to a single mongoDB instance
    - Attach a schema (currently there are no schema validations being performed)
    - Insert
    - Paginated Fetch
    - Update (whole document)
    - Delete
    - Add Indexs
* Version 0.1 has the following methods
	- func Init() *Goose
	- func (goose *Goose) Connect(connectionString, dbName string) (*mgo.Session, *mgo.Database)
	- func (goose *Goose) Definition(name string) schema
	- func (sch *schema) Save(doc interface{}) error
	- func (sch *schema) Index(page, count int, selectData []string, filter []byte) []bson.M
	- func (sch *schema) Count(filter []byte) int
	- func (sch *schema) Get(id string, selectData []string) bson.M
	- func (sch *schema) Update(id string, _doc interface{}) (interface{}, bson.M)
	- func (sch *schema) Delete(id string) error
	- func (sch *schema) SetIndex(keys []string, unique, dropDups, backgroud, sparse bool)