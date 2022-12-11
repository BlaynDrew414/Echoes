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

func UpdatePost(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	var DB = database.ConnectDB()
	var postCollection = getcollection.GetCollection(DB, "Posts")

	postId := c.Param("postId")
	var post model.Posts
	
	defer cancel()


	objId, _ := primitive.ObjectIDFromHex(postId)

	if err := c.BindJSON(&post); err != nil {
		c.JSON (http.StatusInternalServerError, gin.H{"message": err})
		return 
	}

	edited := bson.M{"title": post.Title, "article": post.Article}
	
	result, err := postCollection.UpdateOne(ctx, bson.M{"id": objId}, bson.M{"$set": edited})

	res := map [string]interface{}{"data": result}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}

	if result.MatchedCount < 1 {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Data doesnt exist"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "data updated sucessfully!", "Data": res})
}
