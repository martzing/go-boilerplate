package middleware

import "github.com/gofiber/fiber/v2"

func Helmet(c *fiber.Ctx) error {
	c.Set(fiber.HeaderXXSSProtection, "0")
	c.Set(fiber.HeaderContentSecurityPolicy, "default-src 'self'")
	c.Set(fiber.HeaderConnection, "keep-alive")
	c.Set(fiber.HeaderContentSecurityPolicy, "default-src 'self'")
	c.Set("Cross-Origin-Embedder-Policy", "require-corp")
	c.Set("Cross-Origin-opener-Policy", "same-origin")
	c.Set(fiber.HeaderCrossOriginResourcePolicy, "same-origin")
	c.Set(fiber.HeaderXDNSPrefetchControl, "off")
	c.Set(fiber.HeaderXFrameOptions, "SAMEORIGIN")
	c.Set(fiber.HeaderStrictTransportSecurity, "max-age=63072000; includeSubDomains")
	c.Set(fiber.HeaderXDownloadOptions, "noopen")
	c.Set(fiber.HeaderXContentTypeOptions, "nosniff")
	c.Set("Origin-Agent-Cluster", "?1")
	c.Set(fiber.HeaderXPermittedCrossDomainPolicies, "none")
	c.Set(fiber.HeaderReferrerPolicy, "no-referrer")
	c.Set(fiber.HeaderXXSSProtection, "0")
	c.Set(fiber.HeaderCacheControl, "no-store")
	c.Set(fiber.HeaderPragma, "no-cache")
	return c.Next() 
}