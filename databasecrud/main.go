// package main

// import (
// 	"database/sql"
// 	"fmt"
// 	"log"

// 	_ "modernc.org/sqlite"
// )

// type User struct {
// 	Id   int16
// 	Name string
// }

// func listDrivers() {
// 	for _, driver := range sql.Drivers() {
// 		fmt.Printf("Driver: %v\n", driver)
// 	}
// }

// func queryDatabase(db *sql.DB) []User {
// 	users := []User{}
// 	rows, err := db.Query("SELECT * FROM names")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer rows.Close()

// 	for rows.Next() {
// 		u := User{}
// 		err := rows.Scan(&u.Id, &u.Name)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		users = append(users, u)
// 	}

// 	return users
// }

// func openDatabase() (*sql.DB, error) {
// 	db, err := sql.Open("sqlite", "names.db")
// 	if err != nil {
// 		return nil, err
// 	}
// 	return db, nil
// }

// func insertDatabase(db *sql.DB, u *User) int64 {
// 	res, err := db.Exec(`INSERT INTO names (id, name) VALUES (?, ?)`, u.Id, u.Name)
// 	if err != nil {
// 		fmt.Println("Insert Error:", err.Error())
// 		return 0
// 	}

// 	id, err := res.LastInsertId()
// 	if err != nil {
// 		fmt.Println("LastInsertId Error:", err.Error())
// 		return 0
// 	}
// 	fmt.Printf("Inserted User with ID %v\n", id)
// 	return id
// }

// func updateUser(db *sql.DB, u *User) []User {
// 	res, err := db.Exec(`UPDATE names SET name = ? WHERE id = ?`, u.Name, u.Id)
// 	if err != nil {
// 		fmt.Println("Update Error:", err.Error())
// 		return queryDatabase(db)
// 	}

// 	numR, err := res.RowsAffected()
// 	if err == nil {
// 		fmt.Printf("Rows affected: %d\n", numR)
// 	} else {
// 		fmt.Println("RowsAffected Error:", err.Error())
// 	}

// 	return queryDatabase(db)
// }

// func deleteUser(db *sql.DB, u *User) []User {
// 	res, err := db.Exec(`DELETE FROM names WHERE id = ?`, u.Id)
// 	if err != nil {
// 		fmt.Println("Delete Error:", err.Error())
// 		return queryDatabase(db)
// 	}

// 	numR, err := res.RowsAffected()
// 	if err == nil {
// 		fmt.Printf("Rows affected: %d\n", numR)
// 	} else {
// 		fmt.Println("RowsAffected Error:", err.Error())
// 	}

// 	return queryDatabase(db)
// }

// func main() {
// 	listDrivers()

// 	db, err := openDatabase()
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer db.Close()

// 	fmt.Println("Opened database successfully")

// 	// Insert a new user
// 	usr := User{Id: 6, Name: "John"}
// 	insertDatabase(db, &usr)

// 	// Update the existing user
// 	usr2 := User{Id: 6, Name: "Johnny"}
// 	updateUser(db, &usr2)

// 	// deleteUser
// 	deleteUser(db, &usr2)

// 	// Query and print all users
// 	users := queryDatabase(db)
// 	for i, u := range users {
// 		fmt.Printf("#%v, User: %v\n", i, u)
// 	}
// }
// --------------------------------------------------

package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "modernc.org/sqlite"
)

type User struct {
	Id   int16
	Name string
}

func listDrivers() {
	for _, driver := range sql.Drivers() {
		fmt.Printf("Driver: %v\n", driver)
	}
}

func queryDatabase(ctx context.Context, db *sql.DB) []User {
	users := []User{}

	// Using the context with the query
	rows, err := db.QueryContext(ctx, "SELECT * FROM names")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		u := User{}
		err := rows.Scan(&u.Id, &u.Name)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, u)
	}

	return users
}

func openDatabase() (*sql.DB, error) {
	db, err := sql.Open("sqlite", "names.db")
	if err != nil {
		return nil, err
	}
	return db, nil
}

func insertDatabase(ctx context.Context, db *sql.DB, u *User) int64 {
	// Using the context with the insert operation
	res, err := db.ExecContext(ctx, `INSERT INTO names (id, name) VALUES (?, ?)`, u.Id, u.Name)
	if err != nil {
		fmt.Println("Insert Error:", err.Error())
		return 0
	}

	id, err := res.LastInsertId()
	if err != nil {
		fmt.Println("LastInsertId Error:", err.Error())
		return 0
	}
	fmt.Printf("Inserted User with ID %v\n", id)
	return id
}

func updateUser(ctx context.Context, db *sql.DB, u *User) []User {
	// Using the context with the update operation
	res, err := db.ExecContext(ctx, `UPDATE names SET name = ? WHERE id = ?`, u.Name, u.Id)
	if err != nil {
		fmt.Println("Update Error:", err.Error())
		return queryDatabase(ctx, db)
	}

	numR, err := res.RowsAffected()
	if err == nil {
		fmt.Printf("Rows affected: %d\n", numR)
	} else {
		fmt.Println("RowsAffected Error:", err.Error())
	}

	return queryDatabase(ctx, db)
}

func deleteUser(ctx context.Context, db *sql.DB, u *User) []User {
	// Using the context with the delete operation
	res, err := db.ExecContext(ctx, `DELETE FROM names WHERE id = ?`, u.Id)
	if err != nil {
		fmt.Println("Delete Error:", err.Error())
		return queryDatabase(ctx, db)
	}

	numR, err := res.RowsAffected()
	if err == nil {
		fmt.Printf("Rows affected: %d\n", numR)
	} else {
		fmt.Println("RowsAffected Error:", err.Error())
	}

	return queryDatabase(ctx, db)
}

func main() {
	listDrivers()

	db, err := openDatabase()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	fmt.Println("Opened database successfully")

	// Create a context with a timeout of 5 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Insert a new user
	usr := User{Id: 6, Name: "John"}
	insertDatabase(ctx, db, &usr)

	// Update the existing user
	usr2 := User{Id: 6, Name: "Johnny"}
	updateUser(ctx, db, &usr2)

	// deleteUser
	deleteUser(ctx, db, &usr2)

	// Query and print all users
	users := queryDatabase(ctx, db)
	for i, u := range users {
		fmt.Printf("#%v, User: %v\n", i, u)
	}
}
