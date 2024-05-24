package campaign

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	name     = "Aprenda PHP em 2024"
	content  = "Seja um dev destaque todas as semanas da sprint."
	contacts = []string{"bruno.sm@gmail.com", "bruno.ms@gmail.com"}
)

func TestNewCampaign(t *testing.T) {
	//given
	assert := assert.New(t)

	//when
	sut, _ := NewCampaign(name, content, contacts)

	//then
	assert.Equal(sut.Name, "Aprenda PHP em 2024.")
	assert.Equal(sut.Content, "Seja um dev destaque todas as semanas da sprint.")
}

func TestShouldErrorWhenCreateNewCampaign(t *testing.T) {
	//given
	assert := assert.New(t)

	//when
	_, err := NewCampaign("", content, contacts)

	//then
	assert.Equal("field <name> is required", err.Error())
}
