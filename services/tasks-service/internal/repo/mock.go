package repo

import (
	"database/sql"
	"errors"
	"log"
	"reflect"
)

type repoMock struct {
	err    error
	values []any
}

func (mock repoMock) Query(query string, args ...any) (*sql.Rows, error) {

	return nil, nil
}

func (mock repoMock) QueryRow(query string, args ...any) row {
	return rowMock{err: mock.err, values: mock.values}
}

type rowMock struct {
	err    error
	values []any
}

func (row rowMock) Scan(dest ...any) error {

	log.Println("exec scan")
	if len(dest) == len(row.values) {
		log.Println("Before cycle")
		for i := range row.values {
			log.Printf("Set value from %v to %v\n", dest[i], row.values[i])
			err := setValue(dest[i], row.values[i])
			if errors.Is(err, NilValue) {
				continue
			}
			if err != nil {
				return err
			}
		}
	}

	return row.err
}

var (
	TypeErr  error = errors.New("Wrong type")
	NilValue error = errors.New("Nil value")
)

func setValue(dest any, value any) error {

	destVal := reflect.ValueOf(dest)
	destType := reflect.TypeOf(dest)
	valueVal := reflect.ValueOf(value)
	valueType := reflect.TypeOf(value)

	if destType != valueType {
		return TypeErr
	}
	if destVal.IsNil() {
		return NilValue
	}

	valueVal.Set(destVal)
	return nil
}
