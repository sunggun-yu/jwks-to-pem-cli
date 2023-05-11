package jwks

import (
	"context"
	"fmt"
	"strings"

	"github.com/lestrrat-go/jwx/v2/jwk"
)

func Convert(ctx context.Context, src []byte) ([]string, error) {
	set, err := jwk.Parse(src)
	if err != nil {
		fmt.Println("failed to parse JWKS")
		return nil, err
	}

	pems := []string{}
	for it := set.Keys(ctx); it.Next(ctx); {
		pair := it.Pair()
		key := pair.Value.(jwk.Key)

		publicKey, _ := key.PublicKey()
		encoded, err := jwk.EncodePEM(publicKey)

		if err != nil {
			fmt.Println("failed to encode JWKS to PEM")
			return nil, err
		}
		pems = append(pems, strings.TrimSuffix(string(encoded), "\n"))
	}
	return pems, nil
}
