package handler

import (
	"encoding/json"
	"net/http"
	"restaurant/api/token"
	pb "restaurant/genproto/users"
	"restaurant/models"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// @Summary Register a new user
// @Description Register a new user with the provided details
// @Tags auth
// @Accept json
// @Produce json
// @Param user body users.RegisterUser true "User registration details"
// @Success 202 {object} nil
// @Failure 400 {object} error
// @Router /users/register [post]
func (h *Handler) Register(ctx *gin.Context) {

	req := pb.RegisterUser{}
	err := json.NewDecoder(ctx.Request.Body).Decode(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
		h.logger.Error(err.Error())
		return
	}

	hashedpassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
		h.logger.Error(err.Error())
		return
	}
	req.Password = string(hashedpassword)

	err = h.UserRepo.Register(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
		h.logger.Error(err.Error())
		return
	}

	ctx.JSON(http.StatusAccepted, nil)
}

// @Summary User login
// @Description Authenticate a user and return JWT token
// @Tags auth
// @Accept json
// @Produce json
// @Param user body users.LoginUser true "User login details"
// @Success 202 {object} users.Token
// @Failure 400 {object} error
// @Failure 500 {object} error
// @Router /users/login [post]
func (h *Handler) Login(ctx *gin.Context) {
	req := pb.LoginUser{}

	if err := json.NewDecoder(ctx.Request.Body).Decode(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
		h.logger.Error(err.Error())
		return
	}

	user, err := h.UserRepo.GetUserByEmail(req.Email)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
		h.logger.Error(err.Error())
		return
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
		h.logger.Error(err.Error())
		return
	}

	token := token.GenerateJWT(&pb.GetUser{
		Id:       user.Id,
		Username: user.Username,
		Email:    req.Email,
		Phone:    user.Phone_number,
	})

	err = h.UserRepo.StoreRefreshToken(&models.RefreshToken{
		UserId:    user.Id,
		Token:     token.RefreshToken,
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	})
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
		h.logger.Error(err.Error())
		return
	}

	ctx.JSON(http.StatusAccepted, token)
}
