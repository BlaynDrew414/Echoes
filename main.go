package main

import (
	routes "example.com/echoes/routes"
	"github.com/gin-gonic/gin"
	
)

func main() {

router := gin.Default()

router.POST("/", routes.CreateEcho)

// called as localhost:3000/getOne/{id}
router.GET("getOne/:echoId", routes.ReadOneEcho)

// called as localhost:3000/update/{id}
router.PUT("/update/:echoId", routes.UpdateEcho)

// called as localhost:3000/delete/{id}
router.DELETE("/delete/:echoId", routes.DeleteEcho)

router.Run("localhost: 3000")
}