package main

import (
	"homeworkGun/db_operation"

	"homeworkGun/handlers"
	"homeworkGun/middleware"
	"log"

	"github.com/gin-gonic/gin"

	_ "github.com/lib/pq"
)

func main() {
	check_config()

	r := gin.Default()
	db := db_operation.ConnectToDB()
	r.Use(func(c *gin.Context) {
		db_operation.SetDBToContext(c, db)
		c.Next()
	})
	r.GET("/crud", db_operation.QueryHandler)

	r.POST("/login", handlers.LoginHandler)
	r.POST("/register", handlers.RegisterHandler)

	authGroup := r.Group("/").Use(middleware.TokenAuthMiddleware())
	authGroup.GET("/user", handlers.UserProfileHandler)

	log.Fatal(r.Run("0.0.0.0:1145")) // listen and serve on 0.0.0.0:1145
}
