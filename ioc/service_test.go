package ioc

import (
	"reflect"
	"testing"
)

func TestService_say(t *testing.T) {
	id := GetContext().getBeanById("test")
	of := reflect.ValueOf(id)
	of.Kind()
}
