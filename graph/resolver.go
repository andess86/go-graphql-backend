package graph

import "github.com/andess86/gqlgen-todos/generator"

//go:generate go run github.com/99designs/gqlgen generate

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.
// Resolver struct will hold references to generators or other dependencies

type Resolver struct {
	PropellerDataChannel <-chan generator.Propeller
	AlarmDataChannel     <-chan generator.Alarm

}

// NewResolver initializes a new resolver with the propeller and alarm generators
func NewResolver() *Resolver {
    propellerGen := generator.NewPropellerGenerator()
    alarmGen := generator.NewAlarmGenerator()

    return &Resolver{
        PropellerDataChannel: propellerGen.StartGenerating(),
        AlarmDataChannel:     alarmGen.StartGenerating(),
    }
}