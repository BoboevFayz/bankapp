package card

import (
	"fmt"
	"bank/pkg/bank/types"
)

func ExampleTotal()  {
	fmt.Println(Total([]types.Card{
		{
			Balance: 1_000_00,
			Active: true,		},
	}))
    fmt.Println(Total([]types.Card{
		{
			Balance: 1_000_00,
			Active: true,		},
			{
				Balance: 2_000_00,
				Active: true,		},
	}))
	fmt.Println(Total([]types.Card{
		{
			Balance: 1_000_00,
			Active: false,		},
	}))
	fmt.Println(Total([]types.Card{
		{
			Balance: -1_000_00,
			Active: true,		},
	}))
	//Output:
	//100000
	//300000
	//0
	//0
}