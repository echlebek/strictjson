package strictjson

import (
	"strings"
	"testing"
)

const input1 = `{"a": 5, "b": "six", "c": [1,2,3]}`
const input2 = `{"a": 5, "B": "six", "c": [1,2,3], "d": 0.5}`
const input3 = `{"a": 5, "b": "six", "c": [1,2,3], "d": 0.5, "e": "foo"}`
const invalid = `:5, "b": 10}`

func TestCheckInvalidJSON(t *testing.T) {
	type Foo struct {
		A int     `json:"a"`
		B string  `json:"b"`
		C []int   `json:"c"`
		D float64 `json:"d"`
	}
	var f Foo
	if err := Check([]byte(invalid), &f); err == nil {
		t.Error("expected error")
	}
}

func TestCheckExtraKey(t *testing.T) {
	type Foo struct {
		A int     `json:"a"`
		B string  `json:"b"`
		C []int   `json:"c"`
		D float64 `json:"d"`
	}
	var f Foo
	if err := Check([]byte(input3), &f); err != nil {
		t.Error(err)
	}
}

func TestCheckMissingKey(t *testing.T) {
	type Foo struct {
		A int     `json:"a"`
		B string  `json:"b"`
		C []int   `json:"c"`
		D float64 `json:"d"`
	}

	var f Foo

	if err := Check([]byte(input1), f); err == nil {
		t.Errorf("expected error")
	} else if !strings.Contains(err.Error(), "[D]") {
		t.Errorf("bad error: %s", err)
	}
	if err := Check([]byte(input2), f); err == nil {
		t.Errorf("expected error")
	} else if !strings.Contains(err.Error(), "[B]") {
		t.Errorf("bad error: %s", err)
	}

	// Repeat with pointer to f
	if err := Check([]byte(input1), &f); err == nil {
		t.Errorf("expected error")
	} else if !strings.Contains(err.Error(), "[D]") {
		t.Errorf("bad error: %s", err)
	}
	if err := Check([]byte(input2), &f); err == nil {
		t.Errorf("expected error")
	} else if !strings.Contains(err.Error(), "[B]") {
		t.Errorf("bad error: %s", err)
	}
}

func TestCheckOmitEmpty(t *testing.T) {

	type Foo struct {
		A int     `json:"a"`
		B string  `json:"b"`
		C []int   `json:"c"`
		D float64 `json:"d,omitempty"`
	}
	var f Foo

	if err := Check([]byte(input1), &f); err != nil {
		t.Error(err)
	}
	if err := Check([]byte(input2), &f); err == nil {
		t.Error("expected error")
	} else if !strings.Contains(err.Error(), "[B]") {
		t.Errorf("bad error: %s", err)
	}
	if err := Check([]byte(input3), &f); err != nil {
		t.Error(err)
	}
}

func TestCheckInvalidTarget(t *testing.T) {
	if err := Check(nil, 1); err == nil {
		t.Error("expected error")
	} else if !strings.Contains(err.Error(), "non-struct") {
		t.Errorf("bad error: %s", err)
	}
	m := make(map[string]interface{})
	if err := Check(nil, &m); err == nil {
		t.Error("expected error")
	} else if !strings.Contains(err.Error(), "non-struct") {
		t.Errorf("bad error: %s", err)
	}

}
