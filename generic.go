package horizon

import "encoding/json"

// Generic

type String struct{ String string }

func (s *String) MarshalJSON() ([]byte, error) {
	if s == nil || s == Delete {
		return json.Marshal(nil)
	}
	return json.Marshal(s.String)
}

func (s *String) UnmarshalJSON(data []byte) error {
	if data == nil || len(data) == 0 {
		return nil
	}
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	if v == "" {
		return nil
	}
	*s = String{v}
	return nil
}

var Delete = &String{""}
