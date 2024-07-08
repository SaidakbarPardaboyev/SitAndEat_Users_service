package handler

import (
	"encoding/json"
	"log"
	"net/http"
	pb "restaurant/genproto/users"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func (h *Handler) Register(ctx *gin.Context) {

	req := pb.RegisterUser{}
	err := json.NewDecoder(ctx.Request.Body).Decode(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
		log.Println(err)
		return
	}

	hashedpassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
		log.Println(err)
		return
	}
	req.Password = string(hashedpassword)

	err = h.UserRepo.Register(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
		log.Println(err)
		return
	}

	ctx.JSON(http.StatusAccepted, nil)
}

func (h *Handler) Login(ctx *gin.Context) {

}
