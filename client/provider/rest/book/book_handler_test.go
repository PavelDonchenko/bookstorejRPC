package rest

import (
	"testing"

	controllers "github.com/PavelDonchenko/bookstorejRPC/client/controllers/book"
	"github.com/gin-gonic/gin"
)

func Test_routerBookHandler_CreateBook(t *testing.T) {
	type fields struct {
		c *controllers.BaseBookHandler
	}
	type args struct {
		ctx *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &routerBookHandler{
				c: tt.fields.c,
			}
			h.CreateBook(tt.args.ctx)
		})
	}
}
