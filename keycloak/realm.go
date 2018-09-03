package keycloak

import (
	"fmt"
)

type Realm struct {
	Id          string `json:"id"`
	Realm       string `json:"realm"`
	Enabled     bool   `json:"enabled"`
	DisplayName string `json:"displayName"`

	// Login Config
	RegistrationAllowed         bool `json:"registrationAllowed"`
	RegistrationEmailAsUsername bool `json:"registrationEmailAsUsername"`
	EditUsernameAllowed         bool `json:"editUsernameAllowed"`
	ResetPasswordAllowed        bool `json:"resetPasswordAllowed"`
	RememberMe                  bool `json:"rememberMe"`
	VerifyEmail                 bool `json:"verifyEmail"`
	LoginWithEmailAllowed       bool `json:"loginWithEmailAllowed"`
	DuplicateEmailsAllowed      bool `json:"duplicateEmailsAllowed"`
}

func (keycloakClient *KeycloakClient) NewRealm(realm *Realm) error {
	return keycloakClient.post("/realms/", realm)
}

func (keycloakClient *KeycloakClient) GetRealm(id string) (*Realm, error) {
	var realm Realm

	err := keycloakClient.get(fmt.Sprintf("/realms/%s", id), &realm)
	if err != nil {
		return nil, err
	}

	return &realm, nil
}

func (keycloakClient *KeycloakClient) UpdateRealm(realm *Realm) error {
	return keycloakClient.put(fmt.Sprintf("/realms/%s", realm.Id), realm)
}

func (keycloakClient *KeycloakClient) DeleteRealm(id string) error {
	return keycloakClient.delete(fmt.Sprintf("/realms/%s", id))
}

func (realm *Realm) Validate() error {
	if realm.RegistrationAllowed == false && realm.RegistrationEmailAsUsername == true {
		return fmt.Errorf("validation error: RegistrationEmailAsUsername cannot be true if RegistrationAllowed is false")
	}

	if realm.DuplicateEmailsAllowed == true && realm.RegistrationEmailAsUsername == true {
		return fmt.Errorf("validation error: DuplicateEmailsAllowed cannot be true if RegistrationEmailAsUsername is true")
	}

	if realm.DuplicateEmailsAllowed == true && realm.LoginWithEmailAllowed == true {
		return fmt.Errorf("validation error: DuplicateEmailsAllowed cannot be true if LoginWithEmailAllowed is true")
	}

	return nil
}