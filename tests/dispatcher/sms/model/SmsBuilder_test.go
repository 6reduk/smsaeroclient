package model

import (
	"errors"
	smsDto "github.com/6reduk/smsaeroclient/internal/dispatcher/sms/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"testing"
)

type SmsBuilderShould struct {
	suite.Suite
}

func TestSmsBuilderShould(t *testing.T) {
	suite.Run(t, &SmsBuilderShould{})
}

func (s *SmsBuilderShould) TestSmsBuilder_InvalidPhoneNumber_ReturnError() {
	invalidPhone := "15tdg"
	sms, buildErrors := smsDto.NewSmsMessageBuilder().
		Number(invalidPhone).
		Build()
	require.Nil(s.T(), sms)
	require.NotNil(s.T(), buildErrors)

	assert.Equal(s.T(), 4, len(buildErrors))
	assert.True(s.T(), errors.Is(buildErrors[0], smsDto.ErrInvalidPhone))
}

func (s *SmsBuilderShould) TestSmsBuilder_WithoutParams_ReturnError() {
	sms, buildErrors := smsDto.NewSmsMessageBuilder().Build()
	require.Nil(s.T(), sms)
	require.NotNil(s.T(), buildErrors)

	assert.Equal(s.T(), 3, len(buildErrors))
	assert.Contains(s.T(), buildErrors, smsDto.ErrPhoneNumberIsRequired)
	assert.Contains(s.T(), buildErrors, smsDto.ErrSignIsRequired)
	assert.Contains(s.T(), buildErrors, smsDto.ErrTextIsRequired)
}

func (s *SmsBuilderShould) TestSmsBuilder_ValidParams_ReturnSuccess() {
	phoneNumbers := []string{"+7-515-213", "8765"}
	expectedPhoneNumbers := []string{"7515213", "8765"}
	expectedSign := "testSign"
	expectedText := "testText"

	sms, buildErrors := smsDto.NewSmsMessageBuilder().
		Numbers(phoneNumbers).
		Sign(expectedSign).
		Text(expectedText).
		Build()
	require.Nil(s.T(), buildErrors)
	require.NotNil(s.T(), sms)

	assert.EqualValues(s.T(), expectedPhoneNumbers, sms.Numbers())
	assert.Equal(s.T(), expectedSign, sms.Sign())
	assert.Equal(s.T(), expectedText, sms.Text())
}
