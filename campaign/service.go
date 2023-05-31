package campaign

type Service interface {
	GetCampaigns(userID int) ([]Campaign, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) Getampaigns(userId int) ([]Campaign, error) {
	if userId != 0 {
		campaigns, err := s.repository.FindByUserID(userId)

		if err != nil {
			return campaigns, err
		}

		return campaigns, nil

	}

	campaigns, err := s.repository.FindByUserID(userId)

	if err != nil {
		return campaigns, err
	}

	return campaigns, nil
}
