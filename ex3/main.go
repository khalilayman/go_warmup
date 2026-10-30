package main

import(
	"fmt"
)

func main(){
	type Note struct {
		Title   string
		Content string
		Tags 	[] string
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

