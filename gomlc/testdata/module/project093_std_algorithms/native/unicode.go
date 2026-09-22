package native

import (
	"sort"
	"unicode"
)

func UnicodeFlags(value rune) uint32 {
	predicates := []func(rune) bool{
		unicode.IsControl, unicode.IsLetter, unicode.IsNumber, unicode.IsSpace,
		unicode.IsLower, unicode.IsUpper, unicode.IsDigit, unicode.IsMark,
		unicode.IsPunct, unicode.IsSymbol, unicode.IsTitle, unicode.IsGraphic, unicode.IsPrint,
	}
	var flags uint32
	for index, predicate := range predicates {
		if predicate(value) {
			flags |= 1 << index
		}
	}
	return flags
}

func UnicodeCases(value rune) (rune, rune, rune, rune) {
	return unicode.ToUpper(value), unicode.ToLower(value), unicode.ToTitle(value), unicode.SimpleFold(value)
}

func UnicodeTurkish(value rune) (rune, rune, rune) {
	return unicode.TurkishCase.ToUpper(value), unicode.TurkishCase.ToLower(value), unicode.TurkishCase.ToTitle(value)
}

func tables(family int) map[string]*unicode.RangeTable {
	switch family {
	case 0:
		return unicode.Categories
	case 1:
		return unicode.Scripts
	case 2:
		return unicode.Properties
	case 3:
		return unicode.FoldCategory
	case 4:
		return unicode.FoldScript
	default:
		panic("unknown Unicode family")
	}
}

func UnicodeNames(family int) []string {
	var result []string
	for name := range tables(family) {
		result = append(result, name)
	}
	sort.Strings(result)
	return result
}

func UnicodeRanges(family int, name string) []uint32 {
	table := tables(family)[name]
	var result []uint32
	for _, value := range table.R16 {
		result = append(result, uint32(value.Lo), uint32(value.Hi), uint32(value.Stride))
	}
	for _, value := range table.R32 {
		result = append(result, value.Lo, value.Hi, value.Stride)
	}
	return result
}

func UnicodeContains(family int, name string, value uint32) bool {
	return unicode.Is(tables(family)[name], rune(value))
}
