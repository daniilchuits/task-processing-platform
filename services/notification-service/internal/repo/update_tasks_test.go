package repo

import (
	"database/sql"
	"notifycation-service/internal/domain"
	"testing"
)

func (f fakeRepo) Exec(query string, args ...any) (sql.Result, error) {
	if f.err != nil {
		return nil, f.err
	}
	return fakeResult{rowsAffected: f.res.rowsAffected}, nil
}

func (f fakeRepo) Close() error {
	return f.err
}

func TestUpdateTasks(t *testing.T) {

	mock := &notifyRepo{db: fakeRepo{res: fakeResult{rowsAffected: 1}}}

	t.Run("success", func(t *testing.T) {
		t.Log(2222222222222222)
		res, err := mock.UpdateTasks(domain.Msg{Status: domain.Finished})
		t.Log(3333333333333333)
		if err != nil {
			t.Fatalf("Expected %v err, got %v", nil, err)
		}
		if res != 1 {
			t.Errorf("Expected res == %d, got res == %d", 1, res)
		}
	})

}
