package adapter

import "time"

// Data

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

type TwoFAValidateData struct {
	Token string `json:"token"`
}

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

type CreateUserData struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Role     string `json:"role"`
	Notify   bool   `json:"notify"`
}

type FilterUsersData struct {
	Id       *[]uint   `json:"id"`
	Username *[]string `json:"username"`
	Email    *[]string `json:"email"`
	Role     *[]string `json:"role"`
}

type CreateRoleData struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type FilterRolesData struct {
	Id   *[]string `json:"id"`
	Name *[]string `json:"name"`
}

type UpdateRoleData struct {
	Name string `json:"name"`
}

// Results

type SigninResult struct {
	Access  string
	Refresh string
	Mfa     bool
}

type SignupResult struct {
	Id       uint
	Created  time.Time
	Username string
	Email    string
	Name     string
	Role     SignupRoleResult
	Mfa      bool
}
type SignupRoleResult struct {
	Id      string
	Name    string
	Created time.Time
}

type ProfileResult struct {
	Id       uint
	Created  time.Time
	Username string
	Email    string
	Name     string
	Role     ProfileRoleResult
	Mfa      bool
	Device   string
}
type ProfileRoleResult struct {
	Id      string
	Name    string
	Created time.Time
}

type TwoFAValidateResult struct {
	Access  string
	Refresh string
}

type TwoFASettingsResult struct {
	Secret string
	Url    string
}

type CreateUserResult struct {
	Id       uint
	Created  time.Time
	Username string
	Email    string
	Name     string
	Role     CreateUserRoleResult
	Mfa      bool
}
type CreateUserRoleResult struct {
	Id      string
	Name    string
	Created time.Time
}

type FilterUsersResult struct {
	Id       uint
	Created  time.Time
	Username string
	Email    string
	Name     string
	Role     FilterUsersRoleResult
	Mfa      bool
}
type FilterUsersRoleResult struct {
	Id      string
	Name    string
	Created time.Time
}

type CreateRoleResult struct {
	Id      string
	Name    string
	Created time.Time
}

type FilterRolesResult struct {
	Id      string
	Name    string
	Created time.Time
}
