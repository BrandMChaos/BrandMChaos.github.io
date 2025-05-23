// package main

// import (
// 	"database/sql"
// 	"fmt"
// 	"log"

// 	_ "github.com/lib/pq"
// )

// func main() {
// 	// Database connection parameters
// 	connStr := "user=yourusername dbname=yourdbname password=yourpassword host=localhost sslmode=disable"

// 	// Open a database connection
// 	db, err := sql.Open("postgres", connStr)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer db.Close()

// 	// Verify the connection
// 	err = db.Ping()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// Specify the table name you want to inspect
// 	tableName := "your_table_name"

// 	// Query to get column names
// 	query := `
// 		SELECT column_name
// 		FROM information_schema.columns
// 		WHERE table_name = $1
// 		ORDER BY ordinal_position
// 	`

// 	// Execute the query
// 	rows, err := db.Query(query, tableName)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer rows.Close()

// 	// Iterate through the results
// 	fmt.Printf("Columns in table '%s':\n", tableName)
// 	for rows.Next() {
// 		var columnName string
// 		if err := rows.Scan(&columnName); err != nil {
// 			log.Fatal(err)
// 		}
// 		fmt.Println(columnName)
// 	}

// 	// Check for errors from iterating over rows
// 	if err = rows.Err(); err != nil {
// 		log.Fatal(err)
// 	}
// }
