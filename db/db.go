package db

import "payment/model"

var Fruits = []model.Product{
	{
		Id:          1,
		Title:       "lemon",
		Description: "fruit with citric acid",
		Price:       5,
	},
	{
		Id:          2,
		Title:       "orange",
		Description: "my favorite fruit",
		Price:       10,
	},
	{
		Id:          3,
		Title:       "Banana",
		Description: "fruit liked by children",
		Price:       5,
	},
}

func GetAllProducts() []model.Product {
	return Fruits
}

func CreateProduct(f model.Product) model.Product {
	Fruits = append(Fruits, model.Product{
		Id:          len(Fruits) + 1,
		Title:       f.Title,
		Description: f.Description,
		Price:       f.Price,
	})

	return Fruits[len(Fruits)-1]
}

func GetAProduct(id int) model.Product {
	for _, fruit := range Fruits {
		if fruit.Id == id {
			return fruit
		}
	}
	return model.Product{}
}
