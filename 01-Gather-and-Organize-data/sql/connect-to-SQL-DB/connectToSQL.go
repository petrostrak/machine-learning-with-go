package main

import (
	"database/sql"
	"fmt"
	"os"

	// pq is the libary that allows us to connect
	// to postgres with databases/sql.
	_ "github.com/lib/pq"
)

func main() {

	// Get the postgres connection URL. I have it stored in
	// an environmental variable.
	pgURL := os.Getenv("PGURL")
	if pgURL == "" {
		fmt.Println("PGURL empty")
		os.Exit(1)
	}

	// Open a database value.  Specify the postgres driver
	// for databases/sql.
	db, err := sql.Open("postgres", pgURL)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer db.Close()

	// sql.Open() does not establish any connections to the
	// database.  It just prepares the database connection value
	// for later use.  To make sure the database is available and
	// accessible, we will use db.Ping().
	if err := db.Ping(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Query the DB
	rows, err := db.Query(`
		SELECT
			sepal_length as sLength,
			sepal_width as sWidth,
			petal_length as pLength,
			petal_width as pWidth
		FROM iris
		WHERE species = $1`, "Iris-setosa")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer rows.Close()

	// Iterate over the rows, sending the results to
	// standard out.
	for rows.Next() {
		var (
			sLength float64
			sWidth  float64
			pLength float64
			pWidth  float64
		)

		err := rows.Scan(
			&sLength,
			&sWidth,
			&pLength,
			&pWidth,
		)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Printf("%.2f, %.2f, %.2f, %.2f\n", sLength, sWidth, pLength, pWidth)

		// Check for errors after we are done iterating over rows.
		err = rows.Err()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}
