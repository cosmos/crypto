package hasher

import "github.com/cosmos/crypto-provider/pkg/components"

type FileHash struct{}

func (FileHash) Hash(input []byte, options components.HasherOpts) (output []byte, err error) {
	//TODO implement me
	panic("implement me")
}