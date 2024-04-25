package wallet_test

import (
	"regexp"
	"testing"

	"github.com/nem0z/LightStash/wallet"
)

const patternP2PKH string = `^1[1-9A-HJ-NP-Za-km-z]{25,34}$`
const patternBench32 string = `^bc1q[0-9a-z]{38}$`

func testPattern(pattern, addr string) bool {
	re := regexp.MustCompile(pattern)
	return re.MatchString(addr)
}

func TestP2PKH(t *testing.T) {
	w, err := wallet.New()
	if err != nil {
		t.Errorf("unexpected error when creating new wallet : %q", err)
		return
	}

	p2pkh, err := w.P2PKH()
	if err != nil {
		t.Errorf("unexpected error when generating p2pkh : %q", err)
		return
	}

	if !testPattern(patternP2PKH, p2pkh) {
		t.Errorf("incorrect address format for p2pkh : %q", err)
	}
}

func TestBench32(t *testing.T) {
	w, err := wallet.New()
	if err != nil {
		t.Errorf("unexpected error when creating new wallet : %q", err)
		return
	}

	bench32, err := w.Bench32()
	if err != nil {
		t.Errorf("unexpected error when generating bench32 : %q", err)
		return
	}

	if !testPattern(patternBench32, bench32) {
		t.Errorf("incorrect address format for bench32 : %q", bench32)
	}
}
