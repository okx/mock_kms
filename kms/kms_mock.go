package kms

import (
	"log"
)

// Global variables, similar to the real library
var (
	secretData = make(map[string]string) // Store mock "secret data"
)

// Init simulates initializing KMS client
func Init() error {
	return nil
}

// GetAwsSecretValue simulates getting a secret value
func GetAwsSecretValue(key string) (string, error) {
	log.Printf("[mock-kms] GetAwsSecretValue called for key=%s", key)
	// If exists, return corresponding value, otherwise return mock
	if val, ok := secretData[key]; ok {
		// Pretend this is already decrypted result (or you can keep it as ciphertext)
		return val, nil
	}
	// If not found, return a fixed string
	return "mock-" + key, nil
}

// GetAwsSecretData returns all currently saved mock secrets
func GetAwsSecretData() map[string]string {
	// Copy or return directly
	dataCopy := make(map[string]string)
	for k, v := range secretData {
		dataCopy[k] = v
	}
	return dataCopy
}
