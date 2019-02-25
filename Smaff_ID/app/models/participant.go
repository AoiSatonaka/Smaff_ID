package models

import "time"

type SmaffMParticipant struct {
  ConferenceId int
  EmployeerId string
  Attendance bool
  CreatedAt time.Time
  UpdatedAt time.Time
  DeletedAt time.Time
}