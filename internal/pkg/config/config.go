package config

type Endpoint struct {
	ID          uint16 `json:"id"`
	EchoMode    bool   `json:"echo"`
	Name        string `json:"name"`
	URI         string `json:"uri"`
	Method      string `json:"method"`
	ContentType string `json:"contentType"`
	Request     string `json:"request"`
	Response    string `json:"response"`
	// multi-method support params
	IgnoreItem     bool
	EnabledMethods uint16 // OR -ed values
	ResponseMap    map[string]string
}

type Service struct {
	ID        uint16     `json:"id"`
	Port      uint16     `json:"port"`
	Name      string     `json:"name"`
	Endpoints []Endpoint `json:"endpoints"`
}

type Config struct {
	Services []Service `json:"services"`
}
