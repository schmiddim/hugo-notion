package data

import (
	"fmt"
	"github.com/flytam/filenamify"
	"strings"
	"time"
)

type Tag struct {
	Name      string
	Permalink string
}

type Author struct {
	Email     string
	FirstName string
	LastName  string
}

type Post struct {
	Title       string
	Content     string
	Tags        []Tag
	DateCreated time.Time
	Author      Author
	Draft       bool
}

func (p *Post) renderMarkDownHead() string {
	return fmt.Sprintf(
		`+++
title = "%s"
date = %s
draft = %t
tags = [%s]
+++
`, p.Title, p.DateCreated.Format(time.RFC3339), p.Draft, p.TagString())

}
func (p *Post) TagString() string {
	str := ""
	for _, tag := range p.Tags {
		str += fmt.Sprintf("\"%s\", ", tag.Name)
	}
	return str
}
func (p *Post) String() string {
	str := p.renderMarkDownHead()
	str += p.Content
	return str
}

func (p *Post) TitleToFilename() string {
	output, err := filenamify.Filenamify(p.Title+".md", filenamify.Options{
		Replacement: "-",
	})
	if err != nil {
		panic(err)
	}
	return strings.TrimSpace(strings.ToLower(strings.Replace(output, " ", "-", -1)))
}

type Blog struct {
	Posts []Post
	Title string
}
