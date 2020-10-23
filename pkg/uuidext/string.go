package uuidext

import "github.com/google/uuid"

func ToString(id uuid.UUID) string {
	if id == uuid.Nil {
		return ""
	}
	return id.String()
}

func FromString(str string) (uuid.UUID, error) {
	if str == "" {
		return uuid.Nil, nil
	}
	return uuid.Parse(str)
}

func FromStringMust(str string) uuid.UUID {
	id, err := FromString(str)
	if err != nil {
		panic(err)
	}
	return id
}
