package parser

import (
	"errors"
	"regexp"
	"sort"
	"strings"
)

type kv struct {
	Key   string
	Value int
}

var ipRegex = regexp.MustCompile(`^\d+\.\d+\.\d+\.\d+`)
var urlRegex = regexp.MustCompile(` (http|https|\/)(\S)*`)

// Extract ip and url from string
func ExtractInfo(str string) (string, string, error) {
	ip := ipRegex.FindString(str)
	url := strings.Replace(urlRegex.FindString(str), " ", "", -1)
	if ip == "" || url == "" {
		return ip, url, errors.New("Unable to extract IP or URL")
	}
	return ip, url, nil
}

// Add new key with default value equal 1 or increase value if the key exists
func AddNewEntry(entryMap map[string]int, entry string) map[string]int {
	if v, ok := entryMap[entry]; ok {
		entryMap[entry] = v + 1
	} else {
		entryMap[entry] = 1
	}
	return entryMap
}

// Return top 3 keys from map with the biggest values
func GetTopThree(entryMap map[string]int) []string {
	var ss []kv
	for k, v := range entryMap {
		ss = append(ss, kv{k, v})
	}
	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value > ss[j].Value
	})
	var top []string
	for _, j := range ss[:3] {
		top = append(top, j.Key)
	}
	return top
}
