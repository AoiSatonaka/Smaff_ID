package controllers

import (
  "crypto/rsa"
  "fmt"
  "github.com/dgrijalva/jwt-go"
  "github.com/revel/revel"
  "io/ioutil"
  "net/http"
  "time"
)

type User struct {
  *revel.Controller
}

var (
  verifyKey *rsa.PublicKey
  signKey *rsa.PrivateKey
)

func (u User) Login() revel.Result {
  email:=u.Params.Form.Get("email")
  password:=u.Params.Form.Get("password")

  signBites,err := ioutil.ReadFile("./smaff.rsa")
  if err != nil {
    panic(err)
  }

  signKey,err := jwt.ParseRSAPrivateKeyFromPEM(signBites)
  if err != nil {
    panic(err)
  }

  if email == "" && password == "" {
    //token生成
    token := jwt.New(jwt.SigningMethodPS256)

    claims := token.Claims.(jwt.MapClaims)
    claims["name"] = ""
    claims["admin"] = true
    claims["exp"] = time.Now().Add(time.Hour * 3).Unix()      //トークンの有効時間を３時間に設定

    tokenString,err := token.SignedString(signKey)
    if err != nil {
      fmt.Println(err)
    }
    u.Response.Status=http.StatusOK
    u.RenderJSON(tokenString);
  }
  return nil
}
