package main

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

func TestProcessLogLine(t *testing.T) {
	expectedIps := map[string]int{"79.125.00.21": 2, "79.125.00.22": 1, "79.125.00.23": 1, "79.125.00.24": 1}
	expectedUrls := map[string]int{"/newsletter/": 2, "/test/": 1, "/": 1, "http://example.com": 1}
	processLogLine(`79.125.00.21 - - [10/Jul/2018:20:03:40 +0200] "GET /newsletter/ HTTP/1.1" 200 3574 "-" "Mozilla/5.0 (compatible; MSIE 10.0; Windows NT 6.1; Trident/5.0)"`)
	processLogLine(`79.125.00.21 - - [10/Jul/2018:20:03:40 +0200] "GET /newsletter/ HTTP/1.1" 200 3574 "-" "Mozilla/5.0 (compatible; MSIE 10.0; Windows NT 6.1; Trident/5.0)"`)
	processLogLine(`79.125.00.22 - - [10/Jul/2018:20:03:40 +0200] "GET /test/ HTTP/1.1" 200 3574 "-" "Mozilla/5.0 (compatible; MSIE 10.0; Windows NT 6.1; Trident/5.0)"`)
	processLogLine(`79.125.00.23 - - [10/Jul/2018:20:03:40 +0200] "GET / HTTP/1.1" 200 3574 "-" "Mozilla/5.0 (compatible; MSIE 10.0; Windows NT 6.1; Trident/5.0)"`)
	processLogLine(`79.125.00.24 - - [10/Jul/2018:20:03:40 +0200] "GET http://example.com HTTP/1.1" 200 3574 "-" "Mozilla/5.0 (compatible; MSIE 10.0; Windows NT 6.1; Trident/5.0)"`)
	assert.Equal(t, expectedIps, ips)
	assert.Equal(t, expectedUrls, urls)
	processLogLine(`79.125.00.25 - - [10/Jul/2018:20:03:40 +0200] "GET "Mozilla/5.0 (compatible; MSIE 10.0; Windows NT 6.1; Trident/5.0)"`)
	assert.Equal(t, expectedIps, ips)
	assert.Equal(t, expectedUrls, urls)
}

func TestGetResults(t *testing.T) {
	uniqueIps, topUrls, topIps := getResults()
	assert.Equal(t, 4, uniqueIps)
	assert.Subset(t, []string{"/newsletter/", "/test/", "/", "http://example.com"}, topUrls)
	assert.Contains(t, topUrls, "/newsletter/")
	assert.Subset(t, []string{"79.125.00.21", "79.125.00.22", "79.125.00.23", "79.125.00.24"}, topIps)
	assert.Contains(t, topIps, "79.125.00.21")
}

func TestReadInput(t *testing.T) {
	ips = make(map[string]int)
	urls = make(map[string]int)
	readInput("../../programming-task/programming-task-example-data.log")
	assert.Contains(t, ips, "168.41.191.34", "168.41.191.40", "168.41.191.41", "168.41.191.43", "168.41.191.9", "177.71.128.21", "50.112.00.11", "50.112.00.28", "72.44.32.10", "72.44.32.11", "79.125.00.21")
	assert.Contains(t, urls, "/", "/asset.css", "/asset.js", "/blog/2018/08/survey-your-opinion-matters/", "/blog/category/community/", "/docs/", "/docs/manage-users/", "/docs/manage-websites/", "/download/counter/", "/faq/", "/faq/how-to-install/", "/faq/how-to/", "/hosting/", "/intranet-analytics/", "/moved-permanently", "/newsletter/", "/temp-redirect", "/this/page/does/not/exist/", "/to-an-error", "/translations/", "http://example.net/blog/category/meta/", "http://example.net/faq/")
}

func TestWriteOutput(t *testing.T) {
	writeOutput()
	content, err := ioutil.ReadFile("../../output.txt")
	if err != nil {
		println(err)
	}
	output := string(content)
	assert.Contains(t, output, "The number of unique IP addresses: 11")
	assert.Contains(t, output, "The top 3 most visited URLs:")
	assert.Contains(t, output, "The top 3 most active IP addresses:")
}
