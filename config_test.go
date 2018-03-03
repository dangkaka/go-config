package config

import (
	"reflect"
	"testing"
)

func TestViperConfig_All(t *testing.T) {
	config := NewViperConfig("testdata/config.json")
	stringVal := config.GetString("string")
	if stringVal != "test" {
		t.Errorf("Error get config string")
	}
	boolVal := config.GetBool("bool")
	if boolVal != true {
		t.Errorf("Error get config bool")
	}
	intVal := config.GetInt("int")
	if intVal != 1 {
		t.Errorf("Error get config int")
	}
	int64Val := config.GetInt64("int64")
	if int64Val != 2 {
		t.Errorf("Error get config int64")
	}
	stringSliceVal := config.GetStringSlice("slice.string")
	if !reflect.DeepEqual(stringSliceVal, []string{"test1", "test2"}) {
		t.Errorf("Error get config string slice")
	}
}
