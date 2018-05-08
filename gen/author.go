package main

import "os/user"

// Author is an author of articles
type Author struct {
	ID   string            `json:"id"`
	Name string            `json:"name"`
	URLs map[string]string `json:"urls"`
	Blog *Blog             `json:"-"`
}

// NewAuthor constructs a new author
func NewAuthor() Author {
	id, name := NewAuthorIDs()
	return Author{
		ID:   id,
		Name: name,
		URLs: map[string]string{"Twitter": "https://twitter.com/example"},
	}
}

// NewAuthorIDs constructs a new author ID and name
func NewAuthorIDs() (id, name string) {
	u, err := user.Current()
	must(err)
	if u.Name == "" {
		name = "Your Name"
	} else {
		name = u.Name
	}
	return u.Username, name
}

// Articles written by the author
func (a *Author) Articles() []Article {
	var articles []Article
	for _, article := range a.Blog.Articles {
		for _, author := range article.AuthorIDs {
			if a.ID == author {
				articles = append(articles, article)
			}
		}
	}

	return articles
}
