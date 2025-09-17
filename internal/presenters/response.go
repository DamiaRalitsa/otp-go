package presenters

type Response struct {
	StatusCode int         `json:"status_code,omitempty"`
	Message    string      `json:"message"`
	Success    bool        `json:"success"`
	Error      string      `json:"error,omitempty"`
	Data       interface{} `json:"data,omitempty"`
	Result     interface{} `json:"result,omitempty"`
}
