package controller

import (
	"net/http"

	"github.com/KPWithCode/hurd/config"
	"github.com/KPWithCode/hurd/models"
	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	users := []models.User{}
	// if err = db.DB.Find(user).Error; err !=nil {
	// 	return err
	// }
	config.DB.Find(&users)
	c.JSON(200, &users)
}
func CreateUser(c *gin.Context) {
	user := models.User{}
	c.ShouldBindJSON(&user)
	config.DB.Create(&user)
	c.JSON(http.StatusOK, gin.H{"data": &user})
}
func DeleteUser(c *gin.Context) {
	user := models.User{}
	config.DB.Where("id = ?", c.Param("id")).Delete(&user)
	c.JSON(http.StatusOK, gin.H{"data": user})
}
func UpdateUser(c *gin.Context) {
	user := []models.User{}
	config.DB.Where("id = ?", c.Param("id")).First(&user)
	config.DB.Save(&user)
	c.JSON(200, &user)
}
