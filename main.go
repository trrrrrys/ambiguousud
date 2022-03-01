package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var stdout = os.Stdout
var stderr = os.Stderr

var offset int
var verbose bool

func init() {
	flag.IntVar(&offset, "offset", 9, "UTC offset (hour)")
	flag.BoolVar(&verbose, "verbose", false, "verbose")
	flag.Parse()
	time.Local = time.FixedZone("", offset*3600)
	log.SetFlags(log.Lshortfile)
}

func main() {
	if err := run(); err != nil {
		fmt.Fprint(stderr, err.Error())
		os.Exit(1)
	}
}

const (
	f1 = "20060102"
	f2 = "2006-01-02"
	f3 = "2006/01/02"
	f4 = "20060102 15:04:05"
	f5 = "2006-01-02 15:04:05"
	f6 = "2006/01/02 15:04:05"
)

var patterns = []struct {
	reg    *regexp.Regexp
	layout string
}{
	{
		regexp.MustCompile(`^[0-9]{4}(0[1-9]|1[0-2])(0[1-9]|[12][0-9]|3[01])\s([01][0-9]|2[0-3]):([0-5][0-9]):([0-5][0-9])$`),
		"20060102 15:04:05",
	},
	{
		regexp.MustCompile(`^[0-9]{4}-(0[1-9]|1[0-2])-(0[1-9]|[12][0-9]|3[01])\s([01][0-9]|2[0-3]):([0-5][0-9]):([0-5][0-9])$`),
		"2006-01-02 15:04:05",
	},
	{
		regexp.MustCompile(`^[0-9]{4}/(0[1-9]|1[0-2])/(0[1-9]|[12][0-9]|3[01])\s([01][0-9]|2[0-3]):([0-5][0-9]):([0-5][0-9])$`),
		"2006/01/02 15:04:05",
	},
	{
		regexp.MustCompile("^[0-9]{4}(0[1-9]|1[0-2])(0[1-9]|[12][0-9]|3[01])$"),
		"20060102",
	},
	{
		regexp.MustCompile("^[0-9]{4}-(0[1-9]|1[0-2])-(0[1-9]|[12][0-9]|3[01])$"),
		"2006-01-02",
	},
	{
		regexp.MustCompile("^[0-9]{4}/(0[1-9]|1[0-2])/(0[1-9]|[12][0-9]|3[01])$"),
		"2006/01/02",
	},
}

func run() error {
	var arg string
	if verbose {
		log.Println(os.Args)
	}
	if n := len(os.Args); n > 2 {
		arg = strings.Join(os.Args[1:], " ")
	} else if n == 2 {
		arg = os.Args[1]
	} else {
		flag.Usage()
		return nil
	}
	result, err := AmbiguousConvert(arg)
	if err != nil {
		return err
	}
	fmt.Fprint(stdout, result)
	return nil
}

// AmbiguousConvert converts unixtime and datetime to each other
func AmbiguousConvert(s string) (any, error) {
	if i, err := strconv.Atoi(s); err == nil {
		return time.Unix(int64(i), 0).Local().Format(f5), nil
	}
	for _, v := range patterns {
		if v.reg.Copy().MatchString(s) {
			d, err := time.ParseInLocation(v.layout, s, time.Local)
			if err != nil {
				return "", fmt.Errorf("format parse error: %w", err)
			}
			return d.Unix(), nil
		}
	}
	return "", nil
}
