package main

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
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
	PrivateValue *big.Int `json:"-"`
	OverTheWire  *big.Int
	SharedSecret *big.Int `json:"-"`
	CreatedDate  time.Time
}

func (dh *dhKey) initPrivateKey() {
	var err error
	dh.PrivateValue, err = rand.Prime(rand.Reader, 256)
	if err != nil {
		log.Printf("failed to generate private key: %v", err)
	}
	dh.OverTheWire = new(big.Int).Exp(big.NewInt(int64(dh.PublicValue.G)), dh.PrivateValue, dh.PublicValue.P)
}

func (dh *dhKey) findSharedKey(otherOverTheWire *big.Int) {
	dh.SharedSecret = new(big.Int).Exp(otherOverTheWire, dh.PrivateValue, dh.PublicValue.P)
	log.Printf("shared secret established for session %s", dh.Id)
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
