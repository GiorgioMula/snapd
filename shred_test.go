package osutil_test

import (
	"os"
	"path/filepath"

	"github.com/snapcore/snapd/osutil"
	. "gopkg.in/check.v1"
)

type ShredTestSuite struct {
	testFileName string
	testFile     *os.File
}

var _ = Suite(&ShredTestSuite{})

func (s *ShredTestSuite) SetUpTest(c *C) {
	s.testFileName = filepath.Join(c.MkDir(), "randomfile")
	s.testFile, _ = os.Create(s.testFileName)
}

func (s *ShredTestSuite) TestShredNotExistFile(c *C) {
	s.testFile.Close()
	os.Remove(s.testFileName)

	err := osutil.Shred(s.testFileName)
	c.Assert(err, NotNil)
}

func (s *ShredTestSuite) TestShredSmallSize(c *C) {
	s.testFile.WriteString("some data")
	s.testFile.Close()

	err := osutil.Shred(s.testFileName)
	c.Assert(err, Equals, nil)
	_, err = os.Stat(s.testFileName)
	c.Assert(err, NotNil)
}

func (s *ShredTestSuite) TestShredBigSize(c *C) {
	nullData := [4096]byte{}
	s.testFile.Write(nullData[:])
	s.testFile.Close()

	err := osutil.Shred(s.testFileName)
	c.Assert(err, Equals, nil)
	_, err = os.Stat(s.testFileName)
	c.Assert(err, NotNil)
}
