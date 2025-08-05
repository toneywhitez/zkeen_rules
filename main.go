package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	router "github.com/v2fly/v2ray-core/v5/app/router/routercommon"
	"google.golang.org/protobuf/proto"
)

func main() {

	fileFlag := flag.String("file", "", "Path to local .dat file")
	urlFlag := flag.String("url", "", "URL to fetch .dat file from")
	outDir := flag.String("out", "out", "Output directory for files")
	flag.Parse()

	var data []byte
	var err error

	switch {
	case *fileFlag != "":
		data, err = os.ReadFile(*fileFlag)
		if err != nil {
			fmt.Println("Error reading file:", err)
			os.Exit(1)
		}
	case *urlFlag != "":
		resp, err := http.Get(*urlFlag)
		if err != nil {
			fmt.Println("Error fetching from URL:", err)
			os.Exit(1)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			fmt.Println("HTTP error:", resp.Status)
			os.Exit(1)
		}

		data, err = io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error reading response body:", err)
			os.Exit(1)
		}
	default:
		fmt.Println("Please provide either -file or -url")
		os.Exit(1)
	}

	list := new(router.GeoSiteList)
	if err := proto.Unmarshal(data, list); err != nil {
		panic(err)
	}

	// 🚫 Remove existing output directory (if it exists)
	if err := os.RemoveAll(*outDir); err != nil {
		fmt.Println("Error removing output directory:", err)
		os.Exit(1)
	}

	// ✅ Recreate output directory
	if err := os.MkdirAll(*outDir, 0755); err != nil {
		fmt.Println("Error creating output directory:", err)
		os.Exit(1)
	}

	for _, entry := range list.Entry {
		//fmt.Printf("%s\n", entry.CountryCode)
		filename := filepath.Join(*outDir, strings.ToLower(entry.CountryCode)+".list")
		file, err := os.Create(filename)
		if err != nil {
			fmt.Printf("Error creating file %s: %v\n", filename, err)
			continue
		}
		for _, domain := range entry.Domain {
			//fmt.Printf("\t%s: %s\n", domain.Type, domain.Value)
			prefix := ""

			switch domain.Type {
			case router.Domain_Plain:
				prefix = "DOMAIN-KEYWORD"
			case router.Domain_RootDomain:
				prefix = "DOMAIN-SUFFIX"
			case router.Domain_Regex:
				prefix = "DOMAIN-REGEX"
			case router.Domain_Full:
				prefix = "DOMAIN"

			}
			fmt.Fprintf(file, "%s:%s\n", prefix, domain.Value)
		}
		file.Close()
	}

}
