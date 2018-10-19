package des

import (
	"testing"
)

func TestDes_1(t *testing.T) {
	plain := "Hello, world!"
	cipher := Encrypt(plain)
	mPlain := Decrypt(cipher)
	t.Log("plain: "+plain)
	t.Log("cipher: "+ cipher)
	if mPlain != plain {
		t.Error(mPlain)
	}
}

func TestDes_2(t *testing.T) {
	plain := "DES (Data Encryption Standard) was developed in the early 1970s at IBM and based on an earlier design by Horst Feistel. The algorithm was submitted to the National Bureau of Standards (NBS) following the agency's invitation to propose a candidate for the protection of sensitive, unclassified electronic government data. In 1976, after consultation with the National Security Agency (NSA), the NBS eventually selected a slightly modified version (strengthened against differential cryptanalysis, but weakened against brute force attacks), which was published as an official Federal Information Processing Standard (FIPS) for the United States in 1977."
	cipher := Encrypt(plain)
	mPlain := Decrypt(cipher)
	t.Log("plain: "+plain)
	t.Log("cipher: "+ cipher)
	if mPlain != plain {
		t.Error(mPlain)
	}
}