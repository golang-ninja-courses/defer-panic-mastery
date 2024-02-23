package registrationservice

import "fmt"

func ExampleNewService_invalidConfig() {
	s, err := NewService(
		`^[a-zA-Z0-9_.+-]+@[a-zA-Z0-9-]+\.[a-zA-Z0-9-.]+$`,
		[]string{`(.*[a-z]{3,}`}, // Broken.
	)
	if err != nil {
		fmt.Println("invalid service configuration")
		return
	}

	err = s.SignUp("sensei@golang-ninja.ru", "Uqc7F4P7qY")
	mustNil(err)

	// Output:
	// invalid service configuration
}

func mustNil(err error) {
	if err != nil {
		panic(err)
	}
}
