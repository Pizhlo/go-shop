package api

import (
	"log"
	"net/http"

	db "github.com/Pizhlo/go-shop/db/sqlc"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

type CreateUserParams struct {
	Username       string `form:"username" binding:"required,alphanum"`
	Email          string `form:"email" binding:"required,email"`
	HashedPassword string `form:"hashed_password" binding:"required,min=6"` // TODO implement password -> hashed password
}


type userResponse struct {
	Username          string    `json:"username"`
	Email             string    `json:"email"`
}

func newUserResponse(user db.User) userResponse {
	return userResponse{
		Username:          user.Username,
		Email:             user.Email,
	}
}

func (server *Server) createUser(ctx *gin.Context) {	
	var req CreateUserParams
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return  
	}  // TODO сделать отображение ошибок на форме

	arg := db.CreateUserParams{
		Username: req.Username,
		Email: req.Email,
		HashedPassword: req.HashedPassword,
	}

	user, err := server.store.CreateUser(ctx, arg) 

	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			log.Println(pqErr.Code.Name())
			switch pqErr.Code.Name() {
			case "unique_violation":
				ctx.JSON(http.StatusForbidden, errorResponse(err))
				return
			}
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	rsp := newUserResponse(user)

	ctx.JSON(http.StatusOK, rsp)
}
