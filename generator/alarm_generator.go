package generator

import (
	"log"
	"math/rand"
	"time"
)

// Alarm represents an individual alarm structure
type Alarm struct {
    Name          string
    SeverityLevel string
}

// AlarmGenerator handles the generation of alarm data
type AlarmGenerator struct{}

// Possible values for AlarmName and SeverityLevel
var alarmNames = []string{
    "OPEN_PITCH_PRESSURE",
    "WRONG_WAY_ALARM",
    "OPEN_LOAD_LIMIT_INPUT",
    "CLUTCH_FAILURE",
    "HUB_LUBRICATION_FAULT",
    "CONTROL_FAULT",
    "NO_HYDRAULIC_PRESSURE",
}

var severityLevels = []string{"INFO", "WARNING", "CRITICAL"}

// NewAlarmGenerator creates a new instance of AlarmGenerator
func NewAlarmGenerator() *AlarmGenerator {
    return &AlarmGenerator{}
}

// StartGenerating returns a channel that emits alarm data periodically
func (ag *AlarmGenerator) StartGenerating() <-chan Alarm {
    alarmChannel := make(chan Alarm)

    go func() {
        log.Println("Starting to send mock alarm data...")

        defer close(alarmChannel)
        for {
            alarmData := Alarm{
                Name:          alarmNames[rand.Intn(len(alarmNames))],           // Random alarm name
                SeverityLevel: severityLevels[rand.Intn(len(severityLevels))],   // Random severity level
            }
            alarmChannel <- alarmData
            log.Println("Sent new alarm data:", alarmData)

            time.Sleep(1 * time.Second) // Simulate data update every 1 second
        }
    }()

    return alarmChannel
}
