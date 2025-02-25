package kms

import (
	"log"
	"sync"

	"github.com/pkg/errors"
)

// Global variables, similar to the real library
var (
	secretData = make(map[string]string) // Store mock "secret data"
	mu         sync.Mutex                // Thread safety
	inited     bool                      // Whether initialized
)

// Init simulates initializing KMS client
func Init() error {
	mu.Lock()
	defer mu.Unlock()

	if inited {
		log.Println("[mock-kms] KMS already initialized.")
		return nil
	}

	log.Println("[mock-kms] Initializing mock KMS...")

	secretData["db_user"] = "root"     // Example
	secretData["db_pass"] = "fakepass" // Example, add prefix to simulate ciphertext

	inited = true
	log.Println("[mock-kms] KMS initialization completed (mock).")
	return nil
}

// GetAwsSecretValue simulates getting a secret value
func GetAwsSecretValue(key string) (string, error) {
	log.Printf("[mock-kms] GetAwsSecretValue called for key=%s", key)
	if !inited {
		return "", errors.New("[mock-kms] KMS not initialized")
	}
	// If exists, return corresponding value, otherwise return mock
	if val, ok := secretData[key]; ok {
		// Pretend this is already decrypted result (or you can keep it as ciphertext)
		return val, nil
	}
	// If not found, return a fixed string
	return "mock-" + key, nil
}

// Decrypt simulates decryption
func Decrypt(origin string) (string, error) {
	log.Printf("[mock-kms] Decrypt called with origin=%s", origin)
	if !inited {
		return "", errors.New("[mock-kms] KMS not initialized")
	}
	// Return fixed string here to simulate "successful decryption"
	return "decrypted-mock-value", nil
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

func initKmsClient(stsEndpoint string, kmsEndpoint string) error {
	log.Printf("[mock-kms] initKmsClient called with sts=%s, kms=%s (mock)", stsEndpoint, kmsEndpoint)
	// No network requests made
	return nil
}

func getAndDecrypt(secretName string) (map[string]string, error) {
	log.Printf("[mock-kms] getAndDecrypt called with secretName=%s", secretName)
	if !inited {
		return nil, errors.New("[mock-kms] not initialized")
	}
	// Return a simple map, pretending it's decrypted
	m := map[string]string{
		"db_user": "root",
		"db_pass": "fakepass",
	}
	return m, nil
}
