package google

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/magnusfarje/gosoc/models"
)

const (
	tokenInfo = "https://www.googleapis.com/oauth2/v3/tokeninfo"
)

// Provider - Provider struct
type Provider struct {
	ClientKey        string
	Secret           string
	CheckTokenOrigin bool
}

// New - Create new provider
func New(clientKey, secret string, checkTokenOrigin bool) *Provider {
	p := &Provider{
		ClientKey:        clientKey,
		Secret:           secret,
		CheckTokenOrigin: checkTokenOrigin,
	}

	return p
}

// Name - Getter for provider name
func (p *Provider) Name() string {
	return "google"
}

// ValidateToken - Validates a token and returns user
func (p *Provider) ValidateToken(token string) (models.User, error) {
	puser, err := p.getUser(token)
	if err != nil {
		return models.User{}, err
	}

	if p.CheckTokenOrigin {
		if err := p.validateOrigin(puser); err != nil {
			return models.User{}, err
		}
	}

	user := p.mapToUser(puser)
	return user, nil
}

// GetUser - Get user with google id_token
func (p *Provider) getUser(token string) (User, error) {
	user := User{}
	resp, err := http.Get(getTokenInfoURL(token))
	if err != nil {
		if resp != nil {
			defer resp.Body.Close()
		}
		return user, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return user, err
	}

	bs := string(body)
	puser := User{}

	if err := json.Unmarshal([]byte(bs), &puser); err != nil {
		return user, err
	}

	if puser.ErrorDescription != "" {
		return user, fmt.Errorf("Token error: %s", puser.ErrorDescription)
	}

	return puser, nil
}

// ValidateOrigin - Validates origin of token is your "app"
func (p *Provider) validateOrigin(puser User) error {
	if p.ClientKey == "" || puser.Aud != p.ClientKey {
		return errors.New("Error verifying token origin")
	}

	return nil
}

func (p *Provider) mapToUser(puser User) models.User {
	t, _ := strconv.ParseInt(puser.Exp, 10, 64)
	tu := time.Unix(t, 0)

	return models.User{
		ID:        puser.Sub,
		Mail:      puser.Email,
		Provider:  p.Name(),
		Expire:    tu,
		FirstName: puser.Name,
		LastName:  puser.GivenName,
		Picture:   puser.Picture,
	}
}

func getTokenInfoURL(token string) string {
	return tokenInfo + "?id_token=" + token
}
