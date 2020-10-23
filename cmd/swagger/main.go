package main

import (
	"flag"
	"github.com/gitfort/goraz/pkg/openapi"
	log "github.com/sirupsen/logrus"
	"strings"
)

var (
	title       string
	description string
	version     string
	servers     string
	mixes       []string
	output      string
)

func init() {
	flag.StringVar(&title, "title", "", "title of document")
	flag.StringVar(&description, "description", "", "description of document")
	flag.StringVar(&version, "version", "", "version of document")
	flag.StringVar(&servers, "servers", "", "server urls of document")
	flag.StringVar(&output, "output", "swagger.json", "mixin swagger output")
	flag.Parse()
	mixes = flag.Args()
	if len(mixes) == 0 {
		log.Fatal("mixes is required")
	}
	if title == "" {
		log.Fatal("title is required")
	}
	if servers == "" {
		log.Fatal("servers is required")
	}
	if version == "" {
		log.Fatal("version is required")
	}
}
func main() {
	primSwagger := &openapi.Swagger{
		Version: "3.0.1",
		Info: &openapi.Info{
			Title:       title,
			Description: description,
			Version:     version,
		},
	}
	for _, server := range strings.Split(servers, ",") {
		primSwagger.Servers = append(primSwagger.Servers, &openapi.Server{
			Url: strings.TrimSpace(server),
		})
	}

	for _, mix := range mixes {
		mixSwagger, err := openapi.ReadFile(mix)
		if err != nil {
			log.WithError(err).WithField("file", mix).Fatal("can not read mix swagger file")
		}
		primSwagger = openapi.Mixin(primSwagger, mixSwagger)
	}

	if err := openapi.WriteFile(primSwagger, output); err != nil {
		log.WithError(err).Fatal("can not write mixin swagger file")
	}
}
