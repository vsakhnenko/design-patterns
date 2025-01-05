package abstract_product

// IMagazine is the abstract product interface for magazines
type IMagazine interface {
	GetTitle() string
	GetPublisher() string
}

// Magazine is a concrete implementation of IMagazine
type Magazine struct {
	Title     string
	Publisher string
}

func (m *Magazine) GetTitle() string {
	return m.Title
}

func (m *Magazine) GetPublisher() string {
	return m.Publisher
}
