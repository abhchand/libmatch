package core

func ExamplePrint__format_csv() {
	res := MatchResult{
		Mapping: map[string]string{
			"A": "B",
			"B": "A",
		},
	}

	res.Print("csv")
	// Unordered output:
	// A,B
	// B,A
}

func ExamplePrint__format_json() {
	res := MatchResult{
		Mapping: map[string]string{
			"A": "B",
			"B": "A",
		},
	}

	res.Print("json")
	// Unordered output:
	// {"mapping":{"A":"B","B":"A"}}
}
