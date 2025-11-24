package ytmsearch

import "net/http"

type YTMSearch struct {
	httpClient *http.Client
}

func NewClient(httpClient *http.Client) *YTMSearch {
	if httpClient == nil {
		httpClient = &http.Client{}
	}
	return &YTMSearch{httpClient}
}
