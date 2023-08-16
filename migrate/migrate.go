package main

import (
	"github.com/Jx2Jorge/TUTORIAL.GO/config"
	"github.com/Jx2Jorge/TUTORIAL.GO/models"
)

func init() {
	config.ConnectToDB()
}

func main() {
	config.DB.AutoMigrate(&models.Person{})
}
