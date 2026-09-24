package native

import "strconv"

func FormatComplex(real, imaginary float64, format byte, precision, width int) string {
	if width == 32 {
		real = float64(float32(real))
		imaginary = float64(float32(imaginary))
	}
	return strconv.FormatComplex(complex(real, imaginary), format, precision, width*2)
}
