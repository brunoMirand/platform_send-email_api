package campaign

import (
	"send-email/internal/contract"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateCampaign(t *testing.T) {
	//given
	assert := assert.New(t)
	repository := new(RepositoryMock)
	input := contract.Campaign{
		Name:    "Aprenda PHP 8",
		Content: "Seja o destaque!",
		Emails:  []string{"phpnet@gmail.com", "odevphp@gmail.com"},
	}
	expectedID := "12345"
	repository.On("Save", mock.AnythingOfType("*campaign.Campaign")).Return(expectedID, nil)

	sut := Service{Repository: repository}

	//when
	id, error := sut.Create(input)

	//then
	assert.Nil(error)
	assert.NotNil(id)
	repository.AssertCalled(t, "Save", mock.AnythingOfType("*campaign.Campaign"))
}
