package test

type PingRequest struct {
	Input string `json:"message"`
}

type PongResponse struct {
	Output string `json:"message"`
}
