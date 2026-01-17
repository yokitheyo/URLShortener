package storage

import (
	"github.com/yokitheyo/URLShortener/internal/application/ports"
)

// Ensure RedisCache implements ports.Cache interface
var _ ports.Cache = (*RedisCache)(nil)
