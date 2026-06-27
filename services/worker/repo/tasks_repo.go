package repo

import "database/sql"

type tasksRepo struct {
	db *sql.DB
}

func NewTasksRepo(db *sql.DB) tasksRepo {
	return tasksRepo{db: db}
}
