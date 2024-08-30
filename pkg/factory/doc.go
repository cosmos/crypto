// Package factory provides a flexible and extensible factory pattern implementation
// for creating and managing crypto providers.
//
// The main components of this package are:
//
//   - Factory: A singleton struct that manages the registration and creation of
//     crypto providers.
//
// - RegisterFactory: A method to register new crypto provider factories.
//
// - CreateCryptoProvider: A method to create crypto providers based on two key parameters:
//  1. providerType: Specifies the type of crypto provider (e.g., "ledger", "file", "memory").
//  2. source: Represents the data source from which the provider should be built.
//     This could be a file path, a hardware device identifier, or any other
//     source-specific information needed to initialize the provider.
//     This approach allows for flexible creation of providers with different
//     backends and data sources.
//
// - GetRegisteredFactories: A method to retrieve all registered factory types.
//
// This package is designed to be thread-safe and allows for easy extension of
// supported crypto provider types and data sources.
package factory
