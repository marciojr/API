package data

type Expenses struct {
	ID       int      `json:"id"`
	Name     string   `json:"name"`
	Price    float64  `json:"price"`
	Category Category `json:"category"`
}

func GetFirstExpense() []Expenses {

	var expenses = []Expenses{
		{ID: 1, Name: "Hamburguer", Price: 30.00, Category: GetFoodCategory()},
		{ID: 2, Name: "Dr. Andrew", Price: 120.00, Category: GetHealthCategory()},
		{ID: 3, Name: "St. Joseph Hospital Emergency", Price: 180.00, Category: GetEmergencyCategory()},
	}

	return expenses
}

func GetSecondExpense() []Expenses {

	var expenses = []Expenses{
		{ID: 1, Name: "Ice Cream", Price: 5.00, Category: GetFoodCategory()},
		{ID: 2, Name: "Dra. Andrea", Price: 130.00, Category: GetHealthCategory()},
		{ID: 3, Name: "Dentist Emergency", Price: 90.00, Category: GetEmergencyCategory()},
	}

	return expenses
}

func GetThirdExpense() []Expenses {

	var expenses = []Expenses{
		{ID: 1, Name: "Lemon Juice", Price: 10.00, Category: GetFoodCategory()},
		{ID: 2, Name: "Dr. Philips", Price: 120.00, Category: GetHealthCategory()},
		{ID: 3, Name: "New Door to my house", Price: 180.00, Category: GetEmergencyCategory()},
	}

	return expenses
}
