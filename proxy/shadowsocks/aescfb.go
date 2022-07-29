package shadowsocks

import (
	"io"

	"github.com/xtls/xray-core/common/buf"
	"github.com/xtls/xray-core/common/crypto"
)

// AesCfb represents all AES-CFB ciphers.
type AesCfb struct {
	KeyBytes int32
}

func (*AesCfb) IsAEAD() bool {
	return false
}

func (v *AesCfb) KeySize() int32 {
	return v.KeyBytes
}

func (v *AesCfb) IVSize() int32 {
	return 16
}

func (v *AesCfb) NewEncryptionWriter(key []byte, iv []byte, writer io.Writer) (buf.Writer, error) {
	stream := crypto.NewAesEncryptionStream(key, iv)
	return &buf.SequentialWriter{Writer: crypto.NewCryptionWriter(stream, writer)}, nil
}

func (v *AesCfb) NewDecryptionReader(key []byte, iv []byte, reader io.Reader) (buf.Reader, error) {
	stream := crypto.NewAesDecryptionStream(key, iv)
	return &buf.SingleReader{
		Reader: crypto.NewCryptionReader(stream, reader),
	}, nil
}

func (v *AesCfb) EncodePacket(key []byte, b *buf.Buffer) error {
	iv := b.BytesTo(v.IVSize())
	stream := crypto.NewAesEncryptionStream(key, iv)
	stream.XORKeyStream(b.BytesFrom(v.IVSize()), b.BytesFrom(v.IVSize()))
	return nil
}

func (v *AesCfb) DecodePacket(key []byte, b *buf.Buffer) error {
	if b.Len() <= v.IVSize() {
		return newError("insufficient data: ", b.Len())
	}
	iv := b.BytesTo(v.IVSize())
	stream := crypto.NewAesDecryptionStream(key, iv)
	stream.XORKeyStream(b.BytesFrom(v.IVSize()), b.BytesFrom(v.IVSize()))
	b.Advance(v.IVSize())
	return nil
}
