# strictjson
Validate JSON objects against Go structs.

[![Build Status](https://api.travis-ci.org/echlebek/strictjson.svg)](https://travis-ci.org/echlebek/strictjson)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](https://godoc.org/github.com/echlebek/strictjson)

`strictjson.Check` allows you to check if all of the fields in a struct are satisfied
by the incoming JSON message.

Example:

    type Foo struct {
    	A int    `json:"a"`
    	B string `json:"b"`
    }

    func (f *Foo) UnmarshalJSON(b []byte) error {
    	type __ Foo
    	var g __
    	if err := Check(b, f); err != nil {
    		return err
    	}
    	if err := json.Unmarshal(b, &g); err != nil {
    		return err
    	}
    	*f = Foo(g)
    	return nil
    }

In the above example, if a user tries to supply a json message that doesn't contain
keys `a` and `b`, `strictjson.Check` will return an error.
