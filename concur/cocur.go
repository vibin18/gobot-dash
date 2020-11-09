package main

import (
	"fmt"
	"time"
)

var Cache = []Book{
	{
		Id:     2,
		Name:   "Mastering Golang",
		Author: "Google",
	},
}

func CheckDataCache(id int, cache []Book) bool {
	for _, item := range cache {
		if item.Id == id {
			fmt.Println("Fetched from cache")
			return true
		}
	}
	return false
}
func CheckDataDb(id int, db []Book) bool {
	for _, item := range db {
		time.Sleep(5 * time.Millisecond)
		if item.Id == id {
			fmt.Println("Fetched from Database")
			Cache = append(Cache, item)
			return true
		}
	}
	return false
}

func displayBooks(book Book) {
	fmt.Printf("Name : %s\n", book.Name)
	fmt.Printf("Author : %s\n", book.Author)
	fmt.Println("==========================================")
}
func display(books []Book) {
	for _, item := range books {
		displayBooks(item)
	}
}
func main() {
	for i, item := range Books {
		i = i + 1
		if CheckDataCache(i, Cache) {
			displayBooks(item)
		} else if CheckDataDb(i, Books) {
			displayBooks(item)
		}

	}

}
