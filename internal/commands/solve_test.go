package commands

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"testing"

	"github.com/abhchand/libmatch/internal/config"
	"github.com/stretchr/testify/assert"
	"github.com/urfave/cli/v2"
)

var testFile = "/tmp/libmatch_test.json"

func TestSolveAction(t *testing.T) {
	body := `
  [
    { "name":"A", "preferences": ["B"] },
    { "name":"B", "preferences": ["A"] }
  ]
	`
	writeToFile(testFile, body)

	globalSet := flag.NewFlagSet("test", 0)
	globalSet.String("algorithm", "SRP", "doc")
	globalSet.String("format", "csv", "doc")
	globalSet.String("file", testFile, "doc")

	app := cli.NewApp()
	ctx := cli.NewContext(app, globalSet, nil)
	err := solveAction(ctx)

	assert.Nil(t, err)
}

func TestSolveAction_ErrorCreatingConfig(t *testing.T) {
	// Force a config error by specifying an invalid file
	_ = os.Remove(testFile)

	globalSet := flag.NewFlagSet("test", 0)
	globalSet.String("algorithm", "SRP", "doc")
	globalSet.String("format", "csv", "doc")
	globalSet.String("file", testFile, "doc")

	app := cli.NewApp()
	ctx := cli.NewContext(app, globalSet, nil)
	err := solveAction(ctx)

	if assert.NotNil(t, err) {
		assert.Equal(t,
			fmt.Sprintf("open %v: no such file or directory", testFile), err.Error())
	}
}

func TestSolveAction_ErrorValidatingConfig(t *testing.T) {
	body := `
  [
    { "name":"A", "preferences": ["B"] },
    { "name":"B", "preferences": ["A"] }
  ]
	`
	writeToFile(testFile, body)

	globalSet := flag.NewFlagSet("test", 0)
	// Force a validation error by specifying a bad `algorithm` value
	globalSet.String("algorithm", "INVALID_VALUE", "doc")
	globalSet.String("format", "csv", "doc")
	globalSet.String("file", testFile, "doc")

	app := cli.NewApp()
	ctx := cli.NewContext(app, globalSet, nil)
	err := solveAction(ctx)

	if assert.NotNil(t, err) {
		assert.Equal(t, "Unknown `--algorithm` value: INVALID_VALUE", err.Error())
	}
}

func TestSolveAction_ErrorSolvingMatching(t *testing.T) {
	t.Skip("Skipping test - no way to force `SolveSRP` to return erro yet")
}

func TestValidateConfig(t *testing.T) {
	body := `
  [
    { "name":"A", "preferences": ["B"] },
    { "name":"B", "preferences": ["A"] }
  ]
	`
	writeToFile(testFile, body)

	globalSet := flag.NewFlagSet("test", 0)
	globalSet.String("algorithm", "SRP", "doc")
	globalSet.String("format", "csv", "doc")
	globalSet.Bool("debug", false, "doc")
	globalSet.String("file", testFile, "doc")

	app := cli.NewApp()
	ctx := cli.NewContext(app, globalSet, nil)
	cfg, err := config.NewConfig(ctx)
	err = validateConfig(*cfg)

	assert.Nil(t, err)
}

func TestValidateConfig_InvalidAlgorithm(t *testing.T) {
	body := `
  [
    { "name":"A", "preferences": ["B"] },
    { "name":"B", "preferences": ["A"] }
  ]
	`
	writeToFile(testFile, body)

	globalSet := flag.NewFlagSet("test", 0)
	globalSet.String("algorithm", "INVALID_VALUE", "doc")
	globalSet.String("format", "csv", "doc")
	globalSet.Bool("debug", false, "doc")
	globalSet.String("file", testFile, "doc")

	app := cli.NewApp()
	ctx := cli.NewContext(app, globalSet, nil)
	cfg, err := config.NewConfig(ctx)
	err = validateConfig(*cfg)

	if assert.NotNil(t, err) {
		assert.Equal(t, "Unknown `--algorithm` value: INVALID_VALUE", err.Error())
	}
}

func TestValidateConfig_InvalidFormat(t *testing.T) {
	body := `
  [
    { "name":"A", "preferences": ["B"] },
    { "name":"B", "preferences": ["A"] }
  ]
	`
	writeToFile(testFile, body)

	globalSet := flag.NewFlagSet("test", 0)
	globalSet.String("algorithm", "srp", "doc")
	globalSet.String("format", "INVALID_VALUE", "doc")
	globalSet.Bool("debug", false, "doc")
	globalSet.String("file", testFile, "doc")

	app := cli.NewApp()
	ctx := cli.NewContext(app, globalSet, nil)
	cfg, err := config.NewConfig(ctx)
	err = validateConfig(*cfg)

	if assert.NotNil(t, err) {
		assert.Equal(t, "Unknown `--format` value: INVALID_VALUE", err.Error())
	}
}

func writeToFile(filename, body string) {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println(err)
	}

	writer := bufio.NewWriter(file)

	_, err = writer.WriteString(body)
	if err != nil {
		fmt.Printf("Could not create file: %s\n", err.Error())
	}

	writer.Flush()
}
