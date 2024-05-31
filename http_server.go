package http_server

import (
	"crypto/tls"
	"net/http"
	"time"
)

const (
	DefaultIdleTimeout       = time.Minute * 2
	DefaultReadHeaderTimeout = DefaultReadTimeout
	DefaultReadTimeout       = time.Second * 5
	DefaultTLSMinVersion     = tls.VersionTLS13
	DefaultWriteTimeout      = time.Second * 5
)

func NewDefaultCipherSuites() []uint16 {
	return []uint16{
		tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
		tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
		tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305, // Go 1.8 only.
		tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
		tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
		tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305, // Go 1.8 only.
	}
}

// NewDefaultCurvePreferences — returns curves that have assembly implementations.
func NewDefaultCurvePreferences() []tls.CurveID {
	return []tls.CurveID{
		tls.CurveP256,
		tls.X25519, // Go 1.8 only.
	}
}

func NewDefaultHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Connection", "close")

		http.Redirect(w, r, "https://"+r.Host+r.URL.String(), http.StatusMovedPermanently)
	}
}

func NewHTTPServer() *http.Server {
	return &http.Server{
		Handler:           NewDefaultHandler(),
		IdleTimeout:       DefaultIdleTimeout,
		ReadHeaderTimeout: DefaultReadHeaderTimeout,
		ReadTimeout:       DefaultReadTimeout,
		TLSConfig: &tls.Config{
			CipherSuites:     NewDefaultCipherSuites(),
			CurvePreferences: NewDefaultCurvePreferences(),
			MinVersion:       DefaultTLSMinVersion,
		},
		WriteTimeout: DefaultWriteTimeout,
	}
}
