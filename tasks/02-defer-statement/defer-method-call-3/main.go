package main

import "fmt"

type User struct {
	email string
}

func NewUser(e string) *User {
	return &User{email: e}
}

func (u *User) SetEmail(e string) { u.email = e }
func (u *User) Email() string     { return u.email }
func (u User) PrintEmail()        { fmt.Printf("user %q was processed", u.Email()) }

func processUser(u *User) {
	defer u.PrintEmail()
	// defer User.PrintEmail(*u)
	defer u.SetEmail("unknown")
	// ...
}

func main() {
	u := NewUser("info@golang-courses.ru")
	processUser(u)
}
