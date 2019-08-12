package config

// Config > Realm Settings > General
type General struct {
	Name    string `json:"name"`
	Enabled bool   `json:"enabled"`
}

// Config > Realm Settings > Login
type Login struct {
	ForgotPassword bool   `json:"forgotPassword"`
	VerifyEmail    bool   `json:"verifyEmail"`
	RequireSSL     string `json:"requireSSL"`
}

// Config > Realm Settings > Keys
type Keys struct {
	Providers Providers `json:"providers"`
}

// Config > Realm Settings > Keys > Providers
type Providers struct {
	Type        string `json:"type"`
	DisplayName string `json:"displayName"`
	Priority    int    `json:"priority"`
	Algorithm   string `json:"algorithm"`
	KeySize     int    `json:"keySize"`
}

// Config > Realm Settings > Email
type Email struct {
	Host            string `json:"host"`
	Port            int    `json:"port"`
	FromDisplayName string `json:"fromDisplayName"`
	From            string `json:"from"`
	EnableStartTLS  bool   `json:"enableStartTLS"`
}

// Config > Realm Settings > Themes
type Themes struct {
	LoginTheme        string `json:"loginTheme"`
	AccountTheme      string `json:"accountTheme"`
	AdminConsoleTheme string `json:"adminConsoleTheme"`
	EmailTheme        string `json:"emailTheme"`
}

// Config > Clients > ClientSettings
type ClientSettings struct {
	ClientID               string   `json:"clientID"`
	LoginTheme             string   `json:"loginTheme"`
	AccessType             string   `json:"accessType"`
	ServiceAccountsEnabled bool     `json:"serviceAccountsEnabled"`
	AuthorizationEnabled   bool     `json:"authorizationEnabled"`
	ValidRedirectURIs      []string `json:"validRedirectURIs"`
}

// Config > Clients > ClientScopes
type DefaultClientScopes struct {
	DefaultClientScopes []string `json:"defaultClientScopes"`
}

// Config > Clients > Mappers
type Mappers struct {
	Name           string `json:"name"`
	MapperType     string `json:"mapperType"`
	UserAttribute  string `json:"userAttribute"`
	TokenClaimName string `json:"tokenClaimName"`
	ClaimJSONType  string `json:"claimJSONType"`
}

// Config > Clients > Scope
type Scope struct {
	FullScopeAllowed bool               `json:"fullScopeAllowed"`
	ScopeClientRoles []ScopeClientRoles `json:"clientRoles"`
}

// Config > Clients > Scope > ClientRoles
type ScopeClientRoles struct {
	RealmManagement []string `json:"realm-management"`
}

// Config > Clients > ServiceAccountRoles
type ServiceAccountRoles struct {
	ServiceAccountRolesClientRoles []ServiceAccountRolesClientRoles `json:"clientRoles"`
}

// Config > Clients > ServiceAccountRoles > ClientRoles
type ServiceAccountRolesClientRoles struct {
	RealmManagement []string `json:"realm-management"`
}

// Config > Clients > SettingsFromSecondClient
// type SettingsFromSecondClient struct {
// 	ClientID               string `json:"clientID"`
// 	AccessType             string `json:"accessType"`
// 	StandardFlowEnabled    bool   `json:"standardFlowEnabled"`
// 	ServiceAccountsEnabled bool   `json:"serviceAccountsEnabled"`
// 	AuthorizationEnabled   bool   `json:"authorizationEnabled"`
// }

// Config > ClientScopes > Settings
type ClientScopesSettings struct {
	Name              string `json:"name"`
	Description       string `json:"description"`
	ConsentScreenText string `json:"consentScreenText"`
}

// Config > ClientScopes > Mappers
type ClientScopesMappers struct {
	Name                   string `json:"name"`
	MapperType             string `json:"mapperType"`
	IncludedClientAudience string `json:"includedClientAudience"`
}

// Config > RealmSettings
type RealmSettings struct {
	General General `json:"general"`
	Login   Login   `json:"login"`
	Keys    Keys    `json:"keys"`
	Email   Email   `json:"email"`
	Themes  Themes  `json:"themes"`
}

// Config > Clients
type Client struct {
	ClientSettings      ClientSettings      `json:"settings,omitempty"`
	DefaultClientScopes DefaultClientScopes `json:"clientScopes"`
	Mappers             []Mappers           `json:"mappers,omitempty"`
	Scope               Scope               `json:"scope"`
	ServiceAccountRoles ServiceAccountRoles `json:"serviceAccountRoles"`
}

// Config > ClientScopes
type ClientScopes struct {
	ClientScopesSettings ClientScopesSettings  `json:"settings"`
	ClientScopesMappers  []ClientScopesMappers `json:"mappers"`
}
type Config struct {
	RealmSettings RealmSettings  `json:"realmSettings"`
	Clients       []Client       `json:"clients"`
	ClientScopes  []ClientScopes `json:"clientScopes"`
}

// Generated:
// type Config struct {
// 	RealmSettings struct {
// 		General struct {
// 			Name    string `json:"name"`
// 			Enabled bool   `json:"enabled"`
// 		} `json:"general"`
// 		Login struct {
// 			ForgotPassword bool   `json:"forgotPassword"`
// 			VerifyEmail    bool   `json:"verifyEmail"`
// 			RequireSSL     string `json:"requireSSL"`
// 		} `json:"login"`
// 		Keys struct {
// 			Providers struct {
// 				Type        string `json:"type"`
// 				DisplayName string `json:"displayName"`
// 				Priority    int    `json:"priority"`
// 				Algorithm   string `json:"algorithm"`
// 				KeySize     int    `json:"keySize"`
// 			} `json:"providers"`
// 		} `json:"keys"`
// 		Email struct {
// 			Host            string `json:"host"`
// 			Port            int    `json:"port"`
// 			FromDisplayName string `json:"fromDisplayName"`
// 			From            string `json:"from"`
// 			EnableStartTLS  bool   `json:"enableStartTLS"`
// 		} `json:"email"`
// 		Themes struct {
// 			LoginTheme        string `json:"loginTheme"`
// 			AccountTheme      string `json:"accountTheme"`
// 			AdminConsoleTheme string `json:"adminConsoleTheme"`
// 			EmailTheme        string `json:"emailTheme"`
// 		} `json:"themes"`
// 	} `json:"realmSettings"`
// 	Clients []struct {
// 		Settings struct {
// 			ClientID               string   `json:"clientID"`
// 			LoginTheme             string   `json:"loginTheme"`
// 			AccessType             string   `json:"accessType"`
// 			ServiceAccountsEnabled bool     `json:"serviceAccountsEnabled"`
// 			AuthorizationEnabled   bool     `json:"authorizationEnabled"`
// 			ValidRedirectURIs      []string `json:"validRedirectURIs"`
// 		} `json:"settings"`
// 		ClientScopes struct {
// 			DefaultClientScopes []string `json:"defaultClientScopes "`
// 		} `json:"clientScopes"`
// 		Mappers []struct {
// 			Name           string `json:"name"`
// 			MapperType     string `json:"mapperType"`
// 			UserAttribute  string `json:"userAttribute"`
// 			TokenClaimName string `json:"tokenClaimName"`
// 			ClaimJSONType  string `json:"claimJSONType"`
// 		} `json:"mappers,omitempty"`
// 		Scope struct {
// 			FullScopeAllowed bool `json:"fullScopeAllowed"`
// 			ClientRoles      []struct {
// 				RealmManagement []string `json:"realm-management"`
// 			} `json:"clientRoles"`
// 		} `json:"scope"`
// 		ServiceAccountRoles struct {
// 			ClientRoles []struct {
// 				RealmManagement []string `json:"realm-management"`
// 			} `json:"clientRoles"`
// 		} `json:"serviceAccountRoles"`
// 	} `json:"clients"`
// 	ClientScopes []struct {
// 		Settings struct {
// 			Name              string `json:"name"`
// 			Description       string `json:"description"`
// 			ConsentScreenText string `json:"consentScreenText"`
// 		} `json:"settings"`
// 		Mappers []struct {
// 			Name                   string `json:"name"`
// 			MapperType             string `json:"mapperType"`
// 			IncludedClientAudience string `json:"includedClientAudience"`
// 		} `json:"mappers"`
// 	} `json:"clientScopes"`
// }
