package models

import "time"

type Smaff_m_room struct {
  Id int
  Name string
  Security_Level int
  Created_At time.Time
  Updated_At time.Time
  Deleted_At time.Time
}
