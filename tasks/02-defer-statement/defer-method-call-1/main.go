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

func processUser(u *User) {
	defer fmt.Printf("user %q was processed", u.Email())
	defer u.SetEmail("unknown")
	// ...
}

func main() {
	u := NewUser("sensei@golang-ninja.ru")
	processUser(u)
}
