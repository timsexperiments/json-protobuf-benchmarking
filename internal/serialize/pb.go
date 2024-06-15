package serialize

import (
	"google.golang.org/protobuf/proto"
)

type PbSerializer[M proto.Message] struct {
	message M
}

func CreatePbSerializer[M proto.Message](defaultMessage M) *PbSerializer[M] {
	return &PbSerializer[M]{message: defaultMessage}
}

func (s *PbSerializer[M]) Serialize(data M) ([]byte, error) {
	return proto.Marshal(data)
}

func (s *PbSerializer[M]) Deserialize(data []byte) (*M, error) {
	message := s.message
	err := proto.Unmarshal(data, message)
	if err != nil {
		return nil, err
	}
	return &message, nil
}
