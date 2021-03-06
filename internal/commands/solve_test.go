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
var otherFile = "/tmp/libmatch_test2.json"

func TestSolveAction(t *testing.T) {
	t.Run("SMP", func(t *testing.T) {
		body := `
	  [
	    { "name":"A", "preferences": ["C", "D"] },
	    { "name":"B", "preferences": ["D", "C"] }
	  ]
		`
		writeToFile(testFile, body)

		body = `
	  [
	    { "name":"C", "preferences": ["A", "B"] },
	    { "name":"D", "preferences": ["B", "A"] }
	  ]
		`
		writeToFile(otherFile, body)

		globalSet := flag.NewFlagSet("test", 0)
		globalSet.String("algorithm", "SMP", "doc")
		globalSet.String("format", "csv", "doc")
		globalSet.Var(cli.NewStringSlice(testFile, otherFile), "file", "doc")

		app := cli.NewApp()
		ctx := cli.NewContext(app, globalSet, nil)
		err := solveAction(ctx)

		assert.Nil(t, err)
	})

	t.Run("SRP", func(t *testing.T) {
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
		globalSet.Var(cli.NewStringSlice(testFile), "file", "doc")

		app := cli.NewApp()
		ctx := cli.NewContext(app, globalSet, nil)
		err := solveAction(ctx)

		assert.Nil(t, err)
	})

	t.Run("error creating Config", func(t *testing.T) {
		// Force a config error by specifying an invalid file
		_ = os.Remove(testFile)

		globalSet := flag.NewFlagSet("test", 0)
		globalSet.String("algorithm", "SRP", "doc")
		globalSet.String("format", "csv", "doc")
		globalSet.Var(cli.NewStringSlice(testFile), "file", "doc")

		app := cli.NewApp()
		ctx := cli.NewContext(app, globalSet, nil)
		err := solveAction(ctx)

		if assert.NotNil(t, err) {
			assert.Equal(t,
				fmt.Sprintf("open %v: no such file or directory", testFile), err.Error())
		}
	})

	t.Run("error validating Config", func(t *testing.T) {
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
		globalSet.Var(cli.NewStringSlice(testFile), "file", "doc")

		app := cli.NewApp()
		ctx := cli.NewContext(app, globalSet, nil)
		err := solveAction(ctx)

		if assert.NotNil(t, err) {
			assert.Equal(t, "Unknown `--algorithm` value: INVALID_VALUE", err.Error())
		}
	})

	t.Run("error solving matching", func(t *testing.T) {
		body := `
	  [
	    {"name":"A", "preferences": ["B", "E", "C", "F", "D"] },
	    {"name":"B", "preferences": ["C", "F", "E", "A", "D"] },
	    {"name":"C", "preferences": ["E", "A", "F", "D", "B"] },
	    {"name":"D", "preferences": ["B", "A", "C", "F", "E"] },
	    {"name":"E", "preferences": ["A", "C", "D", "B", "F"] },
	    {"name":"F", "preferences": ["C", "A", "E", "B", "D"] }
	  ]
		`

		writeToFile(testFile, body)

		globalSet := flag.NewFlagSet("test", 0)
		globalSet.String("algorithm", "SRP", "doc")
		globalSet.String("format", "csv", "doc")
		globalSet.Var(cli.NewStringSlice(testFile), "file", "doc")

		app := cli.NewApp()
		ctx := cli.NewContext(app, globalSet, nil)
		err := solveAction(ctx)

		if assert.NotNil(t, err) {
			assert.Equal(t, "No stable solution exists", err.Error())
		}
	})
}

func TestValidateConfig(t *testing.T) {
	t.Run("success", func(t *testing.T) {
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
		globalSet.Var(cli.NewStringSlice(testFile), "file", "doc")

		app := cli.NewApp()
		ctx := cli.NewContext(app, globalSet, nil)
		cfg, err := config.NewConfig(ctx)
		err = validateConfig(*cfg)

		assert.Nil(t, err)
	})

	t.Run("invalid algorithm", func(t *testing.T) {
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
		globalSet.Var(cli.NewStringSlice(testFile), "file", "doc")

		app := cli.NewApp()
		ctx := cli.NewContext(app, globalSet, nil)
		cfg, err := config.NewConfig(ctx)
		err = validateConfig(*cfg)

		if assert.NotNil(t, err) {
			assert.Equal(t, "Unknown `--algorithm` value: INVALID_VALUE", err.Error())
		}
	})

	t.Run("invalid number of tables", func(t *testing.T) {
		testFile2 := "/tmp/libmatch_test_2.json"

		body := `
	  [
	    { "name":"A", "preferences": ["B"] },
	    { "name":"B", "preferences": ["A"] }
	  ]
		`
		writeToFile(testFile, body)
		writeToFile(testFile2, body)

		globalSet := flag.NewFlagSet("test", 0)
		globalSet.String("algorithm", "SRP", "doc")
		globalSet.String("format", "csv", "doc")
		globalSet.Bool("debug", false, "doc")
		globalSet.Var(cli.NewStringSlice(testFile, testFile2), "file", "doc")

		app := cli.NewApp()
		ctx := cli.NewContext(app, globalSet, nil)
		cfg, err := config.NewConfig(ctx)
		err = validateConfig(*cfg)

		if assert.NotNil(t, err) {
			assert.Equal(t, "Expected --file to be specified exactly 1 time(s)", err.Error())
		}
	})

	t.Run("invalid format", func(t *testing.T) {
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
		globalSet.Var(cli.NewStringSlice(testFile), "file", "doc")

		app := cli.NewApp()
		ctx := cli.NewContext(app, globalSet, nil)
		cfg, err := config.NewConfig(ctx)
		err = validateConfig(*cfg)

		if assert.NotNil(t, err) {
			assert.Equal(t, "Unknown `--format` value: INVALID_VALUE", err.Error())
		}
	})
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
