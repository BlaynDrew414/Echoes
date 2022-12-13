package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	)
	
	type Echoes struct {
	ID       primitive.ObjectID  
	Echo    string 			 `json:"echo" bson:"echo"`
	Book  	string  		 `json:"book" bson:"book"`
	Author  string  		 `json:"author" bson:"author"`
	}


