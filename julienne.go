package main

import (
	"encoding/csv"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

// Rows Struct used to contain all the rows in the JSON object
type Rows struct {
	Rows []Nodes `json:"rows"`
}

// Nodes map which contains the root key of each of the server elements. This is a random server name, hence this approach
type Nodes map[string]Node

// Node struct contains the actual server details
type Node struct {
	Name      string   `json:"name"`
	Os        string   `json:"os"`
	Osversion string   `json:"os_version"`
	Pkg       Packages `json:"packages"`
}

// Packages contains the root key of each package object, similar to the Nodes above
type Packages map[string]Package

// Package contains the package details
type Package struct {
	Version   string `json:"version"`
	Publisher string `json:"publisher"`
}

// Records is an array of all the record objects
type Records []Record

// Record contains the sane structure of the node information
type Record struct {
	Hostname  string
	Os        string
	Osversion string
	Packages  []string
}

func main() {
	var jsonPath = flag.String("path", "", "Specify a json file exported from knife")
	var output = flag.String("output", "apps_output.csv", "Specify a name for the exported CSV")
	flag.Parse()
	if *jsonPath == "" {
		fmt.Println("USAGE: chef-json-parse -path path_to_json_file")
		os.Exit(1)
	}
	jsonFile, err := os.Open(*jsonPath)
	if err != nil {
		fmt.Println(err)
	}
	csvOutput := [][]string{}
	csvRow := []string{}
	cleanResults := cleanJSON(jsonFile)

	for _, v := range cleanResults {
		for _, s := range v.Packages {
			csvRow = []string{v.Hostname, v.Os, v.Osversion, s}
			csvOutput = append(csvOutput, csvRow)
		}

	}

	csvfile, err := os.Create(*output)
	if err != nil {
		fmt.Println(err)
	}

	csvwriter := csv.NewWriter(csvfile)
	for _, row := range csvOutput {
		_ = csvwriter.Write(row)
	}
	csvwriter.Flush()
	csvfile.Close()
}

func cleanJSON(jsonFile *os.File) Records {
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var rows Rows
	json.Unmarshal(byteValue, &rows)
	records := Records{}

	for k := range rows.Rows {
		record := Record{}
		for _, b := range rows.Rows[k] {
			record.Os = b.Os
			record.Hostname = b.Name
			record.Osversion = b.Osversion
			for pkgname := range b.Pkg {
				record.Packages = append(record.Packages, pkgname)
			}
		}
		records = append(records, record)

	}
	return records
}
