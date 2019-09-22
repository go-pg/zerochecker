package zerochecker_test

import (
	"reflect"
	"testing"

	"github.com/go-pg/zerochecker"
)

func TestChecker(t *testing.T) {
	isZero := zerochecker.Checker(reflect.TypeOf(int(0)))
	if !isZero(reflect.ValueOf(0)) {
		t.Fatal("0 is not zero")
	}
	if isZero(reflect.ValueOf(1)) {
		t.Fatal("1 is zero")
	}
}
