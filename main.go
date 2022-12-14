package main

import (
	"flag"
	"fmt"
	"github.com/FrauElster/gopenApiToGraphQL/parser"
	"github.com/FrauElster/gopenApiToGraphQL/util"
	"log"
	"os"
)

type opts struct {
	oasFile string
	gqlFile string
}

func parseFlags() (opts, error) {
	// parse oas flag
	oasFile := flag.String("oas", "", "the openapi spec file")
	gqlRawFile := flag.String("gql", "", "the output file")
	flag.Parse()

	// check if set
	if *oasFile == "" {
		return opts{}, fmt.Errorf("\"oas\" is not set")
	}

	// check if set
	if *gqlRawFile == "" {
		return opts{}, fmt.Errorf("\"gql\" is not set")
	}
	// relative -> absolute filepath
	gqlFile, err := util.ToAbsolutePath(*gqlRawFile)
	if err != nil {
		return opts{}, err
	}

	return opts{oasFile: *oasFile, gqlFile: gqlFile}, nil
}

func main() {
	// parse flags, this will also download a http oas file to TmpDir
	opts, err := parseFlags()
	if err != nil {
		log.Fatal(err)
	}

	// parse OAS to GraphQL
	gqlSpec, err := parser.Parse(opts.oasFile)
	if err != nil {
		log.Fatalf("parsing err: %s", err)
	}

	// write it to file
	err = os.WriteFile(opts.gqlFile, []byte(gqlSpec.String()), 0644)
	if err != nil {
		log.Fatalf("could not save %s: %s", opts.gqlFile, err)
	}
	println("Here you go: " + opts.gqlFile)
}
