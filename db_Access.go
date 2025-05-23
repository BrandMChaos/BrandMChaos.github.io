package main

import (
	"database/sql"
	"fmt"
	"log"

	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "dvdrental"
)

func main() {

	dbTest()

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}

func dbTest() {

	// connection string
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	// psqlconn := fmt.Sprintf("host=%s port=%d dbname=%s sslmode=disable", host, port, dbname)

	// open database
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)

	getDVDInfo(db)

	// close database
	defer db.Close()

	// check db
	err = db.Ping()
	CheckError(err)

	fmt.Println("Connected!")
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

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

func getProjects() {

}
