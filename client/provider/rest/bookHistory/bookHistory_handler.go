package rest

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	controllers2 "github.com/PavelDonchenko/bookstorejRPC/client/controllers/bookHandler"
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

func (h *RouterBookHistoryHandler) SearchBookHistory(ctx *gin.Context) {
	var query string
	if query, _ = ctx.GetQuery("q"); query == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "no search query present"})
		return
	}

	body := fmt.Sprintf(
		`{"query": {"multi_match": {"query": "%s", "fields": ["title", "body"]}}}`,
		query)
	res, err := h.ESClient.Search(
		h.ESClient.Search.WithContext(context.Background()),
		h.ESClient.Search.WithIndex("posts"),
		h.ESClient.Search.WithBody(strings.NewReader(body)),
		h.ESClient.Search.WithPretty(),
	)
	if err != nil {
		fmt.Errorf("elasticsearch error")
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer res.Body.Close()

	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			fmt.Errorf("error parsing the response body")
		} else {
			// Print the response status and error information.
			fmt.Errorf("failed to search query: [%s] %s: %s",
				res.Status(),
				e["error"].(map[string]interface{})["type"],
				e["error"].(map[string]interface{})["reason"],
			)
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": e["error"].(map[string]interface{})["reason"]})
		return
	}

	fmt.Printf("res", res.Status())

	var r map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		fmt.Println("elasticsearch error")
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": r["hits"]})
}
