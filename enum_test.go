package enum

import (
	"reflect"
	"testing"
)

// TestStringEnum tests the New function for initializing a struct with string fields.
func TestStringEnum(t *testing.T) {
	HttpStatus := New[struct {
		StatusOK                  string
		StatusNotFound            string
		StatusInternalServerError string
	}]()
	if HttpStatus.StatusOK != "StatusOK" || HttpStatus.StatusNotFound != "StatusNotFound" || HttpStatus.StatusInternalServerError != "StatusInternalServerError" {
		t.Errorf("got %+v, want {StatusOK: StatusOK, StatusNotFound: StatusNotFound, StatusInternalServerError: StatusInternalServerError}", HttpStatus)
	}
}

// TestIntegerEnum tests the New function for initializing a struct with integer fields.
func TestIntegerEnum(t *testing.T) {
	HttpStatus := New[struct {
		StatusOK                  int `enum:"200"`
		StatusNotFound            int `enum:"404"`
		StatusInternalServerError int `enum:"500"`
	}]()
	if HttpStatus.StatusOK != 200 || HttpStatus.StatusNotFound != 404 || HttpStatus.StatusInternalServerError != 500 {
		t.Errorf("got %+v, want {StatusOK: 200, StatusNotFound: 404, StatusInternalServerError: 500}", HttpStatus)
	}
}

// TestNestedEnum tests the functionality of a nested enum-like structure
// created using the New function. It verifies that the Code field contains
// the correct integer values for HTTP status codes (200, 404, 500) and that
// the Type field contains the corresponding string representations
// ("StatusOK", "StatusNotFound", "StatusInternalServerError").
func TestNestedEnum(t *testing.T) {
	HttpStatus := New[struct {
		Code struct {
			StatusOK                  int `enum:"200"`
			StatusNotFound            int `enum:"404"`
			StatusInternalServerError int `enum:"500"`
		}
		Type struct {
			StatusOK                  string
			StatusNotFound            string
			StatusInternalServerError string
		}
	}]()
	if HttpStatus.Code.StatusOK != 200 || HttpStatus.Code.StatusNotFound != 404 || HttpStatus.Code.StatusInternalServerError != 500 {
		t.Errorf("got %+v, want {Code: {StatusOK: 200, StatusNotFound: 404, StatusInternalServerError: 500}}", HttpStatus.Code)
	}
	if HttpStatus.Type.StatusOK != "StatusOK" || HttpStatus.Type.StatusNotFound != "StatusNotFound" || HttpStatus.Type.StatusInternalServerError != "StatusInternalServerError" {
		t.Errorf("got %+v, want {Type: {StatusOK: StatusOK, StatusNotFound: StatusNotFound, StatusInternalServerError: StatusInternalServerError}}", HttpStatus.Type)
	}
}

// TestContains tests the Contains function with the HttpStatus enum.
func TestContains(t *testing.T) {
	HttpStatus := New[struct {
		StatusOK                  string
		StatusNotFound            string
		StatusInternalServerError string
	}]()

	if !Contains(HttpStatus, "StatusOK") {
		t.Errorf("Contains(%v, %q) = false; want true", HttpStatus, "StatusOK")
	}
	if Contains(HttpStatus, "Unknown") {
		t.Errorf("Contains(%v, %q) = true; want false", HttpStatus, "Unknown")
	}
}

// TestKeys tests the Keys function with the HttpStatus enum.
func TestKeys(t *testing.T) {
	HttpStatus := New[struct {
		StatusOK                  string
		StatusNotFound            string
		StatusInternalServerError string
	}]()

	got := Keys(HttpStatus)
	want := []string{"StatusOK", "StatusNotFound", "StatusInternalServerError"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Keys(%v) = %v; want %v", HttpStatus, got, want)
	}

	if got := Keys(123); got != nil {
		t.Errorf("Keys(123) = %v; want nil", got)
	}
}

// TestValues tests the Values function with the HttpStatus enum.
func TestValues(t *testing.T) {
	HttpStatus := New[struct {
		StatusOK                  string
		StatusNotFound            string
		StatusInternalServerError string
	}]()

	got := Values[string](HttpStatus)
	want := []string{"StatusOK", "StatusNotFound", "StatusInternalServerError"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Values[string](%v) = %v; want %v", HttpStatus, got, want)
	}

	if got := Values[int](HttpStatus); len(got) > 0 {
		t.Errorf("Values[int](%v) = %v; want []", HttpStatus, got)
	}
}
