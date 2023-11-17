package schema

import (
	"database/sql"
	"fmt"
	"os"
)

func CreateSchema(db *sql.DB) error {
	var schemaName = os.Getenv("SCHEMA")
	var createSchemaStr = fmt.Sprintf("CREATE SCHEMA IF NOT EXISTS %s", schemaName)

	_, err := db.Exec(createSchemaStr)
	if err != nil {
		return err
	}
	fmt.Println("Schema created successfully")
	return nil
}
