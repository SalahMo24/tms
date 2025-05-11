package assert

import (
	"fmt"
	"reflect"
	"runtime"
)

// True asserts that the condition is true
func True(condition bool, message string) {
	if !condition {
		panicWithMessage(message)
	}
}

// Equal asserts that two values are equal
func Equal(expected, actual interface{}, message string) {
	if !reflect.DeepEqual(expected, actual) {
		panicWithMessage(fmt.Sprintf("%s: expected %v, got %v", message, expected, actual))
	}
}

// Nil asserts that the value is nil
func Nil(value any, message string) {
	if value != nil {
		panicWithMessage(fmt.Sprintf("%s: expected nil, got %v", message, value))
	}
}

// NotNil asserts that the value is not nil
func NotNil(value any, message string) {
	if value == nil {
		panicWithMessage(fmt.Sprintf("%s: expected not nil", message))
	}
}

// Type asserts that the value is of the expected type
func Type(expectedType any, value any, message string) {
	expected := reflect.TypeOf(expectedType)
	actual := reflect.TypeOf(value)

	if actual != expected {
		panicWithMessage(fmt.Sprintf("%s: expected type %v, got %v", message, expected, actual))
	}
}

func panicWithMessage(message string) {
	_, file, line, _ := runtime.Caller(2)
	panic(fmt.Sprintf("Assertion failed at %s:%d: %s", file, line, message))
}
