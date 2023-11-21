package article

import (
    "time"
)

// Article represents the structure of an article
type Article struct {
    ID          string
    Title       string
    Description string
    ExpireAt    time.Time
    ImagePaths  []string
}

// NewArticle creates a new Article object
func NewArticle(id, title, description string, expireAt time.Time) *Article {
    return &Article{
        ID:          id,
        Title:       title,
        Description: description,
        ExpireAt:    expireAt,
        ImagePaths:  make([]string, 0),
    }
}

// AddImage adds an image path to the article //By using pointers, ensure that these methods can modify the actual instance.
func (a *Article) AddImage(imagePath string) {
    a.ImagePaths = append(a.ImagePaths, imagePath)
}
