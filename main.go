package main

import (
	routes "example.com/echoes/routes"
	"github.com/gin-gonic/gin"
	"github.com/gofiber/fiber/v2"
)

func main() {

	//app := fiber.New()

   // app.Get("/", func(c *fiber.Ctx) error {
       // return c.SendString("Hello, World 👋!")
   // })

   //app.Listen(":3000")

router := gin.Default()

router.POST("/", routes.CreatePost)

// called as localhost:3000/getOne/{id}
router.GET("getOne/:postId", routes.ReadOnePost)

// called as localhost:3000/update/{id}
router.PUT("/update/:postId", routes.UpdatePost)

// called as localhost:3000/delete/{id}
router.DELETE("/delete/:postId", routes.DeletePost)

router.Run("localhost: 3000")
}