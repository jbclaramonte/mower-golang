package service

import (
	"testing"

	"github.com/jbclaramonte/mower-golang/core/domain"
	"github.com/stretchr/testify/assert"
)

func TestCli(t *testing.T) {

	t.Run("string content should be parsed into differents parts", func(t *testing.T) {
		contentString := `4 4
2 2 N
A`
		g, m, c := TranslateContentToCommands(contentString)

		assert.Equal(t, domain.Garden{Width: 4, Height: 4}, *g)
		assert.Equal(t, domain.Mower{X: 2, Y: 2, Orientation: domain.North}, *m)
		assert.Equal(t, "A", c)
	})
}
