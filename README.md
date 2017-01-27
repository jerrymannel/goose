# Goose v1.0

Version history
|Version|Date|
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
> Connect(<connection string>, <db name>)

To connect first we have to initialize goose and then use the `Connect()` function
```go
goose := goose.Init()
goose.Connect("localhost", "golang")
```
Calling `goose.Init()` will return you the global goose instance.

### Attach Schema
> Definitions(<schema name>)

The `Definitions()` function is used to attach a schema.
```go
schema := goose.Definition("people")
```
Please note, currently goose doesn't perform any schema validations.

### Insert document
> Save(<interface>)

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
> Index(<page number>, <number of documents per page>, <fields to retrieve>, <filters>)

```go
filter := []byte(`{"age":20}`)
	selectQuery := []string{"name", "_id"}
	resultSet := schema.Index(1, 1, selectQuery, filter)
```