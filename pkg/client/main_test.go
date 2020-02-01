package client

import (
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/moutend/go-backlog/internal/testutil"
)

func TestMain(m *testing.M) {
	log.SetPrefix("TEST: ")
	log.SetFlags(0)

	if !testutil.EnableLoggerOutput {
		log.SetOutput(ioutil.Discard)
	}

	// call flag.Parse() here if TestMain uses flags
	os.Exit(m.Run())
}
