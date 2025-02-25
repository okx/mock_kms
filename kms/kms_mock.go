package kms

// Init simulates initializing KMS client
func Init() error {
	return nil
}

// GetAwsSecretValue simulates getting a secret value
func GetAwsSecretValue(key string) (string, error) {
	// If not found, return a fixed string
	return key, nil
}

// GetAwsSecretData returns all currently saved mock secrets
func GetAwsSecretData() map[string]error {
	// Copy or return directly
	return nil
}
