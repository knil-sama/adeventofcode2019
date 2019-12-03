package main

import (
	"testing"
	. "gopkg.in/check.v1"
)

func Test(t *testing.T) { TestingT(t) }

type MySuite struct{}

var _ = Suite(&MySuite{})

func (s *MySuite) TestHelloWorld(c *C) {
  c.Assert(2, Equals, computeFuel(12))
  c.Assert(2, Equals, computeFuel(14))
  c.Assert(654, Equals, computeFuel(1969))
  c.Assert(33583, Equals, computeFuel(100756))
}
