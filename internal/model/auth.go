package model

type Auth struct {
	Id            string `json:"id,omitempty"`
	Email         string `json:"email,omitempty"`
	EmailVerified bool   `json:"email_verified,omitempty"`
	FullName      string `json:"full_name,omitempty"`
	Image         string `json:"image,omitempty"`
}
