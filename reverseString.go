package main

func RS(str string) string {
	r := []rune(str)
	f := 0
	l := len(r) - 1
	for f < l {
		r[f], r[l] = r[l], r[f]
		f++
		l--
	}
	return string(r)
}
