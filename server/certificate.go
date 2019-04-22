package server

import (
    "crypto"
    "crypto/ecdsa"
    "crypto/elliptic"
    "crypto/rand"
    "fmt"
    "github.com/go-acme/lego/certcrypto"
    "github.com/go-acme/lego/certificate"
    "github.com/go-acme/lego/challenge/http01"
    "github.com/go-acme/lego/lego"
    "github.com/go-acme/lego/registration"
    "log"
)

func IssueCertificate(server WalleServer) {

    // Todo load from a file:
    privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
    if err != nil {
        log.Fatal(err)
    }

    myUser := letsencryptUser{
        Email: "you@yours.com",
        key:   privateKey,
    }

    config := lego.NewConfig(&myUser)

    // This CA URL is configured for a local dev instance of Boulder running in Docker in a VM.
    config.CADirURL = "https://acme-staging-v02.api.letsencrypt.org/directory"
    config.Certificate.KeyType = certcrypto.RSA4096

    // A client facilitates communication with the CA server.
    client, err := lego.NewClient(config)
    if err != nil {
        log.Fatal(err)
    }

    err = client.Challenge.SetHTTP01Provider(http01.NewProviderServer("", "80"))
    if err != nil {
        log.Fatal(err)
    }

    reg, err := client.Registration.Register(registration.RegisterOptions{TermsOfServiceAgreed: true})
    if err != nil {
        log.Fatal(err)
    }

    myUser.Registration = reg

    request := certificate.ObtainRequest{
        Domains: []string{"hanft.xyz"},
        Bundle:  true,
    }
    certificates, err := client.Certificate.Obtain(request)
    if err != nil {
        log.Fatal(err)
    }

    // Each certificate comes back with the cert bytes, the bytes of the client's
    // private key, and a certificate URL. SAVE THESE TO DISK.
    fmt.Printf("%#v\n", certificates)
}

func RenewCertificate(server WalleServer) {

}

type letsencryptUser struct {
    Email        string
    Registration *registration.Resource
    key          crypto.PrivateKey
}

func (u *letsencryptUser) GetEmail() string {
    return u.Email
}
func (u letsencryptUser) GetRegistration() *registration.Resource {
    return u.Registration
}
func (u *letsencryptUser) GetPrivateKey() crypto.PrivateKey {
    return u.key
}
