package domain

import "time"

type Level int

const (
	_ Level = iota
	Info
	Warning
	Critical
)

type LogTask struct {
	ID      int
	Level   Level
	Message string
}

type LogResult struct {
	WorkerID   int
	LogID      int
	DetectedAt time.Time
}
