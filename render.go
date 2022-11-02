package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"os"
	"sort"
	"strings"

	"github.com/gomarkdown/markdown"
)

type Entity struct {
	Avatar string `json:"avatar"`
	Description template.HTML `json:"description"`
	Github string `json:"github"`
	Gitlab string `json:"gitlab"`
	Mastodon string `json:"mastodon"`
	Matrix string `json:"matrix"`
	Name string `json:"name"`
	Reddit string `json:"reddit"`
	Tags []string `json:"tags"`
	Twitter string `json:"twitter"`
	Website string `json:"website"`
	Youtube string `json:"youtube"`
	Icons template.HTML
}

type indexData struct {
	Projects []Entity
	People []Entity
}

func main() {
	indexTempl, err := template.ParseFiles("template/index.html")
	if err != nil {
		log.Fatal(err)
	}
	f, err := os.Create("static/index.html")
	if err != nil {
		log.Fatal(err)
	}
	projects := parseFiles("./projects/")
	people := parseFiles("./people/")

	sort.Slice(projects, func(i, j int) bool {
		return strings.ToLower(projects[i].Name) < strings.ToLower(projects[j].Name)
	})

	sort.Slice(people, func(i, j int) bool {
		return strings.ToLower(people[i].Name) < strings.ToLower(people[j].Name)
	})

	i := indexData{
		Projects: projects,
		People: people,
	}
	if err = indexTempl.Execute(f, i); err != nil {
		log.Fatal(err)
	}
	f.Close()
}

func parseFiles(dir string) []Entity {
	files, err := os.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}
	var e []Entity
	for _, file := range files {
		b, err := os.ReadFile(dir + file.Name())
		if err != nil {
			log.Fatal("Error when reading file %s: %s\n",
			    file.Name(), err)
		}
		var temp Entity
		if err := json.Unmarshal(b, &temp); err != nil {
			log.Fatalf("Error when rendering file %s: %s\n",
			    file.Name(), err)
		}
		temp.Description = template.HTML(markdown.ToHTML([]byte(temp.Description), nil, nil))
		if temp.Avatar == "" {
			temp.Avatar = "./assets/default/128x128.png"
		}
		renderIcons(&temp)
		e = append(e, temp)
	}
	return e
}

func renderIcons(e *Entity) {
	if (e.Matrix != "") {
		e.Icons += template.HTML(fmt.Sprintf(`
			<a href="%s"><span class="icon is-large">
			  <i class="fa fa-matrix-org fa-2x"></i>
			</span></a>`, e.Matrix))
	}
	if (e.Gitlab != "") {
		e.Icons += template.HTML(fmt.Sprintf(`
			<a href="https://gitlab.com/%s"><span class="icon is-large">
			  <i class="fa fa-gitlab fa-2x"></i>
			</span></a>`, e.Gitlab))
	}
	if (e.Github != "") {
		e.Icons += template.HTML(fmt.Sprintf(`
			<a href="https://github.com/%s"><span class="icon is-large">
			  <i class="fa fa-github fa-2x"></i>
			</span></a>`, e.Github))
	}
	if (e.Twitter != "") {
		e.Icons += template.HTML(fmt.Sprintf(`
			<a href="https://twitter.com/%s"><span class="icon is-large">
			  <i class="fa fa-twitter fa-2x"></i>
			</span></a>`, e.Twitter))
	}
	if (e.Youtube != "") {
		e.Icons += template.HTML(fmt.Sprintf(`
			<a href="https://youtube.com/%s"><span class="icon is-large">
			  <i class="fa fa-youtube fa-2x"></i>
			</span></a>`, e.Youtube))
	}
	if (e.Mastodon != "") {
		e.Icons += template.HTML(fmt.Sprintf(`
			<a href="%s"><span class="icon is-large">
			  <i class="fa fa-mastodon fa-2x"></i>
			</span></a>`, e.Mastodon))
	}
	if (e.Website != "") {
		e.Icons += template.HTML(fmt.Sprintf(`
			<a href="%s"><span class="icon is-large">
			  <i class="fa fa-external-link fa-2x"></i>
			</span></a>`, e.Website))
	}
}
