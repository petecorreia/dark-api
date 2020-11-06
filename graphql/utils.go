package graphql

import (
	char "github.com/petecorreia/dark-api/characters"
	"github.com/petecorreia/dark-api/graphql/model"
)

func resolveCharacter(c *char.Character) *model.Character {
	return &model.Character{
		ID:      c.ID,
		Name:    c.Name,
		Worlds:  c.Worlds,
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
