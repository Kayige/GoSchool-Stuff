package main

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type AddTestSuite struct {
	suite.Suite
}

type SubtractTestSuite struct {
	suite.Suite
}

type MultiplyTestSuite struct {
	suite.Suite
}

type DivideTestSuite struct {
	suite.Suite
}

func (suite *SubtractTestSuite) TestSubtract() {
	assert.Equal(suite.T(), 0, Subtract(1, 1))
	assert.Equal(suite.T(), 1, Subtract(2, 1))
	assert.Equal(suite.T(), 0, Subtract(2, -2))
}

func (suite *AddTestSuite) TestAdd() {
	assert.Equal(suite.T(), 3, Add(1, 2))
	assert.Equal(suite.T(), 1, Add(1, 0))
	assert.Equal(suite.T(), 0, Add(2, -2))
}

func (suite *MultiplyTestSuite) TestMultiply() {
	assert.Equal(suite.T(), 2, Multiply(1, 2))
	assert.Equal(suite.T(), 4, Multiply(2, 2))
	assert.Equal(suite.T(), 6, Multiply(3, 2))
}
