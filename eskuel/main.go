package main

import (
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite"
)

func Delete(db *sql.DB) {
	s, err := db.Prepare("DELETE FROM stuff WHERE id = ?")
	if err != nil {
		panic(err)
	}

	r, err := s.Exec(3)
	if err != nil {
		panic(err)
	}
	a, _ := r.RowsAffected()
	fmt.Println(a)
}

func Update(db *sql.DB) {
	s, err := db.Prepare("UPDATE stuff SET name = ? WHERE id = ?")
	if err != nil {
		panic(err)
	}
	r, err := s.Exec("Pretty Kitty", 1)
	if err != nil {
		panic(err)
	}

	a, err := r.RowsAffected()
	if err != nil {
		panic(err)
	}
	fmt.Println(a)
}

func Multiple(db *sql.DB) {
	r, err := db.Query("SELECT id, qty, name FROM stuff")
	if err != nil {
		panic(err)
	}
	defer r.Close()
	for r.Next() {
		var id, qty int
		var name string
		if err := r.Scan(&id, &qty, &name); err != nil {
			panic(err)
		}
		fmt.Println(id, qty, name)
	}
}

func Single(db *sql.DB) {
	var name string
	var id, qty int
	if err := db.QueryRow("SELECT id, name, qty FROM stuff where id = ?", 1).Scan(&id, &name, &qty); err != nil {
		panic(err)
	}
	fmt.Println(id, name, qty)
}

func Create(db *sql.DB) {
	stmt, err := db.Prepare("INSERT INTO stuff VALUES(null, ?, ?)")
	if err != nil {
		// Run away!
		panic(err)
	}
	res, err := stmt.Exec(2, "Hello Kitty")
	if err != nil {
		panic(err)
	}
	a, _ := res.RowsAffected()
	fmt.Println("Writes", a)
}

func main() {
	db, err := sql.Open("sqlite", "../dev.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	// Create stuff
	Create(db)
	// Hent 1 række
	Single(db)
	// Hent flere
	Multiple(db)
	// Opdater
	Update(db)
	// Slet
	Delete(db)
}
