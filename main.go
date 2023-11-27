package main

import (
	"encoding/hex"
	"flag"
	"fmt"

	"github.com/spacemeshos/post/shared"
    "github.com/spacemeshos/go-spacemesh/activation"
)

var (
	postNonce     uint
	postPow       uint64
	indicesString string
)

func parseFlags() {
	flag.UintVar(&postNonce, "postNonce", 0, "post proof nonce")
	flag.Uint64Var(&postPow, "postPow", 0, "post proof pow")
	flag.StringVar(&indicesString, "postIndices", "", "post proof indices"

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
    if err := activation.savePost("./", post); err != nil {
		b.log.With().Warning("failed to save initial post: %w", log.Err(err))
	}
	fmt.Printf("quit.")
}
