package main

import (
	"bufio"
	"fmt"
	"os"
)

var testFile = "/tmp/libmatch_test.json"

func ExampleMain__solve_success() {
	body := `
  [
    { "name":"A", "preferences": ["B", "C", "D"] },
    { "name":"B", "preferences": ["A", "C", "D"] },
    { "name":"C", "preferences": ["A", "B", "D"] },
    { "name":"D", "preferences": ["A", "B", "C"] }
  ]
  `

	writeToFile(testFile, body)

	os.Args = []string{
		"libmatch", "solve", "-a", "srp", "-o", "csv", "-f", testFile,
	}

	main()

	// Unordered output:
	// A,B
	// B,A
	// C,D
	// D,C
}

func ExampleMain__solve_error() {
	body := `
  [
	  {"name":"A","preferences":["B","E","C","F","D"]},
	  {"name":"B","preferences":["C","F","E","A","D"]},
	  {"name":"C","preferences":["E","A","F","D","B"]},
	  {"name":"D","preferences":["B","A","C","F","E"]},
	  {"name":"E","preferences":["A","C","D","B","F"]},
	  {"name":"F","preferences":["C","A","E","B","D"]}
	]

  `

	writeToFile(testFile, body)

	os.Args = []string{
		"libmatch", "solve", "-a", "srp", "-o", "csv", "-f", testFile,
	}

	main()

	// Output:
	// No stable solution exists
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
