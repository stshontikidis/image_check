package docker

type authResponse struct {
	Token     string `json:"token"`
	ExpiresIn int    `json:"expires_in"`
	IssuedAt  string `json:"issued_at"`
}

type config struct {
	Digest    string `json:"digest"`
	MediaType string `json:"mediaType"`
	Size      int    `json:"size"`
}

type layer struct {
	Digest    string `json:"digest"`
	MediaType string `json:"mediaType"`
	Size      int    `json:"Size"`
}
type manifest struct {
	Config        config  `json:"config"`
	Layers        []layer `json:"layers"`
	MediaType     string  `json:"mediaType"`
	SchemaVersion int     `json:"schemaVersion"`
}
