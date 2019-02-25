package models

import "time"

type SmaffMEmployeer struct {
  Email  string
  Department_Id int
  Password string
  SecurityLevel int
  FirstName string
  LastName string
  FirstEnglishName string
  LastEnglishName string
  Tel string
  HireDate time.Time `sql:"not null;type:date"`
  CreatedAt time.Time
  UpdatedAt time.Time
  DeletedAt time.Time
}
