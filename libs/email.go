package libs

import (
	"github.com/markbates/pkger"
	"github.com/thoas/go-funk"
	"io/ioutil"
	"path"
	"regexp"
	"runtime"
	"strings"
)

var freeDomains []string
var roles []string
var wildcardDisposableDomains []string
var disposableDomains []string

type EmailDetails struct {
	Email        string
	Domain       string
	IsValid      bool
	IsDisposable bool
	IsRoleBased  bool
	IsFree       bool
	Username     string
}

func readFile(filename string) ([]string, error) {
	content, err := ioutil.ReadFile(path.Join(getCurrDirectory(), filename))
	if err != nil {
		return nil, err
	}
	items := strings.Split(string(content), "\n")
	for i := 0; i < len(items); i++ {
		items[i] = strings.TrimSpace(items[i])
	}
	return items, nil
}

func getCurrDirectory() string {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("Something went wrong!")
	}
	return path.Dir(filename)
}

func initialize() (err error) {
	if len(disposableDomains) == 0 {
		_ = pkger.Include("/../resources")
		disposableDomains, err = readFile("../resources/disposable-email-providers.txt")
		if err != nil {
			return err
		}
	}
	if len(wildcardDisposableDomains) == 0 {
		wildcardDisposableDomains, err = readFile("../resources/wildcard-disposable-email-providers.txt")
		if err != nil {
			return err
		}
	}
	if len(freeDomains) == 0 {
		freeDomains, err = readFile("../resources/free-email-providers.txt")
		if err != nil {
			return err
		}
	}
	if len(roles) == 0 {
		roles, err = readFile("../resources/roles.txt")
		if err != nil {
			return err
		}
	}
	return err
}

func isFree(domain string) bool {
	// TODO: handle subdomains for wildcard disposable domains
	return funk.ContainsString(freeDomains, domain) || funk.ContainsString(wildcardDisposableDomains, domain)
}

func isRoleBased(username string) bool {
	return funk.ContainsString(roles, username)
}

func isDisposable(domain string) bool {
	return funk.ContainsString(disposableDomains, domain)
}

func isValid(email string) bool {
	// Regex source - https://docs.isitarealemail.com/how-to-validate-email-addresses-in-golang
	pattern := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	return pattern.MatchString(email)
}

func GetEmailDetails(email string) (EmailDetails, error) {
	email = strings.ToLower(strings.TrimSpace(email))
	if !isValid(email) {
		return EmailDetails{Email: email, IsValid: false}, nil
	}
	err := initialize()
	if err != nil {
		return EmailDetails{}, err
	}
	splitEmail := strings.SplitN(email, "@", 2)
	username := splitEmail[0]
	domain := splitEmail[1]
	return EmailDetails{
		Email:        email,
		Domain:       domain,
		Username:     username,
		IsValid:      true,
		IsDisposable: isDisposable(domain),
		IsFree:       isFree(domain),
		IsRoleBased:  isRoleBased(username),
	}, nil
}
