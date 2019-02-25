package controllers

import (
  "Smaff_ID/app/models"
  "fmt"
  "github.com/revel/revel"
)

type Interceptor struct {
}

func (i *Interceptor) DbConnect() revel.Result{
  fmt.Println("DB Connecting...")
  models.InitDB()
  return nil
}

func (i *Interceptor) DbClose() revel.Result{
  fmt.Println("DB Closing...")
  models.DB.Close()
  return nil
}
