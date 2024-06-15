package serialize

import "encoding/json"

type JsonSerializer[T any] struct {
	value T
}

func CreateJsonSerializer[T any](defaultValue T) *JsonSerializer[T] {
	return &JsonSerializer[T]{value: defaultValue}
}

func (s *JsonSerializer[T]) Serialize(data T) ([]byte, error) {
	return json.Marshal(data)
}

func (s *JsonSerializer[T]) Deserialize(data []byte) (*T, error) {
	value := s.value
	err := json.Unmarshal(data, value)
	if err != nil {
		return nil, err
	}
	return &value, nil
}
