package main

import(
	"fmt"
	"bufio"
	"encoding/json"
	"os"
)

func main(){
	type Note struct {
	Title   string   `json:"title"`
	Content string   `json:"content"`
	Tags    []string `json:"tags"`
}

NewNote(title, content string) *Note{
	return &Note{
		Title: title,
		Content: content,
		Tags: []tags{},
}
}

func (n *Note) preview() string  {
	const max_content = 80 
	if n.content <= max_content
	{
		return n.content
	}

	return n.content[max_content]	
}

func (n *Note) addtag(tag string)
{
	for _ ,t:= range n.Tags{
		if t ==n.Tags {
			return
		}
	} 

}

func LoadFromFile(path string) ([]*Note, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("open %s: %w", path, err)
	}
	defer f.Close()

	var notes []*Note
	if err := json.NewDecoder(bufio.NewReader(f)).Decode(&notes); err != nil {
		return nil, fmt.Errorf("decode %s: %w", path, err)
	}
	return notes, nil
}

func main() {
	n := NewNote("Intro Go", "Go est un langage compilé, simple et rapide, idéal pour les APIs.")
	n.AddTag("go")
	n.AddTag("go")
	n.AddTag("backend")

	fmt.Println("Titre   :", n.Title)
	fmt.Println("Preview :", n.Preview())
	fmt.Println("Tags    :", n.Tags)

	notes, err := LoadFromFile("notes.json")
	if err != nil {
		fmt.Println("Erreur chargement:", err)
		return
	}

	fmt.Println("\nNotes taguées 'go' :")
	for _, note := range notes {
		for _, tag := range note.Tags {
			if tag == "go" {
				fmt.Printf("- %s : %s\n", note.Title, note.Preview())
				break
			}
		}
	}
}