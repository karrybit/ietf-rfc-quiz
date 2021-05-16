package main

import (
	"encoding/json"
	"encoding/xml"
	"math/rand"
	"os"
	"time"
)

type RFC struct {
	Entries []RFCEntry `xml:"rfc-entry"`
}

type RFCEntry struct {
	DocID   string `xml:"doc-id"`
	Title   string `xml:"title"`
	Authors struct {
		Name []string `xml:"name"`
	} `xml:"author"`
	Date struct {
		Month string `xml:"month"`
		Year  int    `xml:"year"`
	} `xml:"date"`
	Keywords struct {
		Kw []string `xml:"kw"`
	} `xml:"keywords"`
	Abstract struct {
		P string `xml:"p"`
	} `xml:"abstract"`
	Obsoletes struct {
		DocID []string `xml:"doc-id"`
	} `xml:"obsoletes"`
	ObsoletedBy struct {
		DocID []string `xml:"doc-id"`
	} `xml:"obsoleted-by"`
	Updates struct {
		DocID []string `xml:"doc-id"`
	} `xml:"updates"`
	UpdatedBy struct {
		DocID []string `xml:"doc-id"`
	} `xml:"updated-by"`
	CurrentStatus     string `xml:"current-status"`
	PublicationStatus string `xml:"publication-status"`
	Stream            string `xml:"stream"`
}

type RFCResponse struct {
	Entries []RFCEntryResponse `json:"entries"`
}

type RFCEntryResponse struct {
	DocID             string    `json:"docID"`
	Title             string    `json:"title"`
	Authors           []string  `json:"authors"`
	Date              time.Time `json:"date"`
	Keywords          []string  `json:"keywords"`
	Abstract          string    `json:"abstract"`
	ObsoletesDocIDs   []string  `json:"obsoletesDocIDs"`
	ObsoletedByDocIDs []string  `json:"obsoletedByDocIDs"`
	UpdatesDocIDs     []string  `json:"updatesDocIDs"`
	UpdatedByDocIDs   []string  `json:"updatedByDocIDs"`
	CurrentStatus     string    `json:"currentStatus"`
	PublicationStatus string    `json:"publicationStatus"`
	Stream            string    `json:"stream"`
}

func toMonthFromStringMonth(month string) time.Month {
	switch month {
	case "January":
		return time.January
	case "February":
		return time.February
	case "March":
		return time.March
	case "April":
		return time.April
	case "May":
		return time.May
	case "June":
		return time.June
	case "July":
		return time.July
	case "August":
		return time.August
	case "September":
		return time.September
	case "October":
		return time.October
	case "November":
		return time.November
	default: // "December"
		return time.December
	}
}

func newRFCResponse(rfcEntries []RFCEntry) []RFCEntryResponse {
	rfcEntryResponse := make([]RFCEntryResponse, 0)
	for _, entry := range rfcEntries {
		rfcEntryResponse = append(rfcEntryResponse, RFCEntryResponse{
			DocID:             entry.DocID,
			Title:             entry.Title,
			Authors:           entry.Authors.Name,
			Date:              time.Date(entry.Date.Year, toMonthFromStringMonth(entry.Date.Month), 1, 0, 0, 0, 0, time.UTC),
			Keywords:          entry.Keywords.Kw,
			Abstract:          entry.Abstract.P,
			ObsoletesDocIDs:   entry.Obsoletes.DocID,
			ObsoletedByDocIDs: entry.ObsoletedBy.DocID,
			UpdatesDocIDs:     entry.Updates.DocID,
			UpdatedByDocIDs:   entry.UpdatedBy.DocID,
			CurrentStatus:     entry.CurrentStatus,
			PublicationStatus: entry.PublicationStatus,
			Stream:            entry.Stream,
		})
	}
	return rfcEntryResponse
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

	entryMap := map[uint64]RFCEntry{}
	for len(entryMap) < 4 {
		i := rand.Uint64() % uint64(len(rfc.Entries))
		entryMap[i] = rfc.Entries[i]
	}
	entries := make([]RFCEntry, 0, 4)
	for _, entry := range entryMap {
		entries = append(entries, entry)
	}
	rfcEntryResponse := newRFCResponse(entries)
	response := RFCResponse{Entries: rfcEntryResponse}
	if err := json.NewEncoder(os.Stdout).Encode(response); err != nil {
		panic(err)
	}
}
