package auth

import "crypto"

type SignerOpts interface {
	crypto.SignerOpts
}

type signer interface {

	Sign(key Key, digest []byte, opts SignerOpts) (signature []byte, err error)

}

type verifier interface {

	Verify(key Key, signature, digest []byte, opts SignerOpts) (valid bool, err error)

}

type KeyGenOpts interface {}

type keyGenerator interface {

	GenerateKey(opts KeyGenOpts) (pri, pub Key, err error)

}