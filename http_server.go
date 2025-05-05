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
		tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
		tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
	}
}

func NewDefaultCurvePreferences() []tls.CurveID {
	return []tls.CurveID{
		tls.CurveP256,
	}
}

func NewDefaultHTTPHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Connection", "close")

		http.Redirect(w, r, "https://"+r.Host+r.URL.String(), http.StatusMovedPermanently)
	}
}

func NewDefaultHTTPServer() *http.Server {
	return &http.Server{
		Handler:           NewDefaultHTTPHandler(),
		IdleTimeout:       DefaultIdleTimeout,
		ReadHeaderTimeout: DefaultReadHeaderTimeout,
		ReadTimeout:       DefaultReadTimeout,
		TLSConfig:         NewDefaultTLSConfig(),
		WriteTimeout:      DefaultWriteTimeout,
	}
}

func NewDefaultTLSConfig() *tls.Config {
	return &tls.Config{
		CipherSuites:     NewDefaultCipherSuites(),
		CurvePreferences: NewDefaultCurvePreferences(),
		MinVersion:       DefaultTLSMinVersion,
	}
}
