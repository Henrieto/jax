package jax

type ServerOption func(*HttpServer)

func UseTls(certificate, key string) ServerOption {
	return func(hs *HttpServer) {
		hs.Cert = certificate
		hs.Key = key
		hs.UseTls = true
	}
}
