package campaign

import (
	"errors"
	"strconv"

	"github.com/gosimple/slug"
)

type Service interface {
	GetCampaigns(UserID int) ([]Campaign, error)
	GetCampaignByID(input GetCampaignDetailInput) (Campaign, error)
	CreateCampaign(input CreateCampaignInput) (Campaign, error)
	UpdateCampaign(inputID GetCampaignDetailInput, inputData CreateCampaignInput) (Campaign, error)
	SaveCampaignImage(input CreateCampaignImageInput, fileLocation string) (CampaignImage, error)
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

func (s *service) GetCampaignByID(input GetCampaignDetailInput) (Campaign, error) {
	campaign, err := s.repository.FindByID(input.ID)
	if err != nil {
		return campaign, err
	}

	return campaign, nil
}

func (s *service) CreateCampaign(input CreateCampaignInput) (Campaign, error) {
	campaign := Campaign{}
	campaign.Name = input.Name
	campaign.ShortDescription = input.ShortDescription
	campaign.Description = input.Description
	campaign.GoalAmount = input.GoalAmount
	campaign.Perks = input.Perks
	campaign.UserID = input.User.ID
	campaign.Slug = slug.Make(input.Name + " " + strconv.Itoa(input.User.ID))

	// make slug

	newCampaign, err := s.repository.Save(campaign)

	if err != nil {
		return newCampaign, err
	}

	return newCampaign, nil
}

func (s *service) UpdateCampaign(inputID GetCampaignDetailInput, inputData CreateCampaignInput) (Campaign, error) {
	campaign, err := s.repository.FindByID(inputID.ID)
	if err != nil {
		return campaign, err
	}

	if campaign.UserID != inputData.User.ID {
		return campaign, errors.New("Not an owner of the campaign")
	}

	campaign.Name = inputData.Name
	campaign.ShortDescription = inputData.ShortDescription
	campaign.Description = inputData.Description
	campaign.GoalAmount = inputData.GoalAmount
	campaign.Perks = inputData.Perks

	updatedCampaign, err := s.repository.Update(campaign)
	if err != nil {
		return updatedCampaign, err
	}

	return updatedCampaign, nil
}

func (s *service) SaveCampaignImage(input CreateCampaignImageInput, fileLocation string) (CampaignImage, error) {
	// get campaign first to check user
	campaign, err := s.repository.FindByID(input.CampaignID)
	if err != nil {
		return CampaignImage{}, err
	}

	if campaign.UserID != input.User.ID {
		return CampaignImage{}, errors.New("Not an owner of the campaign")
	}

	if input.IsPrimary != nil && *input.IsPrimary {
		_, err := s.repository.MarkAllImagesAsNonPrimary(input.CampaignID)
		if err != nil {
			return CampaignImage{}, err
		}
	}

	campaignImage := CampaignImage{}

	campaignImage.CampaignID = input.CampaignID
	campaignImage.IsPrimary = input.IsPrimary != nil && *input.IsPrimary
	campaignImage.FileName = fileLocation

	newCampaignImage, err := s.repository.CreateImage(campaignImage)

	if err != nil {
		return newCampaignImage, err
	}

	return newCampaignImage, nil
}
