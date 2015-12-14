package strictjson

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

func checkErr(e error) error {
	return fmt.Errorf("strictjson: Check(%s)", e)
}

// Check checks that all of obj's struct fields are provided by the JSON
// object encoded in b.
//
// If obj is not a struct, an error will be returned.
//
// If a struct field has an empty `json` tag, or an omitempty tag, the absence
// of a corresponding key in b will not be considered an error.
//
// Example:
//
//   type Foo struct {
//       A int    `json:"a"`
//       B string `json:"b"`
//   }
//
//   func (f *Foo) UnmarshalJSON(b []byte) error {
//       if err := strictjson.Check(b, f); err != nil {
//           return err
//       }
//       return json.Unmarshal(b, f)
//   }
//
func Check(b []byte, obj interface{}) error {
	v := reflect.ValueOf(obj)
	for v.Kind() == reflect.Ptr {
		v = reflect.Indirect(v)
	}
	if v.Kind() != reflect.Struct {
		return checkErr(fmt.Errorf("non-struct %s", v.Kind()))
	}
	m := make(map[string]*json.RawMessage)
	if err := json.Unmarshal(b, &m); err != nil {
		return checkErr(err)
	}
	var errors []string
	vt := v.Type()
	for i := 0; i < v.NumField(); i++ {
		field := vt.Field(i)
		if tag := field.Tag.Get("json"); tag == "" || strings.Contains(tag, ",omitempty") {
			continue
		} else if _, ok := m[tag]; !ok {
			errors = append(errors, field.Name)
		}
	}
	if errors != nil {
		return fmt.Errorf("missing fields: %v", errors)
	}
	return nil
}
