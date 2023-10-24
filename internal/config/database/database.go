package database

import (
	"database/sql"
	"fmt"

	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func InitDb() *sql.DB {

	var err error

	godotenv.Load()

	var (
		user    = os.Getenv("USER")
		pass    = os.Getenv("PASS")
		dbname  = os.Getenv("DB")
		host    = os.Getenv("HOST")
		port    = os.Getenv("PORT")
		sslmode = os.Getenv("SSLMODE")
	)

	connectString := fmt.Sprintf("user=%s dbname=%s password=%s host=%s port=%s sslmode=%s", user, dbname, pass, host, port, sslmode)

	db, err := sql.Open("postgres", connectString)

	if err != nil {
		fmt.Println("Error opening database connection", err.Error())
		panic(err)
	}

	fmt.Println("Database connected!")

	return db
}

func CloseDb(db *sql.DB) {
	defer db.Close()
}
