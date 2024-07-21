package data

type Category struct {
	ID   int    `json:"id"`
	Type string `json:"type"`
}

func GetFoodCategory() Category {
	return Category{ID: 1, Type: "Food"}
}

func GetHealthCategory() Category {
	return Category{ID: 2, Type: "Health"}
}

func GetEmergencyCategory() Category {
	return Category{ID: 3, Type: "Emergency"}
}
