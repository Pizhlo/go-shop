package api

import (
	"net/http"
	db "github.com/Pizhlo/go-shop/db/sqlc"
	"github.com/lib/pq"
	"github.com/gin-gonic/gin"
	"log"
)

type CreateUserParams struct {
	Username       string `json:"username"`
	Email          string `json:"email"`
	HashedPassword string `json:"hashed_password"`
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
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

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
