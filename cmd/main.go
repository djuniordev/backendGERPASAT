package main

import (

	"github.com/djuniordev/SAST-MA/database"
	"github.com/djuniordev/SAST-MA/routes"

)

func main() {
	database.ConectaComBancoDeDados()
	routes.Initialize()
}