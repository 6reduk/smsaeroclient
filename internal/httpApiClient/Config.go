package httpApiClient

type HttpClientConfig struct {
	Timeout               int
	KeepAlive             int
	MaxIdleConns          int
	IdleConnTimeout       int
	TLSHandshakeTimeout   int
	ExpectContinueTimeout int
}

func GetDefaultConfig() *HttpClientConfig {
	return &HttpClientConfig{
		Timeout:               10,
		KeepAlive:             10,
		MaxIdleConns:          5,
		IdleConnTimeout:       5,
		TLSHandshakeTimeout:   5,
		ExpectContinueTimeout: 5,
	}
}
