package postgreStorage

import (
	"app/pkg/model"
	"database/sql"
)

// PostgreStorage реализует интерфейс Storage для PostgreSQL
type PostgreStorage struct {
	db *sql.DB
}

// NewPostgreStorage создает новый экземпляр PostgreStorage
func NewPostgreStorage(db *sql.DB) *PostgreStorage {
	return &PostgreStorage{db: db}
}

// SelectUser возвращает пользователя по ID
func (s *PostgreStorage) SelectUser(userID int) (model.User, error) {
	var user model.User
	query := `SELECT user_id, user_login FROM user WHERE id = $1`

	err := s.db.QueryRow(query, userID).Scan(&user.UserID, &user.UserLogin)
	if err != nil {
		return user, err
	}
	return user, nil
}

// InsertTask добавляет новую задачу и возвращает её ID
func (s *PostgreStorage) InsertTask(task model.Task) (int, error) {
	var id int
	query := `INSERT INTO task (task_id, task_name) VALUES ($1, $2) RETURNING id`

	err := s.db.QueryRow(query, task.TaskID, task.TaskName).Scan(&id)
	if err != nil {
		return -1, err
	}

	return id, nil
}

// DeleteTask удаляет задачу по ID и возвращает количество удаленных строк
func (s *PostgreStorage) DeleteTask(taskID int) (int, error) {
	query := `DELETE FROM task WHERE id = $1`
	res, err := s.db.Exec(query, taskID)
	if err != nil {
		return -1, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return -1, err
	}

	return int(rowsAffected), nil
}
