package repo

import (
	"database/sql"
	"errors"

	"task-manager/internal/models"
)

type TaskRepository struct {
	db *sql.DB
}

func NewTaskRepository(db *sql.DB) *TaskRepository {
	return &TaskRepository{db: db}
}

func (r *TaskRepository) CreateTask(task *models.Task) error {
	query := `
        INSERT INTO tasks (user_id, title, description, status, created_at, updated_at)
        VALUES ($1, $2, $3, $4, NOW(), NOW())
        RETURNING id, created_at, updated_at
    `
	return r.db.QueryRow(query, task.UserID, task.Title, task.Description, task.Status).
		Scan(&task.ID, &task.CreatedAt, &task.UpdatedAt)
}

func (r *TaskRepository) GetTaskByID(id int) (*models.Task, error) {
	task := &models.Task{}
	query := `
        SELECT id, user_id, title, description, status, created_at, updated_at
        FROM tasks
        WHERE id = $1
    `
	err := r.db.QueryRow(query, id).Scan(
		&task.ID,
		&task.UserID,
		&task.Title,
		&task.Description,
		&task.Status,
		&task.CreatedAt,
		&task.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return task, nil
}

func (r *TaskRepository) GetAllTasks() ([]*models.Task, error) {
	query := `
        SELECT id, user_id, title, description, status, created_at, updated_at
        FROM tasks
        ORDER BY created_at DESC
    `
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tasks := []*models.Task{}
	for rows.Next() {
		t := &models.Task{}
		if err := rows.Scan(&t.ID, &t.UserID, &t.Title, &t.Description, &t.Status, &t.CreatedAt, &t.UpdatedAt); err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}
	return tasks, nil
}

func (r *TaskRepository) UpdateTask(task *models.Task) error {
	query := `
        UPDATE tasks
        SET title = $1, description = $2, status = $3, updated_at = NOW()
        WHERE id = $4
    `
	_, err := r.db.Exec(query, task.Title, task.Description, task.Status, task.ID)
	return err
}

func (r *TaskRepository) DeleteTask(id int) error {
	query := `DELETE FROM tasks WHERE id = $1`
	_, err := r.db.Exec(query, id)
	return err
}
