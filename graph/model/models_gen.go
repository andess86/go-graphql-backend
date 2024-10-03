// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
)

type Alarm struct {
	Name          AlarmName     `json:"name"`
	SeverityLevel SeverityLevel `json:"severityLevel"`
}

type AlarmState struct {
	Alarms []*Alarm `json:"alarms"`
}

type Mutation struct {
}

type Propeller struct {
	ID     string  `json:"id"`
	Pitch  float64 `json:"pitch"`
	Rpm    int     `json:"rpm"`
	Degree int     `json:"degree"`
}

type Query struct {
}

type SteeringState struct {
	Mode SteeringMode `json:"mode"`
}

type Subscription struct {
}

type AlarmName string

const (
	AlarmNameOpenPitchPressure   AlarmName = "OPEN_PITCH_PRESSURE"
	AlarmNameWrongWayAlarm       AlarmName = "WRONG_WAY_ALARM"
	AlarmNameOpenLoadLimitInput  AlarmName = "OPEN_LOAD_LIMIT_INPUT"
	AlarmNameClutchFailure       AlarmName = "CLUTCH_FAILURE"
	AlarmNameHubLubricationFault AlarmName = "HUB_LUBRICATION_FAULT"
	AlarmNameControlFault        AlarmName = "CONTROL_FAULT"
	AlarmNameNoHydraulicPressure AlarmName = "NO_HYDRAULIC_PRESSURE"
)

var AllAlarmName = []AlarmName{
	AlarmNameOpenPitchPressure,
	AlarmNameWrongWayAlarm,
	AlarmNameOpenLoadLimitInput,
	AlarmNameClutchFailure,
	AlarmNameHubLubricationFault,
	AlarmNameControlFault,
	AlarmNameNoHydraulicPressure,
}

func (e AlarmName) IsValid() bool {
	switch e {
	case AlarmNameOpenPitchPressure, AlarmNameWrongWayAlarm, AlarmNameOpenLoadLimitInput, AlarmNameClutchFailure, AlarmNameHubLubricationFault, AlarmNameControlFault, AlarmNameNoHydraulicPressure:
		return true
	}
	return false
}

func (e AlarmName) String() string {
	return string(e)
}

func (e *AlarmName) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = AlarmName(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid AlarmName", str)
	}
	return nil
}

func (e AlarmName) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type SeverityLevel string

const (
	SeverityLevelInfo     SeverityLevel = "INFO"
	SeverityLevelWarning  SeverityLevel = "WARNING"
	SeverityLevelCritical SeverityLevel = "CRITICAL"
)

var AllSeverityLevel = []SeverityLevel{
	SeverityLevelInfo,
	SeverityLevelWarning,
	SeverityLevelCritical,
}

func (e SeverityLevel) IsValid() bool {
	switch e {
	case SeverityLevelInfo, SeverityLevelWarning, SeverityLevelCritical:
		return true
	}
	return false
}

func (e SeverityLevel) String() string {
	return string(e)
}

func (e *SeverityLevel) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = SeverityLevel(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid SeverityLevel", str)
	}
	return nil
}

func (e SeverityLevel) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type SteeringMode string

const (
	SteeringModeStart SteeringMode = "START"
	SteeringModeAuto  SteeringMode = "AUTO"
	SteeringModeStop  SteeringMode = "STOP"
	SteeringModeReset SteeringMode = "RESET"
)

var AllSteeringMode = []SteeringMode{
	SteeringModeStart,
	SteeringModeAuto,
	SteeringModeStop,
	SteeringModeReset,
}

func (e SteeringMode) IsValid() bool {
	switch e {
	case SteeringModeStart, SteeringModeAuto, SteeringModeStop, SteeringModeReset:
		return true
	}
	return false
}

func (e SteeringMode) String() string {
	return string(e)
}

func (e *SteeringMode) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = SteeringMode(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid SteeringMode", str)
	}
	return nil
}

func (e SteeringMode) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
