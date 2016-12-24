package main

import (
	"strconv"
	"strings"
)

func ApplyFilter(link string, mode int) string {
	if mode == 1 {
		return filterSliceExtension(link)
	} else if mode == 999 {
		// hard-coding
		return strings.Replace(link, "s.jpg", ".jpg", -1)
	}
	// else
	return link
}

// http://img9.bcyimg.com/coser/47840/post/17813/fd80a60032f211e6941727176b80f3f3.jpg/w650
// URL의 끝에 추가 인자가 들어가는 경우가 있다
// 여기에서는 /w650이 추가인자이며 이미지파일의 크기를 잡을떄 쓰는거로 추정
// 확장자로 추정되는 부분을 찾아서 잘라내는게 목적이다
func filterSliceExtension(link string) string {
	extensions := []string{
		".jpg",
		".jpeg",
		".png",
		".bmp",
		".gif",
	}
	founds := make([]int, len(extensions))
	lowerlink := strings.ToLower(link)

	for i, ext := range extensions {
		found := strings.LastIndex(lowerlink, ext)
		if found < 0 {
			founds[i] = found
		} else {
			founds[i] = found + len(ext)
		}
	}

	found := -1
	for _, f := range founds {
		if f > found {
			found = f
		}
	}
	return link[:found]
}

func padLeft(str, pad string, length int) string {
	if len(str) >= length {
		return str
	}

	for {
		if len(str) >= length {
			return str[0:length]
		}
		str = pad + str
	}
}

func MakeFilename(basename string, idx int) string {
	idxstr := strconv.Itoa(idx)
	idxstr = padLeft(idxstr, "0", 3)
	return idxstr + "-" + basename
}
