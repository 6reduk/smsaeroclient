package testUtil

import (
	"fmt"
	"os"
	"strconv"
	"testing"
)

const (
	EnableIntegrationTestEnv = "INTEGRATION_TEST_ENABLE"
	UserEnv                  = "USER"
	PasswordEnv              = "PASSWORD"
	DebugEnv                 = "DEBUG"
)

type TestConfig struct {
	User     string
	Password string
	Debug    bool
}

func GetTestConfigFromEnv() *TestConfig {
	user := os.Getenv(UserEnv)
	password := os.Getenv(PasswordEnv)
	debugStr := os.Getenv(DebugEnv)

	conf := &TestConfig{
		User:     user,
		Password: password,
		Debug:    true,
	}

	if debugStr != "" {
		debug, err := strconv.ParseBool(debugStr)
		if err != nil {
			panic(fmt.Sprintf("can't parse %s to bool: %s", DebugEnv, err))
		}

		conf.Debug = debug
	}

	return conf
}

func (c *TestConfig) IsEmpty() bool {
	return c.User == "" || c.Password == ""
}

func IntegrationTestEnabled() bool {
	enableStr := os.Getenv(EnableIntegrationTestEnv)
	if enableStr == "" {
		return false
	}

	enable, err := strconv.ParseBool(enableStr)
	if err != nil {
		panic(fmt.Sprintf("can't parse %s, reason: %s", EnableIntegrationTestEnv, err))
	}

	return enable
}

func SkipIntegrationTestIfNeed(T *testing.T) {
	if IntegrationTestEnabled() {
		return
	}

	T.Skip("Integration test is disabled")
}
