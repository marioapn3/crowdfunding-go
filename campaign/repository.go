package campaign

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]Campaign, error)
	FindByUserId(UserID int) ([]Campaign, error)
	FindByID(ID int) (Campaign, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(DB *gorm.DB) *repository {
	return &repository{DB}
}

func (r *repository) FindAll() ([]Campaign, error) {
	var campaigns []Campaign
	err := r.db.Preload("CampaignImages", "campaign_images.is_primary = true").Find(&campaigns).Error
	if err != nil {
		return campaigns, err
	}

	return campaigns, nil
}

func (r *repository) FindByUserId(UserID int) ([]Campaign, error) {
	var campaigns []Campaign
	err := r.db.Where("user_id = ?", UserID).Preload("CampaignImages", "campaign_images.is_primary = true").Find(&campaigns).Error
	if err != nil {
		return campaigns, err
	}

	return campaigns, nil

}

func (r *repository) FindByID(ID int) (Campaign, error) {
	var campaign Campaign

	err := r.db.Preload("CampaignImages").Preload("User").Where("id = ?", ID).Find(&campaign).Error

	if err != nil {
		return campaign, err
	}

	return campaign, nil
}
