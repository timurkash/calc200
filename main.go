package main

import (
	"fmt"
	"github.com/apaxa-go/eval"
	"log"
	"math"
	"strconv"
	"strings"
)

const (
	r = 3
)

func main() {
	fmt.Println()
	for i := 0; i < int(math.Pow(r, 9)); i++ {
		expr := fmt.Sprintf("9%s8%s7%s6%s5%s4%s3%s2%s1%s0", getSign(getDigits(i))...)
		calc := calc(expr)
		if calc == 200 {
			fmt.Println(calc, "=", expr)
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

func mod(i int) (int, int) {
	m := i % r
	return (i - m) / r, m
}

func getSign(s string) []interface{} {
	t := strings.Split(s, "")
	b := make([]interface{}, 9)
	for i := range t {
		if t[i] == "0" {
			b[i] = ""
		} else if t[i] == "1" {
			b[i] = " - "
		} else if t[i] == "2" {
			b[i] = " + "
		}
	}
	return b
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
