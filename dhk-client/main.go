package main

import "fmt"

func main() {

	dh, err := requestPublicKey()
	if err != nil {
		fmt.Println(err)
	}
	dh.initPrivateKey()
	otherPrivateValue := intiateExchange(dh)
	dh.findSharedKey(otherPrivateValue)
}
