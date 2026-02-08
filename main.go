package main

import (
	"io"
	"net/http"
	"os"
)

func download(url, file string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	out, err := os.Create(file)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}

func main() {
	download(
		"https://stats-ectreport69.ect.go.th/data/records/stats_cons.json",
		"data/stats_cons.json",
	)
	download(
		"https://static-ectreport69.ect.go.th/data/data/refs/info_party_overview.json",
		"data/info_party_overview.json",
	)
	download(
		"https://stats-ectreport69.ect.go.th/data/records/stats_party.json",
		"data/stats_party.json",
	)
	download(
		"https://static-ectreport69.ect.go.th/data/data/refs/info_province.json",
		"data/info_province.json",
	)
}
