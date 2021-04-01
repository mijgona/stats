package stats

import (
	"github.com/mijgona/bank/pkg/types"
	"fmt"
	
)


func ExampleAvg_positive() {
	fmt.Println(Avg([]types.Payment{
		{
			ID: 		1,
			Amount:   	10_000,
			Category: 	"auto",
		},
		{
			ID:       2,
			Amount:   1_000,
			Category: "auto",
		},{
			ID:       3,
			Amount:   5_000,
			Category: "mobile",
		},
	}))
	//Output:
	//5333
}

func ExampleTotalInCategory_positive() {
	sum:=TotalInCategory([]types.Payment{
		{
			ID: 		1,
			Amount:   	10_000,
			Category: 	"auto",
		},
		{
			ID:       2,
			Amount:   1_000,
			Category: "auto",
		},{
			ID:       3,
			Amount:   5_000,
			Category: "mobile",
		},
	},"auto")
	fmt.Println(sum)
	//Output:
	//11000
}