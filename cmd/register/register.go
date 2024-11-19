package register

import (
	// Import all provider packages here
	"github.com/cosmos/crypto-provider/pkg/factory"
	_ "github.com/cosmos/crypto-provider/pkg/impl/file"
	_ "github.com/cosmos/crypto-provider/pkg/impl/file/cmd"
	// Add other providers as needed
	// _ "github.com/cosmos/crypto-provider/pkg/impl/someprovider"
)

// Init is a dummy function to ensure this package is imported
func Init() {
	_ = factory.GetGlobalFactory()
}
