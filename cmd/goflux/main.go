package main

import (
	"log"
	"os"

	"github.com/Nerzal/goflux"
	"github.com/urfave/cli/v2"
)

var gofluxClient = goflux.New()

func main() {
	app := &cli.App{
		Name:  "goflux",
		Usage: "Used to automatically generate flux files",
		Commands: []*cli.Command{
			{
				Name:        "init",
				HelpName:    "Initialize",
				Description: "Initializes a new folder structure to work in",
				Usage:       "Initialize new project",
				Flags: []cli.Flag{
					cli.BashCompletionFlag,
					cli.HelpFlag,
					&cli.StringFlag{
						Required: true,
						Name:     "component",
					},
				},
				Action: Initialize,
			},
			{
				Name:        "namespace",
				HelpName:    "namespace",
				Description: "Creates a namespace file",
				Usage:       "Create a namespace file",
				Flags: []cli.Flag{
					cli.BashCompletionFlag,
					cli.HelpFlag,
					&cli.StringFlag{
						Required: true,
						Name:     "namespace",
					},
				},
				Action: NameSpace,
			},
			{
				Name:        "service",
				HelpName:    "service",
				Description: "Creates a service file",
				Usage:       "Create a service file",
				Flags: []cli.Flag{
					cli.BashCompletionFlag,
					cli.HelpFlag,
					&cli.StringFlag{
						Required: true,
						Name:     "namespace",
					},
					&cli.StringFlag{
						Required: true,
						Name:     "component",
					},
				},
				Action: Service,
			},
			{
				Name:        "backend",
				HelpName:    "backend",
				Description: "Creates folders and files for a backend service",
				Usage:       "Create files for a backend service",
				Flags: []cli.Flag{
					cli.BashCompletionFlag,
					cli.HelpFlag,
					&cli.StringFlag{
						Required: true,
						Name:     "component",
					},
					&cli.StringFlag{
						Required: true,
						Name:     "namespace",
					},
				},
				Action: Backend,
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

// Initialize is used as cli command
func Initialize(c *cli.Context) error {
	projectName := c.String("component")
	err := gofluxClient.Initialize(projectName)
	if err != nil {
		return err
	}

	return nil
}

// Backend is a cli command
func Backend(c *cli.Context) error {
	projectName := c.String("component")
	namespace := c.String("namespace")

	err := gofluxClient.CreateNameSpace(projectName, namespace)
	if err != nil {
		return err
	}

	err = gofluxClient.CreateBase(projectName, namespace)
	if err != nil {
		return err
	}

	return nil
}

// NameSpace is a cli command
func NameSpace(c *cli.Context) error {
	projectName := c.String("component")
	namespace := c.String("namespace")

	err := gofluxClient.CreateNameSpace(projectName, namespace, ".")
	if err != nil {
		return err
	}

	return nil
}

// Service is a cli command
func Service(c *cli.Context) error {
	projectName := c.String("component")
	namespace := c.String("namespace")

	err := gofluxClient.CreateService(projectName, namespace, ".")
	if err != nil {
		return err
	}

	return nil
}
