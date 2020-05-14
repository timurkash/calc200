package main

import (
	"fmt"
	"github.com/apaxa-go/eval"
	"log"
	"strconv"
)

func main() {
	fmt.Println()
	for i1 := -1; i1 <= 1; i1++ {
		for i2 := -1; i2 <= 1; i2++ {
			for i3 := -1; i3 <= 1; i3++ {
				for i4 := -1; i4 <= 1; i4++ {
					for i5 := -1; i5 <= 1; i5++ {
						for i6 := -1; i6 <= 1; i6++ {
							for i7 := -1; i7 <= 1; i7++ {
								for i8 := -1; i8 <= 1; i8++ {
									for i9 := -1; i9 <= 1; i9++ {
										exp := expr(i1, i2, i3, i4, i5, i6, i7, i8, i9)
										calc := calc(exp)
										if calc == 200 {
											fmt.Println(exp, "=", calc)
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
}

func expr(i ...int) string {
	return fmt.Sprintf("9%s8%s7%s6%s5%s4%s3%s2%s1%s0", getSign(i[0]), getSign(i[1]), getSign(i[2]), getSign(i[3]), getSign(i[4]), getSign(i[5]), getSign(i[6]), getSign(i[7]), getSign(i[8]))
}

func getSign(s int) string {
	if s == -1 {
		return "-"
	}
	if s == 1 {
		return "+"
	}
	return ""
}

func calc(exp string) int {
	expr, _ := eval.ParseString(exp, "")
	//handleErr(err)
	r, _ := expr.EvalToInterface(nil)
	calced_ := fmt.Sprintf("%v", r)
	calced, _ := strconv.Atoi(calced_)
	return calced
}

func handleErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
