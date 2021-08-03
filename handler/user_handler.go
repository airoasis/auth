package handler

import (
	"fmt"
	"github.com/airoasis/auth/model"
	"github.com/airoasis/auth/model/entity"
	"github.com/airoasis/auth/repository"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"net/http"
)

//CreateUser ... Create User
func CreateUser(c *gin.Context) {
	var userRequestDTO model.UserRequestDTO
	if err := c.ShouldBindJSON(&userRequestDTO); err == nil {
		var user entity.User
		copier.Copy(&user, userRequestDTO)

		err := repository.CreateUser(&user)
		if err != nil {
			fmt.Println(err.Error())
			c.AbortWithStatus(http.StatusNotFound)
		} else {
			var userResponseDTO model.UserResponseDTO
			copier.Copy(&userResponseDTO, user)
			c.JSON(http.StatusOK, userResponseDTO)
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

//GetUserByUsernameAndPassword ... Request from OAuth2
func GetUserByUsernameAndPassword(c *gin.Context) {
	var userAuthRequestDTO model.UserAuthRequestDTO
	if err := c.ShouldBindJSON(&userAuthRequestDTO); err == nil {
		var user entity.User
		err := repository.GetUserByUsername(&user, userAuthRequestDTO.Username)
		if err != nil {
			c.AbortWithStatus(http.StatusNotFound)
		} else {
			if user.Password == userAuthRequestDTO.Password {
				var userResponseDTO model.UserResponseDTO
				copier.Copy(&userResponseDTO, user)
				c.JSON(http.StatusOK, userResponseDTO)
			} else {
				c.AbortWithStatus(http.StatusNotFound)
			}
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

//GetUserByID ... Get the user by id
func GetUserByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var user entity.User
	err := repository.GetUserByID(&user, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		var userResponseDTO model.UserResponseDTO
		copier.Copy(&userResponseDTO, user)
		c.JSON(http.StatusOK, userResponseDTO)
	}
}

//DeleteUser ... Delete the user
func DeleteUser(c *gin.Context) {
	var user entity.User
	id := c.Params.ByName("id")
	err := repository.DeleteUser(&user, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id" + id: "is deleted"})
	}
}