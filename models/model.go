package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	)
	
	type Posts struct {
	ID       primitive.ObjectID  
	Title    string 			 `json:"title" bson:"title"`
	Article  string  		     `json:"article" bson:"article"`
	}


