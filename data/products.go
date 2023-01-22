package data

import "time"

type Product struct {
	Id        int
	Name      string
	Desc      string
	Price     float32
	SKU       string //internal id
	CreatedOn string
	UpdatedOn string
	DeletedOn string
}

var productList = []*Product{
	{
		Id:        1,
		Name:      "green tea",
		Desc:      "the most green tea you have ever tasted",
		Price:     9.25,
		SKU:       "tea31",
		CreatedOn: time.Now().UTC().GoString(),
		UpdatedOn: time.Now().UTC().GoString(),
	},
	{
		Id:        2,
		Name:      "british black tea",
		Desc:      "the most british tea you have ever seen",
		Price:     11.9,
		SKU:       "brits21",
		CreatedOn: time.Now().UTC().GoString(),
		UpdatedOn: time.Now().UTC().GoString(),
	},
}
