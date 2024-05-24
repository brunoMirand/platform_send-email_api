package campaign

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCampaign(t *testing.T) {
	//given
	assert := assert.New(t)
	name := "Aprenda PHP em 2024"
	content := "Seja um dev destaque todas as semanas da sprint."
	contacts := []string{"bruno.sm@gmail.com", "bruno.ms@gmail.com"}

	//when
	sut := NewCampaign(name, content, contacts)

	//then
	assert.Equal(sut.Name, "Aprenda PHP em 2024.")
	assert.Equal(sut.Content, "Seja um dev destaque todas as semanas da sprint.")
}
