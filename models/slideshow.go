package models

type Picture struct {
	ID 		string `json:"id"`
	Name 	string `json:"name"`
	URL		string `json:"url"`
}

func GetSlideshow() ([]*Picture , error) {

	var pictures  = make([]*Picture,0)

	if err := DB.Find(&pictures).Error;err != nil {
		return nil, err
	}
	return pictures , nil

}
