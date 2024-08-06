package main

import (
	"fmt"

	"github.com/olezhek28/code-base/cmd/3_testify_unit_test/credit_score"
)

func main() {
	client := credit_score.Client{
		Gender:        "male",
		Age:           30,
		Profession:    "engineer",
		Experience:    7,
		AverageSalary: 60000,
	}

	creditScore, err := credit_score.CalculateCreditScore(client)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Credit Score: %d\n", creditScore)
	}
}
