package kms

type KMS interface {
	Init() error
	GetSecretValue(key string) (string, error)
}
