package main

import (
	"encoding/json"
	"net/http"
)

func fetchPosts() ([]Post, error) {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var posts []Post
	if err := json.NewDecoder(resp.Body).Decode(&posts); err != nil {
		return nil, err
	}
	return posts, nil
}

func fetchPhotos() (map[int]string, error) {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/photos")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var photos []struct {
		ID     int    `json:"id"`
		URL    string `json:"url"`
		AlbumID int   `json:"albumId"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&photos); err != nil {
		return nil, err
	}

	photoMap := make(map[int]string)
	for _, photo := range photos {
		photoMap[photo.ID] = photo.URL
	}
	return photoMap, nil
}

func fetchPostsWithPhotos() ([]Post, error) {
	posts, err := fetchPosts()
	if err != nil {
		return nil, err
	}

	photos, err := fetchPhotos()
	if err != nil {
		return nil, err
	}

	for i := range posts {
		if photoURL, ok := photos[posts[i].ID]; ok {
			posts[i].Photo = photoURL
		}
	}

	return posts, nil
}
