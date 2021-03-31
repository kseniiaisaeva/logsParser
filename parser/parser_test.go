package parser

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestExtractInfo(t *testing.T) {
	ip, url, _ := ExtractInfo(`177.71.128.21 - - [10/Jul/2018:22:21:28 +0200] "GET /intranet-analytics/ HTTP/1.1" 200 357`)
	assert.Equal(t, "177.71.128.21", ip)
	assert.Equal(t, "/intranet-analytics/", url)

	ip, url, err := ExtractInfo(`177.71.128. - - [10/Jul/2018:22:21:28 +0200] "GET intranet-analytics/ HTTP/1.1" 200 357`)
	assert.Empty(t, ip)
	assert.Empty(t, url)
	assert.Error(t, err)
}

func TestAddNewEntry(t *testing.T) {
	mp := make(map[string]int)
	mp = AddNewEntry(mp, "177.71.128.21")
	assert.Len(t, mp, 1)

	mp = AddNewEntry(mp, "177.71.128.21")
	assert.Len(t, mp, 1)
	assert.Equal(t, 2, mp["177.71.128.21"])
}

func TestGetTopThree(t *testing.T) {
	mp := map[string]int{"one": 1, "two": 2, "three": 3, "four": 4, "five": 5}
	top := GetTopThree(mp)
	assert.Len(t, top, 3)
	assert.Contains(t, top, "five", "four", "three")

	mp = map[string]int{"one": 1, "two": 2, "three": 3, "four": 2, "five": 2}
	top = GetTopThree(mp)
	assert.Len(t, top, 3)
	assert.Contains(t, top, "three")
	assert.Subset(t, []string{"two", "four", "five", "three"}, top)
}
