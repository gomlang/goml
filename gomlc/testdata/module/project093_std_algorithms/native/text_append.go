package native

import "strconv"

func AppendConversions(prefix string, signed int64, unsigned uint64, floating float64, flag bool, value string) string {
	output := []byte(prefix)
	output = strconv.AppendInt(output, signed, 16)
	output = append(output, '|')
	output = strconv.AppendUint(output, unsigned, 36)
	output = append(output, '|')
	output = strconv.AppendFloat(output, floating, 'g', -1, 64)
	output = append(output, '|')
	output = strconv.AppendBool(output, flag)
	output = append(output, '|')
	output = strconv.AppendQuoteToGraphic(output, value)
	return string(output)
}
