package database

import (
	"database/sql"
	"fmt"

	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func InitDb() (*sql.DB, error) {
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
		return nil, err
	}

	fmt.Println("Database connected!")
	return db, nil
}

//func CreateTables(db *sql.DB) error {
//	createUserTable := `
//		CREATE TABLE IF NOT EXISTS tdlist.users (
//			id UUID PRIMARY KEY,
//			name VARCHAR(255) NOT NULL,
//			email VARCHAR(255) NOT NULL,
//			password VARCHAR(255) NOT NULL
//		);
//	`
//
//	createTaskTable := `
//		CREATE TABLE IF NOT EXISTS tdlist.tasks (
//			id UUID PRIMARY KEY,
//			title VARCHAR(255) NOT NULL,
//			description VARCHAR(255) NOT NULL,
//			user_id UUID REFERENCES tdlist.users(id)
//		);
//	`
//
//	_, err := db.Exec(createUserTable)
//	if err != nil {
//		return fmt.Errorf("Error creating the User table: %s", err.Error())
//	}
//
//	_, err = db.Exec(createTaskTable)
//	if err != nil {
//		return fmt.Errorf("Error creating the Task table: %s", err.Error())
//	}
//
//	return nil
//}
