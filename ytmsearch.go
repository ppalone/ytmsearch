package ytmsearch

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

const (
	ytMusicBaseURL = "https://music.youtube.com/youtubei/v1/search"
)

type YTMSearch struct {
	httpClient *http.Client
}

type MusicItem struct {
	VideoID    string
	Title      string
	Thumbnails []Thumbnail
	Views      string
	Duration   string
}

// Thumbnail
type Thumbnail struct {
	URL    string `json:"url"`
	Width  uint   `json:"width"`
	Height uint   `json:"height"`
}

type SearchResults struct {
	Results      []MusicItem
	HasNext      bool
	Continuation string
}

func NewClient(httpClient *http.Client) *YTMSearch {
	if httpClient == nil {
		httpClient = &http.Client{}
	}
	return &YTMSearch{httpClient}
}

func (c *YTMSearch) Search(ctx context.Context, q string, filters ...SearchFilter) (SearchResults, error) {
	f := defaultSearchFilters()

	for _, filter := range filters {
		filter(f)
	}

	return c.search(ctx, q, f)
}

func (c *YTMSearch) search(ctx context.Context, q string, filters *searchFilters) (SearchResults, error) {
	searchFilter, ok := searchFiltersMap[filters.searchType]
	if !ok {
		return SearchResults{}, fmt.Errorf("invalid search type provided")
	}

	body := map[string]string{
		"query":  q,
		"params": searchFilter,
	}

	req, err := makeRequest(ctx, body, nil)
	if err != nil {
		return SearchResults{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return SearchResults{}, err
	}

	apiResponse := new(intertubeSearchResponse)
	err = json.NewDecoder(resp.Body).Decode(apiResponse)
	if err != nil {
		return SearchResults{}, err
	}

	return apiResponse.toResults()
}

func makeRequest(ctx context.Context, body map[string]string, _ url.Values) (*http.Request, error) {
	payload := map[string]any{
		"context": map[string]any{
			"client": map[string]any{
				"clientName":    "WEB_REMIX",
				"clientVersion": "1.20251119.03.01",
				"hl":            "en",
				"gl":            "IN",
			},
		},
	}

	// add to payload
	for k, v := range body {
		payload[k] = v
	}

	data, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, ytMusicBaseURL, bytes.NewReader(data))
	if err != nil {
		return nil, err
	}

	// TODO: add params to request

	req.Header.Set("Content-Type", "application/json")
	return req, nil
}
