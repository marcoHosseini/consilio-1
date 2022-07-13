package local

import (
	"flag"
	"os"
	"testing"

	_ "github.com/kevinklinger/open_terraform/noninternal/logging"
)

func TestMain(m *testing.M) {
	flag.Parse()
	os.Exit(m.Run())
}
