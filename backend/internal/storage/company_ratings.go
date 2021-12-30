package storage

import "github.com/elhmn/camerdevs/pkg/models/v1beta"

//GetCompanyRatings get company ratings
func (db DB) GetCompanyRatings() ([]v1beta.CompanyRating, error) {
	ratings := []v1beta.CompanyRating{}
	ret := db.c.Find(&ratings)
	if ret.Error != nil {
		return ratings, ret.Error
	}

	return ratings, nil
}

//GetCompanyRatingsByID get company ratings by ID
func (db DB) GetCompanyRatingsByID(id int64) (v1beta.CompanyRating, error) {
	rating := v1beta.CompanyRating{}
	ret := db.c.First(&rating, "id = ?", id)
	if ret.Error != nil {
		return rating, ret.Error
	}

	return rating, nil
}