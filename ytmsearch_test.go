package ytmsearch_test

import (
	"testing"

	"github.com/ppalone/ytmsearch"
	"github.com/stretchr/testify/assert"
)

func Test_NewClient(t *testing.T) {
	c := ytmsearch.NewClient(nil)
	assert.NotNil(t, c)
}
