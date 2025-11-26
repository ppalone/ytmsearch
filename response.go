package ytmsearch

import (
	"errors"
	"strings"
)

var (
	ErrNoResults = errors.New("no results")
)

type intertubeSearchResponse struct {
	Contents struct {
		TabbedSearchResultsRenderer struct {
			Tabs []struct {
				TabRenderer struct {
					Content struct {
						SectionListRenderer struct {
							Contents []intertubeContent `json:"contents"`
						} `json:"sectionListRenderer"`
					} `json:"content"`
				} `json:"tabRenderer"`
			} `json:"tabs"`
		} `json:"tabbedSearchResultsRenderer"`
	} `json:"contents"`
}

type intertubeContent struct {
	MusicShelfRenderer struct {
		Contents []struct {
			MusicResponsiveListItemRenderer struct {
				Thumbnail struct {
					MusicThumbnailRenderer struct {
						Thumbnail struct {
							Thumbnails []struct {
								URL    string `json:"url"`
								Width  int    `json:"width"`
								Height int    `json:"height"`
							} `json:"thumbnails"`
						} `json:"thumbnail"`
					} `json:"musicThumbnailRenderer"`
				} `json:"thumbnail"`
				FlexColumns []struct {
					MusicResponsiveListItemFlexColumnRenderer struct {
						Text struct {
							Runs []struct {
								Text               string `json:"text"`
								NavigationEndpoint struct {
									WatchEndpoint struct {
										VideoID string `json:"videoId"`
									} `json:"watchEndpoint"`
								} `json:"navigationEndpoint"`
							} `json:"runs"`
						} `json:"text"`
					} `json:"musicResponsiveListItemFlexColumnRenderer"`
				} `json:"flexColumns"`
			} `json:"musicResponsiveListItemRenderer"`
		} `json:"contents"`
		Continuations []struct {
			NextContinuationData struct {
				Continuation string `json:"continuation"`
			} `json:"nextContinuationData"`
		} `json:"continuations"`
	} `json:"musicShelfRenderer"`
}

func (r *intertubeSearchResponse) toResults() (SearchResults, error) {
	tabs := r.Contents.TabbedSearchResultsRenderer.Tabs
	if len(tabs) == 0 {
		return SearchResults{}, ErrNoResults
	}
	tab := tabs[0]

	contents := tab.TabRenderer.Content.SectionListRenderer.Contents
	if len(contents) == 0 {
		return SearchResults{}, ErrNoResults
	}

	return extract(contents[0]), nil
}

func extract(content intertubeContent) SearchResults {
	items := make([]MusicItem, 0)
	for _, c := range content.MusicShelfRenderer.Contents {

		// get thumbnails
		thumbnails := make([]Thumbnail, 0)
		for _, t := range c.MusicResponsiveListItemRenderer.Thumbnail.MusicThumbnailRenderer.Thumbnail.Thumbnails {
			thumbnails = append(thumbnails, Thumbnail{
				URL:    t.URL,
				Width:  uint(t.Width),
				Height: uint(t.Height),
			})
		}

		cols := c.MusicResponsiveListItemRenderer.FlexColumns
		if len(cols) < 2 {
			continue
		}

		// get title and id
		info := cols[0].MusicResponsiveListItemFlexColumnRenderer.Text.Runs
		if len(info) == 0 {
			continue
		}
		title := info[0].Text
		id := info[0].NavigationEndpoint.WatchEndpoint.VideoID

		// skip
		if len(id) == 0 {
			continue
		}

		// other metadata
		meta := cols[1].MusicResponsiveListItemFlexColumnRenderer.Text.Runs
		var views, duration string

		n := len(cols)
		m := len(meta)
		switch n {
		case 3: // songs
			duration = meta[m-1].Text
			count := cols[2].MusicResponsiveListItemFlexColumnRenderer.Text.Runs
			if len(count) > 0 {
				views = strings.TrimSuffix(count[0].Text, " plays")
			}
		case 2: // videos
			duration = meta[m-1].Text
			views = strings.TrimSuffix(meta[m-3].Text, " views")
		default:
			continue
		}

		items = append(items, MusicItem{
			VideoID:    id,
			Title:      title,
			Thumbnails: thumbnails,
			Views:      views,
			Duration:   duration,
		})
	}

	continuation := ""
	c := content.MusicShelfRenderer.Continuations
	if len(c) > 0 {
		continuation = c[0].NextContinuationData.Continuation
	}

	return SearchResults{
		Results:      items,
		HasNext:      len(continuation) != 0,
		Continuation: continuation,
	}
}
