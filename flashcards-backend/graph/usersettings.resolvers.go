package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"flashcards-backend/ent"
	"flashcards-backend/graph/model"
	"flashcards-backend/modelconv"
	"fmt"
)

func (r *mutationResolver) SetSettings(ctx context.Context, input model.SetSettings) (*model.UserSettings, error) {
	settings, err := r.DB.UserSettings.Query().First(ctx)
	if err != nil && !ent.IsNotFound(err) {
		return nil, fmt.Errorf("finding old settings: %v", err)
	}
	if settings == nil {
		// Create settings
		settings, err = r.DB.UserSettings.Create().SetNewCardsPerDay(input.NewCardsPerDay).Save(ctx)
		if err != nil {
			return nil, fmt.Errorf("creating new settings: %v", err)
		}
	} else {
		// Update settings
		settings, err = r.DB.UserSettings.UpdateOne(settings).SetNewCardsPerDay(input.NewCardsPerDay).Save(ctx)
		if err != nil {
			return nil, fmt.Errorf("updating settings for id %d: %v", settings.ID, err)
		}
	}
	return modelconv.UserSettings(settings), nil
}

func (r *queryResolver) UserSettings(ctx context.Context) (*model.UserSettings, error) {
	settings, err := r.DB.UserSettings.Query().First(ctx)
	if err != nil && !ent.IsNotFound(err) {
		return nil, fmt.Errorf("finding old settings: %v", err)
	}
	if settings == nil {
		// Create settings
		settings, err = r.DB.UserSettings.Create().Save(ctx)
		if err != nil {
			return nil, fmt.Errorf("creating new settings: %v", err)
		}
	}
	return modelconv.UserSettings(settings), nil
}
