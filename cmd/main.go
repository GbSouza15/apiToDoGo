package main

import (
	"fmt"
	"os"

	"github.com/GbSouza15/apiToDoGo/internal/app/routers"
	"github.com/GbSouza15/apiToDoGo/internal/database"
	"github.com/GbSouza15/apiToDoGo/internal/database/schema"
)

func main() {
	db, err := database.InitDb()
	if err != nil {
		fmt.Printf("Error starting the database, %v", err.Error())
		os.Exit(1)
	}

	defer db.Close()

	err = schema.CreateSchema(db)
	if err != nil {
		fmt.Printf("Schema error: %v", err)
	}

	err = routers.RoutesApi(db)
	if err != nil {
		fmt.Printf("Error starting the server: %v", err)
		os.Exit(1)
	}
}
