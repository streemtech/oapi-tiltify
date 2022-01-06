package tiltifyApi

import (
	"net/url"
	"strconv"
)

type LinkParams struct {
	Before int
	Count  int
	After  int
}

func ParseLinks(linkString string) (params LinkParams) {
	res, err := url.Parse(linkString)
	if err != nil {
		return LinkParams{}
	}

	before := -1
	beforeStr := res.Query().Get("before")

	v, err := strconv.Atoi(beforeStr)
	if err == nil {
		before = v
	}

	count := -1
	countStr := res.Query().Get("count")

	v, err = strconv.Atoi(countStr)
	if err == nil {
		count = v
	}

	after := -1
	afterStr := res.Query().Get("after")

	v, err = strconv.Atoi(afterStr)
	if err == nil {
		after = v
	}

	return LinkParams{
		Before: before,
		Count:  count,
		After:  after,
	}
}
