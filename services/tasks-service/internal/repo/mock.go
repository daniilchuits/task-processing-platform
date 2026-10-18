package repo

import (
	"database/sql"
	"errors"
	"fmt"
	"reflect"
	"testing"
)

type repoMock struct {
	err    error
	values []any
	t      *testing.T
}

func (mock repoMock) Query(query string, args ...any) (*sql.Rows, error) {

	return nil, nil
}

func (mock repoMock) QueryRow(query string, args ...any) row {
	return rowMock{err: mock.err, values: mock.values, t: mock.t}
}

type rowMock struct {
	err    error
	values []any
	t      *testing.T
}

func (row rowMock) Scan(dest ...any) error {

	t := row.t
	t.Log("exec scan")
	// if len(dest) == len(row.values) {
	for i := range row.values {
		// dest is []any, which is []ptr
		// i want to get value for dest[i]
		f := forTesting{t: t}
		err := f.setValue(dest[i], row.values[i])
		if err != nil {
			return fmt.Errorf("%w in %d index", err, i)
		}
	}

	return row.err
}

type forTesting struct {
	t *testing.T
}

var (
	TypeErr   error = errors.New("Type error")
	NilValue  error = errors.New("Nil value")
	CannotSet error = errors.New("Cannot set")
)

func (f forTesting) setValue(dest any, value any) error {

	// dest is ptr type
	// value is non-ptr

	ptr := reflect.ValueOf(dest)

	if !ptr.IsValid() {
		return NilValue
	}
	if ptr.Kind() != reflect.Pointer {
		return TypeErr
	}
	if ptr.IsNil() {
		return NilValue
	}

	target := ptr.Elem()
	if !target.CanSet() {
		return CannotSet
	}

	if value == nil {
		if target.Kind() != reflect.Pointer {
			return TypeErr
		}
		target.SetZero()
		return nil
	}

	source := reflect.ValueOf(value)
	if !source.Type().AssignableTo(target.Type()) {
		return TypeErr
	}

	target.Set(source)
	return nil
}
