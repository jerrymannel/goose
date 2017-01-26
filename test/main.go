package main

import (
	"log"

	"github.com/jerrymannel/goose"

	"gopkg.in/mgo.v2/bson"
)

type People struct {
	Name string
	Age  int
}

func main() {

	goose := goose.Init()
	goose.Connect("localhost", "golang")
	schema := goose.Definition("people")

	// INSERT
	schema.Save(&People{"Apple", 10})
	schema.Save(&People{"A", 20})

	// GET values using filter
	filter := []byte(`{"age":20}`)
	selectQuery := []string{"name", "_id"}
	resultSet := schema.Index(1, 1, selectQuery, filter)
	log.Printf("Result length : %d\n", len(resultSet))
	log.Printf("Results (bson array) : %s\n", resultSet)
	result1 := resultSet[0]

	filter = []byte(`{"name":/ppl/}`)
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

	// GET by ID
	singleEntry1 := schema.Get(id1, []string{})
	log.Printf("Get by ID 1 : %s\n", singleEntry1)

	singleEntry2 := schema.Get(id2, []string{})
	log.Printf("Get by ID 2 : %s\n", singleEntry2)

	// UPDATE
	log.Println("Updating age to 100")
	d := []byte(`{"age":100}`)
	var data interface{}
	bson.UnmarshalJSON(d, &data)
	change, data := schema.Update(id1, data)
	log.Println(change)
	log.Println(data)

	singleEntry1 = schema.Get(id1, []string{})
	log.Printf("Get by ID 1: %s\n", singleEntry1)

	singleEntry2 = schema.Get(id2, []string{})
	log.Printf("Get by ID 2: %s\n", singleEntry2)

	// DELETE
	log.Println("Delete documents")
	schema.Delete(id1)
	schema.Delete(id2)

	resultSet = schema.Index(1, 1, selectQuery, filter)
	log.Printf("Result length : %d\n", len(resultSet))
	log.Printf("Results (bson array) : %s\n", resultSet)

}
