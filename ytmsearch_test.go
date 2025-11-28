package ytmsearch_test

import (
	"context"
	"testing"

	"github.com/ppalone/ytmsearch"
	"github.com/stretchr/testify/assert"
)

func Test_NewClient(t *testing.T) {
	c := ytmsearch.NewClient(nil)
	assert.NotNil(t, c)
}

func Test_Search(t *testing.T) {
	t.Run("without any filters", func(t *testing.T) {
		c := ytmsearch.NewClient(nil)
		q := "nocopyrightsounds"
		res, err := c.Search(context.Background(), q)

		assert.NoError(t, err)
		assert.NotEmpty(t, res.Results)
		assert.True(t, res.HasNext)
		assert.NotEmpty(t, res.Continuation)

		for _, i := range res.Results {
			assert.NotEmpty(t, i.VideoID)
			assert.NotEmpty(t, i.Thumbnails)
			assert.NotEmpty(t, i.Title)
			assert.NotEmpty(t, i.Views)
			assert.NotEmpty(t, i.Duration)
			assert.Contains(t, i.Duration, ":")
		}
	})

	t.Run("with songs filters", func(t *testing.T) {
		c := ytmsearch.NewClient(nil)
		q := "nocopyrightsounds"
		res, err := c.Search(context.Background(), q, ytmsearch.WithSearchType(ytmsearch.SONGS))

		assert.NoError(t, err)
		assert.NotEmpty(t, res.Results)
		assert.True(t, res.HasNext)
		assert.NotEmpty(t, res.Continuation)

		for _, i := range res.Results {
			assert.NotEmpty(t, i.VideoID)
			assert.NotEmpty(t, i.Thumbnails)
			assert.NotEmpty(t, i.Title)
			assert.NotEmpty(t, i.Views)
			assert.NotEmpty(t, i.Duration)
			assert.Contains(t, i.Duration, ":")
		}
	})

	t.Run("with videos filters", func(t *testing.T) {
		c := ytmsearch.NewClient(nil)
		q := "nocopyrightsounds"
		res, err := c.Search(context.Background(), q, ytmsearch.WithSearchType(ytmsearch.VIDEOS))

		assert.NoError(t, err)
		assert.NotEmpty(t, res.Results)
		assert.True(t, res.HasNext)
		assert.NotEmpty(t, res.Continuation)

		for _, i := range res.Results {
			assert.NotEmpty(t, i.VideoID)
			assert.NotEmpty(t, i.Thumbnails)
			assert.NotEmpty(t, i.Title)
			assert.NotEmpty(t, i.Views)
			assert.NotEmpty(t, i.Duration)
			assert.Contains(t, i.Duration, ":")
		}
	})
}

func Test_SearchNext(t *testing.T) {
	t.Run("with valid token from no filters search", func(t *testing.T) {
		c := ytmsearch.NewClient(nil)
		q := "nocopyrightsounds"
		res, err := c.Search(context.Background(), q)
		assert.NoError(t, err)
		assert.NotEmpty(t, res.Results)
		assert.True(t, res.HasNext)

		resNext, err := c.SearchNext(context.Background(), res.Continuation)
		assert.NoError(t, err)
		assert.NotEmpty(t, resNext.Results)
		assert.True(t, resNext.HasNext)

		for _, i := range resNext.Results {
			assert.NotEmpty(t, i.VideoID)
			assert.NotEmpty(t, i.Thumbnails)
			assert.NotEmpty(t, i.Title)
			assert.NotEmpty(t, i.Views)
			assert.NotEmpty(t, i.Duration)
			assert.Contains(t, i.Duration, ":")
		}
	})

	t.Run("with valid token from songs search", func(t *testing.T) {
		c := ytmsearch.NewClient(nil)
		q := "nocopyrightsounds"
		res, err := c.Search(context.Background(), q, ytmsearch.WithSearchType(ytmsearch.SONGS))
		assert.NoError(t, err)
		assert.NotEmpty(t, res.Results)
		assert.True(t, res.HasNext)

		resNext, err := c.SearchNext(context.Background(), res.Continuation)
		assert.NoError(t, err)
		assert.NotEmpty(t, resNext.Results)
		assert.True(t, resNext.HasNext)

		for _, i := range resNext.Results {
			assert.NotEmpty(t, i.VideoID)
			assert.NotEmpty(t, i.Thumbnails)
			assert.NotEmpty(t, i.Title)
			assert.NotEmpty(t, i.Views)
			assert.NotEmpty(t, i.Duration)
			assert.Contains(t, i.Duration, ":")
		}
	})

	t.Run("with valid token from videos search", func(t *testing.T) {
		c := ytmsearch.NewClient(nil)
		q := "nocopyrightsounds"
		res, err := c.Search(context.Background(), q, ytmsearch.WithSearchType(ytmsearch.VIDEOS))
		assert.NoError(t, err)
		assert.NotEmpty(t, res.Results)
		assert.True(t, res.HasNext)

		resNext, err := c.SearchNext(context.Background(), res.Continuation)
		assert.NoError(t, err)
		assert.NotEmpty(t, resNext.Results)
		assert.True(t, resNext.HasNext)

		for _, i := range resNext.Results {
			assert.NotEmpty(t, i.VideoID)
			assert.NotEmpty(t, i.Thumbnails)
			assert.NotEmpty(t, i.Title)
			assert.NotEmpty(t, i.Views)
			assert.NotEmpty(t, i.Duration)
			assert.Contains(t, i.Duration, ":")
		}
	})

	t.Run("with invalid token", func(t *testing.T) {
		c := ytmsearch.NewClient(nil)
		res, err := c.SearchNext(context.Background(), "xxxxxxx")
		assert.NoError(t, err)
		assert.Empty(t, res.Results)
		assert.False(t, res.HasNext)
	})
}
