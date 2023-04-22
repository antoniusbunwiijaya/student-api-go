package web

type StudentCreateRequest struct {
	Name    string   `validate:"required" json:"name"`
	Age     int      `validate:"required" json:"age"`
	Gender  int      `json:"gender"`
	Major   string   `validate:"required" json:"major"`
	Hobbies []string `json:"hobbies"`
}
