package storage

import (
    "article-management-service/pkg/article"
    "sync"
    "time"
)

/*
    sync.RWMutex helps synchronize write operations without synchronizing 
    read operations, while ensuring that there are no active read operations 
    when there is a write operation in progress.
*/
// ArticleStorage holds the articles in memory
type ArticleStorage struct {
    mu       sync.RWMutex
    articles map[string]*article.Article
}


// NewArticleStorage creates a new instance of ArticleStorage
func NewArticleStorage() *ArticleStorage {
    return &ArticleStorage{
        articles: make(map[string]*article.Article),
    }
}

// AddArticle adds a new article to the storage
func (s *ArticleStorage) AddArticle(a *article.Article) {
    s.mu.Lock()
    defer s.mu.Unlock()
    s.articles[a.ID] = a
}
/*In the AddArticle method, an article is safely added to the articles map in ArticleStorage.
 The use of a mutex lock and defer ensures that this operation is thread-safe and the lock is 
 appropriately managed. This handling concurrent access to shared resources.*/


// GetArticle retrieves an article by ID
func (s *ArticleStorage) GetArticle(id string) (*article.Article, bool) {
    s.mu.RLock()
    defer s.mu.RUnlock()
    a, exists := s.articles[id]
    return a, exists
}

// GetAllArticles returns all articles from the storage
func (s *ArticleStorage) GetAllArticles() []*article.Article {
    s.mu.RLock()
    defer s.mu.RUnlock()

    var articles []*article.Article
    for _, art := range s.articles {
        articles = append(articles, art)
    }
    return articles
}

// RemoveArticle removes an article from storage
func (s *ArticleStorage) RemoveArticle(id string) {
    s.mu.Lock()
    defer s.mu.Unlock()
    delete(s.articles, id)
}

// RemoveExpiredArticles removes articles that have passed their expiration date
func (s *ArticleStorage) RemoveExpiredArticles() {
    s.mu.Lock()
    defer s.mu.Unlock()

    for id, article := range s.articles {
        if article.ExpireAt.Before(time.Now()) {
            delete(s.articles, id)
        }
    }
}
