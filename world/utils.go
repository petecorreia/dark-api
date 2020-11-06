package world

import "github.com/petecorreia/dark-api/character"

func getWorldCharacterIds(id string) []string {
	ids := []string{}

	for _, char := range character.GetAll() {
		for _, w := range char.Worlds {
			if w == id {
				ids = append(ids, char.ID)
				break
			}
		}
	}

	return ids
}
