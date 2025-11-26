package ytmsearch

type SearchType int

const (
	SONGS SearchType = iota
	VIDEOS
)

type searchFilters struct {
	searchType SearchType
}

// https://github.com/raitonoberu/ytmusic/blob/master/enums.go
var searchFiltersMap map[SearchType]string = map[SearchType]string{
	SONGS:  "EgWKAQIIAWoMEA4QChADEAQQCRAF",
	VIDEOS: "EgWKAQIQAWoMEA4QChADEAQQCRAF",
}

func defaultSearchFilters() *searchFilters {
	return &searchFilters{
		searchType: SONGS,
	}
}

type SearchFilter func(f *searchFilters)

func WithSearchType(searchType SearchType) SearchFilter {
	return func(f *searchFilters) {
		f.searchType = searchType
	}
}
