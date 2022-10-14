package rest

import (
	"net/http"
	"strconv"

	controllers "github.com/PavelDonchenko/bookstorejRPC/client/controllers/user"
	pb "github.com/PavelDonchenko/bookstorejRPC/client/gen/proto"
	"github.com/gin-gonic/gin"
)

type UserInput struct {
	Id       uint32 `json:"id"`
	Nickname string `json:"nickname"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type routerUserHandler struct {
	c *controllers.BaseUserHandler
}

func NewRouterUserHandler(c *controllers.BaseUserHandler) *routerUserHandler {
	return &routerUserHandler{
		c: c,
	}
}

func (h *routerUserHandler) GetAllUsers(ctx *gin.Context) {
	var page uint64 = 0
	var err error
	p := ctx.Param("page")

	if p != "" {
		page, err = strconv.ParseUint(p, 10, 32)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}

	res, err := h.c.GetAllUsers(uint32(page))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"response": res})
}

func (h *routerUserHandler) GetUser(ctx *gin.Context) {
	id := ctx.Param("id")
	Id, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.c.GetUser(uint32(Id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"response": res})
}

func (h *routerUserHandler) CreateUser(ctx *gin.Context) {
	var input UserInput

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if input.Nickname == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Title is required"})
		return
	}

	user := pb.UserItem{Nickname: input.Nickname, Email: input.Email, Password: input.Password}

	res, err := h.c.CreateUser(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"response": res})
}

func (h *routerUserHandler) UpdateUser(ctx *gin.Context) {
	_, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var input UserInput

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	book := pb.UserItem{Id: uint32(input.Id), Nickname: input.Nickname, Email: input.Email, Password: input.Password}

	res, err := h.c.UpdateUser(&book)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"response": res})
}

func (h *routerUserHandler) DeleteUser(ctx *gin.Context) {
	userId, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.c.DeleteUser(uint32(userId))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"response": res})
}
