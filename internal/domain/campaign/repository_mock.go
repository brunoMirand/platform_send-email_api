package campaign

import (
	"github.com/stretchr/testify/mock"
)

type RepositoryMock struct {
	mock.Mock
}

func (r *RepositoryMock) Save(campaign *Campaign) (string, error) {
	args := r.Called(campaign)
	return args.String(0), args.Error(1)
}
