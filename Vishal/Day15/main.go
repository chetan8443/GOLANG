
package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"github.com/vishal/caloryTracker-app/routes"
	
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8082"
	}

	router := gin.New()
	router.Use(gin.Logger())
	// router.Use(cors.Default())
	router. Use (cors.Default())

	router.POST("/entry/create",routes.AddEntry)
	router.GET("/entries",routes.GetEntries)
	router.GET("/entry/:id",routes.GetEntryById)
	router.GET("/ingrediants/:ingrediants",routes.GetEntryByIngerdients)

	// update
	router.PUT("/entry/update/:id",routes.UpdateById)
	router.PUT("/ingrediants/update/:ingrediants",routes.UpdateByIngredients)
	router.DELETE("/entry/delete/:id",routes.DeleteById)

	router.Run(":"+port)
}
