package rest

import (
	"net/http"
	"strconv"

	controllers2 "github.com/PavelDonchenko/bookstorejRPC/client/controllers/bookHistory"
	pb "github.com/PavelDonchenko/bookstorejRPC/client/gen/proto"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/gin-gonic/gin"
)

type BookHistoryInput struct {
	Id         uint64 `json:"id"`
	BookAuthor string `json:"book_author"`
	BookName   string `json:"book_name"`
	LastUser   string `json:"last_user"`
}

type RouterBookHistoryHandler struct {
	c        *controllers2.BaseBookHistoryHandler
	ESClient *elasticsearch.Client
}

func NewRouterBookHistoryHandler(c *controllers2.BaseBookHistoryHandler, esClient *elasticsearch.Client) *RouterBookHistoryHandler {
	return &RouterBookHistoryHandler{
		c:        c,
		ESClient: esClient,
	}
}

func (h *RouterBookHistoryHandler) GetOneBookHistory(ctx *gin.Context) {
	BHId, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.c.GetOneBookHistory(BHId)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"response": res})
}
func (h *RouterBookHistoryHandler) CreateBookHistory(ctx *gin.Context) {
	var input BookHistoryInput

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	BookHistory := pb.BookHistoryItem{
		Id:         input.Id,
		BookAuthor: input.BookAuthor,
		BookName:   input.BookName,
		LastUser:   input.LastUser,
	}

	res, err := h.c.InsertBookHistory(&BookHistory)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"response": res})
}
func (h *RouterBookHistoryHandler) DeleteBookHistory(ctx *gin.Context) {
	BHId, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.c.DeleteBookHistory(BHId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"response": res})
}
