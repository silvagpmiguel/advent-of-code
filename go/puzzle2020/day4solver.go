package puzzle2020

import (
	"aoc/solver"
	"regexp"
	"strconv"
	"strings"
)

var hclRe *regexp.Regexp
var pidRe *regexp.Regexp

// Passport struct
type Passport struct {
	Fields      map[string]string
	ContainsCid bool
	Length      int
	Valid       bool
}

// Day4 structure
type Day4 struct {
	Passports []Passport
}

func init() {
	pidRe = regexp.MustCompile(`^[0-9]{9}$`)
	hclRe = regexp.MustCompile(`^#([0-9]|[a-f]){6}$`)
}

// NewDay4Solver constructs a new solver for day 4
func NewDay4Solver() solver.Solver {
	return &Day4{}
}

// ProcessInput of day 4
func (d *Day4) ProcessInput(content string) error {
	passport := Passport{Fields: make(map[string]string), ContainsCid: false, Length: 0}
	lines := strings.Split(content, "\n")

	for _, line := range lines {
		if line == "" {
			passport.Valid = checkPassport(passport)
			d.Passports = append(d.Passports, passport)
			passport = Passport{Fields: make(map[string]string), ContainsCid: false, Length: 0}
			continue
		}
		fields := strings.Split(line, " ")
		for _, field := range fields {
			splitted := strings.Split(field, ":")
			key := splitted[0]
			val := splitted[1]
			passport.Fields[key] = val

			if key == "cid" {
				passport.ContainsCid = true
			}

			passport.Length++
		}
	}

	return nil
}

// Part1 of day 4
func (d *Day4) Part1() (string, error) {
	valid := 0

	for _, passport := range d.Passports {
		len := passport.Length
		if len == 8 || len == 7 && !passport.ContainsCid {
			valid++
		}
	}

	return strconv.Itoa(valid), nil
}

// Part2 of day 4
func (d *Day4) Part2() (string, error) {
	valid := 0

	for _, passport := range d.Passports {
		if passport.Valid {
			valid++
		}
	}

	return strconv.Itoa(valid), nil
}

func checkPassport(p Passport) bool {
	fields := p.Fields
	length := p.Length
	containsCid := p.ContainsCid
	byr := fields["byr"]
	iyr := fields["iyr"]
	eyr := fields["eyr"]
	hgt := fields["hgt"]
	ecl := fields["ecl"]
	hcl := fields["hcl"]
	pid := fields["pid"]
	isNum := false

	byrInt, err := strconv.Atoi(byr)
	isNum = isNum || err == nil
	byrCheck := (byrInt >= 1920 && byrInt <= 2002)

	iyrInt, err := strconv.Atoi(iyr)
	isNum = isNum || err == nil
	iyrCheck := (iyrInt >= 2010 && iyrInt <= 2020)

	eyrInt, err := strconv.Atoi(eyr)
	isNum = isNum || err == nil
	eyrCheck := (eyrInt >= 2020 && eyrInt <= 2030)

	eclTypes := map[string]bool{"amb": true, "blu": true, "brn": true, "gry": true, "grn": true, "hzl": true, "oth": true}
	_, eclCheck := eclTypes[ecl]

	hgtCheck := checkGht(hgt)

	return isNum && (length == 8 || length == 7 && !containsCid) && byrCheck && iyrCheck && eyrCheck && hgtCheck && hclRe.MatchString(hcl) && eclCheck && pidRe.MatchString(pid)
}

func checkGht(hgt string) bool {
	var hgtType string
	var aux string
	hgtLen := len(hgt)

	if hgtLen == 4 {
		hgtType = hgt[2:4]
		aux = hgt[:2]
	} else if hgtLen == 5 {
		hgtType = hgt[3:5]
		aux = hgt[:3]
	} else {
		return false
	}

	hgtNum, err := strconv.Atoi(aux)
	isNum := err == nil

	return isNum && (hgtType == "in" && hgtNum >= 59 && hgtNum <= 76) || (hgtType == "cm" && hgtNum >= 150 && hgtNum <= 193)
}
