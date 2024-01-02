package main

import (
	"fmt"

	"github.com/sharan-aithal/filesize"
)

func main() {
	var fz = filesize.FileSize{
		Size: 120,
		Unit: filesize.KB,
	}

	fmt.Println(fz.ToBytes())
}
