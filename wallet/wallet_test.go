package wallet_test

import (
	"bytes"
	"testing"

	"github.com/nem0z/LightStash/wallet"
)

func TestImportExport(t *testing.T) {
	const walletPath = "./tmp/test/wallet.pem"
	w, err := wallet.New()
	if err != nil {
		t.Errorf("unexpected error when creating new wallet : %q", err)
		return
	}

	err = w.Export(walletPath)
	if err != nil {
		t.Errorf("unexpected error when exporting wallet : %q", err)
		return
	}

	w2, err := wallet.Load(walletPath)
	if err != nil {
		t.Errorf("unexpected error when laoding wallet : %q", err)
		return
	}

	if !bytes.Equal(w.PubKey(), w2.PubKey()) {
		t.Errorf("pubKey from exported wallet and imported wallet are diffentes : (%q, %q)",
			w.PubKey(), w2.PubKey())
	}
}
