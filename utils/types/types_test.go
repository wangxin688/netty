package types

import (
	"reflect"
	"testing"
)

func TestSetResponse(t *testing.T) {
	code := 200
	data := "test data"
	message := "Success"

	result := SetResponse(code, data, message)

	if result.Code != code {
		t.Errorf("Expected code to be %d, but got %d", code, result.Code)
	}

	if result.Data != data {
		t.Errorf("Expected data to be %v, but got %v", data, result.Data)
	}

	if result.Message != message {
		t.Errorf("Expected message to be %s, but got %s", message, result.Message)
	}
}

func TestSetListResponse(t *testing.T) {
	code := 200
	data := ListT[int]{Count: 3, Results: []int{1, 2, 3}}
	message := "Success"

	result := SetListResponse(code, data, message)

	if result.Code != code {
		t.Errorf("Expected code to be %d, but got %d", code, result.Code)
	}

	if !reflect.DeepEqual(result.Data, data) {
		t.Errorf("Expected data to be %v, but got %v", data, result.Data)
	}

	if result.Message != message {
		t.Errorf("Expected message to be %s, but got %s", message, result.Message)
	}
}