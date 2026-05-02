# Format validators

Standalone format checks backed by go-playground/validator. Useful for validating individual values outside of struct tags.

---

## API reference

### IsEmail

```go
func IsEmail(s string) bool
```

Reports whether `s` is a valid email address.

```go
ligovalidator.IsEmail("user@example.com")            // true
ligovalidator.IsEmail("user+tag@sub.example.co.uk")  // true
ligovalidator.IsEmail("notanemail")                  // false
ligovalidator.IsEmail("")                            // false
```

### IsURL

```go
func IsURL(s string) bool
```

Reports whether `s` is a valid URL (any scheme).

```go
ligovalidator.IsURL("https://example.com")           // true
ligovalidator.IsURL("http://example.com/path?q=1")   // true
ligovalidator.IsURL("not a url")                     // false
ligovalidator.IsURL("")                              // false
```

### IsUUID

```go
func IsUUID(s string) bool
```

Reports whether `s` is a valid UUID (any version).

```go
ligovalidator.IsUUID("550e8400-e29b-41d4-a716-446655440000") // true
ligovalidator.IsUUID("f47ac10b-58cc-4372-a567-0e02b2c3d479") // true
ligovalidator.IsUUID("not-a-uuid")                           // false
ligovalidator.IsUUID("")                                     // false
```

### IsIP

```go
func IsIP(s string) bool
```

Reports whether `s` is a valid IP address (v4 or v6).

```go
ligovalidator.IsIP("192.168.1.1") // true
ligovalidator.IsIP("2001:db8::1") // true
ligovalidator.IsIP("999.999.999.999") // false
```

### IsIPv4

```go
func IsIPv4(s string) bool
```

Reports whether `s` is a valid IPv4 address.

```go
ligovalidator.IsIPv4("192.168.1.1") // true
ligovalidator.IsIPv4("2001:db8::1") // false
```

### IsIPv6

```go
func IsIPv6(s string) bool
```

Reports whether `s` is a valid IPv6 address.

```go
ligovalidator.IsIPv6("2001:db8::1") // true
ligovalidator.IsIPv6("192.168.1.1") // false
```
