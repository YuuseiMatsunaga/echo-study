package controllers

import (
  "echo/database"
  "net/http"
  "github.com/labstack/echo/v4"
  "time"
  "echo/entities"
)

func GetUsers(c echo.Context) error {
  users := []entities.User{}
  database.DB.Find(&users)
  return c.JSON(http.StatusOK, users)
}

func GetUser(c echo.Context) error {
  user := setUser(c.Param("id"))

  if user.Name == "" {
    return c.JSON(http.StatusOK, nil) 
  } else {
    return c.JSON(http.StatusOK, user) 
  }
}

func CreateUser(c echo.Context) error {
  user := new(entities.User)
  if err := c.Bind(&user); err != nil {
    return err
  }
  user.CreatedAt = time.Now()
  database.DB.Create(&user)
  return c.JSON(http.StatusOK, user)
}

func UpdateUser(c echo.Context) error {
  user := setUser(c.Param("id"))
  if err := c.Bind(&user); err != nil {
    return err
  }
  database.DB.Save(&user)
  return c.JSON(http.StatusOK, user)
}

func DeleteUser(c echo.Context) error {
  id := c.Param("id")
  database.DB.Delete(&entities.User{}, id)
  return c.NoContent(http.StatusNoContent)
}

func setUser(id string) entities.User {
  user := entities.User{}
  database.DB.First(&user, id)
  return user
}