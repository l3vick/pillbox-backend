package handler

import (
	"database/sql"
	"encoding/json"
)

type JsonNullInt64 struct {
	sql.NullInt64
}

type JsonNullFloat64 struct {
	sql.NullFloat64
}

type JsonNullString struct {
	sql.NullString
}

type JsonNullBool struct {
	sql.NullBool
}

func (v *JsonNullBool) MarshalJSON() ([]byte, error) {
	if v.Valid {
		return json.Marshal(v.Bool)
	} else {
		return json.Marshal(nil)
	}
}

func (v *JsonNullBool) UnmarshalJSON(data []byte) error {
	var b *bool
	if err := json.Unmarshal(data, &b); err != nil {
		return err
	}

	if b != nil {
		v.Valid = true
		v.Bool = *b
	} else {
		v.Valid = false
	}

	return nil
}

func (v *JsonNullString) MarshalJSON() ([]byte, error) {
	if v.Valid {
		return json.Marshal(v.String)
	} else {
		return json.Marshal(nil)
	}
}

func (v *JsonNullString) UnmarshalJSON(data []byte) error {
	var str *string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}

	if str != nil {
		v.Valid = true
		v.String = *str
	} else {
		v.Valid = false
	}

	return nil
}

func (v JsonNullFloat64) MarshalJSON() ([]byte, error) {
	if v.Valid {
		return json.Marshal(v.Float64)
	} else {
		return json.Marshal(0)
	}
}

func (v *JsonNullFloat64) UnmarshalJSON(data []byte) error {
	// Unmarshalling into a pointer will let us detect null

	var x *float64
	err := json.Unmarshal(data, &x)
	if err != nil {
		return err
	}
	if x != nil {
		v.Valid = true
		v.Float64 = *x
	} else {
		v.Valid = false
	}
	return nil
}

func (v JsonNullInt64) MarshalJSON() ([]byte, error) {
	if v.Valid {
		return json.Marshal(v.Int64)
	} else {
		return json.Marshal(0)
	}
}

func (v JsonNullInt64) UnmarshalJSON(data []byte) error {
	// Unmarshalling into a pointer will let us detect null
	var x *int64
	if err := json.Unmarshal(data, &x); err != nil {
		return err
	}
	if x != nil {
		v.Valid = true
		v.Int64 = *x
	} else {
		v.Valid = false
	}
	return nil
}
