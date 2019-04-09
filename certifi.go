// Code generated by go generate; DO NOT EDIT
// 2019-01-04 18:09:52.717022 -0800 PST m=+1.021926376
// https://mkcert.org/generate/

package gocertifi

//go:generate go run gen.go

import (
	"crypto/x509"
	"errors"
	"io/ioutil"
)

var ErrReadFailed = errors.New("gocertifi: error when reading certificates file")
var ErrParseFailed = errors.New("gocertifi: error when parsing certificates")

// CACerts builds an X.509 certificate pool containing the Mozilla CA
// Certificate bundle. Returns nil on error along with an appropriate error
// code.
func CACerts() (*x509.CertPool, error) {
	certBytes, err := ioutil.ReadFile("cacert.pem")
	if err != nil {
		return nil, ErrReadFailed
	}
	pool := x509.NewCertPool()
	ok := pool.AppendCertsFromPEM(certBytes)
	if !ok {
		return nil, ErrParseFailed
	}
	return pool, nil
}
