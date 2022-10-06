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

type routerHandler struct {
	c *controllers.BaseHandler
}

func NewRouterHandler(c *controllers.BaseHandler) *routerHandler {
	return &routerHandler{
		c: c,
	}
}

func (h routerHandler) GetAll(ctx *gin.Context) {
	res, err := h.c.GetAll()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"response": res})
}

func (h routerHandler) GetOne(ctx *gin.Context) {
	userId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.c.Get(uint32(userId))

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"response": res})
}

func (h routerHandler) Create(ctx *gin.Context) {
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

	res, err := h.c.Create(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"response": res})
}

func (h routerHandler) Update(ctx *gin.Context) {
	var input UserInput

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := pb.UserItem{Nickname: input.Nickname, Email: input.Email, Password: input.Password}

	res, err := h.c.Update(&user)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"response": res})
}

func (h routerHandler) Delete(ctx *gin.Context) {
	var input UserInput

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := pb.UserItem{Id: input.Id}

	res, err := h.c.Delete(user.Id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"response": res})
}
