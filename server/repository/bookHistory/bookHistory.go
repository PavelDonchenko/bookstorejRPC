package bookHistory

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	model "github.com/PavelDonchenko/bookstorejRPC/server/models"
	"github.com/PavelDonchenko/bookstorejRPC/server/utils"
	"github.com/elastic/go-elasticsearch/v8/esapi"
)

type BookHistoryESStorage struct {
	elastic ElasticSearch
	timeout time.Duration
}

func NewBookHistoryStorage(elastic ElasticSearch) *BookHistoryESStorage {
	return &BookHistoryESStorage{
		elastic: elastic,
		timeout: time.Second * 10,
	}
}

func (bs *BookHistoryESStorage) GetOneBookHistory(ctx context.Context, id uint64) (model.BookHistory, error) {
	// res, err := p.elastic.client.Get()
	req := esapi.GetRequest{
		Index:      bs.elastic.alias,
		DocumentID: strconv.FormatUint(id, 10),
	}

	ctx, cancel := context.WithTimeout(ctx, bs.timeout)
	defer cancel()

	res, err := req.Do(ctx, bs.elastic.client)
	if err != nil {
		return model.BookHistory{}, fmt.Errorf("find one: request: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode == 404 {
		return model.BookHistory{}, utils.ErrNotFound
	}

	if res.IsError() {
		return model.BookHistory{}, fmt.Errorf("find one: response: %s", res.String())
	}

	var (
		bookHistory model.BookHistory
		body        document
	)
	body.Source = &bookHistory

	if err := json.NewDecoder(res.Body).Decode(&body); err != nil {
		return model.BookHistory{}, fmt.Errorf("find one: decode: %w", err)
	}

	return bookHistory, nil
}

func (bs *BookHistoryESStorage) InsertBookHistory(ctx context.Context, bh model.BookHistory) error {
	body, err := json.Marshal(bh)
	if err != nil {
		return fmt.Errorf("insert: marshall: %w", err)
	}
	// res, err := p.elastic.client.Create()
	req := esapi.CreateRequest{
		Index:      bs.elastic.alias,
		DocumentID: strconv.FormatUint(bh.Id, 10),
		Body:       bytes.NewReader(body),
	}

	ctx, cancel := context.WithTimeout(ctx, bs.timeout)
	defer cancel()

	res, err := req.Do(ctx, bs.elastic.client)
	if err != nil {
		if err != nil {
			return fmt.Errorf("insert: request: %w", err)
		}
	}
	defer res.Body.Close()

	if res.StatusCode == 409 {
		return utils.ErrConflict
	}

	if res.IsError() {
		return fmt.Errorf("insert: response: %s", res.String())
	}

	return nil
}

func (bs *BookHistoryESStorage) DeleteBookHistory(ctx context.Context, id uint64) (bool, error) {
	// res, err := p.elastic.client.Delete()
	req := esapi.DeleteRequest{
		Index:      bs.elastic.alias,
		DocumentID: strconv.FormatUint(id, 10),
	}

	ctx, cancel := context.WithTimeout(ctx, bs.timeout)
	defer cancel()

	res, err := req.Do(ctx, bs.elastic.client)
	if err != nil {
		return false, fmt.Errorf("delete: request: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode == 404 {
		return false, utils.ErrNotFound
	}

	if res.IsError() {
		return false, fmt.Errorf("delete: response: %s", res.String())
	}

	return true, nil
}

// document represents a single document in Get API response body.
type document struct {
	Source interface{} `json:"_source"`
}
