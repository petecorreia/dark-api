package character

// Character represents one of the Dark characters
type Character struct {
	ID            string   `json:"id"`
	Name          string   `json:"name"`
	Worlds        []string `json:"worlds"`
	Aliases       []string `json:"aliases"`
	Alternates    []string `json:"alternates"`
	Parents       []string `json:"parents"`
	Children      []string `json:"children"`
	Relationships []string `json:"relationships"`
}
