package main

import (
	"encoding/json"
	"encoding/xml"
	"os"
)

type RFC struct {
	Entries []struct {
		DocID   string `xml:"doc-id", json:"docID"`
		Title   string `xml:"title", json:"title"`
		Authors struct {
			Name []string `xml:"name", "json:name"`
		} `xml:"author", json:"authors"`
		Date struct {
			Month string `xml:"month", json:"month"`
			Year  int16  `xml:"year", json:"year"`
		} `xml:"date", json:"date"`
		Keywords struct {
			Kw []string `xml:"kw", json:"kw"`
		} `xml:"keywords", json:"keywords"`
		Abstract struct {
			P string `xml:"p", json:"p"`
		} `xml:"abstract", json:"abstract"`
		Obsoletes struct {
			DocID []string `xml:"doc-id", json:"docID"`
		} `xml:"obsoletes", json:"obsoletes"`
		ObsoletedBy struct {
			DocID []string `xml:"doc-id", json:"docID"`
		} `xml:"obsoleted-by", json:"obsoletedBy"`
		Updates struct {
			DocID []string `xml:"doc-id", json:"docID"`
		} `xml:"updates", json:"updates"`
		UpdatedBy struct {
			DocID []string `xml:"doc-id", json:"docID"`
		} `xml:"updated-by", json:"updatedBy"`
		CurrentStatus     string `xml:"current-status", json:"currentStatus"`
		PublicationStatus string `xml:"publication-status", json:"publicationStatus"`
		Stream            string `xml:"stream", json:"stream"`
	} `xml:"rfc-entry", json:"rfcEntries"`
}

func main() {
	f, err := os.ReadFile("./rfc-index.xml")
	if err != nil {
		panic(err)
	}
	var rfc RFC
	if err := xml.Unmarshal(f, &rfc); err != nil {
		panic(err)
	}
	if err := json.NewEncoder(os.Stdout).Encode(rfc.Entries); err != nil {
		panic(err)
	}
}
