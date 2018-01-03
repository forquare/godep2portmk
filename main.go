package main

import (
	"fmt"
	"flag"
	"bytes"
	"strings"
	"io"
	"os"
	"github.com/pelletier/go-toml"
)

type Projects struct {
	SolveMeta	SolveMeta `toml:"solve-meta"`
	Project		[]Project `toml:"projects"`
}

type SolveMeta struct {
	InputsDigest	[]byte
	AnalyzerName	string
	AnalyzerVersion	int
	SolverName	string
	SolverVersion	int
}

type Project struct {
	Name     string   `toml:"name"`
	Branch   string   `toml:"branch,omitempty"`
	Revision string   `toml:"revision"`
	Version  string   `toml:"version,omitempty"`
	Source   string   `toml:"source,omitempty"`
	Packages []string `toml:"packages"`
}

func main() {
	flag.Usage = func() {
		fmt.Println(`Reads Gopkg.lock files from either stdin or as an argument`)
	}
	flag.Parse()

	var data []byte

	// Read from stdin
	if flag.NArg() == 0 {
		data = readData(os.Stdin)
	} else {
		// Read from args
		if len(os.Args[:1]) < 1 {
			fmt.Fprintln(os.Stderr, "No file given")
			os.Exit(1)
		}
		if _, err := os.Stat(os.Args[1]); os.IsNotExist(err) {
			fmt.Fprintln(os.Stderr, "File not found")
			os.Exit(2)
		}
		file, err := os.Open(os.Args[1])
		if err != nil {
			panic(err)
		}

		data = readData(file)
	}
	
	projects := Projects{}
	err := toml.Unmarshal(data, &projects)
	if err != nil {
		panic(err)
	}
	
//	fmt.Println(projects)
	for _, p := range projects.Project {
		name := strings.ToLower(p.Name)
		revision := "XXXXXXX"
		
		if len([]rune(p.Revision)) == 40 {
			revision = strings.Join(strings.Split(p.Revision, "")[0:7], "")
		} else {
			fmt.Println("-- MANUALLY DO", name, "--")
		}

		splitted := strings.Split(name, "/")

		site := splitted[0]
		account := splitted[1]
		repo := ""

		if strings.Contains(site, "golang.org") {
			account = "golang"
		}

		if strings.Contains(site, "gopkg.in") {
			if strings.Contains(account, "yaml.v2"){
				account = "go-yaml"
				repo = "yaml"
			} else {
				fmt.Println("-- MANUALLY DO", name, "--")
			}
		}

		if repo == "" {
			repo = splitted[2]
		}

		if repo == "jwalterweatherman" {
			repo = "jWalterWeatherman"
		}

		repo2 := repo

		if strings.Contains(repo2, "-") {
			repo2 = strings.Replace(repo2, "-", "_", -1)
		}

		fmt.Printf("%s:%s:%s:%s/src/%s \\\n", account, repo, revision, repo2, name)


	}
}

func readData(d io.Reader) ([]byte) {
	buf := &bytes.Buffer{}
	_, err := buf.ReadFrom(d)
	if err != nil {
		panic(err)
	}

	return buf.Bytes()
}


/*
Required output:
account:repo:commit:repo(use underscore)/src/sitename/account/repo
*/
