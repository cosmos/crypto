package register

import (
	// Import all provider packages here
	"crypto-provider/pkg/factory"
	_ "crypto-provider/pkg/provider/file"
	_ "crypto-provider/pkg/provider/file/cmd"
	// Add other providers as needed
	// _ "crypto-provider/pkg/provider/someprovider"
)

// Init is a dummy function to ensure this package is imported
func Init() {
	_ = factory.GetFactory()
}
