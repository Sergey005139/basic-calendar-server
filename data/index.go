package data

import "errors"

type UserStorage struct {
	users []User
}

func (us *UserStorage) FindByUsername(v string) *User {
	for _, r := range us.users {
		if r.Username == v {
			return &r
		}
	}
	return nil
}

func (us *UserStorage) FindByUUID(v string) *User {
	for _, r := range us.users {
		if r.UUID == v {
			return &r
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
	UUID string
	Username string
	Password string
}

type Expense struct {
	Date string
	Category string
	Amount float64
}
