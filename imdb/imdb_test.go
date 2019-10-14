package imdb

import (
	"reflect"
	"testing"
)

func TestGet(t *testing.T) {
	db := NewDB()

	db.Set("test", []byte("valueeee"))

	val, err := db.Get("test")

	expectedVal := []byte("valueeee")

	if err != nil || !reflect.DeepEqual(expectedVal, val.data) {
		t.Errorf("Expected value %s but actual value is %s or error returned", expectedVal, val.data)
	}

	val, err = db.Get("errortest")
	expectedErr := "Key not found in database"
	if err == nil || err.Error() != expectedErr {
		t.Errorf("Expected Get to return %s but found %s", expectedErr, err.Error())
	}
}

func TestSet(t *testing.T) {
	db := NewDB()

	err := db.Set("test", []byte("testValue"))

	if err != nil {
		t.Error("Expected error to be nil but found ", err.Error())
	}

	err = db.Set("", []byte("testValue"))

	if err == nil {
		t.Error("Expected error Key to be inserted should be of atleast length 1 but found nil")
	}

}

func TestUnset(t *testing.T) {
	db := NewDB()
	err := db.Set("test", []byte("testValue"))

	if err != nil {
		t.Error(err.Error())
	}

	db.Unset("test")

	if !reflect.DeepEqual(db.entries["test"], value{}) {
		t.Error("Expected value to be unset")
	}
}
