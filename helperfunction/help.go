package format

import (
	"fmt"
	"strconv"
	"strings"
)

func All(str string) string {
	return trqnstostr((cot(punc(casea(flag(ee(split((str)))))))))
}

func trqnstostr(slice []string) string {
	result := ""
	for index, val := range slice {
		result += val
		if index != len(slice)-1 && val != "(" && slice[index+1] != ")" && (slice[index+1] != "," && slice[index+1] != ";" && slice[index+1] != ":" && slice[index+1] != "." && slice[index+1] != "!" && slice[index+1] != "?") {
			result += " "
		}
	}
	return result
}

func isflag(str string) bool {
	exmpl := []string{"hex", "bin", "low", "cap", "up"}

	for _, value := range exmpl {
		if str == value {
			return true
		}
	}
	return false
}

func ispunc(punc rune) bool {
	exmpl := []rune{',', ';', '.', '!', '?', ':'}
	for _, val := range exmpl {
		if val == punc {
			return true
		}
	}
	return false
}

func split3(slice []string) ([]string, int) {
	counter := 0
	result := []string{}
	checker := false
	//	str := ""
	// loop for counting cot
	for i := 0; i < len(slice); i++ {
		for index, value := range slice[i] {
			if value == '\'' {
				if index == 0 || index == len(slice[i])-1 {
					counter++
				}
			}
		}
	}
	ee := counter
	if counter%2 == 0 && counter != 0 {
		checker = false
	} else if counter%2 != 0 && counter != 1 { // ' abc    'abc
		checker = true
	} else {
		return slice, ee
	}
	// return counter
	checker2 := false
	for i := 0; i < len(slice); i++ {
		if len(slice[i]) == 1 {
			if slice[i] == "'" {
				result = append(result, slice[i])
				counter--
			} else {
				result = append(result, slice[i])
			}
		} else {
			for index, value := range slice[i] {
				if value == '\'' {
					if slice[i][0] == '\'' {
						if checker && counter != 1 {
							result = append(result, string(value))
							slice[i] = slice[i][1:]
							counter--
							checker2 = true
							if slice[i][len(slice[i])-1] == '\'' && counter != 1 {
								result = append(result, slice[i][:len(slice[i])-1])
								result = append(result, string(slice[i][len(slice[i])-1]))
								counter--
								checker2 = true
								break
							} else { // '''' ;;;; ','
								result = append(result, slice[i])
								checker2 = true
								break
							}
						} else if !checker && counter != 0 {
							result = append(result, string(value))
							slice[i] = slice[i][1:]
							counter--
							checker2 = true
							if slice[i][len(slice[i])-1] == '\'' && counter != 0 {
								result = append(result, slice[i][:len(slice[i])-1])
								result = append(result, string(slice[i][len(slice[i])-1]))
								counter--
								checker2 = true
								break
							} else {
								result = append(result, slice[i])
								checker2 = true
								break
							}
						}
					} else if slice[i][len(slice[i])-1] == '\'' {
						if checker && counter != 1 {
							result = append(result, slice[i][:index])
							result = append(result, string(slice[i][len(slice[i])-1]))
							counter--
							fmt.Print("hhhh")
							checker2 = true
							break
						} else if !checker && counter != 0 {
							result = append(result, slice[i][:len(slice[i])-1])
							result = append(result, string(slice[i][len(slice[i])-1]))
							counter--
							fmt.Print("hhhh")
							checker2 = true
							break
						} else {
							result = append(result, slice[i])
							checker2 = true
							break
						}
					}
				}
			}
			if checker2 {
				checker2 = false
			} else {
				result = append(result, slice[i])
			}
		}
	}
	return result, ee
}


func iscot(cot rune) bool {
	if cot == '\'' {
		return true
	}
	return false
}

func isbracket(bracket rune) bool {
	exmpl := []rune{'(', '[', '{', '}', ']', ')'}
	for _, val := range exmpl {
		if val == bracket {
			return true
		}
	}
	return false
}

func casehex(str string) string {
	str1 := ""
	dec, err := strconv.ParseInt(str, 16, 64)
	if err == nil {
		str1 = fmt.Sprintf("%d", dec)
		return str1
	}

	return str
}

func casebin(str string) string {
	str1 := ""
	dec, err := strconv.ParseInt(str, 2, 64)
	if err == nil {
		str1 = fmt.Sprintf("%d", dec)
		return str1
	}
	return str
}

func casecap(str string) string {
	capstr := strings.ToUpper(string(str[0]))
	capstr += strings.ToLower(str[1:])
	return capstr
}

func caelow(str string) string {
	lowstr := strings.ToLower(str)
	return lowstr
}

func caseup(str string) string {
	upstr := strings.ToUpper(str)
	return upstr
}

func split(str string) []string {
	expl := []rune{'\t', ' '}
	result := []string{}
	strrune := []rune(str)
	str2 := ""
	for index, value := range strrune {
		if value != expl[0] && value != expl[1] && !isbracket(value) && !iscot(value) && !ispunc(value) {
			str2 += string(value)
		} else {
			if isbracket(value) || ispunc(value) {
				if str2 != "" {
					result = append(result, str2)
					str2 = ""
					str2 += string(value)
					result = append(result, str2)
					str2 = ""
				} else {
					str2 += string(value)
					result = append(result, str2)
					str2 = ""
				}
			} else if iscot(value) {
				if index < len(strrune)-2 && iscot(strrune[index+1]) {
					if str2 != "" {
						result = append(result, str2)
						str2 = ""
						str2 += string(value)
						result = append(result, str2)
						str2 = ""
					} else {
						str2 += string(value)
						result = append(result, str2)
						str2 = ""
					}
				} else {
					str2 += string(value)
					continue
				}
			} else {
				if str2 != "" {
					result = append(result, str2)
					str2 = ""
				}
			}
		}
	}

	if str2 != "" {
		result = append(result, str2)
		str2 = ""
	}
	return result
}

func flag(slice []string) []string {
	checker := false
	index := 0

	for i := 1; i < len(slice)-1; i++ {
		if len(slice) > 3 {
			if isflag(slice[i]) && slice[i-1] == "(" && slice[i+1] == ")" {
				if slice[i] == "hex" {
					for j := i - 2; j >= 0; j-- {
						if casehex(slice[j]) != slice[j] {
							checker = true
							index = j
							break
						}
					}
					if checker {
						slice[index] = casehex(slice[index])
						slice = append(slice[:i-1], slice[i+2:]...)
						i = 0
					} else {
						slice = append(slice[:i-1], slice[i+2:]...)
						i = 0
					}
					checker = false
				} else if slice[i] == "bin" {
					for j := i - 2; j >= 0; j-- {
						if casebin(slice[j]) != slice[j] {
							checker = true
							index = j
							break
						}
					}
					if checker {
						slice[index] = casebin(slice[index])
						slice = append(slice[:i-1], slice[i+2:]...)
						i = 0
					} else {
						slice = append(slice[:i-1], slice[i+2:]...)
						i = 0
					}
					checker = false
				} else if slice[i] == "cap" {
					for j := i - 2; j >= 0; j-- {
						if caelow(slice[j]) != caseup(slice[j]) {
							checker = true
							index = j
							break
						}
					}
					if checker {
						strrune := []rune(slice[index])

						if len(strrune) > 1 {
							slice[index] = caseup(string(strrune[0])) + caelow(string(strrune[1:]))
						} else {
							slice[index] = caseup(slice[index])
						}
						slice = append(slice[:i-1], slice[i+2:]...)
						i = 0

					} else {
						slice = append(slice[:i-1], slice[i+2:]...)
						i = 0
					}
					checker = false
				} else if slice[i] == "low" {
					for j := i - 2; j >= 0; j-- {
						if caelow(slice[j]) != caseup(slice[j]) {
							checker = true
							index = j
							break
						}
					}
					if checker {
						slice[index] = caelow(slice[index])
						slice = append(slice[:i-1], slice[i+2:]...)
						i = 0
					} else {
						slice = append(slice[:i-1], slice[i+2:]...)
						i = 0
					}
					checker = false
				} else if slice[i] == "up" {
					for j := i - 2; j >= 0; j-- {
						if caelow(slice[j]) != caseup(slice[j]) {
							checker = true
							index = j
							break
						}
					}
					if checker {
						slice[index] = caseup(slice[index])
						slice = append(slice[:i-1], slice[i+2:]...)
						i = 0
					} else {
						slice = append(slice[:i-1], slice[i+2:]...)
						i = 0
					}
					checker = false
				}
			} else if isflag(slice[i]) && slice[i-1] == "(" && slice[i+1] == "," && slice[i+3] == ")" {
				if slice[i] == "cap" {
					val, err := strconv.Atoi(slice[i+2])
					if err == nil && val > 0 {
						for j := i - 2; j >= 0; j-- {
							if caelow(slice[j]) != caseup(slice[j]) {
								strrune := []rune(slice[j])

								if len(strrune) > 1 {
									slice[j] = caseup(string(strrune[0])) + caelow(string(strrune[1:]))
									val--
								} else {
									slice[j] = caseup(slice[j])
									val--
								}

							}
							if val == 0 {
								break
							}
						}
						slice = append(slice[:i-1], slice[i+4:]...)
						i = 0

					}
				} else if slice[i] == "up" {
					val, err := strconv.Atoi(slice[i+2])
					if err == nil && val > 0 {
						for j := i - 2; j >= 0; j-- {
							if caseup(slice[j]) != caelow(slice[j]) {
								slice[j] = caseup(slice[j])
								val--
							}
							if val == 0 {
								break
							}
						}
						slice = append(slice[:i-1], slice[i+4:]...)
						i = 0

					}
				} else if slice[i] == "low" {
					val, err := strconv.Atoi(slice[i+2])
					if err == nil && val > 0 {
						for j := i - 2; j >= 0; j-- {
							if caseup(slice[j]) != caelow(slice[j]) {
								slice[j] = caelow(slice[j])
								val--
							}
							if val == 0 {
								break
							}
						}
						slice = append(slice[:i-1], slice[i+4:]...)
						i = 0
					}
				}
			} else {
				continue
			}
		}
	}
	return slice
}

func casea(slice []string) []string {
	for i := 0; i < len(slice)-1; i++ {
		if slice[i] == "a" || slice[i] == "A" {
			if strings.ToLower(string(slice[i+1][0])) == "a" || strings.ToLower(string(slice[i+1][0])) == "e" || strings.ToLower(string(slice[i+1][0])) == "i" || strings.ToLower(string(slice[i+1][0])) == "u" || strings.ToLower(string(slice[i+1][0])) == "o" || strings.ToLower(string(slice[i+1][0])) == "h" {
				slice[i] = slice[i] + "n"
			}
		}
	}
	return slice
}

func punc(slice []string) []string {
	for i := 1; i < len(slice); i++ {
		if slice[i] == "," || slice[i] == ";" || slice[i] == ":" || slice[i] == "." || slice[i] == "!" || slice[i] == "?" {
			if i == 0 {
				continue
			} else {
				if slice[i-1] != "'" {
					slice[i-1] = slice[i-1] + slice[i]
					slice = append(slice[:i], slice[i+1:]...)
					i--
				}
			}
		}
	}
	return slice
}

func cot(slice []string) []string {
	_, counetr := split3(slice)
	chek := false
	if counetr%2 != 0 {
		chek = true
	}
	cheker := false
	for i := 0; i < len(slice); i++ {
		if slice[i] == "'" {
			if !cheker {
				if i < len(slice)-1 {
					if chek && counetr == 1 {
						break
					} else {
						slice[i] = slice[i] + slice[i+1]
						if slice[i+1] != "'" {
							counetr--
							cheker = true
						}
						slice = append(slice[:i+1], slice[i+2:]...)
						i -= 1
					}
				} else {
					continue
				}
			} else {
				slice[i-1] = slice[i-1] + slice[i]

				slice = append(slice[:i], slice[i+1:]...)
				i -= 2
				cheker = false

			}
		}
	}
	cheker = false
	return slice
}

func ee(slice []string) []string {
	t, _ := split3(slice)
	return t
}
