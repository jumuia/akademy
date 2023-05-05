package handler

type createUserRequest struct {
	FirstName   string `json:"first_name,omitempty"`
	MiddleName  string `json:"middle_name,omitempty"`
	LastName    string `json:"last_name,omitempty"`
	DOB         string `json:"dob,omitempty" time_format:"2006-01-02" time_utc:"1"`
	PhoneNumber string `json:"phone_number,omitempty"`
	Email       string `json:"email,omitempty"`
}

type fetchUserRequest struct {
	ID string `json:"id,omitempty" uri:"user_id" binding:"required"`
}
