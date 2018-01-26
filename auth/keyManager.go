package auth

import (
	"errors"
	"path/filepath"
	"io/ioutil"
	"encoding/hex"
	"os"
	"strings"
	"crypto/rsa"
	"crypto/ecdsa"
)

type keyManager struct {
	path string
}

func (km *keyManager) Init(path string) {

	if len(path) == 0 {
		km.path = "./KeyRepository"
	} else {
		if !strings.HasPrefix(path, "./") {
			km.path = "./" + path
		} else {
			km.path = path
		}
	}

}

func (km *keyManager) Store(keys... Key) (err error) {

	if len(keys) == 0 {
		return errors.New("Input values should not be NIL")
	}

	for _, key := range keys {
		switch k := key.(type) {
		case *rsaPrivateKey:
			err = km.storePrivateKey(k)
		case *rsaPublicKey:
			err = km.storePublicKey(k)
		case *ecdsaPrivateKey:
			err = km.storePrivateKey(k)
		case *ecdsaPublicKey:
			err = km.storePublicKey(k)
		default:
			return errors.New("Unspported Key Type.")
		}
	}

	return nil
}

func (km *keyManager) storePublicKey(key Key) (err error) {

	data, err := PublicKeyToPEM(key)
	if err != nil {
		return
	}

	path, err := km.getFullPath(hex.EncodeToString(key.SKI()), "pub")
	if err != nil {
		return
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		err = ioutil.WriteFile(path, data, 0700)
		if err != nil {
			return err
		}
	}

	return nil
}

func (km *keyManager) storePrivateKey(key Key) (err error) {

	data, err := PrivateKeyToPEM(key)
	if err != nil {
		return
	}

	path, err := km.getFullPath(hex.EncodeToString(key.SKI()), "pri")
	if err != nil {
		return
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		err = ioutil.WriteFile(path, data, 0700)
		if err != nil {
			return err
		}
	}

	return nil
}

func (km *keyManager) LoadKey() (pri, pub Key, err error) {

	if _, err := os.Stat(km.path); os.IsNotExist(err) {
		return nil, nil, errors.New("Keys are not exist")
	}

	files, err := ioutil.ReadDir(km.path)
	if err != nil {
		return nil, nil, errors.New("Failed to read key repository directory")
	}

	for _, file := range files {

		suffix, valid := km.getSuffix(file.Name())
		if valid == true {
			alias := strings.Split(file.Name(), "_")[0]
			switch suffix {
			case "pri":
				key, err := km.loadPrivateKey(alias)
				if err != nil {
					return nil, nil, err
				}

				switch key.(type) {
				case *rsa.PrivateKey:
					pri = &rsaPrivateKey{key.(*rsa.PrivateKey)}
				case *ecdsa.PrivateKey:
					pri = &ecdsaPrivateKey{key.(*ecdsa.PrivateKey)}
				default:
					return nil, nil, errors.New("Failed to load Key")
				}

			case "pub":
				key, err := km.loadPublicKey(alias)
				if err != nil {
					return nil, nil, err
				}

				switch key.(type) {
				case *rsa.PublicKey:
					pub = &rsaPublicKey{key.(*rsa.PublicKey)}
				case *ecdsa.PublicKey:
					pub = &ecdsaPublicKey{key.(*ecdsa.PublicKey)}
				default:
					return nil, nil, errors.New("Failed to load Key")
				}
			}
		}
	}

	if pri == nil || pub == nil {
		return nil, nil, errors.New("Failed to load Key")
	}

	return

}

func (km *keyManager) loadPrivateKey(alias string) (key interface{}, err error) {

	if len(alias) == 0 {
		return nil, errors.New("Input value should not be blank")
	}

	path, err := km.getFullPath(alias, "pri")
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	key, err = PEMToPrivateKey(data)
	if err != nil {
		return nil, err
	}

	return

}

func (km *keyManager) loadPublicKey(alias string) (key interface{}, err error) {

	if len(alias) == 0 {
		return nil, errors.New("Input value should not be blank")
	}

	path, err := km.getFullPath(alias, "pub")
	if err != nil {
		return nil, err
	}
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	key, err = PEMToPublicKey(data)
	if err != nil {
		return nil, err
	}

	return

}

func (km *keyManager) getSuffix(name string) (suffix string, valid bool) {

	if strings.HasSuffix(name, "pri") {
		return "pri", true
	} else if strings.HasSuffix(name, "pub") {
		return "pub", true
	}

	return "", false

}

func (km *keyManager) getFullPath(alias, suffix string) (path string, err error) {
	if _, err := os.Stat(km.path); os.IsNotExist(err) {
		err = os.MkdirAll(km.path, 0755)
		if err != nil {
			return "", err
		}
	}

	return filepath.Join(km.path, alias + "_" + suffix), nil
}


