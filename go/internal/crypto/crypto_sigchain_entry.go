package crypto

import (
	p2pcrypto "github.com/libp2p/go-libp2p-core/crypto"
	mh "github.com/multiformats/go-multihash"
	"github.com/pkg/errors"
)

func (m *SigChainEntry) Sign(key p2pcrypto.PrivKey) error {
	entryToSign := &SigChainEntry{
		EntryTypeCode:         m.EntryTypeCode,
		ParentEntryHash:       m.ParentEntryHash,
		CreatedAt:             m.CreatedAt,
		ExpiringAt:            m.ExpiringAt,
		SignerPublicKeyBytes:  m.SignerPublicKeyBytes,
		SubjectPublicKeyBytes: m.SubjectPublicKeyBytes,
	}

	entryBytes, err := entryToSign.Marshal()
	if err != nil {
		return errors.Wrap(err, "unable to marshal entry to sign")
	}

	h, err := mh.Sum(entryBytes, mh.SHA2_256, -1)
	if err != nil {
		return errors.Wrap(err, "unable to make hash of entry to sign")
	}

	entryHash := []byte(h.B58String())

	sig, err := key.Sign(entryHash)
	if err != nil {
		return errors.Wrap(err, "unable to sign entry to add")
	}

	m.EntryHash = entryHash
	m.Signature = sig

	return nil
}

func (m *SigChainEntry) GetSignedBy() (p2pcrypto.PubKey, error) {
	pubKey, err := p2pcrypto.UnmarshalPublicKey(m.SignerPublicKeyBytes)
	if err != nil {
		return nil, errors.Wrap(err, "unable to unmarshal entry signer")
	}

	return pubKey, nil
}

func (m *SigChainEntry) GetSubject() (p2pcrypto.PubKey, error) {
	pubKey, err := p2pcrypto.UnmarshalPublicKey(m.SubjectPublicKeyBytes)
	if err != nil {
		return nil, errors.Wrap(err, "unable to unmarshal entry subject")
	}

	return pubKey, nil
}

func (m *SigChainEntry) Check() error {
	// TODO: implement me

	// TODO: Check parent present in sigchain and valid (except 1st item)
	// TODO: Check signer pubkey valid
	// TODO: Check subject pubkey valid
	// TODO: Check signature valid
	// TODO: Check not self signed (except 1st item)
	// TODO: Check signer not previously revoked

	return nil
}

var _ *SigChainEntry = (*SigChainEntry)(nil)