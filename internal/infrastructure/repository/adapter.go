package repository

import (
	"github.com/yokitheyo/URLShortener/internal/application/ports"
)

// PostgresURLRepository implements ports.URLRepository interface
var _ ports.URLRepository = (*PostgresURLRepository)(nil)
