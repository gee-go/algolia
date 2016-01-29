package algolia_test

import (
	"testing"

	"github.com/gee-go/algolia/algolia"
	"github.com/stretchr/testify/require"
)

func Test_Client(t *testing.T) {
	t.Parallel()

	assert := require.New(t)
	c := algolia.FromEnv()
}
