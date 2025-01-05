package factory

import (
	"abstract_factory/abstract_product"
	"abstract_factory/product"
)

// FictionFactory is a concrete factory that implements IBookShopFactory
// It creates concrete products: FictionBook and ScienceMagazine
type FictionFactory struct{}

func (f *FictionFactory) MakeBook() abstract_product.IBook {
	return &product.FictionBook{
		Book: abstract_product.Book{
			Title:  "Fictional Book",
			Author: "John Doe",
		},
	}
}

func (f *FictionFactory) MakeMagazine() abstract_product.IMagazine {
	return &product.ScienceMagazine{
		Magazine: abstract_product.Magazine{
			Title:     "Science Magazine",
			Publisher: "Science Publisher",
		},
	}
}
