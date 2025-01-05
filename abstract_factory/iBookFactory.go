package main

import (
	"abstract_factory/abstract_product"
	"abstract_factory/factory"
	"fmt"
)

// IBookShopFactory is the abstract factory interface
// It declares methods for creating abstract products
type IBookShopFactory interface {
	MakeBook() abstract_product.IBook
	MakeMagazine() abstract_product.IMagazine
}

func GetBookShopFactory(category string) (IBookShopFactory, error) {
	if category == "fiction" {
		return &factory.FictionFactory{}, nil
	} else if category == "non-fiction" {
		return &factory.NonFictionFactory{}, nil
	}

	return nil, fmt.Errorf("Wrong category type passed")
}
