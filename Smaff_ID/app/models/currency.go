package models

import "time"

type SmaffTCurrency struct {
  EmployeerId string
  Amount int
  CreatedAt time.Time
  UpdatedAt time.Time
  DeletedAt time.Time
}
