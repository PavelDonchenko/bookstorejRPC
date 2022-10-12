package rest

import (
	"net/http"
	"strconv"

	controllers "github.com/PavelDonchenko/bookstorejRPC/client/controllers/book"
	pb "github.com/PavelDonchenko/bookstorejRPC/client/gen/proto"
	"github.com/gin-gonic/gin"
)

type BookInput struct {
	Id         uint32 `json:"id"`
	Name       string `json:"name"`
	BookAuthor string `json:"bookAuthor"`
}

type routerBookHandler struct {
	c *controllers.BaseBookHandler
}

func NewRouterBookHandler(c *controllers.BaseBookHandler) *routerBookHandler {
	return &routerBookHandler{
		c: c,
	}
}

func (h *routerBookHandler) GetAll(ctx *gin.Context) {
	res, err := h.c.GetAll()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"response": res})
}

func (h *routerBookHandler) GetOne(ctx *gin.Context) {
	bookId, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.c.Get(uint32(bookId))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"response": res})
}

func (h *routerBookHandler) Create(ctx *gin.Context) {
	var input BookInput

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if input.Name == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Title is required"})
		return
	}

	book := pb.BookItem{Name: input.Name, BookAuthor: input.BookAuthor}

	res, err := h.c.Create(&book)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"response": res})
}

func (h *routerBookHandler) Update(ctx *gin.Context) {

	var input BookInput

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	book := pb.BookItem{Name: input.Name, BookAuthor: input.BookAuthor}

	resUpdate, err := h.c.Update(&book)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"response": resUpdate})
}

func (h *routerBookHandler) Delete(ctx *gin.Context) {
	bookId, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.c.Delete(uint32(bookId))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"response": res})
}
