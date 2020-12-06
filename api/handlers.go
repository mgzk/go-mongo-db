package api

import (
	"../mongodb"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"net/http"
)

type person struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	FirstName string             `json:"firstName"`
	LastName  string             `json:"lastName"`
	Age       int                `json:"age"`
	Address   string             `json:"address"`
}

var client *mongo.Client
var collection *mongo.Collection

func init() {
	client = mongodb.Client()
	collection = mongodb.Collection(client)
}

func handleGetAll(c *gin.Context) {
	cursor, err := collection.Find(context.TODO(), bson.D{})

	if err != nil {
		log.Fatal(err)
	}

	defer cursor.Close(context.TODO())

	var persons []person
	if err = cursor.All(context.TODO(), &persons); err != nil {
		log.Fatal(err)
	}

	result, err := json.Marshal(persons)

	c.String(http.StatusOK, fmt.Sprint(string(result)))
}

func handleGet(c *gin.Context) {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))

	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	var person person

	err = collection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&person)
	if err != nil {
		c.Status(http.StatusNotFound)
		return
	}

	result, err := json.Marshal(person)

	c.String(http.StatusOK, fmt.Sprint(string(result)))
}

func handlePost(c *gin.Context) {
	var person person

	err := c.ShouldBind(&person)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	insertOneResult, err := collection.InsertOne(context.TODO(), person)
	if err != nil {
		log.Fatal(err)
	}

	c.String(http.StatusOK, insertOneResult.InsertedID.(primitive.ObjectID).Hex())
}

func handlePut(c *gin.Context) {
	var person person

	err := c.ShouldBind(&person)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	person.ID = primitive.NilObjectID

	id, err := primitive.ObjectIDFromHex(c.Param("id"))

	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	update := bson.D{{"$set",
		bson.D{
			{"firstName", person.FirstName},
			{"lastName", person.LastName},
			{"age", person.Age},
			{"address", person.Address},
		},
	}}

	_, err = collection.UpdateOne(context.TODO(), bson.M{"_id": id}, update)
	if err != nil {
		log.Fatal(err)
	}

	err = collection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&person)
	if err != nil {
		c.Status(http.StatusNotFound)
		return
	}

	result, err := json.Marshal(person)
	c.String(http.StatusOK, fmt.Sprint(string(result)))
}

func handleDelete(c *gin.Context) {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))

	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	_, err = collection.DeleteOne(context.TODO(), bson.M{"_id": id})

	if err != nil {
		c.Status(http.StatusNotAcceptable)
	}

	c.Status(http.StatusNoContent)
}
