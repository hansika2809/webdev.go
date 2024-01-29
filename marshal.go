package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type BlogPost struct {
	Title string `json:"title"`
	URL   string `json:"url"`
	// Add other fields as needed
}

func fetchRecentBlogPosts() ([]BlogPost, error) {
	url := "https://api.example.com/recent_blog_posts" // Replace with actual API URL
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	var posts []BlogPost
	err = json.NewDecoder(response.Body).Decode(&posts)
	if err != nil {
		return nil, err
	}

	// Remove unnecessary data, if any
	for i := range posts {
		// For example, you can remove blog posts with specific criteria
		// if posts[i].Title == "Some title" {
		//     posts = append(posts[:i], posts[i+1:]...)
		//     i-- // Adjust index after removing element
		// }
	}

	return posts, nil
}

func main() {
	posts, err := fetchRecentBlogPosts()
	if err != nil {
		fmt.Println("Error fetching recent blog posts:", err)
		return
	}

	fmt.Println("Recent Blog Posts:")
	for _, post := range posts {
		fmt.Println("Title:", post.Title)
		fmt.Println("URL:", post.URL)
		fmt.Println()
	}
}
