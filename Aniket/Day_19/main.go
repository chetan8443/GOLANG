package main

import (
	"log"
)

func main() {
	// connect to DB first
	// env := new(Env)
	var err error
	db, err := ConnectDB()
	if err != nil {
		log.Fatalf("failed to start the server: %v", err)
	}

	createStmt := "CREATE TABLE IF NOT EXISTS artist (name VARCHAR (50), title VARCHAR (50), price VARCHAR (255)"

	_, err = db.Exec(createStmt)
	CheckError(err)

	getStmt := "select * from artist"
	_, err = db.Exec(getStmt)
	CheckError(err)

	// router := gin.Default()
	// // router.GET("/albums/:id", env.GetAlbumByID)
	// router.GET("/albums", env.GetAlbums)
	// router.POST("/albums", env.PostAlbum)
	// // router.PUT("/albums", env.UpdateAlbum)
	// // router.DELETE("/albums/:id", env.DeleteAlbumByID)

	// router.Run("localhost:8080")
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
