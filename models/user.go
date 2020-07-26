package models

import (
	"errors"
	"fmt"
)

// User represents a user
type User struct {
	ID        int
	FirstName string
	LastName  string
}

var (
	users  []*User
	nextID = 1
)

// GetUsers returns the list of users
func GetUsers() []*User {
	return users
}

// AddUser adds a new user to the list of users
func AddUser(u User) (User, error) {
	if u.ID != 0 {
		return User{}, errors.New("The new user can't have an ID")
	}
	u.ID = nextID
	nextID++
	users = append(users, &u)
	return u, nil
}

// GetUserByID returns an user passing the Id
func GetUserByID(id int) (User, error) {
	for _, u := range users {
		if u.ID == id {
			return *u, nil
		}
	}
	return User{}, fmt.Errorf("User non existent: %v", id)
}

// UpdateUser updates user
func UpdateUser(u User) (User, error) {
	for i, candidate := range users {
		if candidate.ID == u.ID {
			users[i] = &u
			return u, nil
		}
	}
	return User{}, fmt.Errorf("User non existent: %v", u.ID)
}

// RemoveUserByID deletes an user passed to the function
func RemoveUserByID(id int) error {
	for i, u := range users {
		if u.ID == id {
			users = append(users[:i], users[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("User non existent: %v", id)
}
