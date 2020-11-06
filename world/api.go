package world

import (
	"errors"
)

func GetAll() []World {
	worldsWithCharacters := []World{}

	for _, w := range worlds {
		worldsWithCharacters = append(worldsWithCharacters, World{
			ID:         w.ID,
			Name:       w.Name,
			Characters: getWorldCharacterIds(w.ID),
		})
	}

	return worldsWithCharacters
}

func GetById(id string) (*World, error) {
	for _, w := range worlds {
		if w.ID == id {
			return &World{
				ID:         w.ID,
				Name:       w.Name,
				Characters: getWorldCharacterIds(w.ID),
			}, nil
		}
	}

	return nil, errors.New("world not found")
}
