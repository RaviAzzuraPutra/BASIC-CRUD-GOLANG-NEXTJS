package helper

import "strings"

func ExtractPublicID(url string) string {
	parts := strings.Split(url, "/")

	publicPath := parts[len(parts)-2] + "/" + strings.Split(parts[len(parts)-1], ".")[0]

	return publicPath
}
