package repo

import (
	"database/sql"
	"notifycation-service/internal/domain"
	"time"
)

type notifyRepo struct {
	db db
}

type fakeRepo struct {
	err error
	res fakeResult
}

type notifier interface {
	UpdateTasks(data domain.Msg) (int, error)
}

type db interface {
	Exec(query string, args ...any) (sql.Result, error)
	Close() error
}

func OpenDataBase(cnnStr string) (*notifyRepo, error) {

	db, err := sql.Open("postgres", cnnStr)
	if err != nil {
		return nil, err
	}

	for range 5 {
		if err = db.Ping(); err == nil {
			return &notifyRepo{db: db}, nil
		}
		time.Sleep(time.Second)
	}
	return nil, err
}

func (repo *notifyRepo) Close() error {
	return repo.db.Close()
}

func newNotifyRepo(db *sql.DB) notifier {
	return &notifyRepo{db: db}
}

// without sql.Result can't test update, nil pointer dereference

type Result interface {
	LastInsertId() (int64, error)
	RowsAffected() (int64, error)
}

type fakeResult struct {
	rowsAffected int64
	err          error
}

func (f fakeResult) LastInsertId() (int64, error) {
	return 0, nil
}

func (f fakeResult) RowsAffected() (int64, error) {
	return f.rowsAffected, f.err
}
