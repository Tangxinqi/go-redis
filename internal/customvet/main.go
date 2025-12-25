package main

import (
	"golang.org/x/tools/go/analysis/multichecker"

	"github.com/Tangxinqi/go-redis/internal/customvet/checks/setval"
)

func main() {
	multichecker.Main(
		setval.Analyzer,
	)
}
