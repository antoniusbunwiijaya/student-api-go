package web

import "antoniusbunwijaya/student-api-go/model/domain"

type StudentResponse struct {
	Id        int          `json:"id"`
	Name      string       `json:"name"`
	Age       int          `json:"age"`
	Gender    int          `json:"gender"`
	CreatedAt string       `json:"createdAt"`
	Major     domain.Major `json:"major"`
	//Hobbies   []domain.Hobby `json:"hobbies"`
}
