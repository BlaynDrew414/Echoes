package routes

import (
	getcollection "example.com/echoes/collection"
	database "example.com/echoes/databases"
	"context"
	"net/http"
	"time"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	)



func DeleteEcho(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var DB = database.ConnectDB()
	echoId := c.Param("echoId")
		
		var echoCollection = getcollection.GetCollection(DB, "Echoes")
		defer cancel()
		objId, _ := primitive.ObjectIDFromHex(echoId)
		result, err := echoCollection.DeleteOne(ctx, bson.M{"id": objId})
		res := map[string]interface{}{"data": result}
		
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err})
			return
			}
		
		if result.DeletedCount < 1 {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "No data to delete"})
			return
			}
		
		c.JSON(http.StatusCreated, gin.H{"message": "Echo deleted", "Data": res})
}
		
	
