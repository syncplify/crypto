package ssh

import cryptoSha256 "crypto/sha256"

var (
	sha256SimdNew    = cryptoSha256.New
	sha256SimdSum256 = cryptoSha256.Sum256
	// sha256SimdNew    = minioSha256.New
	// sha256SimdSum256 = minioSha256.Sum256
)
