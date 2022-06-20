package stub

import (
	"crypto/rand"
	"math/big"
	"strconv"
)

const (
	defaultMinNumber = int64(1000000)
	defaultMaxNumber = int64(9999999)
)

func RandomDefaultStr() string {
	randNumber := RandomDefaultNumber()
	return strconv.Itoa(randNumber)
}

func RandomDefaultNumber() int {
	return RandomNumber(defaultMinNumber, defaultMaxNumber)
}

func RandomNumber(minNumber, maxNumber int64) int {
	nBig, err := rand.Int(rand.Reader, big.NewInt(maxNumber-minNumber))
	if err != nil {
		panic(err)
	}

	return int(minNumber + nBig.Int64())
}
