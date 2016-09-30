package facebook

type User struct {
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Picture   struct {
		Data struct {
			IsSilhouette bool   `json:"is_silhouette"`
			URL          string `json:"url"`
		} `json:"data"`
	} `json:"picture"`
	ID        string `json:"id"`
	ExpiresAt string

	Error struct {
		Message   string `json:"message"`
		Type      string `json:"type"`
		Code      int    `json:"code"`
		FbtraceID string `json:"fbtrace_id"`
	} `json:"error"`
}

type Debug struct {
	Data struct {
		AppID       string   `json:"app_id"`
		Application string   `json:"application"`
		ExpiresAt   int      `json:"expires_at"`
		IsValid     bool     `json:"is_valid"`
		Scopes      []string `json:"scopes"`
		UserID      string   `json:"user_id"`
	} `json:"data"`
}
