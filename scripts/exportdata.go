package main

import (
	"encoding/json"
	"github.com/thoas/go-funk"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"net/http"
	"path"
	"runtime"
	"sort"
	"strings"
)

type domainType string
type fileType string

const (
	DISPOSABLE          domainType = "DISPOSABLE"
	FREE                           = "FREE"
	DISPOSABLE_WILDCARD            = "DISPOSABLE_WILDCARD"
)

const (
	JSON fileType = "JSON"
	TEXT fileType = "TEXT"
)

type domainDetails struct {
	Url        string     `yaml:"url"`
	DomainType domainType `yaml:"type"`
	FileType   fileType   `yaml:"fileType"`
}

type yamlData struct {
	Domains []domainDetails `yaml:"domains"`
}

func readYaml(filename string) yamlData {
	t := yamlData{}
	content, err := ioutil.ReadFile(filename)
	check(err)
	err = yaml.Unmarshal(content, &t)
	check(err)
	return t
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func readUrl(url string) []byte {
	resp, err := http.Get(url)
	check(err)
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	check(err)
	return data
}

func readJson(url string) []string {
	var data []string
	err := json.Unmarshal(readUrl(url), &data)
	check(err)
	sort.Strings(data)
	return data
}

func readText(url string) []string {
	data := readUrl(url)
	items := strings.Split(string(data), "\n")
	for i := 0; i < len(items); i++ {
		items[i] = strings.TrimSpace(items[i])
	}
	sort.Strings(items)
	return items
}

func readData(domain domainDetails) []string {
	if domain.FileType == JSON {
		return readJson(domain.Url)
	}
	return readText(domain.Url)
}

func readFile(filePath string) []string {
	data, err := ioutil.ReadFile(path.Join(getCurrDirectory(), filePath))
	check(err)
	items := strings.Split(string(data), "\n")
	for i := 0; i < len(items); i++ {
		items[i] = strings.TrimSpace(items[i])
	}
	sort.Strings(items)
	return items
}

func getCurrDirectory() string {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("Something went wrong!")
	}
	return path.Dir(filename)
}

func saveData(filePath string, data []string) {
	data = funk.UniqString(data)
	sort.Strings(data)
	println(len(data))
	err := ioutil.WriteFile(path.Join(getCurrDirectory(), filePath), []byte(strings.Join(data, "\n")), 0644)
	check(err)
}

func fetchData() (wildcards []string, free []string, disposable []string) {
	y := readYaml(path.Join(getCurrDirectory(), "../resources/data.yml"))
	for _, domain := range y.Domains {
		if domain.DomainType == DISPOSABLE {
			disposable = append(disposable, readData(domain)...)
		}
		if domain.DomainType == FREE {
			free = append(free, readData(domain)...)
		}
		if domain.DomainType == DISPOSABLE_WILDCARD {
			wildcards = append(wildcards, readData(domain)...)
		}
	}
	return
}

func loadData() (wildcards []string, free []string, disposable []string) {
	wildcards = readFile("../resources/wildcard-disposable-email-providers.txt")
	free = readFile("../resources/free-email-providers.txt")
	disposable = readFile("../resources/disposable-email-providers.txt")
	return
}

func main() {
	wildcards, free, disposable := fetchData()
	println(len(wildcards))
	println(len(free))
	println(len(disposable))

	pWildcards, pFree, pDisposable := loadData()

	wildcards = append(wildcards, pWildcards...)
	free = append(free, pFree...)
	disposable = append(disposable, pDisposable...)
	saveData("../resources/wildcard-disposable-email-providers.txt", wildcards)
	saveData("../resources/free-email-providers.txt", free)
	saveData("../resources/disposable-email-providers.txt", disposable)
}
