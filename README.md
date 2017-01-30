# Goose v1.0

Version history
|Version|Date  |
|:----:|------|
|1.0|27th Jan 2017|

## Introduction
Goose is an ODM(Object Data Mapper) for Golang built on top of the [mgo](https://labix.org/mgo) library. 

## Goal
Mirror all the features of [Mongoose](http://mongoosejs.com/). 

## Features
Currently Goose supports the following.
- Connect to a single mongoDB instance
- Attach a schema (currently there are no schema validations being performed)
- Insert
- Paginated Fetch
- Update (whole document)
- Delete
- Add Indexs

### Connect
`func Init() *Goose`
`func (goose *Goose) Connect(connectionString, dbName string) (*mgo.Session, *mgo.Database)`

To connect first we have to initialize goose and then use the `Connect()` function
```go
goose := goose.Init()
goose.Connect("localhost", "golang")
```
Calling `goose.Init()` will return you the global goose instance.

### Attach Schema
`func (goose *Goose) Definition(name string) schema`

The `Definitions()` function is used to attach a schema.
```go
schema := goose.Definition("people")
```
Please note, currently goose doesn't perform any schema level validations.

### Insert document
`func (sch *schema) Save(doc interface{}) error`

The `Save()` function is used to insert a new document.
```go
type People struct {
	Name string
	Age  int
}

goose := goose.Init()
schema.Save(&People{"Apple", 10})
```

## Paginated fetch
`func (sch *schema) Index(page, count int, selectData []string, filter []byte) []bson.M`

```go
filter := []byte(`{"age":20}`)
selectQuery := []string{"name", "_id"}
resultSet := schema.Index(1, 1, selectQuery, filter)
```
This returns an array of bsons.

## Count
`func (sch *schema) Count(filter []byte) int`
Given a filter, this function returns the number of matching documents.

## Get
`func (sch *schema) Get(id string, selectData []string) bson.M`

Fetches the document based on the `id`. If `selectData` is provided, then only those attributes of the document is retrieved.

## Update
`func (sch *schema) Update(id string, _doc interface{}) (interface{}, bson.M)`

Updates a document with `_doc` where the _id matches `id`

## Delete
`func (sch *schema) Delete(id string) error`

Deletes the document where _id is the same as `id`

## Set Index
`func (sch *schema) SetIndex(keys []string, unique, dropDups, backgroud, sparse bool)`

Sets the indexes for the collection.
Read more about how to set the indexes under *mgo package* [func (*Collection) EnsureIndex](https://godoc.org/gopkg.in/mgo.v2#Collection.EnsureIndex)

