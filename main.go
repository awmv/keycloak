package main

import (
	"encoding/json"
	"io/ioutil"
	"keycloak/pkg/config"
	"log"
	"os"
	"strconv"

	"github.com/Nerzal/gocloak"
)

func main() {
	data, err := ioutil.ReadFile("./credentials.json")
	if err != nil {
		log.Fatalln("Failed to load credentials.json: ", err)
	}

	type Credentials struct {
		URL      string
		USERNAME string
		PASSWORD string
		BRANCH   string
	}
	var obj Credentials

	if err = json.Unmarshal(data, &obj); err != nil {
		log.Fatalln("Failed to unmarshal Json file: ", err)
	}

	client := gocloak.NewClient(obj.URL)
	configuration, err := GetConfig("./config.json")
	if err != nil {
		return
	}
	accessToken, err := Login(obj.USERNAME, obj.PASSWORD, obj.BRANCH, client)
	if err != nil {
		return
	}
	DeleteRealm(configuration.RealmSettings.General.Name, accessToken, client)

	err = CreateRealm(&configuration.RealmSettings, accessToken, client)
	if err != nil {
		return
	}
	err = CreateClients(accessToken, configuration.Clients, configuration.RealmSettings.General.Name, client)
	if err != nil {
		return
	}

	err = CreateClientScope(configuration.ClientScopes, configuration.RealmSettings.General.Name, accessToken, client)
	if err != nil {
		return
	}
}

// CreateClientScope creates instances of ClientScope
func CreateClientScope(scopes []config.ClientScopes, name string, accessToken string, client gocloak.GoCloak) error {
	for _, scope := range scopes {
		newScope := gocloak.ClientScope{
			// ID:   scope.ClientScopesSettings.Name,
			// Name: scope.ClientScopesSettings.Name,
			// Description: scope.ClientScopesSettings.Description,
		}
		err := client.CreateClientScope(accessToken, name, newScope)
		if err != nil {
			log.Println("Failed to create clientScope: ", scope.ClientScopesSettings.Name)
			return err
		}
	}
	return nil
}

// CreateClients loops over client instances
func CreateClients(accessToken string, clients []config.Client, name string, client gocloak.GoCloak) error {
	var newClient gocloak.Client
	for _, configClient := range clients {
		newClient = gocloak.Client{
			ID: configClient.ClientSettings.ClientID,
			// Access: map[string]interface{}{
			// 	"":) "",
			// 	// configClient.ClientSettings.AccessType,
			// },

			ServiceAccountsEnabled:       configClient.ClientSettings.ServiceAccountsEnabled,
			AuthorizationServicesEnabled: configClient.ClientSettings.AuthorizationEnabled,
			// LoginTheme:                   configClient.ClientSettings.LoginTheme
		}
		// if configClient.Mappers != nil {
		// 	newClient.ProtocolMappers = getProtocolMapperRepresentation(configClient)
		// }
		// if configClient.DefaultClientScopes != nil {
		newClient.DefaultClientScopes = configClient.DefaultClientScopes.DefaultClientScopes
		// }
		// CreateClient(token.AccessToken, conf.RealmSettings.General.Name, newClient)
		err := client.CreateClient(accessToken, name, newClient)
		if err != nil {
			newClient := gocloak.Client{ID: configClient.ClientSettings.ClientID}
			log.Println("Failed to create client: ", newClient.Name, err)
			return err
		}
	}
	return nil
}

// GetClients gets the clients in the realm
func GetClients(token string, realm string, client gocloak.GoCloak, clientParams gocloak.GetClientsParams) ([]*gocloak.Client, error) {
	clients, err := client.GetClients(token, realm, clientParams)
	if err != nil {
		log.Println("Failed to load clients: ", err)
	}
	return clients, err
}

// GetProtocolMapperRepresentation returns an Array of ProtocolMapperRepresentation
func GetProtocolMapperRepresentation(configClient config.Client) []gocloak.ProtocolMapperRepresentation {
	representations := []gocloak.ProtocolMapperRepresentation{}

	for _, clientMapper := range configClient.Mappers {
		singleMapper := gocloak.ProtocolMapperRepresentation{
			Config: map[string]string{
				"userinfo.token.claim": "true",
				"user.attribute":       clientMapper.UserAttribute,
				"id.token.claim":       "true",
				"access.token.claim":   "true",
				"claim.name":           clientMapper.Name,
				"jsonType.label":       "String",
			},
			Name:           clientMapper.MapperType,
			Protocol:       "openid-connect",
			ProtocolMapper: "oidc-usermodel-attribute-mapper",
		}
		representations = append(representations, singleMapper)
	}
	return representations
}

// GetConfig deserializes a JSON file
func GetConfig(path string) (*config.Config, error) {
	var conf config.Config
	jsonFile, err := os.Open(path)
	if err != nil {
		log.Println("Failed to read config file: ", err)
		return nil, err
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)

	err = json.Unmarshal(byteValue, &conf)
	if err != nil {
		log.Println("Failed to deserialize config: ", err)
		return nil, err
	}
	return &conf, err
}

// Login logs into the client using the provided credentials
func Login(username string, password string, realm string, client gocloak.GoCloak) (string, error) {
	token, err := client.LoginAdmin(username, password, realm)
	if err != nil {
		log.Println("Failed to login: ", err)
		return "", err
	}
	return token.AccessToken, err
}

// DeleteRealm deletes a realm
func DeleteRealm(realm string, accessToken string, client gocloak.GoCloak) {
	if RealmExists(realm, accessToken, client) {
		err := client.DeleteRealm(accessToken, realm)
		if err != nil {
			log.Println("Failed to delete realm: ", err)
		}
	}
}

// RealmExists determines whether or not the requested realm exists
func RealmExists(realm string, accessToken string, client gocloak.GoCloak) bool {
	res, err := client.GetRealm(accessToken, realm)
	if err != nil {
		log.Println("Failed to load realm: ", err)
	}
	if res != nil {
		return true
	}
	return false
}

// CreateRealm creates a realm
func CreateRealm(realmSettings *config.RealmSettings, accessToken string, client gocloak.GoCloak) error {
	realmRepresentation := gocloak.RealmRepresentation{
		Realm:                realmSettings.General.Name,
		DisplayName:          realmSettings.General.Name,
		Enabled:              realmSettings.General.Enabled,
		ResetPasswordAllowed: realmSettings.Login.ForgotPassword,
		VerifyEmail:          realmSettings.Login.VerifyEmail,
		// SslRequired:          realmSettings.Login.RequireSSL,     //500		// email
		SMTPServer: map[string]string{
			"startls":         strconv.FormatBool(realmSettings.Email.EnableStartTLS), // Neyyyy
			"auth":            "",
			"port":            strconv.Itoa(realmSettings.Email.Port),
			"host":            realmSettings.Email.Host,
			"from":            realmSettings.Email.From,
			"fromDisplayName": realmSettings.Email.FromDisplayName,
			"ssl":             "",
		},
		LoginTheme:   realmSettings.Themes.LoginTheme,
		AccountTheme: realmSettings.Themes.AccountTheme,
		AdminTheme:   realmSettings.Themes.AdminConsoleTheme,
		EmailTheme:   realmSettings.Themes.EmailTheme,
	}
	err := client.CreateRealm(accessToken, realmRepresentation)
	if err != nil {
		log.Println("Failed to create realm: ", err)
		return err
	}
	return nil
}
