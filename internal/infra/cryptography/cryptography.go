package cryptography

type CryptographyInterface interface {
	Encrypt(value string) (*string, error)
	Decrypt(token string) (*string, error)
}
