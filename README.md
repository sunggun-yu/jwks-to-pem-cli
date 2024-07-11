# jwks to pem cli

![GitHub Downloads (all assets, all releases)](https://img.shields.io/github/downloads/sunggun-yu/jwks-to-pem-cli/total)

A Simple CLI for converting jwks to pem that is implementing cli functionality using:

- <https://github.com/spf13/cobra>
- <https://github.com/lestrrat-go/jwx>

Main purpose of this cli is to convert jwks from kubernetes to setup Hashicorp Vault JWK auth: <https://developer.hashicorp.com/vault/docs/auth/jwt/oidc-providers/kubernetes#using-jwt-validation-public-keys>

examples:

```bash
kubectl get --raw /openid/v1/jwks | go run main.go --file -
```

```bash
vault write auth/jwt/config \
   jwt_validation_pubkeys="$(kubectl get --raw /openid/v1/jwks | jwks-to-pem -f -)"
```

## Installation

brew:

```bash
brew install sunggun-yu/tap/jwks-to-pem
```

go install:

```bash
go install github.com/sunggun-yu/jwks-to-pem@<version>
```

docker:

```bash
docker pull ghcr.io/sunggun-yu/jwks-to-pem:latest
```

shell script:

```bash
curl -sfL https://raw.githubusercontent.com/sunggun-yu/jwks-to-pem-cli/main/install.sh | sh
```

it place `jwks-to-pem` in `/tmp` directory. please set `INSTALL_PATH` env var to change directory

```bash
export INSTALL_PATH=/<some-dir>
curl -sfL https://raw.githubusercontent.com/sunggun-yu/jwks-to-pem-cli/main/install.sh | sh
```

or,

```bash
curl -sfL https://raw.githubusercontent.com/sunggun-yu/jwks-to-pem-cli/main/install.sh | INSTALL_PATH=/<some-dir> sh
```
