package financeconnector

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"time"
)

func (client LiveClient) httpClientForCredential(credential credentialEnvelope) (*http.Client, error) {
	if (credential.MTLSCertificatePEM == "") != (credential.MTLSPrivateKeyPEM == "") {
		return nil, fmt.Errorf("credential bundle mTLS certificate and private key must be provided together")
	}
	if client.HTTP != nil {
		return client.HTTP, nil
	}
	if credential.MTLSCertificatePEM == "" {
		return &http.Client{Timeout: 15 * time.Second}, nil
	}
	certificate, err := tls.X509KeyPair(
		[]byte(credential.MTLSCertificatePEM), []byte(credential.MTLSPrivateKeyPEM),
	)
	if err != nil {
		return nil, fmt.Errorf("credential bundle mTLS material is invalid")
	}
	transport, ok := http.DefaultTransport.(*http.Transport)
	if !ok {
		return nil, fmt.Errorf("default HTTP transport does not support mTLS")
	}
	transport = transport.Clone()
	tlsConfig := transport.TLSClientConfig
	if tlsConfig == nil {
		tlsConfig = &tls.Config{}
	} else {
		tlsConfig = tlsConfig.Clone()
	}
	if tlsConfig.MinVersion < tls.VersionTLS12 {
		tlsConfig.MinVersion = tls.VersionTLS12
	}
	tlsConfig.Certificates = []tls.Certificate{certificate}
	transport.TLSClientConfig = tlsConfig
	return &http.Client{Transport: transport, Timeout: 15 * time.Second}, nil
}
