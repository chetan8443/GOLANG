package routes

import (
	"context"

	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/vishal/caloryTracker-app/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// var entryCollection *mongo.Collection = openCollection(Client, "calories")
var entryCollection *mongo.Collection = OpenCollection(Client, "calories")
var validate = validator.New()

func AddEntry(c *gin.Context) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

	var entry models.Entry

	err := c.BindJSON(&entry)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		fmt.Println("error in line 32 entries.go")
		return
	}

	entry.Id = primitive.NewObjectID()
	valEr := validate.Struct(entry)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": valEr.Error()})
		fmt.Println(err)
		fmt.Println("error in line 42 entries.go")
		return
	}

	result, err := entryCollection.InsertOne(ctx, entry)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Record not created"})
		fmt.Println(err)
		fmt.Println("error in line 50 entries.go")
		return
	}
	defer cancel()
	c.JSON(http.StatusOK, result)

}

func GetEntries(c *gin.Context) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	var entries []bson.M
	curser, err := entryCollection.Find(ctx, bson.M{})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		fmt.Println("error in Find all")
		return
	}
	err = curser.All(ctx, &entries)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println("error in curser all")
		fmt.Println(err)
		return
	}
	defer cancel()
	fmt.Println(entries)
	c.JSON(http.StatusOK, entries)
}

func GetEntryById(c *gin.Context) {
	entryId := c.Params.ByName("id")

	DocId, _ := primitive.ObjectIDFromHex(entryId)
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

	var entry bson.M

	err := entryCollection.FindOne(ctx, bson.M{"_id": DocId}).Decode(&entry)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}
	defer cancel()
	fmt.Println(entry)
	c.JSON(http.StatusOK, entry)
	return
}

func GetEntryByIngerdients(c *gin.Context) {

}

func UpdateById(c *gin.Context) {
	entryId := c.Params.ByName("id")
	fmt.Println(entryId)

	DocId, _ := primitive.ObjectIDFromHex(entryId)
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	var entry models.Entry

	err := c.BindJSON(&entry)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}
	valEr := validate.Struct(entry)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": valEr.Error()})
		fmt.Println(err)
		return
	}

	result, rrr := entryCollection.ReplaceOne(ctx, bson.M{"id": DocId},
		bson.M{"dish": entry.Dish,
			"fat":         entry.Fat,
			"ingrediants": entry.Ingrediants,
			"calories":    entry.Calories,
		},
	)
	if rrr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": rrr.Error()})
		fmt.Println(err)
		return
	}
	defer cancel()
	c.JSON(http.StatusOK, result.ModifiedCount)
}

func UpdateByIngredients(c *gin.Context) {

}

func DeleteById(c *gin.Context) {
	entryId := c.Params.ByName("id")
	fmt.Println("entryId")

	DocId, _ := primitive.ObjectIDFromHex(entryId)
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

	result, err := entryCollection.DeleteOne(ctx, bson.M{"id":DocId})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		fmt.Println("error in entris.go line no 160")
		return
	}
	defer cancel()
	c.JSON(http.StatusOK, result.DeletedCount)
}
