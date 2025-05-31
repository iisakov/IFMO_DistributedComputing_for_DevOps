package storage

import (
	"app/pkg/model"
)

var DB Storage

type Storage interface {
	SelectUsers() (model.Users, error)

	InsertTask(model.Task) (int, error)
	DeleteTask(int) (int, error)
}
