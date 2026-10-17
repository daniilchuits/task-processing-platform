package database

import (
	"database/sql"
	"errors"
	"testing"
)

type mockDB struct {
	err    error
	called bool
	query  string
}

func (mock *mockDB) Exec(query string, args ...any) (sql.Result, error) {
	if mock.err != nil {
		return nil, mock.err
	}
	mock.called = true
	mock.query = query
	return nil, nil
}

func TestCreateTable(t *testing.T) {

	mock := &mockDB{}
	dbmanager := NewDbManager(mock)

	data := []struct {
		name      string
		expErr    error
		expCalled bool
		expQuery  string
	}{
		{
			"no err at all", nil, true, " ",
		},
		{
			"same no error", nil, true, "",
		},
	}

	for _, tt := range data {
		t.Run(tt.name, func(t *testing.T) {
			err := dbmanager.CreateTable()
			if err != tt.expErr {
				t.Fatalf("Expected err %v, got %v", tt.expErr, err)
			}
			if mock.called != tt.expCalled {
				t.Errorf("Expected to be called %t, got called %t", mock.called, tt.expCalled)
			}
			if mock.query == "" {
				t.Errorf("Expected to execute query")
			}
		})
	}
}

func TestCreateTableWErr(t *testing.T) {

	var notNilError error = errors.New("xz")

	mock := &mockDB{err: notNilError}
	dbmanager := NewDbManager(mock)

	data := []struct {
		name      string
		expErr    error
		expCalled bool
		expQuery  string
	}{
		{
			"not nil err", notNilError, false, "",
		},
		{
			"again not nil err", notNilError, false, "",
		},
	}

	for _, tt := range data {
		t.Run(tt.name, func(t *testing.T) {
			err := dbmanager.CreateTable()
			if err != tt.expErr {
				t.Fatalf("Expected %v err, got %v", tt.expErr, err)
			}
			if mock.called != tt.expCalled {
				t.Errorf("Expected to be called %t, got called %t", mock.called, tt.expCalled)
			}
			if mock.query != "" {
				t.Errorf("Expected not to execute query")
			}
		})
	}
}
