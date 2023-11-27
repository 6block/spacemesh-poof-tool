package main

import (
	"encoding/hex"
	"flag"
	"fmt"

	"github.com/spacemeshos/go-spacemesh/activation"
	"github.com/spacemeshos/go-spacemesh/common/types"
	"github.com/spacemeshos/post/shared"
)

var (
	postNonce     uint
	postPow       uint64
	indicesString string
	dir           string
)

func parseFlags() {
	flag.UintVar(&postNonce, "postNonce", 0, "post proof nonce")
	flag.Uint64Var(&postPow, "postPow", 0, "post proof pow")
	flag.StringVar(&indicesString, "postIndices", "", "post proof indices")
	flag.StringVar(&dir, "dir", "./", "post.bin store dir")
	flag.Parse()
}

func main() {
	parseFlags()

	indices, err := hex.DecodeString(indicesString)
	if err != nil {
		fmt.Printf("invalid commitmentAtxId: %v", err)
	}
	fmt.Printf("indices: %v", indices)

	post := &shared.Proof{Nonce: uint32(postNonce), Indices: indices, Pow: postPow}
	if err := activation.SavePost(dir, (*types.Post)(post)); err != nil {
		fmt.Printf("failed to save initial post: %v", err)
	}
	fmt.Printf("quit.")
}
