package model

type AvailableItem struct {
	Slug        string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
}

var AvailableItems []AvailableItem = []AvailableItem{
	{
		Slug:        "java",
		Name:        "Java",
		Description: "A best-seller coffee in Jogja",
		Price:       22000.0,
	},
	{
		Slug:        "black",
		Name:        "Black",
		Description: "A healthier option rather than Java",
		Price:       18000.0,
	},
	{
		Slug:        "white",
		Name:        "White",
		Description: "Not healthier than black, but healthier than Java",
		Price:       20000.0,
	},
}
