package campaign

import (
	"send-email/internal/validators"
	"time"

	"github.com/google/uuid"
)

type Contact struct {
	Email string `validate:"email"`
}

type Campaign struct {
	ID        string    `validate:"required"`
	Name      string    `validate:"min=5,max=24"`
	Content   string    `validate:"min=5,max=100"`
	Contacts  []Contact `validate:"min=1,dive"`
	CreatedAt time.Time `validate:"required"`
}

func NewCampaign(name string, content string, emails []string) (*Campaign, error) {
	id := uuid.New()
	contacts := make([]Contact, len(emails))
	for index, email := range emails {
		contacts[index].Email = email
	}

	campaign := &Campaign{
		ID:        id.String(),
		Name:      name + ".",
		Content:   content,
		Contacts:  contacts,
		CreatedAt: time.Now(),
	}

	err := validators.ValidateStruct(campaign)
	if err == nil {
		return campaign, nil
	}

	return nil, err
}
