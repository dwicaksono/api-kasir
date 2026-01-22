package category

type Category struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

var DataCategories = []Category{
	{ID: 1, Name: "Makanan", Description: "Berbagai jenis makanan"},
	{ID: 2, Name: "Minuman", Description: "Berbagai jenis minuman"},
}
