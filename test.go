package main

import yahoo "github.com/r-4981/yahoo-client-go"

fnc main(){
	auction := yahoo.NewAuction("dj00aiZpPVVrZXJ4Wm5GcnVlaCZzPWNvbnN1bWVyc2VjcmV0Jng9NTk-")

	resultset, _ := auction.CategoryLeaf(&yahoo.CategoryLeafParam{
		Category: "2084193603",
	})

	for _, v := range resultset.Result.Item {
		fmt.Println(v.Seller.ID)
	}
}