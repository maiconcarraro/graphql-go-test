package graphql_go_test

import (
	"errors"
	"strconv"

	"github.com/maiconcarraro/graphql-go-test/models"
)

type generatorID struct {
	count int
}

func (generator *generatorID) next() string {
	generator.count++
	return strconv.Itoa(generator.count)
}

var userID = generatorID{0}

var users = []*models.User{
	{
		ID:   userID.next(),
		Name: "Maicon",
	},
	{
		ID:   userID.next(),
		Name: "Carraro",
	},
}

func getMockUsers() []*models.User {
	return users
}

func findMockUser(ID string) *models.User {
	users := getMockUsers()
	for i := range users {
		if users[i].ID == ID {
			return users[i]
		}
	}
	return nil
}

var taskID = generatorID{0}

var tasks = []*models.Task{
	{
		ID:   taskID.next(),
		Text: "Create project",
		Done: true,
		User: findMockUser("1"),
	},
	{
		ID:   taskID.next(),
		Text: "Run project",
		Done: false,
		User: findMockUser("2"),
	},
}

func getMockTasks() []*models.Task {
	return tasks
}

func createMockTask(newTask NewTask) (*models.Task, error) {
	user := findMockUser(newTask.UserID)

	if user == nil {
		return nil, errors.New("Invalid UserID")
	}

	task := &models.Task{
		ID:   taskID.next(),
		Text: newTask.Text,
		Done: false,
		User: findMockUser(newTask.UserID),
	}

	tasks = append(tasks, task)
	return task, nil
}
