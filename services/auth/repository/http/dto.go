package adapter

import "time"

// Data

type TokenRenewData struct {
	RefreshToken string `json:"refresh_token"`
}

type TokenValidateData struct {
	AccessToken string `json:"access_token"`
}

type LogoutDeviceData struct {
	Device string `json:"device"`
}

type AuthData struct {
	User               uint      `json:"user"`
	Role               string    `json:"role"`
	Mfa                bool      `json:"mfa"`
	Device             string    `json:"device"`
	MetaLocation       string    `json:"meta_location"`
	MetaIp             string    `json:"meta_ip"`
	MetaUserAgent      string    `json:"meta_user_agent"`
	MetaOsFullName     string    `json:"meta_os_full_name"`
	MetaOsName         string    `json:"meta_os_name"`
	MetaOsVersion      string    `json:"meta_os_version"`
	MetaPlatform       string    `json:"meta_platform"`
	MetaModel          string    `json:"meta_model"`
	MetaBrowserName    string    `json:"meta_browser_name"`
	MetaBrowserVersion string    `json:"meta_browser_version"`
	MetaEngineName     string    `json:"meta_engine_name"`
	MetaEngineVersion  string    `json:"meta_engine_version"`
	Ttl                time.Time `json:"ttl"`
}

type Auth2faData struct {
	User   uint      `json:"user"`
	Role   string    `json:"role"`
	Device string    `json:"device"`
	Ttl    time.Time `json:"ttl"`
}

// Results

type TokenRenewResult struct {
	Access  string `json:"access_token"`
	Refresh string `json:"refresh_token"`
	Mfa     bool   `json:"mfa_required"`
}

type TokenValidateResult struct {
	Id       string   `json:"id"`
	Device   string   `json:"device"`
	User     uint     `json:"user"`
	Role     string   `json:"role"`
	Mfa      bool     `json:"mfa"`
	Expires  int64    `json:"expires"`
	Issued   int64    `json:"issued"`
	Issuer   string   `json:"issuer"`
	Audience []string `json:"audience"`
}

type DeviceResult struct {
	Id      string        `json:"id"`
	Session SessionResult `json:"session"`
}
type SessionResult struct {
	IssuedAt       string `json:"issued_at"`
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

type AuthResult struct {
	Access  string `json:"access"`
	Refresh string `json:"refresh"`
	Mfa     bool   `json:"mfa"`
}

type Auth2faResult struct {
	Access  string `json:"access"`
	Refresh string `json:"refresh"`
}
