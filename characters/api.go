package characters

import "errors"

func GetAll() []Character {
	return characters
}

func GetById(id string) (*Character, error) {
	for _, c := range characters {
		if c.ID == id {
			return &c, nil
		}
	}

	return nil, errors.New("asd")
}
