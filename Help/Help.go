package Help

import (
	"errors"
	"regexp"
	"strconv"
)

func GetQueryByKey(key string, query map[string]string) string{
	if value, ok := query[key]; ok {
		return value;
	}
	return "";
}

func GetSecond(path string) (int64, error) {
	reg := `/(\d+)`;
	matched, err := regexp.MatchString(reg, path)
	if err != nil {
		return 0, err;
	}

	if matched {
		re, err := regexp.Compile(reg);
		if err != nil {
			return 0, err
		}
		results := re.FindAllStringSubmatch(path, -1);
		value, _ := strconv.ParseInt(results[0][1], 10, 64);
		return value, nil;
	}
	return 0, errors.New("Not Match");
}