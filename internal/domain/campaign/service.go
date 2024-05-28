package campaign

import (
	"send-email/internal/contract"
	internalerrors "send-email/internal/internal-errors"
)

type Service struct {
	Repository Repository
}

func (s *Service) Create(input contract.Campaign) (string, error) {
	campaign, err := NewCampaign(input.Name, input.Content, input.Emails)
	if err != nil {
		return "", err
	}
	id, err := s.Repository.Save(campaign)
	if err != nil {
		return "", internalerrors.ErrInternal
	}

	return id, nil
}
