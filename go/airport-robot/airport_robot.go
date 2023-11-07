package airportrobot

import "fmt"

// Write your code here.
// This exercise does not have tests for each individual task.
// Try to solve all the tasks first before running the tests.
type Greeter interface {
	LanguageName() string
	Greet(string) string
}

// --------------------------
type Italian struct{}

func (i Italian) LanguageName() string {
	return "Italian"
}
func (g Italian) Greet(name string) string {
	return fmt.Sprintf("Ciao %s!", name)
}

// --------------------------

type Portuguese struct {
}

func (i Portuguese) LanguageName() string {
	return "Portuguese"
}
func (g Portuguese) Greet(name string) string {
	return fmt.Sprintf("Olá %s!", name)
}

// --------------------------

type germanGreeter struct{}

func (g germanGreeter) LanguageName() string {
	return "German"
}

func (g germanGreeter) Greet(name string) string {
	return fmt.Sprintf("Hallo %s!", name)
}

// --------------------------

func SayHello(name string, g Greeter) string {

	return fmt.Sprintf("I can speak %s: %s", g.LanguageName(), g.Greet(name))
}
