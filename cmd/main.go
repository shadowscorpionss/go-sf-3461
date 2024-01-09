package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func printErr(err error) {
	fmt.Println("Error:", err)
}

func calc(op1, op2, oper string) (res string) {
	res = "?"
	opr := 0

	o1, err := strconv.Atoi(op1)
	if err != nil {
		printErr(err)
		return
	}
	o2, err := strconv.Atoi(op2)
	if err != nil {
		printErr(err)
		return
	}

	switch oper {
	case "+":
		opr = o1 + o2
	case "-":
		opr = o1 - o2
	case "*":
		opr = o1 * o2
	case "/":
		opr = o1 / o2
	case "^":
		opr = o1 ^ o2
	default:
		return
	}
	res = strconv.Itoa(opr)
	return
}

func main() {
	inFile := os.Args[1]
	outFile := os.Args[2]
	re := regexp.MustCompile(`(?m)^(?P<o1>[0-9]+)(?P<op>[+-/*^])(?P<o2>[0-9]+)=(?P<res>\?)$`)

	bts, err := ioutil.ReadFile(inFile)
	if err != nil {
		log.Fatal(err)
	}
	sbts := string(bts)
	sbts = strings.ReplaceAll(sbts, "\r", "")

	f, err := os.Create(outFile)
	if err != nil {
		log.Fatal(err)
	}

	w := bufio.NewWriter(f)
	fnd := re.FindAllString(sbts, -1)
	
	for _, r := range fnd {
		mts := re.FindStringSubmatch(r)
		op1, op2, oper := mts[re.SubexpIndex("o1")], mts[re.SubexpIndex("o2")], mts[re.SubexpIndex("op")]
		res := calc(op1, op2, oper)
		sbt := re.ReplaceAllString(r, fmt.Sprintf("$1$2$3=%s\n", res))
		w.WriteString(sbt)
	}
	err = w.Flush()
	if err != nil {
		log.Fatal(err)
	}
}
