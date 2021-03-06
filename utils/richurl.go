package utils

import (
	"regexp"
	"strings"
)

/**
 * Converts rich url to a grouped regular expression.
 * 
 * Ex simple:   "/users/{id}/books" => "^\/users\/(?P<id>.+)\/books$"
 * Ex typed:    "/users/{id:[0-9]+}/books" => "^\/users\/(?P<id>[0-9]+)\/books$"
 * Ex multiple: "/users/{userId}/books/{bookId}" => "^\/users\/(?P<userId>.+)\/books\/(?P<bookId>.+)$"
 */
func ConvertRichUrlToRegex(path string, isComplete bool) string {

	parts := strings.Split(path, "/")
	for i, part := range parts {

		if strings.Index(part, "{") == 0 && strings.Index(part, "}") == (len(part)-1) {

			partValue := part[1 : len(part)-1]
			partValues := strings.Split(partValue, ":")
			name := "?P<" + partValues[0] + ">"

			// accept anything but '/' char in group
			regex := "[^/]+"
			if len(partValues) > 1 {
				regex = partValues[1]
			}
			parts[i] = "(" + name + regex + ")"
		}
	}

	url := "^" + strings.Join(parts, "\\/")
	if isComplete {
		url += "$"
	}
	return url
}

/**
 * Parses url with the given regular expression and returns the 
 * group values defined in the expression.
 *
 */
func GetParamsFromRichUrl(regEx, url string) (params map[string]string, matches bool) {

	var compRegEx = regexp.MustCompile(regEx)
	match := compRegEx.FindStringSubmatch(url)

	matches = len(match) > 0

	params = make(map[string]string)
	for i, name := range compRegEx.SubexpNames() {
		if i > 0 && i <= len(match) {
			params[name] = match[i]
		}
	}
	return
}
