package campaign

type Repository interface {
	Save(campaign *Campaign) (string, error)
}
