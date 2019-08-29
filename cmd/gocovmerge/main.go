package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/shabbyrobe/gocovmerge"
	"golang.org/x/tools/cover"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	flag.Parse()

	var merged []*cover.Profile

	for _, file := range flag.Args() {
		profiles, err := cover.ParseProfiles(file)
		if err != nil {
			return fmt.Errorf("failed to parse profiles: %v", err)
		}
		for _, p := range profiles {
			merged = gocovmerge.AddProfile(merged, p)
		}
	}

	return gocovmerge.DumpProfiles(merged, os.Stdout)
}
