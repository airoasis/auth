package handler

import (
	"fmt"
	"github.com/airoasis/auth/model"
	"github.com/airoasis/auth/model/entity"
	"github.com/airoasis/auth/repository"
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"github.com/jinzhu/copier"
	"github.com/rs/zerolog/log"
	"github.com/tidwall/gjson"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

//CreateUser ... Create User
func CreateUser(c *gin.Context) {
	var userRequestDTO model.UserRequestDTO
	if err := c.ShouldBindJSON(&userRequestDTO); err == nil {
		if token, err := provisionAcaPyAgent(userRequestDTO.Username, userRequestDTO.Username); err == nil {
			hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userRequestDTO.Password), 8)
			userRequestDTO.Password = string(hashedPassword)

			var user entity.User
			copier.Copy(&user, userRequestDTO)
			user.AcapyToken = token

			err = repository.CreateUser(&user)
			if err != nil {
				fmt.Println(err.Error())
				c.AbortWithStatus(http.StatusNotFound)
			} else {
				var userResponseDTO model.UserResponseDTO
				copier.Copy(&userResponseDTO, user)
				c.JSON(http.StatusOK, userResponseDTO)
			}
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

func provisionAcaPyAgent(walletName, label string) (token string, err error){
	client := resty.New()

	resp, err := client.R().
		SetBody(map[string]interface{}{
			"wallet_name": walletName,
			"label": label,
		}).Post("http://localhost:8081/wallet")
	if err != nil {
		log.Error().Err(err).Msg("ERROR sending the request")
		return
	}

	if resp.StatusCode() == 200 {
		token = gjson.Get(resp.String(), "token").String()
	}

	return
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
			if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userAuthRequestDTO.Password)); err != nil {
				c.AbortWithStatus(http.StatusUnauthorized)
			} else {
				var userResponseDTO model.UserResponseDTO
				copier.Copy(&userResponseDTO, user)
				c.JSON(http.StatusOK, userResponseDTO)
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