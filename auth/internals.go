package auth

import "crypto"

type SignerOpts interface {
	crypto.SignerOpts
}

type Signer interface {

	Sign(key Key, digest []byte, opts SignerOpts) (signature []byte, err error)

}

type Verifier interface {

	Verify(key Key, signature, digest []byte, opts SignerOpts) (valid bool, err error)

}

type KeyGenOpts interface {}

type KeyGenerator interface {

	KeyGenerate(opts KeyGenOpts) (key Key, err error)

}