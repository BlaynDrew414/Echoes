package routes

import (
	"context"
	"net/http"
	"time"

	getcollection "example.com/echoes/collection"
	database "example.com/echoes/databases"
	model "example.com/echoes/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func UpdateEcho(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	var DB = database.ConnectDB()
	var echoCollection = getcollection.GetCollection(DB, "Echoes")

	echoId := c.Param("echoId")
	var echo model.Echoes
	
	defer cancel()


	objId, _ := primitive.ObjectIDFromHex(echoId)

	if err := c.BindJSON(&echo); err != nil {
		c.JSON (http.StatusInternalServerError, gin.H{"message": err})
		return 
	}

	edited := bson.M{"echo": echo.Echo, "book": echo.Book, "author": echo.Author}
	
	result, err := echoCollection.UpdateOne(ctx, bson.M{"id": objId}, bson.M{"$set": edited})

	res := map [string]interface{}{"data": result}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}

	if result.MatchedCount < 1 {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Data doesnt exist"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "echo updated", "Data": res})
}
