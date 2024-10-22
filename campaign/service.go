package campaign

type Service interface {
	GetCampaigns(UserID int) ([]Campaign, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetCampaigns(UserID int) ([]Campaign, error) {
	// KALAU ADA USER ID DIA BAKAL PANGGIL FIND BY USER ID NYA
	if UserID != 0 {
		campaigns, err := s.repository.FindByUserId(UserID)
		if err != nil {
			return campaigns, err
		}
		return campaigns, nil
	}

	// KALAU TIDAK ADA USER ID DIA BAKAL PANGGIL FIND ALL
	campaigns, err := s.repository.FindAll()
	if err != nil {
		return campaigns, err
	}

	return campaigns, nil
}
