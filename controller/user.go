package controller

import (
	"github.com/KPWithCode/hurd/db"
	"github.com/KPWithCode/hurd/models"
	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	users := []models.User{}
	db.DB.Find(&users)
	c.JSON(200, &users)
}
func CreateUser(c *gin.Context) {
	user := []models.User{}
	c.BindJSON(&user)
	db.DB.Create(&user)
	c.JSON(200, &user)
}
func DeleteUser(c *gin.Context) {
	user := models.User{}
	db.DB.Where("id = ?", c.Param("id")).Delete(&user)
	c.JSON(200, &user)

}
func UpdateUser(c *gin.Context) {
	user := []models.User{}
	db.DB.Where("id = ?", c.Param("id")).First(&user)
	db.DB.Save(&user)
	c.JSON(200, &user)
}
