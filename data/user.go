package data

type User struct {
	ID       int        `json:"id"`
	Name     string     `json:"name"`
	Balance  float64    `json:"balance"`
	Expences []Expenses `json:"expenses"`
}

func GetUsers() []User {

	var users = []User{
		{ID: 1, Name: "John", Balance: 1000.00, Expences: GetFirstExpense()},
		{ID: 2, Name: "Megan", Balance: 5000.00, Expences: GetSecondExpense()},
		{ID: 3, Name: "Michael", Balance: 15000.00, Expences: GetThirdExpense()},
	}

	return users
}
