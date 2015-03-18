package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	const letter = `
Dear {{.Type}} {{.LName}},
{{if .Attending}}
I would like to extend my gratitude for attending our 11th Art Gallery. It was really great to see you and your family and I hope you enjoyed it.{{else}}
I am sorry that you could not be able to attend at our 11th Art Gallery event.{{end}}
{{if .Contribution}}
We would like to Thank you for your contribution. We greatly appreciate it.
{{end}}
I would like to remind you of our upcoming events:
{{range .Events}} {{.}}
{{end}}
Best wishes,
Mr. Sanchez
`

	// Prepare some data to insert into the template.
	type Recipient struct {
		Type, LName string
		Attend      bool
		Donate      bool
		Events      string
	}
	var recipients = []Recipient{
		{"Mr.", "Wislow", true},
		{"Mrs.", "Bennet", false},
		{"Ms.", "McCarthy", false},
	}

	var upcomingEvents = []string{"Event A", "Event B", "Event C"}

	// STEP 1 & STEP 2
	// Create a new template and parse the letter into it.
	t := template.Must(template.New("letter").Parse(letter))

	//STEP 3
	// Execute the template for each recipient.
	for _, r := range recipients {
		err := t.Execute(os.Stdout, r)
		if err != nil {
			log.Println("executing template:", err)
		}
	}

}
