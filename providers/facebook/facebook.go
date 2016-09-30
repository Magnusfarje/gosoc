package facebook

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/magnusfarje/gosoc/models"
)

const (
	profile         string = "https://graph.facebook.com/me?fields=email,first_name,last_name,picture"
	appAcccessToken string = "https://graph.facebook.com/oauth/access_token"
	debugToken      string = "https://graph.facebook.com/debug_token"
)

// Provider - Provider struct
type Provider struct {
	ClientKey        string
	Secret           string
	CheckTokenOrigin bool
}

// New - Create new provider
func New(clientKey, secret string, CheckTokenOrigin bool) *Provider {
	p := &Provider{
		ClientKey:        clientKey,
		Secret:           secret,
		CheckTokenOrigin: CheckTokenOrigin,
	}

	return p
}

// Name - Getter for provider name
func (p *Provider) Name() string {
	return "facebook"
}

// ValidateToken - Validates a token and returns user
func (p *Provider) ValidateToken(token string) (models.User, error) {
	user := models.User{}
	puser, err := p.getUser(token)
	if err != nil {
		return user, err
	}

	if p.CheckTokenOrigin {
		dObj, err := p.validateOrigin(token)
		if err != nil {
			return user, err
		}
		puser.ExpiresAt = strconv.Itoa(dObj.Data.ExpiresAt)
	}
	user = p.mapToUser(puser)
	return user, nil
}

// GetUser - Get user with access token
func (p *Provider) getUser(token string) (User, error) {
	puser := User{}

	// Get facebook profile
	resp, err := http.Get(getProfileURL(token))
	if err != nil {
		if resp != nil {
			defer resp.Body.Close()
		}
		return puser, err
	}
	defer resp.Body.Close()

	// Read body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return puser, err
	}

	// Unmarshal body
	bs := string(body)
	if err := json.Unmarshal([]byte(bs), &puser); err != nil {
		return puser, err
	}

	if puser.Error.Message != "" {
		return puser, fmt.Errorf("Token error: %s", puser.Error.Message)
	}

	return puser, nil

}

// ValidateOrigin - Validates origin of token is your "app"
func (p *Provider) validateOrigin(token string) (Debug, error) {
	dObj := Debug{}

	apptoken, err := p.getApptoken(token)
	if err != nil {
		return dObj, err
	}

	// Facebook debug request
	resp, err := http.Get(getDebugTokenURL(token, apptoken))
	if err != nil {
		if resp != nil {
			defer resp.Body.Close()
		}
		return dObj, err
	}
	defer resp.Body.Close()

	// Read body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return dObj, err
	}
	bs := string(body)

	// Unmarshal body
	if err := json.Unmarshal([]byte(bs), &dObj); err != nil {
		return dObj, err
	}

	if !dObj.Data.IsValid {
		return dObj, errors.New("Could not verify access token origin")
	}

	return dObj, nil

}

func (p *Provider) getApptoken(token string) (string, error) {
	// Facebook get app access token request
	resp, err := http.Get(getApptokenURL(p))
	if err != nil {
		if resp != nil {
			defer resp.Body.Close()
		}
		return "", err
	}
	defer resp.Body.Close()

	// Read body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	bs := strings.Split(string(body), "=")
	if len(bs) < 2 {
		return "", err
	}

	return bs[1], nil
}

func (p *Provider) mapToUser(puser User) models.User {
	t, _ := strconv.ParseInt(puser.ExpiresAt, 10, 64)
	tu := time.Unix(t, 0)
	return models.User{
		Mail:      puser.Email,
		Expire:    tu,
		FirstName: puser.FirstName,
		ID:        puser.ID,
		Picture:   puser.Picture.Data.URL,
		LastName:  puser.LastName,
	}
}

func getApptokenURL(p *Provider) string {
	return appAcccessToken + "?client_id=" + p.ClientKey + "&client_secret=" + p.Secret + "&grant_type=client_credentials"
}

func getDebugTokenURL(token, apptoken string) string {
	return debugToken + "?input_token=" + token + "&access_token=" + apptoken
}

func getProfileURL(token string) string {
	return profile + "&access_token=" + token
}
