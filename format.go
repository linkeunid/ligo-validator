package ligovalidator

import "github.com/go-playground/validator/v10"

var fv = validator.New()

// IsEmail reports whether s is a valid email address.
func IsEmail(s string) bool { return fv.Var(s, "email") == nil }

// IsURL reports whether s is a valid URL (any scheme).
func IsURL(s string) bool { return fv.Var(s, "url") == nil }

// IsUUID reports whether s is a valid UUID (any version).
func IsUUID(s string) bool { return fv.Var(s, "uuid") == nil }

// IsIP reports whether s is a valid IP address (v4 or v6).
func IsIP(s string) bool { return fv.Var(s, "ip") == nil }

// IsIPv4 reports whether s is a valid IPv4 address.
func IsIPv4(s string) bool { return fv.Var(s, "ipv4") == nil }

// IsIPv6 reports whether s is a valid IPv6 address.
func IsIPv6(s string) bool { return fv.Var(s, "ipv6") == nil }
