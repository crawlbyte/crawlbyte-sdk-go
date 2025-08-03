package crawlbyte

type Task struct {
	ID                string            `json:"id"`
	Template          string            `json:"template"`
	Status            string            `json:"status"`
	Url               string            `json:"url,omitempty"`
	Input             []string          `json:"input,omitempty"`
	Fields            []string          `json:"fields,omitempty"`
	Method            string            `json:"method,omitempty"`
	Body              string            `json:"body,omitempty"`
	Proxy             string            `json:"proxy,omitempty"`
	CustomHeaders     map[string]string `json:"customHeaders,omitempty"`
	CustomHeaderOrder []string          `json:"customHeaderOrder,omitempty"`
	JSRendering       bool              `json:"jsRendering,omitempty"`
	CustomSelector    string            `json:"customSelector,omitempty"`
	UserAgentPreset   string            `json:"userAgentPreset,omitempty"`
	UserAgentCustom   string            `json:"userAgentCustom,omitempty"`
	DataType          string            `json:"dataType,omitempty"`
	Location          string            `json:"location,omitempty"`
	SortBy            string            `json:"sortBy,omitempty"`
	Result            interface{}       `json:"result,omitempty"`
	StartTime         string            `json:"startTime,omitempty"`
	EndTime           string            `json:"endTime,omitempty"`
	CreatedAt         string            `json:"createdAt"`
	UpdatedAt         string            `json:"updatedAt,omitempty"`
}

type TaskPayload struct {
	Template          string            `json:"template"`
	Url               string            `json:"url,omitempty"`
	Location          string            `json:"location,omitempty"`
	SortBy            string            `json:"sortBy,omitempty"`
	DataType          string            `json:"dataType,omitempty"`
	Method            string            `json:"method,omitempty"`
	JSRendering       bool              `json:"jsRendering,omitempty"`
	CustomSelector    string            `json:"customSelector,omitempty"`
	UserAgentPreset   string            `json:"userAgentPreset,omitempty"`
	UserAgentCustom   string            `json:"userAgentCustom,omitempty"`
	Proxy             string            `json:"proxy,omitempty"`
	CustomHeaderOrder []string          `json:"customHeaderOrder,omitempty"`
	Input             []string          `json:"input,omitempty"`
	CustomHeaders     map[string]string `json:"customHeaders,omitempty"`
	Fields            []string          `json:"fields,omitempty"`
	Body              string            `json:"body,omitempty"`
	Multithread       bool              `json:"multithread,omitempty"`
}

type PollOptions struct {
	IntervalSeconds int
	TimeoutSeconds  int
}
