package server

import (
    "crypto"
    "crypto/ecdsa"
    "crypto/elliptic"
    "crypto/rand"
    "crypto/x509"
    "github.com/go-acme/lego/certcrypto"
    "github.com/go-acme/lego/certificate"
    "github.com/go-acme/lego/challenge/http01"
    "github.com/go-acme/lego/lego"
    "github.com/go-acme/lego/log"
    "github.com/go-acme/lego/registration"
    "io/ioutil"
    "os"
)

const LETSENCRYPT_DIRECTORY_URL = "https://acme-v02.api.letsencrypt.org/directory"
const CIPHER_STRENGTH = certcrypto.RSA4096

func IssueOrSkipCertificate(server WalleServer) (*certificate.Resource, error) {

    if _, err := os.Stat("./data/" + server.Domain); os.IsNotExist(err) {
        _ = os.MkdirAll("./data/" + server.Domain, os.ModePerm)
    } else {
        log.Infof("%s has already a certificate", server.Domain)
        // Todo do not return nil on success
        return nil, nil
    }

    // Todo load from a file:
    privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)

    if err != nil {
        log.Fatal(err)
    }

    myUser := letsencryptUser{
        Email: server.Ssl.Email,
        key:   privateKey,
    }

    saveIdentityToDisk(*privateKey, server.Domain)

    config := lego.NewConfig(&myUser)

    config.CADirURL = LETSENCRYPT_DIRECTORY_URL
    config.Certificate.KeyType = CIPHER_STRENGTH

    client, err := lego.NewClient(config)

    if err != nil {
        log.Fatal(err)
    }

    // Todo let walle handle all challenges and remove this line:
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
        Domains: []string{server.Domain},
        Bundle:  true,
    }

    certificates, err := client.Certificate.Obtain(request)

    if err != nil {
        log.Fatal(err)
    }

    saveCertificatesToDisk(certificates, server.Domain)

    return certificates, nil

}

func RenewCertificate(server WalleServer) {

}

type letsencryptUser struct {
    Email string
    Registration *registration.Resource
    key crypto.PrivateKey
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


func saveIdentityToDisk(key ecdsa.PrivateKey, name string) {
    marshalledID, err := x509.MarshalECPrivateKey(&key)

    if err != nil {
        log.Fatal(err)
    }

    writeError := ioutil.WriteFile("./data/" + name + "/id_ecdsa.pem", marshalledID, 0644)

    if writeError != nil {
        log.Fatal(writeError)
    }
}

func saveCertificatesToDisk(resource *certificate.Resource, name string) {

    writeError := ioutil.WriteFile("./data/" + name + "/cert.pem", resource.Certificate, 0644)

    if writeError != nil {
        log.Fatal(writeError)
    }

    writeError = ioutil.WriteFile("./data/" + name + "/key.pem", resource.PrivateKey, 0644)

    if writeError != nil {
        log.Fatal(writeError)
    }

}
