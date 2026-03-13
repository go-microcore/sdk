package adapter

import "time"

// Data

type LogoutDeviceData struct {
	Device string `json:"device"`
}

type CreateRoleData struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	SystemFlag  bool   `json:"systemFlag"`
	ServiceFlag bool   `json:"serviceFlag"`
}

type FilterRolesData struct {
	ID          *[]string `json:"id,omitempty"`
	Name        *[]string `json:"name,omitempty"`
	SystemFlag  *bool     `json:"systemFlag,omitempty"`
	ServiceFlag *bool     `json:"serviceFlag,omitempty"`
}

type UpdateRoleData struct {
	ID          *string `json:"id,omitempty"`
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
}

type CreateHttpRuleData struct {
	RoleID  string   `json:"roleId"`
	Path    string   `json:"path"`
	Methods []string `json:"methods"`
	Mfa     bool     `json:"mfa"`
}

type FilterHttpRulesData struct {
	ID      *[]uint   `json:"id,omitempty"`
	RoleID  *[]string `json:"roleId,omitempty"`
	Path    *[]string `json:"path,omitempty"`
	Methods *[]string `json:"methods,omitempty"`
	Mfa     *bool     `json:"mfa,omitempty"`
}

type UpdateHttpRuleData struct {
	RoleID  *uint     `json:"roleId,omitempty"`
	Path    *string   `json:"path,omitempty"`
	Methods *[]string `json:"methods,omitempty"`
	Mfa     *bool     `json:"mfa,omitempty"`
}

type AuthData struct {
	User               uint      `json:"user"`
	Roles              []string  `json:"roles"`
	Mfa                bool      `json:"mfa"`
	Device             string    `json:"device"`
	MetaLocation       string    `json:"metaLocation"`
	MetaIP             string    `json:"metaIp"`
	MetaUserAgent      string    `json:"metaUserAgent"`
	MetaOsFullName     string    `json:"metaOsFullName"`
	MetaOsName         string    `json:"metaOsName"`
	MetaOsVersion      string    `json:"metaOsVersion"`
	MetaPlatform       string    `json:"metaPlatform"`
	MetaModel          string    `json:"metaModel"`
	MetaBrowserName    string    `json:"metaBrowserName"`
	MetaBrowserVersion string    `json:"metaBrowserVersion"`
	MetaEngineName     string    `json:"metaEngineName"`
	MetaEngineVersion  string    `json:"metaEngineVersion"`
	Ttl                time.Time `json:"ttl"`
}

type Auth2faData struct {
	User   uint      `json:"user"`
	Roles  []string  `json:"roles"`
	Device string    `json:"device"`
	Ttl    time.Time `json:"ttl"`
}

type TokenRenewData struct {
	RefreshToken string `json:"refreshToken"`
}

type TokenAuthorizeHttpData struct {
	Path   string `json:"path"`
	Method string `json:"method"`
}

type CreateStaticAccessTokenData struct {
	ID          string   `json:"id"`
	Roles       []string `json:"roles"`
	Description string   `json:"description"`
}

type FilterStaticAccessTokenData struct {
	ID *[]string `json:"id,omitempty"`
}

// Results

type DeviceResult struct {
	ID      string        `json:"id"`
	Session SessionResult `json:"session"`
}
type SessionResult struct {
	IssuedAt       string `json:"issuedAt"`
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

type CreateRoleResult struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	SystemFlag  bool      `json:"systemFlag"`
	ServiceFlag bool      `json:"serviceFlag"`
	Created     time.Time `json:"created"`
	Updated     time.Time `json:"updated"`
}

type FilterRolesResult struct {
	ID          string                      `json:"id"`
	Name        string                      `json:"name"`
	Description string                      `json:"description"`
	SystemFlag  bool                        `json:"systemFlag"`
	ServiceFlag bool                        `json:"serviceFlag"`
	Created     time.Time                   `json:"created"`
	Updated     time.Time                   `json:"updated"`
	HttpRules   []FilterRolesHttpRuleResult `json:"httpRules"`
}
type FilterRolesHttpRuleResult struct {
	ID      uint      `json:"id"`
	RoleID  string    `json:"roleId"`
	Path    string    `json:"path"`
	Methods []string  `json:"methods"`
	Mfa     bool      `json:"mfa"`
	Created time.Time `json:"created"`
	Updated time.Time `json:"updated"`
}

type CreateHttpRuleResult struct {
	ID      uint      `json:"id"`
	RoleID  string    `json:"roleId"`
	Path    string    `json:"path"`
	Methods []string  `json:"methods"`
	Mfa     bool      `json:"mfa"`
	Created time.Time `json:"created"`
	Updated time.Time `json:"updated"`
}

type FilterHttpRulesResult struct {
	ID      uint      `json:"id"`
	RoleID  string    `json:"roleId"`
	Path    string    `json:"path"`
	Methods []string  `json:"methods"`
	Mfa     bool      `json:"mfa"`
	Created time.Time `json:"created"`
	Updated time.Time `json:"updated"`
}

type AuthResult struct {
	Access    string `json:"access"`
	Refresh   string `json:"refresh"`
	Mfa       bool   `json:"mfa"`
	NewDevice bool   `json:"newDevice"`
}

type Auth2faResult struct {
	Access  string `json:"access"`
	Refresh string `json:"refresh"`
}

type TokenRenewResult struct {
	Access  string `json:"accessToken"`
	Refresh string `json:"refreshToken"`
	Mfa     bool   `json:"mfaRequired"`
}

type TokenValidateResult struct {
	ID       string   `json:"id"`
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
	ID       string   `json:"id"`
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
	ID          string    `json:"id"`
	Token       string    `json:"token"`
	UserID      uint      `json:"userId"`
	Device      string    `json:"device"`
	Roles       []string  `json:"roles"`
	Description string    `json:"description"`
	Created     time.Time `json:"created"`
}
