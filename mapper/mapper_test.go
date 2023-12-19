package mapper

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

type Person struct {
	Firstname  string `json:"firstname"`
	SecondName string `json:"secondname"`
	LastName   string `json:"lastname"`
}

type PartialPerson struct {
	Firstname string `json:"firstname"`
	LastName  string `json:"lastname"`
}

func TestMap(t *testing.T) {

	person := Person{
		Firstname:  "First name",
		SecondName: "Second name",
		LastName:   "Last name",
	}

	var partialPerson PartialPerson

	if err := Map(context.Background(), person, &partialPerson); err != nil {
		fmt.Println("error occured while mapping person:", err.Error())
	}

	require.Equal(t, person.Firstname, partialPerson.Firstname, "firstname needs to be equal")
	require.Equal(t, person.LastName, partialPerson.LastName, "lastname needs to be equal")
}
