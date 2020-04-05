package main

import (
	"fmt"
	"html/template"
	"net/http"
	"crypto/sha256"
	"encoding/hex"

	"./models"
)

var posts map[string]*models.Post

func indexHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles(
		"templates/index.html",
		"templates/header.html",
		"templates/write.html",
		"templates/footer.html",
	)
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	fmt.Println(posts)

	t.ExecuteTemplate(w, "index", posts)
}

func savePostHandler(w http.ResponseWriter, r *http.Request) {
	id := Generate()
	content := r.FormValue("content")
	hash := SetHash(content)

	post := models.NewPost(id, content, hash)
	posts[post.Id] = post

	//test hash string ~console
	fmt.Printf("hash: %x\n", post.Hash)

	http.Redirect(w, r, "/", 302)
}

//hashing func 
func SetHash(b_content string) string {
	//headers := bytes.Join([][]byte{b_id, b_content}, []byte{})
	hash := sha256.Sum256([]byte(b_content))
	str := hex.EncodeToString(hash[:])

	return str
}

func main() {
	fmt.Println("Status: server is runnig.")

	posts = make(map[string]*models.Post, 0)

	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets/"))))
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/SavePost", savePostHandler)

	http.ListenAndServe(":3000", nil)
}