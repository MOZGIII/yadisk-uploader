package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	yadisk "github.com/MOZGIII/yadisk-api"
)

var (
	oAuthToken = flag.String("token", "", "OAuth token to use")
	overwrite  = flag.Bool("overwrite", false, "do you want to overwrite the file if it exists")
	verbose    = flag.Bool("verbose", false, "print info during execution")
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [options] local_file.txt disk:/remote_file.txt\n", os.Args[0])
		flag.PrintDefaults()
	}

	flag.Parse()

	if oAuthToken == nil || *oAuthToken == "" {
		flag.Usage()
		fmt.Println("\nError:\n",
			"You must provide an OAuth token.\n",
			"See https://tech.yandex.ru/disk/rest/\n",
			"and https://tech.yandex.ru/oauth/\n",
			"") // last line is for the linter
		os.Exit(2)
		return
	}

	if flag.NArg() != 2 {
		flag.Usage()
		fmt.Println("\nError:\n",
			"You must provide a local file path and a remote file path.\n",
			"") // last line is for the linter
		os.Exit(2)
		return
	}

	src := flag.Arg(0)
	dst := flag.Arg(1)

	if *verbose {
		fmt.Printf("Using token \"%s\"...\n", *oAuthToken)
		fmt.Printf("Uploading \"%s\" to \"%s\"...\n", src, dst)
	}

	log.SetPrefix("error: ")

	file, err := os.Open(src)
	if err != nil {
		log.Fatalln(err)
		return
	}
	defer file.Close()

	api := yadisk.NewAPI(*oAuthToken)

	if err := api.Upload(file, dst, *overwrite); err != nil {
		log.Fatalln(err)
		return
	}

	fmt.Println("File uploaded successfully!")
}
