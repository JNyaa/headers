package headers

import "unicode"

// Version is this package's version
const Version = "2.0.0"

// HTTP headers
const (
	Accept                        = "Accept"
	AcceptCharset                 = "Accept-Charset"
	AcceptEncoding                = "Accept-Encoding"
	AcceptLanguage                = "Accept-Language"
	Authorization                 = "Authorization"
	CacheControl                  = "Cache-Control"
	ContentLength                 = "Content-Length"
	ContentMD5                    = "Content-MD5"
	ContentType                   = "Content-Type"
	IfMatch                       = "If-Match"
	IfModifiedSince               = "If-Modified-Since"
	IfNoneMatch                   = "If-None-Match"
	IfRange                       = "If-Range"
	IfUnmodifiedSince             = "If-Unmodified-Since"
	MaxForwards                   = "Max-Forwards"
	ProxyAuthorization            = "Proxy-Authorization"
	Pragma                        = "Pragma"
	Range                         = "Range"
	Referer                       = "Referer"
	UserAgent                     = "User-Agent"
	TE                            = "TE"
	Via                           = "Via"
	Warning                       = "Warning"
	Cookie                        = "Cookie"
	Origin                        = "Origin"
	AcceptDatetime                = "Accept-Datetime"
	XRequestedWith                = "X-Requested-With"
	AccessControlAllowOrigin      = "Access-Control-Allow-Origin"
	AccessControlAllowMethods     = "Access-Control-Allow-Methods"
	AccessControlAllowHeaders     = "Access-Control-Allow-Headers"
	AccessControlAllowCredentials = "Access-Control-Allow-Credentials"
	AccessControlExposeHeaders    = "Access-Control-Expose-Headers"
	AccessControlMaxAge           = "Access-Control-Max-Age"
	AccessControlRequestMethod    = "Access-Control-Request-Method"
	AccessControlRequestHeaders   = "Access-Control-Request-Headers"
	AcceptPatch                   = "Accept-Patch"
	AcceptRanges                  = "Accept-Ranges"
	Allow                         = "Allow"
	ContentEncoding               = "Content-Encoding"
	ContentLanguage               = "Content-Language"
	ContentLocation               = "Content-Location"
	ContentDisposition            = "Content-Disposition"
	ContentRange                  = "Content-Range"
	ETag                          = "ETag"
	Expires                       = "Expires"
	LastModified                  = "Last-Modified"
	Link                          = "Link"
	Location                      = "Location"
	P3P                           = "P3P"
	ProxyAuthenticate             = "Proxy-Authenticate"
	Refresh                       = "Refresh"
	RetryAfter                    = "Retry-After"
	Server                        = "Server"
	SetCookie                     = "Set-Cookie"
	StrictTransportSecurity       = "Strict-Transport-Security"
	TransferEncoding              = "Transfer-Encoding"
	Upgrade                       = "Upgrade"
	Vary                          = "Vary"
	WWWAuthenticate               = "WWW-Authenticate"

	// Non-Standard
	XFrameOptions          = "X-Frame-Options"
	XXSSProtection         = "X-XSS-Protection"
	ContentSecurityPolicy  = "Content-Security-Policy"
	XContentSecurityPolicy = "X-Content-Security-Policy"
	XWebKitCSP             = "X-WebKit-CSP"
	XContentTypeOptions    = "X-Content-Type-Options"
	XPoweredBy             = "X-Powered-By"
	XUACompatible          = "X-UA-Compatible"
	XForwardedProto        = "X-Forwarded-Proto"
	XHTTPMethodOverride    = "X-HTTP-Method-Override"
	XForwardedFor          = "X-Forwarded-For"
	XRealIP                = "X-Real-IP"
	XCSRFToken             = "X-CSRF-Token"
)

// Normalize formats the input header to the formation of "Xxx-Xxx".
func Normalize(header string) string {
	last := '-'
	normalized := make([]rune, len(header))

	for i, b := range header {
		if last != '-' {
			normalized[i] = unicode.ToLower(b)
		} else {
			normalized[i] = unicode.ToUpper(b)
		}

		last = b
	}

	return string(normalized)
}
