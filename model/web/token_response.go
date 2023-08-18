package web

type TokenResponse struct {
	Token         string `json:"token"`
	RefreshTocken string `json:"refresh_token"`
}
