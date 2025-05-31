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

// SelectUsers возвращает пользователя по ID
func (s *PostgreStorage) SelectUsers() (model.Users, error) {
	var users model.Users
	query := `SELECT id, login FROM app_user`

	rows, err := s.db.Query(query)
	if err != nil {
		return users, err
	}
	for rows.Next() {
		u := model.User{}
		err = rows.Scan(&u.UserID, &u.UserLogin)
		if err != nil {
			return users, err
		}
		users = append(users, u)
	}

	return users, nil
}

// InsertTask добавляет новую задачу и возвращает её ID
func (s *PostgreStorage) InsertTask(task model.Task) (int, error) {
	var id int
	query := `INSERT INTO task (id, name) VALUES ($1, $2) RETURNING id`

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
