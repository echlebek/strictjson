# strictjson
Validate JSON objects against Go structs.

[![Build Status](https://api.travis-ci.org/echlebek/strictjson.svg)](https://api.travis-ci.org/echlebek/strictjson)

`strictjson.Check` allows you to check if all of the fields in a struct are satisfied
by the incoming JSON message.

Example:

    type Foo struct {
        A int    `json:"a"`
        B string `json:"b"`
    }

    func (f *Foo) UnmarshalJSON(b []byte) error {
        if err := strictjson.Check(b, f); err != nil {
            return err
        }
        return json.Unmarshal(b, f)
    }

In the above example, if a user tries to supply a json message that doesn't contain
keys `a` and `b`, `json.Unmarshal` will return an error.
