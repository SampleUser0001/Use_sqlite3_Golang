package main

import (
	db "app/controller"
	"fmt"
)

var filepath string = "sqlite.db"

func main() {
	setup()
	printAll()

	fmt.Println("")
	db.Insert(filepath, "test")
	fmt.Println("Inserted test")
	fmt.Println("")

	printAll()
}

func printAll() {
	fmt.Println("Printing all:")
	for _, v := range db.SelectAll(filepath) {
		fmt.Println(v.ToString())
	}
	fmt.Println("Printing all: finish.")
}

func setup() {
	db.DeleteAll(filepath)
	db.Insert(filepath, "init data")
}
