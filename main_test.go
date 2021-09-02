package jwtware

import (
	"testing"

	"github.com/gofiber/fiber/v2"
)

var defaultTokenLookup = "header:" + fiber.HeaderAuthorization

func TestPanicOnMissingConfiguration(t *testing.T) {
	defer func() {
		// Assert
		if err := recover(); err == nil {
			t.Fatalf("Middleware should panic on missing configuration")
		}
	}()

	// Arrange
	config := make([]Config, 0)

	// Act
	initCfg(config)
}

func TestDefaultConfiguration(t *testing.T) {
	defer func() {
		// Assert
		if err := recover(); err != nil {
			t.Fatalf("Middleware should not panic")
		}
	}()

	// Arrange
	config := append(make([]Config, 0), Config{
		SigningKey: "",
	})

	// Act
	cfg := initCfg(config)

	// Assert
	if cfg.SigningMethod != hs256 {
		t.Fatalf("Default signing method should be 'HS256'")
	}
	if cfg.ContextKey != "user" {
		t.Fatalf("Default context key should be 'user'")
	}
	if cfg.Claims == nil {
		t.Fatalf("Default claims should not be 'nil'")
	}

	if cfg.TokenLookup != defaultTokenLookup {
		t.Fatalf("Default token lookup should be '%v'", defaultTokenLookup)
	}
	if cfg.AuthScheme != "Bearer" {
		t.Fatalf("Default auth scheme should be 'Bearer'")
	}
}

func TestExtractorsInitialization(t *testing.T) {
	defer func() {
		// Assert
		if err := recover(); err != nil {
			t.Fatalf("Middleware should not panic")
		}
	}()

	// Arrange
	cfg := Config{
		SigningKey:  "",
		TokenLookup: defaultTokenLookup + ",query:token,param:token,cookie:token,something:something",
	}

	// Act
	extractors := initExtractors(cfg)

	// Assert
	if len(extractors) != 4 {
		t.Fatalf("Extractors should not be created for invalid lookups")
	}
}
