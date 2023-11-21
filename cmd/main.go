package main

import (
	"article-management-service/pkg/storage"
    "article-management-service/api"
    "net/http"
    "time"
    "log"
)

func main() {
    gStorage := storage.NewArticleStorage() // Shared instance with sharing same data accross the appllication

    // Pass the storage instance to handlers
    apiHandler := api.NewAPIHandler(gStorage)

	// Define your HTTP routes here
	http.HandleFunc("/create-article", apiHandler.CreateArticleHandler)
	http.HandleFunc("/add-image", apiHandler.AddImageHandler)
	http.HandleFunc("/list-articles", apiHandler.ListArticlesHandler)

    // cleaning up expired articles
    go startCleanupRoutine(gStorage)

	// Start the HTTP server
	log.Println("Starting server on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

/*
	startCleanupRoutine will trigger the RemoveExpiredArticles function of in ArticleStorage every hour. 
	This setup ensures that service will regularly clean up expired articles without needing to restart the program.
*/
func startCleanupRoutine(storage *storage.ArticleStorage) {
    ticker := time.NewTicker(1 * time.Hour) // hourly check // time.NewTicker is a standard way in Go to run repeated tasks at regular intervals,
    for {
        select {
        case <-ticker.C:
            storage.RemoveExpiredArticles() // removes expired articles
        }
    }
}
