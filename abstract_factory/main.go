package main

import (
	"abstract_factory/abstract_product"
	"fmt"
)

func main() {
	fictionFactory, _ := GetBookShopFactory("fiction")
	nonFictionFactory, _ := GetBookShopFactory("non-fiction")

	book1 := fictionFactory.MakeBook()
	magazine1 := fictionFactory.MakeMagazine()

	book2 := nonFictionFactory.MakeBook()
	magazine2 := nonFictionFactory.MakeMagazine()

	printBookDetails(book1)
	printMagazineDetails(magazine1)

	printBookDetails(book2)
	printMagazineDetails(magazine2)
}

func printBookDetails(b abstract_product.IBook) {
	fmt.Printf("Title: %s", b.GetTitle())
	fmt.Println()
	fmt.Printf("Author: %s", b.GetAuthor())
	fmt.Println()
}

func printMagazineDetails(m abstract_product.IMagazine) {
	fmt.Printf("Title: %s", m.GetTitle())
	fmt.Println()
	fmt.Printf("Publisher: %s", m.GetPublisher())
	fmt.Println()
}
