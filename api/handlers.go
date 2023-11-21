package api

import (
	"article-management-service/pkg/article"
	"article-management-service/pkg/constants"
	"article-management-service/pkg/storage"
	"encoding/json"
	"net/http"
	"strconv"
	"log"
	"os"
)

type ImageRequest struct {
	ArticleID string `json:"articleID"`
	ImagePath string `json:"imagePath"`
}

type APIHandler struct {
    Storage *storage.ArticleStorage
}

func NewAPIHandler(storage *storage.ArticleStorage) *APIHandler {
    return &APIHandler{Storage: storage}
}

// var gStorage = storage.NewArticleStorage() // Global storage instance

// CreateArticleHandler creates a new article
func (h *APIHandler)CreateArticleHandler(w http.ResponseWriter, r *http.Request) {
	// JSON payload for creating an article
	var newArticle article.Article
	err := json.NewDecoder(r.Body).Decode(&newArticle)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Validate the article's description length
	if len(newArticle.Description) > constants.MaxDescriptionLength {
		http.Error(w, "Description exceeds maximum length", http.StatusBadRequest)
		return
	}

	// Add the article to storage
	h.Storage.AddArticle(&newArticle)

	// Respond with the created article
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newArticle)
}

// AddImageHandler adds an image to an article
func (h *APIHandler)AddImageHandler(w http.ResponseWriter, r *http.Request) {
    var req ImageRequest
    err := json.NewDecoder(r.Body).Decode(&req)
    if err != nil {
        http.Error(w, "Invalid request", http.StatusBadRequest)
        return
    }

    article, exists := h.Storage.GetArticle(req.ArticleID)
    if !exists {
        http.Error(w, "Article not found", http.StatusNotFound)
        return
    }

    if len(article.ImagePaths) >= constants.MaxImagesPerArticle {
        http.Error(w, "Maximum number of images reached", http.StatusBadRequest)
        return
    }

	// Check if the image file exists and its size
	fileInfo, err := os.Stat(req.ImagePath)
	log.Println("Trying to access image at path:", req.ImagePath)
	cwd, _ := os.Getwd()
	log.Println("Current working directory:", cwd)
	if os.IsNotExist(err) {
		http.Error(w, "Image file does not exist", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, "Error accessing image file", http.StatusInternalServerError)
		return
	}

	// Check if the file size exceeds the limit (5MB)
	if fileInfo.Size() > constants.MaxImageSize {
		http.Error(w, "Image file size exceeds the maximum limit of 5MB", http.StatusBadRequest)
		return
	}

    // Add the image to the article
    article.AddImage(req.ImagePath)

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(article)
}

// ListArticlesHandler handles requests to list articles
func (h *APIHandler)ListArticlesHandler(w http.ResponseWriter, r *http.Request) {
	// Check if the request method is GET
    if r.Method != http.MethodGet {
        http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
        return
    }

	queryValues := r.URL.Query()
	withImagesFilter, _ := strconv.ParseBool(queryValues.Get("withImages"))

	articles := h.Storage.GetAllArticles()
	var filteredArticles []*article.Article

	for _, art := range articles {
		hasImages := len(art.ImagePaths) > 0
		if withImagesFilter && hasImages || !withImagesFilter && !hasImages || queryValues.Get("withImages") == "" {
			filteredArticles = append(filteredArticles, art)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(filteredArticles)
}
