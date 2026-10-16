package rabbitmq

import (
	"testing"
)

func TestConnInvalidCnnStr(t *testing.T) {

	conn, err := Conn("")
	if err == nil {
		t.Fatal("Expected error, got nil")
	}

	if conn != nil {
		t.Fatal("Expected nil connection")
	}
}
