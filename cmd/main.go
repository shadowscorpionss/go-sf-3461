package main

import (
	"bufio"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// prints error
func printErr(err error) {
	fmt.Println("Error:", err)
}

// calculates op1 oper op2
// oper +-/*^
// op1, op2 must be int
// returns int converted to string or '?' if something wrong
func calc(op1, op2, oper string) (res string) {
	res = "?"
	//int result
	opr := 0

	//converting operands to int
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

	//switching operation
	switch oper {		
	case "+"://addition
		opr = o1 + o2
	case "-"://subtraction
		opr = o1 - o2
	case "*"://multiplication
		opr = o1 * o2
	case "/"://division
		if o2==0{
			printErr(errors.New("division by zero"))
		}
		opr = o1 / o2
	case "^"://power
		opr = o1 ^ o2
	default:
		//if we don't know operation
		return
	}

	//converting result to string
	res = strconv.Itoa(opr)
	return
}

// reads input file (args[1]) and creates output file (args[2])
// prints calculated result into output file
func main() {
	//we need 2 arguments input and output files
	//but args[0] is executable filename
	if len(os.Args) < 3 {
		log.Fatal("too few arguments")
	}

	//filenames (input, output)
	inFile := os.Args[1]
	outFile := os.Args[2]

	//multiline regex with group names to decode groups and calc result
	re := regexp.MustCompile(`(?m)^(?P<o1>[0-9]+)(?P<op>[+-/*^])(?P<o2>[0-9]+)=(?P<res>\?)$`)

	//reading bytes from file
	bts, err := ioutil.ReadFile(inFile)

	//error of reading file
	if err != nil {
		log.Fatal(err)
	}

	//we will work with strings, thats why converting from bytes
	sbts := string(bts)

	//windows split lines with \r\n, go works with \n
	sbts = strings.ReplaceAll(sbts, "\r", "")

	//creating or replacing filecontent
	f, err := os.Create(outFile)
	if err != nil {
		log.Fatal(err)
	}

	//buffered writer
	w := bufio.NewWriter(f)

	//-1 to find all regex matched strings in sbts
	fnd := re.FindAllString(sbts, -1)

	//for each found string
	for _, r := range fnd {
		//splitting submatch to get values of operands and operation
		mts := re.FindStringSubmatch(r)
		op1, op2, oper := mts[re.SubexpIndex("o1")], mts[re.SubexpIndex("o2")], mts[re.SubexpIndex("op")]
		//calculating result. if err - ?
		res := calc(op1, op2, oper)
		//replacing string
		sbt := re.ReplaceAllString(r, fmt.Sprintf("$1$2$3=%s\n", res))
		//writing to output file with bufio
		w.WriteString(sbt)
	}
	//flush it! =)
	err = w.Flush()

	//logging error if not nil
	if err != nil {
		log.Fatal(err)
	}
}
