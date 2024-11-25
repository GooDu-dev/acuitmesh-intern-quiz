package middlewares

type HeaderRequest struct {
	ContentType   string `json:"Content-Type"`
	ContentCode   string `json:"X-Content-Code"`   // public key for decrypt
	ClientVersion string `json:"X-Client-Version"` // web version
	AccessCtrl    string `json:"X-Access-Ctrl"`    // user auth token
	SourceCtrl    string `json:"X-Source-Ctrl"`    // api token
}

type UserHeaderRequest struct {
	HeaderRequest
	DevIden string `json:"X-Dev-Iden"` // devfle identity (user-id)
	DevEm   string `json:"X-Dev-Em"`   // devfle email (user email)
}
