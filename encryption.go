package mcwss

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/ecdsa"
	"crypto/sha256"
	"fmt"
)

// MinecraftWSEncryptSubprotocol is the subprotocol used for encrypted websocket connections in Minecraft.
const MinecraftWSEncryptSubprotocol = "com.microsoft.minecraft.wsencrypt"

// encryptionSession is a session unique to a player, that handles encryption between the server and the
// client.
type encryptionSession struct {
	serverPrivateKey *ecdsa.PrivateKey
	clientPublicKey  *ecdsa.PublicKey
	salt             []byte

	sharedSecret   []byte
	secretKeyBytes [32]byte
	cipherBlock    cipher.Block

	encryptIV []byte
	decryptIV []byte

	sendCounter int64
}

// init initialises the encryption session, computing the shared secret and secret key bytes as required to
// initialise the cipher blocks for encryption and decryption.
func (session *encryptionSession) init() error {
	session.computeSharedSecret()
	return session.computeIVs()
}

// computeSharedSecret computes the shared secret required for encryption and decryption.
func (session *encryptionSession) computeSharedSecret() {
	// We only care about the 'x' part of this.
	x, _ := session.clientPublicKey.Curve.ScalarMult(session.clientPublicKey.X, session.clientPublicKey.Y, session.serverPrivateKey.D.Bytes())
	session.sharedSecret = x.Bytes()
}

// computeIVs computes the IVs and cipher required to start encrypting and decrypting.
func (session *encryptionSession) computeIVs() error {
	var err error

	// First compute the secret key bytes, which is a hash of the salt and the shared secret computed using
	// the method above.
	session.secretKeyBytes = sha256.Sum256(append(session.salt, session.sharedSecret...))
	session.cipherBlock, err = aes.NewCipher(session.secretKeyBytes[:])
	if err != nil {
		return fmt.Errorf("error creating AES cipher: %v", err)
	}

	session.encryptIV = append([]byte{}, session.secretKeyBytes[:aes.BlockSize]...)
	session.decryptIV = append([]byte{}, session.secretKeyBytes[:aes.BlockSize]...)
	return nil
}

// encrypt encrypts the data passed in the slice itself.
func (session *encryptionSession) encrypt(data []byte) {
	for i := range data {
		cipherFeedback := cipher.NewCFBEncrypter(session.cipherBlock, session.encryptIV)
		cipherFeedback.XORKeyStream(data[i:i+1], data[i:i+1])
		session.encryptIV = append(session.encryptIV[1:], data[i])
	}
}

// decrypt decrypts the data passed in the slice itself.
func (session *encryptionSession) decrypt(data []byte) {
	for i, b := range data {
		cipherFeedback := cipher.NewCFBDecrypter(session.cipherBlock, session.decryptIV)
		cipherFeedback.XORKeyStream(data[i:i+1], data[i:i+1])
		session.decryptIV = append(session.decryptIV[1:], b)
	}
}
