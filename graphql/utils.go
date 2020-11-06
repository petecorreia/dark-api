package graphql

import (
	char "github.com/petecorreia/dark-api/character"
	"github.com/petecorreia/dark-api/graphql/model"
	"github.com/petecorreia/dark-api/world"
)

func resolveCharacter(c *char.Character) *model.Character {
	return &model.Character{
		ID:      c.ID,
		Name:    c.Name,
		Aliases: c.Aliases,
	}
}

func resolveCharacterIDs(ids []string) []*model.Character {
	var characters []*model.Character

	for _, id := range ids {
		character, err := char.GetById(id)
		if err == nil && character != nil {
			characters = append(characters, resolveCharacter(character))
		}
	}

	return characters
}

func resolveWorld(w *world.World) *model.World {
	return &model.World{
		ID:   w.ID,
		Name: w.Name,
	}
}

func resolveWorldIDs(ids []string) []*model.World {
	var worlds []*model.World

	for _, id := range ids {
		w, err := world.GetById(id)
		if err == nil && w != nil {
			worlds = append(worlds, resolveWorld(w))
		}
	}

	return worlds
}
