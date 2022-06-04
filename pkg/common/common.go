package common

import "math/rand"

var StandartPatchSize = []int{8, 3}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func GenerateVoucherCode(patchSize ...int) string {
	voucherCode := []rune{}
	if len(patchSize) == 0 {
		patchSize = StandartPatchSize
	}
	for k, v := range patchSize {
		b := make([]rune, v)
		for i := range b {
			b[i] = letterRunes[rand.Intn(len(letterRunes))]
		}
		voucherCode = append(voucherCode, b...)
		if k != len(patchSize)-1 {
			voucherCode = append(voucherCode, '-')
		}
	}
	return string(voucherCode)
}
