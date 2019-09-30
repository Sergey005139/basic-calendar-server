package data

import "errors"

type UserStorage struct {
	users []User
}

func (us *UserStorage) FindByUsername(v string) *User {
	for i, r := range us.users {
		if r.Username == v {
			return &(us.users[i])
		}
	}
	return nil
}

func (us *UserStorage) FindByUUID(v string) *User {
	for i, r := range us.users {
		if r.UUID == v {
			return &(us.users[i])
		}
	}
	return nil
}

func (us *UserStorage) IsExist(u *User) bool {
	if user := us.FindByUsername(u.Username); user != nil {
		return true
	}
	return false
}

func (us *UserStorage) Insert(u *User) (bool, error) {
	if us.IsExist(u) {
		return false, errors.New("can't insert user to storage because user with same username already exist")
	}
	us.users = append(us.users, *u)
	return true, nil
}

type User struct {
	UUID string `json:"-"`
	Username string `json:"username"`
	Password string `json:"-"`
	Expenses []Expense `json:"-"`
}

func (u *User) AddExpense(e Expense) {
	u.Expenses = append(u.Expenses, e)
}

type Expense struct {
	Date string `json:"date"`
	Category string `json:"category"`
	Amount float64 `json:"amount"`
}
