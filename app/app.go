package main

import (
	db "app/controller"
	"fmt"
)

func main() {
	setup()
	printAll()

	fmt.Println("")
	db.Insert("test")
	fmt.Println("Inserted test")
	fmt.Println("")

	printAll()
}

func printAll() {
	fmt.Println("Printing all:")
	for _, v := range db.SelectAll() {
		fmt.Println(v.ToString())
	}
	fmt.Println("Printing all: finish.")
}

func setup() {
	db.Filepath = "sqlite.db"
	db.DeleteAll()
	db.Insert("init data")
}
