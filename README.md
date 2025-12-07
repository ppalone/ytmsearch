# YTMSearch

Search Songs & Videos from Youtube Music

## Installing

```bash
go get github.com/ppalone/ytmsearch
```

## Usage

### Creating a client

```go
// pass your http.Client if required
c := ytmsearch.NewClient(nil)
```

### Searching songs

Use `.Search()` with no filters or explicitly providing `ytmsearch.WithSearchType(ytmsearch.SONGS)` filter

```go
q := "nocopyrightsounds"
res, err := c.Search(context.Background(), q)
// or with explicit "songs" filter
// res, err := c.Search(context.Background(), q, ytmsearch.WithSearchType(ytmsearch.SONGS))

if err != nil {
  // handle error
}

// res.Results contains search results
for _, song := range res.Results {
  fmt.Println(song)
}
```

### Searching video songs

Use `.Search()` by providing `ytmsearch.WithSearchType(ytmsearch.VIDEOS)` filter

```go
q := "nocopyrightsounds"
res, err := c.Search(context.Background(), q, ytmsearch.WithSearchType(ytmsearch.VIDEOS))
```

### Pagination

If the result set has more results, use `.HasNext` and `.Continuation`:

```go
if res.HasNext {
    resNext, err := c.SearchNext(context.Background(), res.Continuation)
    // Use resNext.Results...
}
```

## Author

Pranjal

## LICENSE

MIT
