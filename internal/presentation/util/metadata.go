package util

import (
	"github.com/wb-go/wbf/ginext"
	"github.com/yokitheyo/URLShortener/internal/application/dto"
	"github.com/yokitheyo/URLShortener/internal/infrastructure/geolocation"
	"github.com/yokitheyo/URLShortener/internal/util"
)

// BuildClickMetadata extracts and builds click metadata from HTTP context
func BuildClickMetadata(c *ginext.Context) dto.ClickMetadata {
	return dto.ClickMetadata{
		IP:        c.ClientIP(),
		UserAgent: c.Request.UserAgent(),
		Browser:   geolocation.GetBrowserFromUserAgent(c.Request.UserAgent()),
		Device:    util.DetectDevice(c.Request.UserAgent()),
	}
}
