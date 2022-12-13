package routes

import (
	getcollection "example.com/echoes/collection"
	database "example.com/echoes/databases"
	model "example.com/echoes/models"
	"context"
	"log"
	"net/http"
	"time"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateEcho(c *gin.Context) {
var DB = database.ConnectDB()
var echoCollection = getcollection.GetCollection(DB, "Echoes")
ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
echo := new(model.Echoes)
defer cancel()

if err := c.BindJSON(&echo); err != nil {
c.JSON(http.StatusBadRequest, gin.H{"message": err})
log.Fatal(err)
return
}

echoPayload := model.Echoes {
	ID: primitive.NewObjectID(),
	Echo: echo.Echo,
	Book: echo.Book,
	Author: echo.Author,
	}

result, err := echoCollection.InsertOne(ctx, echoPayload)

if err != nil {
c.JSON(http.StatusInternalServerError, gin.H{"message": err})
return
}

c.JSON(http.StatusCreated, gin.H{"message": "Echo created", "Data": map[string]interface{}{"data": result}})

}


