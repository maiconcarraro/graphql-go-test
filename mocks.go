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

func findMockUser(ID string) (*models.User, error) {
	users := getMockUsers()
	for i := range users {
		if users[i].ID == ID {
			return users[i], nil
		}
	}
	return nil, errors.New("Invalid UserID")
}

var taskID = generatorID{0}

var tasks = []*models.Task{
	{
		ID:     taskID.next(),
		Text:   "Create project",
		Done:   true,
		UserID: "1",
	},
	{
		ID:     taskID.next(),
		Text:   "Run project",
		Done:   false,
		UserID: "2",
	},
}

func getMockTasks() []*models.Task {
	return tasks
}

func createMockTask(newTask NewTask) (*models.Task, error) {
	_, err := findMockUser(newTask.UserID)

	if err != nil {
		return nil, err
	}

	task := &models.Task{
		ID:     taskID.next(),
		Text:   newTask.Text,
		Done:   false,
		UserID: newTask.UserID,
	}

	tasks = append(tasks, task)
	return task, nil
}
