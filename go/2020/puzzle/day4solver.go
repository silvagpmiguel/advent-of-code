package puzzle

import (
	"aoc/solver"
	"regexp"
	"strconv"
	"strings"
)

/*
byr (Birth Year) - four digits; at least 1920 and at most 2002.
iyr (Issue Year) - four digits; at least 2010 and at most 2020.
eyr (Expiration Year) - four digits; at least 2020 and at most 2030.
hgt (Height) - a number followed by either cm or in:
If cm, the number must be at least 150 and at most 193.
If in, the number must be at least 59 and at most 76.
hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
ecl (Eye Color) - exactly one of: amb blu brn gry grn hzl oth.
pid (Passport ID) - a nine-digit number, including leading zeroes.
cid (Country ID) - ignored, missing or not.
*/

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

	byr, err := strconv.Atoi(fields["byr"])
	if err != nil {
		return false
	}
	byrCheck := (byr >= 1920 && byr <= 2002)

	iyr, err := strconv.Atoi(fields["iyr"])
	if err != nil {
		return false
	}
	iyrCheck := (iyr >= 2010 && iyr <= 2020)

	eyr, err := strconv.Atoi(fields["eyr"])
	if err != nil {
		return false
	}
	eyrCheck := (eyr >= 2020 && eyr <= 2030)

	ecl := fields["ecl"]
	eclTypes := map[string]bool{"amb": true, "blu": true, "brn": true, "gry": true, "grn": true, "hzl": true, "oth": true}
	_, eclCheck := eclTypes[ecl]

	hgt := fields["hgt"]
	hgtLen := len(hgt)
	if hgtLen < 4 {
		return false
	}
	var hgtType string
	var aux string
	if hgtLen == 4 {
		hgtType = hgt[2:4]
		aux = hgt[:2]
	} else {
		hgtType = hgt[3:5]
		aux = hgt[:3]
	}
	hgtNum, err := strconv.Atoi(aux)
	if err != nil {
		return false
	}
	hgtCheck := (hgtType == "in" && hgtNum >= 59 && hgtNum <= 76) || (hgtType == "cm" && hgtNum >= 150 && hgtNum <= 193)

	hcl := fields["hcl"]
	hclRe := regexp.MustCompile(`^#([0-9]|[a-f]){6}$`)

	pid := fields["pid"]
	pidRe := regexp.MustCompile(`^[0-9]{9}$`)

	return (length == 8 || length == 7 && !containsCid) && byrCheck && iyrCheck && eyrCheck && hgtCheck && hclRe.MatchString(hcl) && eclCheck && pidRe.MatchString(pid)
}
