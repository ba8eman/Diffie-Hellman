package main

import "fmt"

func main() {

	dh, err := requestPublicKey()
	if err != nil {
		fmt.Println(err)
	}
	dh.initPrivateKey()
	otherOverTheWire := intiateExchange(dh)
	dh.findSharedKey(otherOverTheWire)
}
