package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	char "github.com/petecorreia/dark-api/characters"
	generated1 "github.com/petecorreia/dark-api/graphql/generated"
	model1 "github.com/petecorreia/dark-api/graphql/model"
)

func (r *characterResolver) Alternates(ctx context.Context, obj *model1.Character) ([]*model1.Character, error) {
	character, err := char.GetById(obj.ID)
	if err != nil {
		return nil, err
	}
	return resolveCharacterIDs(character.Alternates), nil
}

func (r *characterResolver) Parents(ctx context.Context, obj *model1.Character) ([]*model1.Character, error) {
	character, err := char.GetById(obj.ID)
	if err != nil {
		return nil, err
	}
	return resolveCharacterIDs(character.Parents), nil
}

func (r *characterResolver) Children(ctx context.Context, obj *model1.Character) ([]*model1.Character, error) {
	character, err := char.GetById(obj.ID)
	if err != nil {
		return nil, err
	}
	return resolveCharacterIDs(character.Children), nil
}

func (r *characterResolver) Relationships(ctx context.Context, obj *model1.Character) ([]*model1.Character, error) {
	character, err := char.GetById(obj.ID)
	if err != nil {
		return nil, err
	}
	return resolveCharacterIDs(character.Relationships), nil
}

func (r *queryResolver) Characters(ctx context.Context) ([]*model1.Character, error) {
	var characters []*model1.Character

	for _, c := range char.GetAll() {
		characters = append(characters, resolveCharacter(&c))
	}

	return characters, nil
}

func (r *queryResolver) Character(ctx context.Context, id string) (*model1.Character, error) {
	character, err := char.GetById(id)

	if err != nil {
		return nil, err
	}

	return resolveCharacter(character), nil
}

// Character returns generated1.CharacterResolver implementation.
func (r *Resolver) Character() generated1.CharacterResolver { return &characterResolver{r} }

// Query returns generated1.QueryResolver implementation.
func (r *Resolver) Query() generated1.QueryResolver { return &queryResolver{r} }

type characterResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
