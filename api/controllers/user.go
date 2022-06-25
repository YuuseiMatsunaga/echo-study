package controllers

import (
  "echo/database"
  "net/http"
  "github.com/labstack/echo/v4"
  "time"
  "strconv"
  // "echo/model"
)

type User struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	CreatedAt string `json:created_at`
  }
  

func GetUsers(c echo.Context) error {
  users := []User{}
  database.DB.Find(&users)
  return c.JSON(http.StatusOK, users)
}

func GetUser(c echo.Context) error {
  userId, _ := strconv.Atoi(c.Param("id"))
  user := User{Id: userId}
  database.DB.Find(&user)

  if user.Name == "" {
    return c.JSON(http.StatusOK, nil) 
  } else {
    return c.JSON(http.StatusOK, user) 
  }
}

func CreateUser(c echo.Context) error {
  user := User{}
  if err := c.Bind(&user); err != nil {
    return err
  }
  user.CreatedAt = setNow()
  database.DB.Create(&user)
  return c.JSON(http.StatusOK, user)
}

func UpdateUser(c echo.Context) error {
  user := User{}
  if err := c.Bind(&user); err != nil {
    return err
  }
  database.DB.Save(&user)
  return c.JSON(http.StatusOK, user)
}

func DeleteUser(c echo.Context) error {
  id := c.Param("id")
  database.DB.Delete(&User{}, id)
  return c.NoContent(http.StatusNoContent)
}

func setNow() string {
  now := time.Now()
  return now.Format("2006-01-02 15:04:05") 
}