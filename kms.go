package kms

type KMS interface {
	Init() error
	GetAwsSecretValue(key string) (string, error)
}
