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

// StartGenerating returns a channel that emits an array of alarm data periodically
func (ag *AlarmGenerator) StartGenerating() <-chan []Alarm {
    alarmChannel := make(chan []Alarm)

    go func() {
        log.Println("Starting to send mock alarm data...")

        defer close(alarmChannel)
        for {
            // Random number of alarms between 1 and 5
            numberOfAlarms := rand.Intn(5) + 1
            alarms := make([]Alarm, numberOfAlarms)

            for i := 0; i < numberOfAlarms; i++ {
                alarms[i] = Alarm{
                    Name:          alarmNames[rand.Intn(len(alarmNames))],          // Random alarm name
                    SeverityLevel: severityLevels[rand.Intn(len(severityLevels))],  // Random severity level
                }
            }

            // Send the generated array of alarms
            alarmChannel <- alarms
            log.Printf("Sent new alarm data: %v", alarms)

            // Random sleep interval between 1 and 5 seconds
            time.Sleep(time.Duration(rand.Intn(5)+1) * time.Second)
        }
    }()

    return alarmChannel
}
