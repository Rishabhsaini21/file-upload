package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	
	fmt.Println("Env Loading")
	godotenv.Load()
	fmt.Println("Env Loaded")
	app := gin.Default()


	PORT := os.Getenv("PORT")

	fmt.Println("PORT:", PORT)
	app.Run(fmt.Sprintf(":%s", PORT))
}