# strictjson
Validate JSON objects against Go structs.

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
