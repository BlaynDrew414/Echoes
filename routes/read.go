package routes

import (
	getcollection "example.com/echoes/collection"
	database "example.com/echoes/databases"
	model "example.com/echoes/models"
	"context"
	"net/http"
	"time"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ReadOneEcho(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var DB = database.ConnectDB()
	var echoCollection = getcollection.GetCollection(DB, "Echoes")
	
	echoId := c.Param("echoId")
	var result model.Echoes
	
	defer cancel()
	
	objId, _ := primitive.ObjectIDFromHex(echoId)
	
	err := echoCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&result)
	
	res := map[string]interface{}{"data": result}
	
	if err != nil {
	c.JSON(http.StatusInternalServerError, gin.H{"message": err})
	return
	}
	
	c.JSON(http.StatusCreated, gin.H{"message": "success!", "Data": res})
   }