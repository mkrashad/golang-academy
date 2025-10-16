package certificates

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"math/big"
	"os"
	"path/filepath"
	"time"
)

func GenerateCA() error {
	certsDir := os.Getenv("CERTS_DIR")
	if certsDir == "" {
		return fmt.Errorf("CERTS_DIR env variable not set")
	}

	priv, _ := rsa.GenerateKey(rand.Reader, 2048)
	template := &x509.Certificate{
		SerialNumber: big.NewInt(time.Now().UnixNano()),
		Subject: pkix.Name{
			Organization: []string{"TTPSC"},
			Country:      []string{"PL"},
		},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().AddDate(10, 0, 0),
		IsCA:                  true,
		KeyUsage:              x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
		BasicConstraintsValid: true,
	}

	certDER, _ := x509.CreateCertificate(rand.Reader, template, template, &priv.PublicKey, priv)

	// Save to CERTS_DIR
	caCertPath := filepath.Join(certsDir, "ca.crt")
	caKeyPath := filepath.Join(certsDir, "ca.key")
	os.WriteFile(caCertPath, pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: certDER}), 0644)
	os.WriteFile(caKeyPath, pem.EncodeToMemory(&pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(priv)}), 0600)

	return nil
}

func GenerateMovieCert(movieTitle string) error {
	certsDir := os.Getenv("CERTS_DIR")
	caCertPath := os.Getenv("CA_CERT_PATH")
	caKeyPath := os.Getenv("CA_KEY_PATH")

	caCertPEM, _ := os.ReadFile(caCertPath)
	caKeyPEM, _ := os.ReadFile(caKeyPath)
	block, _ := pem.Decode(caCertPEM)
	caCert, _ := x509.ParseCertificate(block.Bytes)
	block, _ = pem.Decode(caKeyPEM)
	caKey, _ := x509.ParsePKCS1PrivateKey(block.Bytes)

	priv, _ := rsa.GenerateKey(rand.Reader, 2048)
	template := &x509.Certificate{
		SerialNumber: big.NewInt(time.Now().UnixNano()),
		Subject:      pkix.Name{CommonName: movieTitle},
		NotBefore:    time.Now(),
		NotAfter:     time.Now().AddDate(1, 0, 0),
		KeyUsage:     x509.KeyUsageDigitalSignature,
	}

	certDER, _ := x509.CreateCertificate(rand.Reader, template, caCert, &priv.PublicKey, caKey)

	os.WriteFile(filepath.Join(certsDir, "movie-"+movieTitle+".crt"), pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: certDER}), 0644)
	os.WriteFile(filepath.Join(certsDir, "movie-"+movieTitle+".key"), pem.EncodeToMemory(&pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(priv)}), 0600)

	return nil
}

func GenerateCharacterCert(characterName, movieTitle string) error {
	certsDir := os.Getenv("CERTS_DIR")

	movieCertPath := filepath.Join(certsDir, "movie-"+movieTitle+".crt")
	movieKeyPath := filepath.Join(certsDir, "movie-"+movieTitle+".key")
	movieCertPEM, _ := os.ReadFile(movieCertPath)
	movieKeyPEM, _ := os.ReadFile(movieKeyPath)
	block, _ := pem.Decode(movieCertPEM)
	movieCert, _ := x509.ParseCertificate(block.Bytes)
	block, _ = pem.Decode(movieKeyPEM)
	movieKey, _ := x509.ParsePKCS1PrivateKey(block.Bytes)

	priv, _ := rsa.GenerateKey(rand.Reader, 2048)
	template := &x509.Certificate{
		SerialNumber: big.NewInt(time.Now().UnixNano()),
		Subject:      pkix.Name{CommonName: characterName},
		NotBefore:    time.Now(),
		NotAfter:     time.Now().AddDate(1, 0, 0),
		KeyUsage:     x509.KeyUsageDigitalSignature,
	}

	certDER, _ := x509.CreateCertificate(rand.Reader, template, movieCert, &priv.PublicKey, movieKey)

	os.WriteFile(filepath.Join(certsDir, "char-"+characterName+".crt"), pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: certDER}), 0644)
	os.WriteFile(filepath.Join(certsDir, "char-"+characterName+".key"), pem.EncodeToMemory(&pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(priv)}), 0600)

	return nil
}
