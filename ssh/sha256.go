package ssh

import cryptoSha256 "crypto/sha256"

var (
	sha256New    = cryptoSha256.New
	sha256Sum256 = cryptoSha256.Sum256
	// sha256New    = minioSha256.New
	// sha256Sum256 = minioSha256.Sum256
)
