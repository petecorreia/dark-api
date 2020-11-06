// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Character struct {
	ID            string       `json:"id"`
	Name          string       `json:"name"`
	Worlds        []string     `json:"worlds"`
	Aliases       []string     `json:"aliases"`
	Alternates    []*Character `json:"alternates"`
	Parents       []*Character `json:"parents"`
	Children      []*Character `json:"children"`
	Relationships []*Character `json:"relationships"`
}