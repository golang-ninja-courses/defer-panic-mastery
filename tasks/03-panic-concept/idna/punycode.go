package idna

import (
	"errors"
	"math"
)

const (
	base        int32 = 36
	damp        int32 = 700
	initialBias int32 = 72
	initialN    int32 = 128
	skew        int32 = 38
	tmax        int32 = 26
	tmin        int32 = 1
)

// punyEncode кодирует строку в соответствии с
// https://datatracker.ietf.org/doc/html/rfc3492#section-6.3 (Encoding procedure).
func punyEncode(s string) (string, error) { //nolint:funlen,gocognit
	output := make([]byte, 0, 2*len(s)+1)
	n, delta, bias := initialN, int32(0), initialBias

	var b, remaining int32
	for _, ch := range s {
		if ch < 128 {
			b++
			output = append(output, byte(ch))
		} else {
			remaining++
		}
	}

	if b > 0 && remaining != 0 {
		output = append(output, '-')
	}

	h := b
	for remaining != 0 { // len(s) не подойдёт, т.к. выдаст длину в байтах, а не рунах!
		m := int32(math.MaxInt32)
		for _, ch := range s {
			if ch < m && ch >= n {
				m = ch
			}
		}

		delta += (m - n) * (h + 1)
		if delta < 0 {
			return "", errors.New("delta overflow")
		}

		n = m
		for _, ch := range s {
			if ch < n {
				delta++
				if delta < 0 {
					return "", errors.New("delta overflow after inc")
				}
				continue
			}
			if ch > n {
				continue
			}

			q := delta
			for k := base; ; k += base {
				t := k - bias
				if t < tmin {
					t = tmin
				} else if t > tmax {
					t = tmax
				}

				if q < t {
					break
				}

				output = append(output, encodeDigit(t+(q-t)%(base-t)))
				q = (q - t) / (base - t)
			}

			output = append(output, encodeDigit(q))
			bias = adaptBias(delta, h+1, h == b)
			delta = 0
			h++
			remaining--
		} // range s

		delta++
		n++
	} // h < len(s)

	return string(output), nil
}

// adaptBias соответствует
// https://datatracker.ietf.org/doc/html/rfc3492#section-6.1 (Bias adaptation function).
func adaptBias(delta, numPoints int32, firstTime bool) int32 {
	if firstTime {
		delta /= damp
	} else {
		delta /= 2
	}
	delta += delta / numPoints

	k := int32(0)
	for delta > ((base-tmin)*tmax)/2 {
		delta /= base - tmin
		k += base
	}
	return k + (base-tmin+1)*delta/(delta+skew)
}

// encodeDigit возвращает символ по коду:
//   - 0..25 отображаются в ASCII a..z
//   - 26..35 отображаются в ASCII 0..9
//
// При любом другом значении digit функция паникует.
func encodeDigit(digit int32) byte {
	// Реализуй меня.
	return '0'
}
