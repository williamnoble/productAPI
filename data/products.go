package data

import (
	"encoding/json"
	"errors"
	"io"
	"time"
)

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	CreatedOn   string  `json:"created_on,omitempty"`
	UpdatedOn   string  `json:"updated_on,omitempty"`
	DeletedOn   string  `json:"deleted_on,omitempty"`
}

var ProductList = []*Product{
	{
		ID:          1,
		Name:        "Bananas",
		Description: "Some Juicy Bananas",
		Price:       0.30,
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
		DeletedOn:   "",
	},
	{
		ID:          2,
		Name:        "Oranges",
		Description: "Some Big Ripe Orange Oranges",
		Price:       0.68,
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
		DeletedOn:   "",
	},
}

type Products []*Product

func AddProduct(p *Product) {
	p.ID = getNextID()
	ProductList = append(ProductList, p)
}

func GetProducts() Products {
	return ProductList
}

func UpdateProduct(id int, p *Product) error {
	_, position, err := findProductByID(id)
	if err != nil {
		return err
	}
	p.ID = id
	ProductList[position] = p
	return nil
}

func getNextID() int {
	l := len(ProductList)
	nextID := ProductList[l-1].ID + 1
	return nextID
}

var ErrNotFound = errors.New("Product not found")

func findProductByID(id int) (*Product, int, error) {
	for i, product := range ProductList {
		if product.ID == id {
			return product, i, nil
		}
	}
	return nil, -1, ErrNotFound

}

func (p *Products) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
	// Note: we can simple return e.encode as it returns an error
}

func (p *Product) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(p)
}
