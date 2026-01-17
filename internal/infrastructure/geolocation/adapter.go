package geolocation

import (
	"github.com/yokitheyo/URLShortener/internal/application/ports"
)

// GeoIPServiceImpl implements ports.GeoService interface
var _ ports.GeoService = (*GeoIPServiceImpl)(nil)
