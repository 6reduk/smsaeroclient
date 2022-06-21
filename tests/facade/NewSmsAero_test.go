package facade

import (
	"github.com/6reduk/smsaeroclient/smsaero"
	"github.com/6reduk/smsaeroclient/smsaero/httpApiClient"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"testing"
)

type NewSMSAeroShould struct {
	suite.Suite
}

func TestNewSMSAeroShould(t *testing.T) {
	suite.Run(t, &NewSMSAeroShould{})
}

func (s *NewSMSAeroShould) TestNewSMSAeroShould_WithConfig_ReturnSuccess() {
	facade := smsaero.NewSmsAeroWithConfig("", "", true, httpApiClient.GetDefaultConfig())
	require.NotNil(s.T(), facade)
	require.NotNil(s.T(), facade.Sms())
	require.NotNil(s.T(), facade.Auth())
}

func (s *NewSMSAeroShould) TestNewSMSAeroShould_WithoutConfig_ReturnSuccess() {
	facade := smsaero.NewSmsAero("", "", true)
	require.NotNil(s.T(), facade)
	require.NotNil(s.T(), facade.Sms())
	require.NotNil(s.T(), facade.Auth())
}
