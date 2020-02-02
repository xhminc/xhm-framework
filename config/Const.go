package config

import "regexp"

const (
	DES_KEY = "XHM_3DES_KEY"
	DES_IV  = "XHM_3DES_IV"
)

var (
	Mobile, _ = regexp.Compile("^1(3|4|5|6|7|8|9)\\d{9}$")
	Email, _  = regexp.Compile("[\\w!#$%&'*+/=?^_`{|}~-]+(?:\\.[\\w!#$%&'*+/=?^_`{|}~-]+)*@(?:[\\w](?:[\\w-]*[\\w])?\\.)+[\\w](?:[\\w-]*[\\w])?")
)
