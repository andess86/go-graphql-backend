package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.54

import (
	"context"
	"log"

	"github.com/andess86/gqlgen-todos/graph/model"
)

// SetSteeringMode is the resolver for the setSteeringMode field.
func (r *mutationResolver) SetSteeringMode(ctx context.Context, mode model.SteeringMode) (*model.SteeringState, error) {
	// Implementing this is beyond the current scope
	return &model.SteeringState{
		Mode: mode,
	}, nil
}

// GetPropellerData is the resolver for the getPropellerData field.
func (r *queryResolver) GetPropellerData(ctx context.Context) (*model.Propeller, error) {
	// Wait for the next value from the propeller generator
	propeller := <-r.PropellerDataChannel

	// Map the generated data to the GraphQL model
	return &model.Propeller{
		ID:     propeller.ID,
		Pitch:  propeller.Pitch,
		Rpm:    propeller.RPM,
		Degree: propeller.Degree,
	}, nil
}

// GetSteeringMode is the resolver for the getSteeringMode field.
func (r *queryResolver) GetSteeringMode(ctx context.Context) (*model.SteeringState, error) {
	// Just providing a mock response for now
	return &model.SteeringState{
		Mode: model.SteeringModeStart,
	}, nil
}

// PropellerDataUpdated is the resolver for the propellerDataUpdated field.
func (r *subscriptionResolver) PropellerDataUpdated(ctx context.Context) (<-chan *model.Propeller, error) {
	log.Printf("Entered start of propellerDataUpdated resolver")
	updateChannel := make(chan *model.Propeller)

	// Use a goroutine to listen for updates
	go func() {
		defer close(updateChannel) // Ensure the channel is closed when done
		log.Printf("Inside goroutine func")

		for {
			select {
			case data, ok := <-r.Resolver.PropellerDataChannel:
				if !ok {
					log.Printf("Propeller data channel closed, exiting goroutine")
					return
				}

				log.Printf("Received data from PropellerDataChannel: %+v", data)

				// Send updates to the subscription channel
				updateChannel <- &model.Propeller{
					ID:     data.ID,
					Pitch:  data.Pitch,
					Rpm:    data.RPM,
					Degree: data.Degree,
				}

				log.Printf("Sent propeller data to updateChannel: %+v", data)

			case <-ctx.Done():
				log.Printf("Context cancelled, exiting goroutine")
				return
			}
		}

	}()

	return updateChannel, nil
}

// SteeringModeUpdated is the resolver for the steeringModeUpdated field.
func (r *subscriptionResolver) SteeringModeUpdated(ctx context.Context) (<-chan *model.SteeringState, error) {
	// Placeholder for steering mode updates
	updateChannel := make(chan *model.SteeringState)

	// Just simulating steering mode changes, as this is beyond current scope
	go func() {
		defer close(updateChannel)
		for {
			updateChannel <- &model.SteeringState{
				Mode: model.SteeringModeAuto,
			}
		}
	}()

	return updateChannel, nil
}

// AlarmStateUpdated is the resolver for the alarmStateUpdated field.
func (r *subscriptionResolver) AlarmStateUpdated(ctx context.Context) (<-chan *model.AlarmState, error) {
	log.Printf("Entered start of alarmStateUpdated resolver")
	updateChannel := make(chan *model.AlarmState)

	// Use a goroutine to listen for updates
	go func() {
		log.Printf("Inside goroutine func for alarm state")

		for alarmArray := range r.Resolver.AlarmDataChannel {
			// Convert generated alarms to GraphQL model alarms
			modelAlarms := make([]*model.Alarm, len(alarmArray))
			for i, alarm := range alarmArray {
				modelAlarms[i] = &model.Alarm{
					Name:          model.AlarmName(alarm.Name),
					SeverityLevel: model.SeverityLevel(alarm.SeverityLevel),
				}
			}

			// Send the array of alarms to the subscription channel
			updateChannel <- &model.AlarmState{
				Alarms: modelAlarms,
			}
		}
		close(updateChannel) // Ensure the channel is closed when done
	}()

	return updateChannel, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

// Subscription returns SubscriptionResolver implementation.
func (r *Resolver) Subscription() SubscriptionResolver { return &subscriptionResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }
