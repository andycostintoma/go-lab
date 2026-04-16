package file_parser

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	"strconv"
)

func ParseData(r io.Reader) ([]string, error) {
	s := bufio.NewScanner(r)
	if !s.Scan() {
		return nil, errors.New("empty")
	}
	countStr := s.Text()
	count, err := strconv.Atoi(countStr)
	//if count > 1000 {
	//	return nil, errors.New("too many")
	//}
	//if count < 0 {
	//	return nil, errors.New("no negative numbers")
	//}
	if err != nil {
		return nil, err
	}
	out := make([]string, 0, count)
	for i := 0; i < count; i++ {
		hasLine := s.Scan()
		if !hasLine {
			return nil, errors.New("too few lines")
		}
		line := s.Text()
		//line = strings.TrimSpace(line)
		//if len(line) == 0 {
		//	return nil, errors.New("blank line")
		//}
		out = append(out, line)
	}
	return out, nil
}

func ToData(s []string) []byte {
	var b bytes.Buffer
	b.WriteString(strconv.Itoa(len(s)))
	b.WriteRune('\n')
	for _, v := range s {
		b.WriteString(v)
		b.WriteRune('\n')
	}
	return b.Bytes()
}
