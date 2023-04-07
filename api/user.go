package api

import (
	"fmt"
	"net/http"

	db "github.com/Pizhlo/go-shop/db/sqlc"
	"github.com/gin-gonic/gin"
)

type CreateUserParams struct {
	FirstName string `form:"first_name" json:"first_name" binding:"required"`
	LastName  string `form:"last_name" json:"last_name" binding:"required"`
	Username  string `form:"username" json:"username" binding:"required,alphanum"`
	Email     string `form:"email" json:"email" binding:"required,email"`
	Password  string `form:"password" json:"password" binding:"required,min=6"`
}

type userResponse struct {
	Username   string `json:"username"`
	Email      string `json:"email"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Favourites string `json: "favourites"`
}

func newUserResponse(user db.User) userResponse {
	return userResponse{
		Username:   user.Username,
		Email:      user.Email,
		FirstName:  user.FirstName,
		LastName:   user.LastName,
		Favourites: string(user.Favourites),
	}
}

func (server *Server) createUser(ctx *gin.Context) {
	var req CreateUserParams
	if err := ctx.Bind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err)) // TODO сделать проверку на то что пароли совпадают
		return
	} // TODO сделать отображение ошибок на форме

	// hashedPassword, err := util.HashPassword(req.Password)
	// if err != nil {
	// 	ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	// 	return
	// }

	// arg := db.CreateUserParams{
	// 	Username:       req.Username,
	// 	Email:          req.Email,
	// 	HashedPassword: hashedPassword,
	// 	FirstName:      req.FirstName,
	// 	LastName:       req.LastName,
	// }

	//user, err := server.store.CreateUser(ctx, arg)

	// if err != nil {
	// 	if pqErr, ok := err.(*pq.Error); ok {
	// 		log.Println(pqErr.Code.Name())
	// 		switch pqErr.Code.Name() {
	// 		case "unique_violation":
	// 			ctx.JSON(http.StatusForbidden, errorResponse(err))
	// 			return
	// 		}
	// 	}
	// 	ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	// 	return
	// }

	// server.loginUser(ctx)

	// rsp := newUserResponse(user)

	//ctx.JSON(http.StatusOK, rsp)

}

type getUserRequest struct {
	ID int32 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getUser(ctx *gin.Context) {
	var req getUserRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	//account, err := server.store.GetUserByID(ctx, req.ID)
	// if err != nil {
	// 	if err == sql.ErrNoRows {
	// 		ctx.JSON(http.StatusNotFound, errorResponse(err))
	// 		return
	// 	}
	// 	ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	// 	return
	// }

	// authPayload := ctx.MustGet(authorizationHeaderKey).(*token.Payload)
	// if account.Username != authPayload.Username {
	// 	err = errors.New("account doesn't belong to the authenticated user")
	// 	ctx.JSON(http.StatusUnauthorized, errorResponse(err))
	// 	return
	// }

	// ctx.JSON(http.StatusOK, account) // TODO Key "authorization" does not exist
}

type loginUserRequest struct {
	Username string `form:"username" binding:"required,alphanum"`
	Password string `form:"password" binding:"required,min=6"`
}

type loginUserResponse struct {
	AccessToken string       `json:"access_token"`
	User        userResponse `json:"user"`
}

func (server *Server) loginUser(ctx *gin.Context) {
	var req loginUserRequest
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	fmt.Println("login user")

	// user, err := server.store.GetUser(ctx, req.Username)
	// if err != nil {
	// 	if err == sql.ErrNoRows {
	// 		ctx.JSON(http.StatusNotFound, errorResponse(err))
	// 		return
	// 	}
	// 	ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	// 	return
	// }

	// err = util.CheckPassword(req.Password, user.HashedPassword)
	// if err != nil {
	// 	ctx.JSON(http.StatusUnauthorized, errorResponse(err))
	// 	return
	// }
	// accessToken, err := server.tokenMaker.CreateToken(user.Username, server.config.AccessTokenDuration)

	// if err != nil {
	// 	ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	// 	return
	// }

	// rsp := loginUserResponse{
	// 	AccessToken: accessToken,
	// 	User:        newUserResponse(user),
	// }
	// ctx.JSON(http.StatusOK, rsp)
	// ctx.Set("user", rsp.User)
	// fmt.Println("ctx: ", ctx.Keys)

	ctx.HTML(http.StatusOK, "account.html", gin.H{"title": "Личный кабинет",
		"auth": ctx.Keys["is_logged_in"], "user": ctx.Keys["user"]})
	fmt.Println("3 keys = ", ctx.Keys["is_logged_in"])
}
