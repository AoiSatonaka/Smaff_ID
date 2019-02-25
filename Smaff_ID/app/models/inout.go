package models

import "time"

type SmaffTInout struct {
  Id int
  EmployeerId string
  RoomId int
  CreatedAt time.Time
}