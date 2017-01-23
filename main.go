package main

import (
	"log"

	"gopkg.in/mgo.v2/bson"
)

func main() {

	goose := Init()
	goose.Connect("localhost", "golang")
	schema := goose.Definition("people", &People{})

	// schema.Save(&People{"A", 10})
	// schema.Save(&People{"A", 20})

	// return

	filter := []byte(`{"age":20}`)
	selectQuery := []string{"name", "_id"}
	resultSet := schema.Index(1, 1, selectQuery, filter)
	log.Printf("Result length : %d\n", len(resultSet))
	log.Printf("Results (bson array) : %s\n", resultSet)
	result1 := resultSet[0]

	filter = []byte(`{"age":10}`)
	selectQuery = []string{"name", "_id"}
	resultSet = schema.Index(1, 1, selectQuery, filter)
	log.Printf("Result length : %d\n", len(resultSet))
	log.Printf("Results (bson array) : %s\n", resultSet)
	result2 := resultSet[0]

	log.Printf("First result (bson) : %s\n", result1)
	objectID1 := result1["_id"].(bson.ObjectId)
	log.Printf("ID 1 : %s\n", objectID1)
	id1 := objectID1.Hex()

	log.Printf("Second result (bson) : %s\n", result2)
	objectID2 := result2["_id"].(bson.ObjectId)
	log.Printf("ID 1 : %s\n", objectID2)
	id2 := objectID2.Hex()

	singleEntry1 := schema.Get(id1, []string{})
	log.Printf("Get by ID 1 : %s\n", singleEntry1)

	singleEntry2 := schema.Get(id2, []string{})
	log.Printf("Get by ID 2 : %s\n", singleEntry2)

	log.Println("Updating age to 100")
	d := []byte(`{"age":100}`)
	var data interface{}
	bson.UnmarshalJSON(d, &data)
	schema.Update(id1, data)

	singleEntry1 = schema.Get(id1, []string{})
	log.Printf("Get by ID 1: %s\n", singleEntry1)

	singleEntry2 = schema.Get(id2, []string{})
	log.Printf("Get by ID 2: %s\n", singleEntry2)

	log.Println("Delete documents")
	schema.Delete(id1)
	schema.Delete(id2)

	resultSet = schema.Index(1, 1, selectQuery, filter)
	log.Printf("Result length : %d\n", len(resultSet))
	log.Printf("Results (bson array) : %s\n", resultSet)

	// id := "5884f8f62bb152b1d73ba010"
	// d := []byte(`{"age":1231231231}`)
	// d := []byte(`{"name":"A 0001"}`)

	// log.Println(p)
	// schema.Update(id, data)
	// schema.Delete(id)

	// s1 := []byte(`{"sourcingPartner":{"$nin": "MPS0"} }`)

	// s3 := []byte(`{"status":{"$in":["Limited Approval", "Approved"]}}`)
	// s4 := []byte(`{"parent":{"$exists":false}}`)
	// log.Println(string(s1))
	// log.Println(string(s2))
	// log.Println(string(s3))
	// log.Println(string(s4))

	// var p interface{}
	// bson.UnmarshalJSON(s1, &p)
	// log.Println(p)
	// log.Println(string(bs1))

	// list := make([]string, 30)
	// for i, _ := range list {
	// 	schema.Save(&People{strings.Join([]string{"A ", strconv.Itoa(i)}, " "), i})
	// }

}
