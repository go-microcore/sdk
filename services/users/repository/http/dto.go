package adapter

import "time"

// Data

type TwoFASettingsData struct {
	Password string `json:"password"`
}

type TwoFAEnableData struct {
	Token string `json:"token"`
}

type TwoFADisableData struct {
	Password string `json:"password"`
	Token    string `json:"token"`
}

type TwoFAValidateData struct {
	Token string `json:"token"`
}

type SigninData struct {
	Login    string          `json:"login"`
	Password string          `json:"password"`
	Device   string          `json:"device"`
	Metadata *SigninMetadata `json:"metadata"`
}
type SigninMetadata struct {
	Location       string `json:"location"`
	Ip             string `json:"ip"`
	UserAgent      string `json:"user_agent"`
	OsFullName     string `json:"os_full_name"`
	OsName         string `json:"os_name"`
	OsVersion      string `json:"os_version"`
	Platform       string `json:"platform"`
	Model          string `json:"model"`
	BrowserName    string `json:"browser_name"`
	BrowserVersion string `json:"browser_version"`
	EngineName     string `json:"engine_name"`
	EngineVersion  string `json:"engine_version"`
}

type SignupData struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

type FilterUsersData struct {
	Id         *[]uint   `json:"id"`
	Username   *[]string `json:"username"`
	Email      *[]string `json:"email"`
	Roles      *[]string `json:"roles"`
	OtpSecret  *[]string `json:"otp_secret"`
	Mfa        *bool     `json:"mfa"`
	SystemFlag *bool     `json:"system_flag"`
}

type UpdateUserData struct {
	Name       string   `json:"name"`
	Username   string   `json:"username"`
	Email      string   `json:"email"`
	Roles      []string `json:"roles"`
	SystemFlag bool     `json:"system_flag"`
}

type CreateUserData struct {
	Username   string   `json:"username"`
	Email      string   `json:"email"`
	Password   string   `json:"password"`
	Name       string   `json:"name"`
	Roles      []string `json:"roles"`
	Notify     bool     `json:"notify"`
	SystemFlag bool     `json:"system_flag"`
}

// Results

type TwoFASettingsResult struct {
	Secret string `json:"secret"`
	Url    string `json:"url"`
}

type TwoFAValidateResult struct {
	Access  string `json:"access_token"`
	Refresh string `json:"refresh_token"`
}

type SigninResult struct {
	Access  string `json:"access_token"`
	Refresh string `json:"refresh_token"`
	Mfa     bool   `json:"mfa_required"`
}

type SignupResult struct {
	Id         uint      `json:"id"`
	Created    time.Time `json:"created"`
	Username   string    `json:"username"`
	Email      string    `json:"email"`
	Name       string    `json:"name"`
	Roles      []string  `json:"roles"`
	Mfa        bool      `json:"mfa"`
	SystemFlag bool      `json:"system_flag"`
}

type ProfileResult struct {
	Id         uint      `json:"id"`
	Created    time.Time `json:"created"`
	Username   string    `json:"username"`
	Email      string    `json:"email"`
	Name       string    `json:"name"`
	Roles      []string  `json:"roles"`
	Mfa        bool      `json:"mfa"`
	SystemFlag bool      `json:"system_flag"`
	Device     string    `json:"device"`
}

type FilterUsersResult struct {
	Id         uint      `json:"id"`
	Created    time.Time `json:"created"`
	Username   string    `json:"username"`
	Email      string    `json:"email"`
	Name       string    `json:"name"`
	Roles      []string  `json:"roles"`
	Mfa        bool      `json:"mfa"`
	SystemFlag bool      `json:"system_flag"`
}

type CreateUserResult struct {
	Id         uint      `json:"id"`
	Created    time.Time `json:"created"`
	Username   string    `json:"username"`
	Email      string    `json:"email"`
	Name       string    `json:"name"`
	Roles      []string  `json:"roles"`
	Mfa        bool      `json:"mfa"`
	SystemFlag bool      `json:"system_flag"`
}
