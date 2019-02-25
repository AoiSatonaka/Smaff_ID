package models

import "time"

type SmaffMConference struct {
  Id int
  Room_Id int
  Name string
  StartTime time.Time
  CreatedAt time.Time
  UpdatedAt time.Time
  DeletedAt time.Time
}
