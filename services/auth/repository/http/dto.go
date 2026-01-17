package adapter

import "time"

// Data

type LogoutDeviceData struct {
	Device string `json:"device"`
}

type CreateRoleData struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	SystemFlag  bool   `json:"system_flag"`
	ServiceFlag bool   `json:"service_flag"`
}

type FilterRolesData struct {
	Id          *[]string `json:"id,omitempty"`
	Name        *[]string `json:"name,omitempty"`
	SystemFlag  *bool     `json:"system_flag,omitempty"`
	ServiceFlag *bool     `json:"service_flag,omitempty"`
}

type UpdateRoleData struct {
	Id          *string `json:"id,omitempty"`
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
}

type CreateHttpRuleData struct {
	RoleId  string   `json:"role_id"`
	Path    string   `json:"path"`
	Methods []string `json:"methods"`
	Mfa     bool     `json:"mfa"`
}

type FilterHttpRulesData struct {
	Id      *[]uint   `json:"id,omitempty"`
	RoleId  *[]string `json:"role_id,omitempty"`
	Path    *[]string `json:"path,omitempty"`
	Methods *[]string `json:"methods,omitempty"`
	Mfa     *bool     `json:"mfa,omitempty"`
}

type UpdateHttpRuleData struct {
	RoleId  *uint     `json:"role_id,omitempty"`
	Path    *string   `json:"path,omitempty"`
	Methods *[]string `json:"methods,omitempty"`
	Mfa     *bool     `json:"mfa,omitempty"`
}

type AuthData struct {
	User               uint      `json:"user"`
	Roles              []string  `json:"roles"`
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
	Roles  []string  `json:"roles"`
	Device string    `json:"device"`
	Ttl    time.Time `json:"ttl"`
}

type TokenRenewData struct {
	RefreshToken string `json:"refresh_token"`
}

type TokenAuthorizeHttpData struct {
	Path   string `json:"path"`
	Method string `json:"method"`
}

type CreateStaticAccessTokenData struct {
	Id          string   `json:"id"`
	Roles       []string `json:"roles"`
	Description string   `json:"description"`
}

type FilterStaticAccessTokenData struct {
	Id *[]string `json:"id,omitempty"`
}

// Results

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

type CreateRoleResult struct {
	Id          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	SystemFlag  bool      `json:"system_flag"`
	ServiceFlag bool      `json:"service_flag"`
	Created     time.Time `json:"created"`
	Updated     time.Time `json:"updated"`
}

type FilterRolesResult struct {
	Id          string                      `json:"id"`
	Name        string                      `json:"name"`
	Description string                      `json:"description"`
	SystemFlag  bool                        `json:"system_flag"`
	ServiceFlag bool                        `json:"service_flag"`
	Created     time.Time                   `json:"created"`
	Updated     time.Time                   `json:"updated"`
	HttpRules   []FilterRolesHttpRuleResult `json:"http_rules"`
}
type FilterRolesHttpRuleResult struct {
	Id      uint      `json:"id"`
	RoleId  string    `json:"role_id"`
	Path    string    `json:"path"`
	Methods []string  `json:"methods"`
	Mfa     bool      `json:"mfa"`
	Created time.Time `json:"created"`
	Updated time.Time `json:"updated"`
}

type CreateHttpRuleResult struct {
	Id      uint      `json:"id"`
	RoleId  string    `json:"role_id"`
	Path    string    `json:"path"`
	Methods []string  `json:"methods"`
	Mfa     bool      `json:"mfa"`
	Created time.Time `json:"created"`
	Updated time.Time `json:"updated"`
}

type FilterHttpRulesResult struct {
	Id      uint      `json:"id"`
	RoleId  string    `json:"role_id"`
	Path    string    `json:"path"`
	Methods []string  `json:"methods"`
	Mfa     bool      `json:"mfa"`
	Created time.Time `json:"created"`
	Updated time.Time `json:"updated"`
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

type TokenRenewResult struct {
	Access  string `json:"access_token"`
	Refresh string `json:"refresh_token"`
	Mfa     bool   `json:"mfa_required"`
}

type TokenValidateResult struct {
	Id       string   `json:"id"`
	Device   string   `json:"device"`
	User     uint     `json:"user"`
	Roles    []string `json:"roles"`
	Mfa      bool     `json:"mfa"`
	Expires  *int64   `json:"expires"`
	Issued   int64    `json:"issued"`
	Issuer   string   `json:"issuer"`
	Audience []string `json:"audience"`
}

type TokenAuthorizeHttpResult struct {
	Token TokenAuthorizeHttpDataResult `json:"token"`
	Auth  TokenAuthorizeHttpAuthResult `json:"auth"`
}
type TokenAuthorizeHttpDataResult struct {
	Id       string   `json:"id"`
	Device   string   `json:"device"`
	User     uint     `json:"user"`
	Roles    []string `json:"roles"`
	Mfa      bool     `json:"mfa"`
	Expires  *int64   `json:"expires"`
	Issued   int64    `json:"issued"`
	Issuer   string   `json:"issuer"`
	Audience []string `json:"audience"`
}
type TokenAuthorizeHttpAuthResult struct {
	Mfa bool `json:"mfa"`
}

type CreateStaticAccessTokenResult struct {
	Token string `json:"token"`
}

type FilterStaticAccessTokenResult struct {
	Id          string    `json:"id"`
	Token       string    `json:"token"`
	UserId      uint      `json:"user_id"`
	Device      string    `json:"device"`
	Roles       []string  `json:"roles"`
	Description string    `json:"description"`
	Created     time.Time `json:"created"`
}
