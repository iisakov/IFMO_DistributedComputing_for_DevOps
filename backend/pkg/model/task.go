package model

type Task struct {
	TaskID   int    `json:"taskId,omitempty" db:"task_id"`
	TaskName string `json:"taskName" db:"task_name"`
}

type Tasks []Task
