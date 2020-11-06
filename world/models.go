package world

// World represents one of the Dark worlds
type World struct {
	ID         string   `json:"id"`
	Name       string   `json:"name"`
	Characters []string `json:"characters"`
}
