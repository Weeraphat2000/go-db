// go mod init github.com/Weeraphat2000/db-go
// go mod tidy
// go run main.go

package main

import (
	"database/sql"
	"fmt"

	_ "github.com/proullon/ramsql/driver"
)

func main() {
	// TODO: sql.Open("ชื่อไดรเวอร์", "ชื่อฐานข้อมูล") เช่น ("mysql", "user")
	db, err := sql.Open("ramsql", "users")
	if err != nil {
		fmt.Println("Error opening database:", err)
		panic(err) // TODO: panic คือ การหยุดการทำงานของโปรแกรมทันที เช่น การเชื่อมต่อฐานข้อมูลล้มเหลว, index out of range, etc.
	}
	// defer db.Close()

	createTable := `
	CREATE TABLE IF NOT EXISTS users (
	id INT AUTO_INCREMENT PRIMARY KEY,
	name VARCHAR(100) NOT NULL,
	email VARCHAR(100) NOT NULL UNIQUE
	);`

	_, err = db.Exec(createTable) // raw SQL query
	if err != nil {
		fmt.Println("Error creating table:", err)
		panic(err)
	}
	fmt.Println("Table created successfully")

	insertUser := `INSERT INTO users (name, email) VALUES (?, ?);`

	// prepareStmt, err := db.Prepare(insertUser) // prepare statement
	// if err != nil {
	// 	fmt.Println("Error preparing statement:", err)
	// 	panic(err)
	// }

	r, err := db.Exec(insertUser, "John Doe", "john@example.com") // raw SQL query with parameters
	if err != nil {
		fmt.Println("Error inserting user:", err)
		panic(err)
	}
	fmt.Println("User inserted successfully")
	lastInsertID, err := r.LastInsertId() // get last insert ID
	if err != nil {
		fmt.Println("Error getting last insert ID:", err)
		panic(err)
	}
	fmt.Println("Last Insert ID:", lastInsertID)

	effectedRows, err := r.RowsAffected() // get affected rows
	if err != nil {
		fmt.Println("Error getting affected rows:", err)
		panic(err)
	}
	fmt.Println("Affected Rows:", effectedRows)

	query := `SELECT * FROM users;`
	rows, err := db.Query(query) // raw SQL query
	if err != nil {
		fmt.Println("Error querying users:", err)
		panic(err)
	}
	for rows.Next() {
		var id int
		var name, email string
		err = rows.Scan(&id, &name, &email) // scan the result into variables
		if err != nil {
			fmt.Println("Error scanning row:", err)
			panic(err)
		}
		fmt.Printf("User: ID=%d, Name=%s, Email=%s\n", id, name, email)
	}
}
