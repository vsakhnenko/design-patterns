package factory

import (
	"abstract_factory/abstract_product"
	"abstract_factory/product"
)

// NonFictionFactory is a concrete factory that implements IBookShopFactory
// It creates concrete products: NonFictionBook and HistoryMagazine
type NonFictionFactory struct{}

func (f *NonFictionFactory) MakeBook() abstract_product.IBook {
	return &product.NonFictionBook{
		Book: abstract_product.Book{
			Title:  "Non-Fictional Book",
			Author: "Jane Doe",
		},
	}
}

func (f *NonFictionFactory) MakeMagazine() abstract_product.IMagazine {
	return &product.HistoryMagazine{
		Magazine: abstract_product.Magazine{
			Title:     "History Magazine",
			Publisher: "History Publisher",
		},
	}
}
