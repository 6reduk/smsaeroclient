package commonModel

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

type Paginator struct {
	Self  string `json:"self"`
	First string `json:"first"`
	Prev  string `json:"prev"`
	Next  string `json:"next"`
	Last  string `json:"last"`
}

func NewPaginator() *Paginator {
	return &Paginator{}
}

func (p *Paginator) extractPageFrom(pageUrl string) (int, error) {
	urlParts := strings.Split(pageUrl, "?")
	if len(urlParts) != 2 {
		return 0, ErrPaginateIsMissing
	}

	query, err := url.ParseQuery(urlParts[1])
	if err != nil {
		return 0, fmt.Errorf("unable parse query for page url: %s", pageUrl)
	}

	pageStr := query.Get("page")
	if pageStr == "" {
		return 0, ErrPaginateIsMissing
	}

	page, err := strconv.Atoi(pageStr)
	if err != nil {
		return 0, fmt.Errorf("unable convert page value to int: %s", pageUrl)
	}

	return page, nil
}

func (p *Paginator) CurrentPage() (int, error) {
	return p.extractPageFrom(p.Self)
}

func (p *Paginator) NextPage() (int, error) {
	return p.extractPageFrom(p.Next)
}

func (p *Paginator) PrevPage() (int, error) {
	return p.extractPageFrom(p.Prev)
}

func (p *Paginator) FirstPage() (int, error) {
	return p.extractPageFrom(p.First)
}

func (p *Paginator) LastPage() (int, error) {
	return p.extractPageFrom(p.Last)
}
