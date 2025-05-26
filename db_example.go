package main

import (
	"database/sql"
	"fmt"
	"log"
)

// This is example code for a self-hosted database called dvdrental, that helps to understand how psql queries can be properly formatted
// You can find the repository here: https://neon.tech/postgresql/postgresql-getting-started/postgresql-sample-database

// dvd rental example
func getDVDInfo(db *sql.DB) {
	rows, err := db.Query(`
							SELECT column_name
							FROM information_schema.columns
							WHERE table_name = 'actor' 
							`)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	defer rows.Close()

	// for _, col := range cols {
	// 	fmt.Println(col)

	// }

	fmt.Println("COLUMNS IN DVDRENTAL TABLE:")
	fmt.Println("-----------------------------")

	count := 1
	for rows.Next() {
		var columnName string
		if err := rows.Scan(&columnName); err != nil {
			log.Fatal("Failed to scan row:", err)
		}
		fmt.Printf("%d. %s\n", count, columnName) // Numbered list
		count++
	}

	fmt.Println("-----------------------------")
	fmt.Printf("Total columns: %d\n", count-1)

}
