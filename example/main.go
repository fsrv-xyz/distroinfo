package main

import (
	"fmt"

	"github.com/fsrv-xyz/distroinfo"
)

func main() {
	info, infoRetreiveError := distroinfo.GetDistroInfoByVersion("debian", "11")
	if infoRetreiveError != nil {
		panic(infoRetreiveError)
	}

	fmt.Printf("%#v\n", info)
	// Output: &distroinfo.DistroInfo{Version:"11", Codename:"Bullseye", Series:"bullseye", Created:"2019-07-06", Release:"2021-08-14", Eol:"2024-08-14", EolLts:"", EolElts:""}
}
