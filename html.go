
package common

import (
	"html"
	"regexp"
	"strings"
)

// encode html chars to string
func HtmlEncode(str string) string {
	return html.EscapeString(str)
}

// decode string to html chars
func HtmlDecode(str string) string {
	return html.UnescapeString(str)
}

// strip tags in html string
func StripTags(src string) string {
	//将HTML标签全转换成小写
	re, _ := regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllStringFunc(src, strings.ToLower)

	//去除STYLE
	re, _ = regexp.Compile("\\<style[\\S\\s]+?\\</style\\>")
	src = re.ReplaceAllString(src, "")

	//去除SCRIPT
	re, _ = regexp.Compile("\\<script[\\S\\s]+?\\</script\\>")
	src = re.ReplaceAllString(src, "")

	//去除所有尖括号内的HTML代码，并换成换行符
	re, _ = regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllString(src, "\n")

	//去除连续的换行符
	re, _ = regexp.Compile("\\s{2,}")
	src = re.ReplaceAllString(src, "\n")

	return strings.TrimSpace(src)
}

// change \n to <br/>
func Nl2br(str string) string {
	return strings.Replace(str, "\n", "<br/>", -1)
}

