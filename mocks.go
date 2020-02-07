package graphql_go_test

import (
	"strconv"
)

type GenerateID struct {
	count int
}

func (generator *GenerateID) next() string {
	generator.count++
	return strconv.Itoa(generator.count)
}

var userID = GenerateID{1}

var users = []*User{
	{
		ID:   userID.next(),
		Name: "Maicon",
	},
	{
		ID:   userID.next(),
		Name: "Carraro",
	},
}

func getMockUsers() []*User {
	return users
}

func findMockUser(ID string) *User {
	users := getMockUsers()
	for i := range users {
		if users[i].ID == ID {
			return users[i]
		}
	}
	return nil
}
