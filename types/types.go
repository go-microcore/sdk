package types

import "encoding/json"

type Nullable[T any] struct {
	Set   bool
	Value *T
}

func (n *Nullable[T]) UnmarshalJSON(b []byte) error {
	n.Set = true
	if string(b) == "null" {
		n.Value = nil
		return nil
	}
	var v T
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	n.Value = &v
	return nil
}

func (n Nullable[T]) MarshalJSON() ([]byte, error) {
	if !n.Set {
		return nil, nil
	}
	if n.Value == nil {
		return []byte("null"), nil
	}
	return json.Marshal(n.Value)
}
