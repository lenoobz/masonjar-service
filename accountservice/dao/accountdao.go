/**
 * Create by Le Trong on 17/Feb/2019
 */

package dao

import (
	"fmt"

	"github.com/letrong/masonjar-service/accountservice/model"
)

// DBNAME Database name
const DBNAME = "masonjar"

// COLLECTION Collection name
const COLLECTION = "user"

// CONNECTIONSTRING DB connection string
const CONNECTIONSTRING = "mongodb://root:4mErAmlAb!@ds119930.mlab.com:19930/masonjar"

// var db *mongo.Database

// Connect establish a connection to database
func init() {
	// client, err := mongo.NewClient(CONNECTIONSTRING)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// err = client.Connect(context.Background())
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// // Collection types can be used to access the database
	// db = client.Database(DBNAME)
}

// InsertOneValue inserts one item from Account model
func InsertOneValue(account model.Account) {
	fmt.Println(account)
	// _, err := db.Collection(COLLECTION).InsertOne(context.Background(), account)
	// if err != nil {
	// 	log.Fatal(err)
	// }
}
