package ligovalidator

import "testing"

func TestIsEmail(t *testing.T) {
	if !IsEmail("user@example.com") {
		t.Fatal("expected true for valid email")
	}
	if !IsEmail("user+tag@sub.example.co.uk") {
		t.Fatal("expected true for valid email with tag and subdomain")
	}
	if IsEmail("notanemail") {
		t.Fatal("expected false for missing @")
	}
	if IsEmail("@nodomain") {
		t.Fatal("expected false for missing local part")
	}
	if IsEmail("") {
		t.Fatal("expected false for empty string")
	}
}

func TestIsURL(t *testing.T) {
	if !IsURL("https://example.com") {
		t.Fatal("expected true for https URL")
	}
	if !IsURL("http://example.com/path?q=1") {
		t.Fatal("expected true for http URL with path and query")
	}
	if IsURL("not a url") {
		t.Fatal("expected false for plain string")
	}
	if IsURL("") {
		t.Fatal("expected false for empty string")
	}
}

func TestIsUUID(t *testing.T) {
	if !IsUUID("550e8400-e29b-41d4-a716-446655440000") {
		t.Fatal("expected true for valid UUID v1")
	}
	if !IsUUID("f47ac10b-58cc-4372-a567-0e02b2c3d479") {
		t.Fatal("expected true for valid UUID v4")
	}
	if IsUUID("not-a-uuid") {
		t.Fatal("expected false for invalid UUID")
	}
	if IsUUID("") {
		t.Fatal("expected false for empty string")
	}
}

func TestIsIP(t *testing.T) {
	if !IsIP("192.168.1.1") {
		t.Fatal("expected true for valid IPv4")
	}
	if !IsIP("2001:db8::1") {
		t.Fatal("expected true for valid IPv6")
	}
	if IsIP("999.999.999.999") {
		t.Fatal("expected false for invalid IP")
	}
	if IsIP("") {
		t.Fatal("expected false for empty string")
	}
}

func TestIsIPv4(t *testing.T) {
	if !IsIPv4("192.168.1.1") {
		t.Fatal("expected true for valid IPv4")
	}
	if IsIPv4("2001:db8::1") {
		t.Fatal("expected false for IPv6")
	}
	if IsIPv4("not-an-ip") {
		t.Fatal("expected false for invalid value")
	}
}

func TestIsIPv6(t *testing.T) {
	if !IsIPv6("2001:db8::1") {
		t.Fatal("expected true for valid IPv6")
	}
	if IsIPv6("192.168.1.1") {
		t.Fatal("expected false for IPv4")
	}
	if IsIPv6("not-an-ip") {
		t.Fatal("expected false for invalid value")
	}
}
