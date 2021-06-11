package ssh

import (
	minioSha256 "github.com/minio/sha256-simd"
)

var (
	sha256SimdNew    = minioSha256.New
	sha256SimdSum256 = minioSha256.Sum256
)
