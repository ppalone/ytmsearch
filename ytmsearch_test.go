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
