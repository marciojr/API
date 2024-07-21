package data

type User struct {
	ID       int        `json:"id"`
	Name     string     `json:"name"`
	Balance  float64    `json:"balance"`
	Expenses []Expenses `json:"expenses"`
}

var Users = []User{
	{ID: 1, Name: "John", Balance: 1000.00, Expenses: GetFirstExpense()},
	{ID: 2, Name: "Megan", Balance: 5000.00, Expenses: GetSecondExpense()},
	{ID: 3, Name: "Michael", Balance: 15000.00, Expenses: GetThirdExpense()},
}

func AddUser(newUser User) {
	Users = append(Users, newUser)
}
