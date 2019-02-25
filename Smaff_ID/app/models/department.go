package models

import "time"

type Smaff_m_department struct {
  Id int
  CompanyId int
  Name string
  DefaultSecurityLevel int
  CreatedAt time.Time
  UpdatedAt time.Time
  DeletedAt time.Time
}
