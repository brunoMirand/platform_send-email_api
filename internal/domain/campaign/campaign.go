package campaign

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Contact struct {
	Email string
}

type Campaign struct {
	ID        string
	Name      string
	Content   string
	Contacts  []Contact
	CreatedAt time.Time
}

func NewCampaign(name string, content string, emails []string) (*Campaign, error) {
	if name == "" {
		return nil, errors.New("field <name> is required")
	}

	id := uuid.New()
	contacts := make([]Contact, len(emails))
	for index, email := range emails {
		contacts[index].Email = email
	}

	return &Campaign{
		ID:        id.String(),
		Name:      name + ".",
		Content:   content,
		Contacts:  contacts,
		CreatedAt: time.Now(),
	}, nil
}
