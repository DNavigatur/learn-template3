package main

import (
	"log"
	"os"
	"text/template"
)

type breakfast struct {
	Bfood []string
}

type lunch struct {
	Lfood []string
}

type dinner struct {
	Dfood []string
}

type menu struct {
	Breakf breakfast
	Lunc   lunch
	Dinne  dinner
}

type menus []menu

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseFiles("tmpl.gohtml"))
}

func main() {

	eat := menus{
		menu{
			Breakf: breakfast{[]string{"Tea & Bread", "Pancake & Ice-cream", "Pap & Akara"}},
			Lunc:   lunch{[]string{"Fried rice & Chicken", "Yam & Stew", "Beans & Garri"}},
			Dinne:  dinner{[]string{"Bread & Egg", "Okpa", "Indomie"}},
		},
	}
	err := tmpl.Execute(os.Stdout, eat)
	if err != nil {
		log.Fatalln(err)
	}
}
