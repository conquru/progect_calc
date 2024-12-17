package calculator

import (
	"fmt"
	"strconv"
)

func Check(symbol string) bool {
	return symbol == "+" || symbol == "-" || symbol == "*" || symbol == "/"
}

func Result(lst_oper, lst_num []string) (float64, error) {
	result, err := strconv.ParseFloat(lst_num[0], 64)
	if err != nil {
		return 0, err
	}
	for i := 0; i < len(lst_oper); i++ {
		b, err := strconv.ParseFloat(lst_num[i+1], 64)
		if err != nil {
			return 0, err
		}
		if lst_oper[i] == "+" {
			result += b
		} else {
			result -= b
		}
	}
	return result, nil
}

func Calc(expression string) (float64, error) {
	operators := ""
	lst_num := []string{}
	lst_oper := []string{}
	lst_num_1 := []string{}
	lst_oper_1 := []string{}
	num := ""
	for i := 0; i < len(expression); i++ {
		if string(expression[i]) == "(" {
			for k := i + 1; k < len(expression); k++ {
				if string(expression[k]) == ")" {
					i++
					break
				} else {
					if Check(string(expression[k])) {
						if string(expression[k]) == ")" {
							break
						} else if string(expression[i]) == "*" || string(expression[i]) == "/" {
							operators = string(expression[i])
							a, err := strconv.ParseFloat(num, 64)
							if err != nil {
								return 0, err
							}
							num = ""
							for j := i + 1; j < len(expression); j++ {
								if Check(string(expression[j])) {
									break
								} else {
									num += string(expression[j])
								}
								i++
								k++
							}

							b, err := strconv.ParseFloat(num, 64)
							if err != nil {
								return 0, err
							}
							if operators == "*" {
								lst_num_1 = append(lst_num_1, fmt.Sprintf("%f", a*b))
							} else {
								lst_num_1 = append(lst_num_1, fmt.Sprintf("%f", a/b))
							}
						} else {
							if num != "" {
								lst_num_1 = append(lst_num_1, num)
							}
							lst_oper_1 = append(lst_oper_1, string(expression[k]))
						}
						num = ""
					} else {
						num += string(expression[k])
					}
				}
				i++
			}
			lst_num_1 = append(lst_num_1, num)
			variable, err := Result(lst_oper_1, lst_num_1)
			if err != nil {
				return 0, err
			}
			variable_s := strconv.FormatFloat(variable, 'g', -1, 64)
			num = variable_s
		} else if Check(string(expression[i])) {
			if string(expression[i]) == "*" || string(expression[i]) == "/" {
				operators = string(expression[i])
				a, err := strconv.ParseFloat(num, 64)
				if err != nil {
					return 0, err
				}
				num = ""
				for j := i + 1; j < len(expression); j++ {
					if Check(string(expression[j])) {
						break
					} else {
						num += string(expression[j])
					}
					i++
				}

				b, err := strconv.ParseFloat(num, 64)
				if err != nil {
					return 0, err
				}
				if operators == "*" {
					lst_num = append(lst_num, fmt.Sprintf("%f", a*b))
				} else {
					lst_num = append(lst_num, fmt.Sprintf("%f", a/b))
				}
			} else {
				if num != "" {
					lst_num = append(lst_num, num)
				}
				lst_oper = append(lst_oper, string(expression[i]))
			}
			num = ""
		} else {
			num += string(expression[i])
		}
	}
	lst_num = append(lst_num, num)
	return Result(lst_oper, lst_num)
}
