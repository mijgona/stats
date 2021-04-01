package stats

import (
	"github.com/mijgona/bank/v2/pkg/types"
	"fmt"
	
)


func ExampleAvg_positive() {
	fmt.Println(Avg([]types.Payment{
		{
			ID: 1,
			Amount: 10_000,
			Category: "auto",
			Status: "INPROGRESS",
		},
		{
			ID: 2,
			Amount: 1_000,
			Category: "auto",
			Status: "OK",
		},{
			ID: 3,
			Amount: 5_000,
			Category: "mobile",
			Status: "FAIL",
		},
	}))
	//Output:
	//5500
}

func ExampleTotalInCategory_positive() {
	sum:=TotalInCategory([]types.Payment{
		{
			ID: 1,
			Amount: 10_000,
			Category: "auto",
			Status: "",
		},
		{
			ID: 2,
			Amount: 1_000,
			Category: "auto",
			Status: "FAIL",
		},{
			ID: 3,
			Amount: 5_000,
			Category: "mobile",
			Status: "",
		},
	},"auto")
	fmt.Println(sum)
	//Output:
	//10000
}