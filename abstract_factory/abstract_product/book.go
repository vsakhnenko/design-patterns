package abstract_product

// IBook is the abstract product interface for books
type IBook interface {
	GetTitle() string
	GetAuthor() string
}

// Book is a concrete implementation of IBook
type Book struct {
	Title  string
	Author string
}

func (b *Book) GetTitle() string {
	return b.Title
}

func (b *Book) GetAuthor() string {
	return b.Author
}
