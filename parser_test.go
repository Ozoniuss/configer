package configer

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/spf13/pflag"
)

var example1 = []byte(`numberr: 13
stringg: hello
booll: true
durationn: 30s
`,
)

type Example1 struct {
	Numberr   int
	Stringg   string
	Booll     bool
	Durationn time.Duration
}

func getyamlopts() []ConfigOption {
	return []ConfigOption{
		{FlagName: "0-numberr", Shorthand: "", Value: 4, ConfigKey: "numberr",
			Usage: ""},
		{FlagName: "0-stringg", Shorthand: "", Value: "aaa", ConfigKey: "stringg",
			Usage: ""},
		{FlagName: "0-booll", Shorthand: "", Value: false, ConfigKey: "booll",
			Usage: ""},
		{FlagName: "0-durationn", Shorthand: "", Value: 2 * time.Second, ConfigKey: "durationn",
			Usage: ""},
	}
}

func initFile(t *testing.T, name string, content []byte) (*os.File, string) {
	dir := t.TempDir()

	f, err := os.OpenFile(fmt.Sprintf("%s/test.yml", dir), os.O_CREATE|os.O_RDWR, 0666)

	if err != nil {
		f.Close()
		t.Fatalf("could not create test file: %s", err.Error())
	}

	_, err = f.Write(content)
	if err != nil {
		f.Close()
		t.Fatalf("could not write yaml data: %s", err.Error())
	}

	return f, dir
}

func checkExample1(t *testing.T, expected, actual Example1) {
	if expected.Numberr != actual.Numberr {
		t.Fatalf("invalid number: want %d, got %d", expected.Numberr, actual.Numberr)
	}
	if expected.Booll != actual.Booll {
		t.Fatalf("invalid bool: want %t, got %t", expected.Booll, actual.Booll)
	}
	if expected.Stringg != actual.Stringg {
		t.Fatalf("invalid string: want %s, got %s", expected.Stringg, actual.Stringg)
	}
	if expected.Durationn != actual.Durationn {
		t.Fatalf("invalid duration: want %v, got %v", expected.Durationn, actual.Durationn)
	}
}

func TestReadYaml(t *testing.T) {

	f, dir := initFile(t, "test.yml", example1)
	defer f.Close()

	ex := Example1{}
	err := NewConfig(&ex, getyamlopts(),
		WithConfigName("test"),
		WithConfigType("yml"),
		WithConfigPath(dir),
		WithSupressLogs())

	// Clear comand-line flags at the end of the test, to allow initializing
	// multiple configs. Since pflag.Parse() can only be called once on a
	// flag set, this is necessary because that call is present inside
	// NewConfig()
	defer func() {
		pflag.CommandLine = pflag.NewFlagSet("", pflag.PanicOnError)
	}()

	if err != nil {
		t.Fatalf("call to new config failed: %s", err.Error())
	}

	checkExample1(t, ex, Example1{
		Numberr:   13,
		Stringg:   "hello",
		Booll:     true,
		Durationn: 30 * time.Second,
	})

}

func TestReadYamlDefaultValues(t *testing.T) {
	ex := Example1{}

	err := NewConfig(&ex, getyamlopts(),
		WithConfigName("garbage"),
		WithConfigType("yml"),
		WithSupressLogs())

	if err != nil {
		t.Fatalf("call to new config failed: %s", err.Error())
	}

	checkExample1(t, ex, Example1{
		Numberr:   4,
		Stringg:   "aaa",
		Booll:     false,
		Durationn: 2 * time.Second,
	})
}
