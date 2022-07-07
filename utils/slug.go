package utils

import "github.com/gosimple/slug"

func SlugArray(Array []string) (slugArray []string) {
	for _, String := range Array {
		slugArray = append(slugArray, slug.Make(String))
	}
	return
}

func SlugString(String string) string {
	return slug.Make(String)
}
