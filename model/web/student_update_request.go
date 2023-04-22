package web

type StudentUpdateRequest struct {
	Id     int    `validate:"required" json:"id"`
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Gender int    `json:"gender"`
	//CreatedAt string       `json:"created_at"`
	//Major   string   `json:"major"`
}
