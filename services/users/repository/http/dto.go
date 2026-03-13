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
	UserAgent      string `json:"userAgent"`
	OsFullName     string `json:"osFullName"`
	OsName         string `json:"osName"`
	OsVersion      string `json:"osVersion"`
	Platform       string `json:"platform"`
	Model          string `json:"model"`
	BrowserName    string `json:"browserName"`
	BrowserVersion string `json:"browserVersion"`
	EngineName     string `json:"engineName"`
	EngineVersion  string `json:"engineVersion"`
}

type SignupData struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

type FilterUsersData struct {
	ID         *[]uint   `json:"id"`
	Username   *[]string `json:"username"`
	Email      *[]string `json:"email"`
	Roles      *[]string `json:"roles"`
	OtpSecret  *[]string `json:"otpSecret"`
	Mfa        *bool     `json:"mfa"`
	SystemFlag *bool     `json:"systemFlag"`
}

type UpdateUserData struct {
	Name       string   `json:"name"`
	Username   string   `json:"username"`
	Email      string   `json:"email"`
	Roles      []string `json:"roles"`
	SystemFlag bool     `json:"systemFlag"`
}

type CreateUserData struct {
	Username   string   `json:"username"`
	Email      string   `json:"email"`
	Password   string   `json:"password"`
	Name       string   `json:"name"`
	Roles      []string `json:"roles"`
	Notify     bool     `json:"notify"`
	SystemFlag bool     `json:"systemFlag"`
}

// Results

type TwoFASettingsResult struct {
	Secret string `json:"secret"`
	Url    string `json:"url"`
}

type TwoFAValidateResult struct {
	Access  string `json:"accessToken"`
	Refresh string `json:"refreshToken"`
}

type SigninResult struct {
	Access  string `json:"accessToken"`
	Refresh string `json:"refreshToken"`
	Mfa     bool   `json:"mfaRequired"`
}

type SignupResult struct {
	ID         uint      `json:"id"`
	Created    time.Time `json:"created"`
	Username   string    `json:"username"`
	Email      string    `json:"email"`
	Name       string    `json:"name"`
	Roles      []string  `json:"roles"`
	Mfa        bool      `json:"mfa"`
	SystemFlag bool      `json:"systemFlag"`
}

type ProfileResult struct {
	ID         uint      `json:"id"`
	Created    time.Time `json:"created"`
	Username   string    `json:"username"`
	Email      string    `json:"email"`
	Name       string    `json:"name"`
	Roles      []string  `json:"roles"`
	Mfa        bool      `json:"mfa"`
	SystemFlag bool      `json:"systemFlag"`
	Device     string    `json:"device"`
}

type FilterUsersResult struct {
	ID         uint      `json:"id"`
	Created    time.Time `json:"created"`
	Username   string    `json:"username"`
	Email      string    `json:"email"`
	Name       string    `json:"name"`
	Roles      []string  `json:"roles"`
	Mfa        bool      `json:"mfa"`
	SystemFlag bool      `json:"systemFlag"`
}

type CreateUserResult struct {
	ID         uint      `json:"id"`
	Created    time.Time `json:"created"`
	Username   string    `json:"username"`
	Email      string    `json:"email"`
	Name       string    `json:"name"`
	Roles      []string  `json:"roles"`
	Mfa        bool      `json:"mfa"`
	SystemFlag bool      `json:"systemFlag"`
}
