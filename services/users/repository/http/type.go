package adapter

import "time"

// Responses

type signinResponse struct {
	Access  string `json:"access_token"`
	Refresh string `json:"refresh_token"`
	Mfa     bool   `json:"mfa_required"`
}

type signupResponse struct {
	Id         uint               `json:"id"`
	Created    time.Time          `json:"created"`
	Username   string             `json:"username"`
	Email      string             `json:"email"`
	Name       string             `json:"name"`
	Role       signupRoleResponse `json:"role"`
	Mfa        bool               `json:"mfa"`
	SystemFlag bool               `json:"system_flag"`
}
type signupRoleResponse struct {
	Id         string    `json:"id"`
	Name       string    `json:"name"`
	SystemFlag bool      `json:"system_flag"`
	Created    time.Time `json:"created"`
}

type profileResponse struct {
	Id         uint                `json:"id"`
	Created    time.Time           `json:"created"`
	Username   string              `json:"username"`
	Email      string              `json:"email"`
	Name       string              `json:"name"`
	Role       profileRoleResponse `json:"role"`
	Mfa        bool                `json:"mfa"`
	SystemFlag bool                `json:"system_flag"`
	Device     string              `json:"device"`
}
type profileRoleResponse struct {
	Id         string    `json:"id"`
	Name       string    `json:"name"`
	SystemFlag bool      `json:"system_flag"`
	Created    time.Time `json:"created"`
}

type twoFAValidateResponse struct {
	Access  string `json:"access_token"`
	Refresh string `json:"refresh_token"`
}

type twoFASettingsResponse struct {
	Secret string `json:"secret"`
	Url    string `json:"url"`
}

type createUserResponse struct {
	Id         uint                   `json:"id"`
	Created    time.Time              `json:"created"`
	Username   string                 `json:"username"`
	Email      string                 `json:"email"`
	Name       string                 `json:"name"`
	Role       createUserRoleResponse `json:"role"`
	Mfa        bool                   `json:"mfa"`
	SystemFlag bool                   `json:"system_flag"`
}
type createUserRoleResponse struct {
	Id         string    `json:"id"`
	Name       string    `json:"name"`
	SystemFlag bool      `json:"system_flag"`
	Created    time.Time `json:"created"`
}

type filterUsersResponse struct {
	Id         uint                    `json:"id"`
	Created    time.Time               `json:"created"`
	Username   string                  `json:"username"`
	Email      string                  `json:"email"`
	Name       string                  `json:"name"`
	Role       filterUsersRoleResponse `json:"role"`
	Mfa        bool                    `json:"mfa"`
	SystemFlag bool                    `json:"system_flag"`
}
type filterUsersRoleResponse struct {
	Id         string    `json:"id"`
	Name       string    `json:"name"`
	SystemFlag bool      `json:"system_flag"`
	Created    time.Time `json:"created"`
}

type createRoleResponse struct {
	Id         string    `json:"id"`
	Name       string    `json:"name"`
	SystemFlag bool      `json:"system_flag"`
	Created    time.Time `json:"created"`
}

type filterRolesResponse struct {
	Id         string    `json:"id"`
	Name       string    `json:"name"`
	SystemFlag bool      `json:"system_flag"`
	Created    time.Time `json:"created"`
}
