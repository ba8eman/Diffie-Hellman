package main

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/big"
	"time"
)

type PublicValue struct {
	P *big.Int
	G int
}

type dhKey struct {
	Id           string
	PublicValue  PublicValue
	PrivateValue *big.Int
	SharedSecret *big.Int `json:"-"`
	CreatedDate  time.Time
}

// returns a big prime and small g value
// g is hard coded here , this is a gap, it is difficult to vet a primitive root
func (dh *dhKey) intiPrimes() {
	p, _ := rand.Prime(rand.Reader, 256)
	g := 5
	pv := PublicValue{P: p, G: g}
	dh.Id = generateId(p)
	dh.PublicValue = pv
	dh.CreatedDate = time.Now()
}

func (dh *dhKey) initPrivateKey() {
	privatePrime, err := rand.Prime(rand.Reader, 256)
	if err != nil {
		fmt.Println(err)
	}
	dh.PrivateValue = new(big.Int).Exp(big.NewInt(int64(dh.PublicValue.G)), privatePrime, dh.PublicValue.P)
}

func (dh *dhKey) findSharedKey(otherPrivateValue *big.Int) {
	fmt.Println(otherPrivateValue)
	dh.SharedSecret = new(big.Int).Exp(otherPrivateValue, dh.PrivateValue, dh.PublicValue.P)
	fmt.Println(dh.SharedSecret, "is the shared secret , congratulations")
}

// generates a hashed id of prime and current time
func generateId(p *big.Int) string {
	now := time.Now()
	hasher := sha256.New()
	hasher.Write(p.Bytes())
	hasher.Write([]byte(fmt.Sprint(now.UnixNano())))
	sum := hasher.Sum(nil)
	return hex.EncodeToString(sum)
}
