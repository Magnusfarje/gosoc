package google

// User - Google user struct
type User struct {
	Iss              string `json:"iss"`
	AtHash           string `json:"at_hash"`
	Aud              string `json:"aud"`
	Sub              string `json:"sub"`
	EmailVerified    string `json:"email_verified"`
	Azp              string `json:"azp"`
	Email            string `json:"email"`
	Iat              string `json:"iat"`
	Exp              string `json:"exp"`
	Name             string `json:"name"`
	Picture          string `json:"picture"`
	GivenName        string `json:"given_name"`
	FamilyName       string `json:"family_name"`
	Locale           string `json:"locale"`
	Alg              string `json:"alg"`
	Kid              string `json:"kid"`
	ErrorDescription string `json:"error_description"`
}
