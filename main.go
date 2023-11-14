package main

import "fmt"

type Book struct {
    Title  string
    Author string
    ISBN   string
}

func displayBook(b Book) {
    fmt.Printf("Title: %s\nAuthor: %s\nISBN: %s\n", b.Title, b.Author, b.ISBN)
}