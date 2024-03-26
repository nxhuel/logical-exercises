package calculator_test

import (
	"golang/calculator"
	"testing"

	"github.com/stretchr/testify/suite"
)

type DividerTestSuite struct {
	suite.Suite
	underTest calculator.Divider
}

func TestDividerTestSuite(t *testing.T) {
	suite.Run(t, new(DividerTestSuite))
}

func (s *DividerTestSuite) SetupTest() {
	s.underTest = calculator.NewDivider()
}

func (s *DividerTestSuite) TestDivide_HappyPath() {
	result, err := s.underTest.Divide(10, 2)

	s.Nil(err)
	s.Equal(5.0, result)
}

func (s *DividerTestSuite) TestDivide_ByZero() {
	result, err := s.underTest.Divide(10, 0)

	s.Equal(calculator.ErrorDivideByZero, err)
	s.Equal(0.0, result)
}
