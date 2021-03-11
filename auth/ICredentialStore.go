package auth

/*
Interface for credential stores which are used to store and lookup credentials to authenticate against external services.
*/
type ICredentialStore interface {
	// Stores credential parameters into the store.
	Store(correlationId string, key string, credential *CredentialParams) error
	// Lookups credential parameters by its key.
	Lookup(correlationId string, key string) (*CredentialParams, error)
}
