package main

import (
	"flag"
	"fmt"
	"github.com/kitt1987/superblock/pkg/xfs"
)

var (
	blkPath = flag.String("blk", "", "Path to the block device")
)

func main() {
	flag.Parse()

	sb, err := xfs.GetSuperBlock(*blkPath)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("%#v\n", sb)
	fmt.Println("Total size:", uint64(sb.SB_blocksize) * uint64(sb.SB_dblocks))
	fmt.Println("Total size:", uint64(sb.SB_blocksize) * uint64(sb.SB_fdblocks))
}
