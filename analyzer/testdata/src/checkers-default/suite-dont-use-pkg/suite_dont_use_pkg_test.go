// Code generated by testifylint/internal/testgen. DO NOT EDIT.

package suitedontusepkg

import (
	"testing"

	"github.com/stretchr/testify/assert"
	a "github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	r "github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type SuiteDontUsePkgCheckerSuite struct {
	suite.Suite
}

func TestSuiteDontUsePkgCheckerSuite(t *testing.T) {
	suite.Run(t, new(SuiteDontUsePkgCheckerSuite))
}

func (s *SuiteDontUsePkgCheckerSuite) TestAll() {
	var result any
	assObj, reqObj := s.Assert(), s.Require()

	s.Equal(42, result)
	s.Equalf(42, result, "msg with args %d %s", 42, "42")
	s.Assert().Equal(42, result)
	s.Assert().Equalf(42, result, "msg with args %d %s", 42, "42")
	s.Require().Equal(42, result)
	s.Require().Equalf(42, result, "msg with args %d %s", 42, "42")
	assObj.Equal(42, result)
	assObj.Equalf(42, result, "msg with args %d %s", 42, "42")
	reqObj.Equal(42, result)
	reqObj.Equalf(42, result, "msg with args %d %s", 42, "42")

	assert.Equal(s.T(), 42, result)                                   // want "suite-dont-use-pkg: use s\\.Equal"
	assert.Equalf(s.T(), 42, result, "msg with args %d %s", 42, "42") // want "suite-dont-use-pkg: use s\\.Equalf"
	a.Equal(s.T(), 42, result)                                        // want "suite-dont-use-pkg: use s\\.Equal"
	a.Equalf(s.T(), 42, result, "msg with args %d %s", 42, "42")      // want "suite-dont-use-pkg: use s\\.Equalf"

	require.Equal(s.T(), 42, result)                                   // want "suite-dont-use-pkg: use s\\.Require\\(\\)\\.Equal"
	require.Equalf(s.T(), 42, result, "msg with args %d %s", 42, "42") // want "suite-dont-use-pkg: use s\\.Require\\(\\)\\.Equalf"
	r.Equal(s.T(), 42, result)                                         // want "suite-dont-use-pkg: use s\\.Require\\(\\)\\.Equal"
	r.Equalf(s.T(), 42, result, "msg with args %d %s", 42, "42")       // want "suite-dont-use-pkg: use s\\.Require\\(\\)\\.Equalf"

	s.T().Run("not detected in order to avoid conflict with suite-subtest-run", func(t *testing.T) {
		var result any
		assObj, reqObj := assert.New(t), require.New(t)

		assert.Equal(t, 42, result)
		assert.Equalf(t, 42, result, "msg with args %d %s", 42, "42")
		assObj.Equal(42, result)
		assObj.Equalf(42, result, "msg with args %d %s", 42, "42")
		require.Equal(t, 42, result)
		require.Equalf(t, 42, result, "msg with args %d %s", 42, "42")
		reqObj.Equal(42, result)
		reqObj.Equalf(42, result, "msg with args %d %s", 42, "42")
	})
}

func TestSuiteDontUsePkgChecker_NoSuiteNoProblem(t *testing.T) {
	var result any
	assObj, reqObj := assert.New(t), require.New(t)

	assert.Equal(t, 42, result)
	assert.Equalf(t, 42, result, "msg with args %d %s", 42, "42")
	assObj.Equal(42, result)
	assObj.Equalf(42, result, "msg with args %d %s", 42, "42")
	require.Equal(t, 42, result)
	require.Equalf(t, 42, result, "msg with args %d %s", 42, "42")
	reqObj.Equal(42, result)
	reqObj.Equalf(42, result, "msg with args %d %s", 42, "42")
}
