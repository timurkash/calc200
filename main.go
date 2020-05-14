package main

import (
	"fmt"
	"github.com/apaxa-go/eval"
	"log"
	"math"
	"strconv"
)

const (
	r = 3
)

func main() {
	fmt.Println()
	for i := 0; i < int(math.Pow(r, 9)); i++ {
		//fmt.Println(getDigits(i))
		exp := expr(getDigits(i))
		calc := calc(exp)
		if calc == 200 {
			fmt.Println(exp, "=", calc)
		}
	}
}

func getDigits(t int) string {
	str := ""
	for {
		i := 0
		t, i = mod(t)
		str = fmt.Sprintf("%d%s", i, str)
		if t == 0 {
			break
		}
	}
	return fmt.Sprintf("%09s", str)
}

func expr(i string) string {
	return fmt.Sprintf("9%s8%s7%s6%s5%s4%s3%s2%s1%s0",
		getSign(i, 0),
		getSign(i, 1),
		getSign(i, 2),
		getSign(i, 3),
		getSign(i, 4),
		getSign(i, 5),
		getSign(i, 6),
		getSign(i, 7),
		getSign(i, 8))
}

func getSign(s string, pos int) string {
	ch := s[pos : pos+1]
	if ch == "1" {
		return "-"
	}
	if ch == "2" {
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

func mod(i int) (int, int) {
	m := i % r
	return (i - m) / r, m
}
