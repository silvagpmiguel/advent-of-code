package service

import (
	"fmt"
)

// Args are the AoC arguments required to run main
type Args struct {
	Cookie   string
	Action   string
	Day      string
	Year     string
	Level    string
	Answer   string
	Filepath string
}

func checkArgs(plen int, params []string) ([]string, error) {
	if plen < 4 || plen > 6 {
		return nil, fmt.Errorf("wrong arguments.\nUsage: ./main \"input|answer|solve\" \"year\" \"day\" [level] [answer]")
	}
	return params, nil
}

// NewArgs returns the arguments/error main needs to run
func NewArgs(plen int, params []string) (*Args, error) {
	arr, err := checkArgs(plen, params)

	if err != nil {
		return nil, err
	}

	env, err := NewEnv("../.env")

	if err != nil {
		return nil, err
	}

	cookie, err := env.GetEnvVariable("COOKIE")

	if err != nil {
		return nil, err
	}

	args := Args{
		Cookie:   cookie,
		Action:   arr[1],
		Year:     arr[2],
		Day:      arr[3],
		Filepath: fmt.Sprintf("../input/%s/%s.in", arr[2], arr[3]),
	}

	if plen == 5 {
		args.Level = arr[4]
	} else if plen == 6 {
		args.Level = arr[4]
		args.Answer = arr[5]
	}

	return &args, nil
}
