package main

import (
	"log"
	"os"

	"github.com/Nerzal/goflux"
	"github.com/Nerzal/goflux/pkg/configmap"
	"github.com/urfave/cli/v2"
)

var gofluxClient = goflux.New()

func main() {
	app := &cli.App{
		Name:  "goflux",
		Usage: "Used to automatically generate kubernetes files",
		Commands: []*cli.Command{
			{
				Name:        "init",
				HelpName:    "Initialize",
				Description: "Initializes a new folder structure to work in",
				Usage:       "Initialize new project",
				Aliases:     []string{"i"},
				Flags: []cli.Flag{
					cli.BashCompletionFlag,
					cli.HelpFlag,
					&cli.StringFlag{
						Required: true,
						Aliases:  []string{"c"},
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
				Aliases:     []string{"n"},
				Flags: []cli.Flag{
					cli.BashCompletionFlag,
					cli.HelpFlag,
					&cli.StringFlag{
						Required: true,
						Aliases:  []string{"n"},
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
				Aliases:     []string{"s"},
				Flags: []cli.Flag{
					cli.BashCompletionFlag,
					cli.HelpFlag,
					&cli.StringFlag{
						Required: true,
						Aliases:  []string{"n"},
						Name:     "namespace",
					},
					&cli.StringFlag{
						Required: true,
						Aliases:  []string{"c"},
						Name:     "component",
					},
				},
				Action: Service,
			},
			{
				Name:        "configmap",
				HelpName:    "configmap",
				Description: "Creates a configmap file",
				Usage:       "Create a configmap file",
				Aliases:     []string{"c"},
				Flags: []cli.Flag{
					cli.BashCompletionFlag,
					cli.HelpFlag,
					&cli.StringFlag{
						Required: true,
						Aliases:  []string{"n"},
						Name:     "namespace",
					},
					&cli.StringFlag{
						Required: true,
						Aliases:  []string{"c"},
						Name:     "component",
					},
				},
				Action: ConfigMap,
			},
			{
				Name:        "deployment",
				HelpName:    "deployment",
				Description: "Creates a deployment file",
				Usage:       "Create a deployment file",
				Aliases:     []string{"d"},
				Flags: []cli.Flag{
					cli.BashCompletionFlag,
					cli.HelpFlag,
					&cli.StringFlag{
						Required: true,
						Aliases:  []string{"n"},
						Name:     "namespace",
					},
					&cli.StringFlag{
						Required: true,
						Aliases:  []string{"c"},
						Name:     "component",
					},
					&cli.StringFlag{
						Required: true,
						Aliases:  []string{"i"},
						Name:     "image-secret",
					},
				},
				Action: Deployment,
			},
			{
				Name:        "kustomize",
				HelpName:    "kustomize",
				Description: "Creates a kustomization file. Automagically finds ressources, and patches and also secrets in a _secret folder",
				Usage:       "Create a kustomization file",
				Aliases:     []string{"k"},
				Flags: []cli.Flag{
					cli.BashCompletionFlag,
					cli.HelpFlag,
					&cli.StringFlag{
						Required: true,
						Aliases:  []string{"n"},
						Name:     "namespace",
					},
					&cli.StringFlag{
						Required: false,
						Aliases:  []string{"p"},
						Name:     "path",
					},
				},
				Action: Kustomize,
			},
			{
				Name:        "hpa",
				HelpName:    "hpa",
				Description: "Creates a hpa file",
				Usage:       "Create a hpa file",
				Aliases:     []string{"h"},
				Flags: []cli.Flag{
					cli.BashCompletionFlag,
					cli.HelpFlag,
					&cli.StringFlag{
						Required: true,
						Aliases:  []string{"n"},
						Name:     "namespace",
					},
					&cli.StringFlag{
						Required: true,
						Aliases:  []string{"c"},
						Name:     "component",
					},
					&cli.IntFlag{
						Required: false,
						Aliases:  []string{"min"},
						Name:     "min-replicas",
					},
					&cli.IntFlag{
						Required: false,
						Aliases:  []string{"max"},
						Name:     "max-replicas",
					},
				},
				Action: HPA,
			},
			{
				Name:        "ingress",
				HelpName:    "ingress",
				Description: "Creates a ingress file",
				Usage:       "Create a ingress file",
				Aliases:     []string{"ig"},
				Flags: []cli.Flag{
					cli.BashCompletionFlag,
					cli.HelpFlag,
					&cli.StringFlag{
						Required: true,
						Aliases:  []string{"n"},
						Name:     "namespace",
					},
					&cli.StringFlag{
						Required: true,
						Aliases:  []string{"c"},
						Name:     "component",
					},
					&cli.StringFlag{
						Required: true,
						Aliases:  []string{"h"},
						Name:     "host",
					},
					&cli.StringFlag{
						Required: true,
						Aliases:  []string{"tlsh"},
						Name:     "tls-host",
					},
					&cli.StringFlag{
						Required: true,
						Aliases:  []string{"s"},
						Name:     "secret-name",
					},
					&cli.StringFlag{
						Required: true,
						Aliases:  []string{"e"},
						Name:     "endpoint",
					},
					&cli.BoolFlag{
						Required: true,
						Aliases:  []string{"acme"},
						Name:     "tls-acme",
					},
				},
				Action: Ingress,
			},
			{
				Name:        "backend",
				HelpName:    "backend",
				Description: "Creates folders and files for a backend service",
				Usage:       "Create files for a backend service",
				Aliases:     []string{"b"},
				Flags: []cli.Flag{
					cli.BashCompletionFlag,
					cli.HelpFlag,
					&cli.StringFlag{
						Required: true,
						Aliases:  []string{"c"},
						Name:     "component",
					},
					&cli.StringFlag{
						Required: true,
						Aliases:  []string{"n"},
						Name:     "namespace",
					},
				},
				Action: Backend,
			},
		},
	}

	app.EnableBashCompletion = true

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

// Deployment is a cli command
func Deployment(c *cli.Context) error {
	projectName := c.String("component")
	namespace := c.String("namespace")
	secret := c.String("image-secret")

	err := gofluxClient.CreateDeployment(projectName, namespace, secret, ".")
	if err != nil {
		return err
	}

	return nil
}

// ConfigMap is a cli command
func ConfigMap(c *cli.Context) error {
	projectName := c.String("component")
	namespace := c.String("namespace")

	exampleData := configmap.Data{}
	exampleData["CHANGE_ME"] = "CHANGE_ME"
	exampleData["EXAMPLE_KEYCLOAK_REALM"] = "CHANGE_ME"

	err := gofluxClient.CreateConfigMap(projectName, namespace, exampleData, ".")
	if err != nil {
		return err
	}

	return nil
}

// Kustomize is a cli command
func Kustomize(c *cli.Context) error {
	namespace := c.String("namespace")
	path := c.Path("path")
	if path == "" {
		path = "."
	}

	err := gofluxClient.CreateKustomization(namespace, path)
	if err != nil {
		return err
	}

	return nil
}

// HPA is a cli command
func HPA(c *cli.Context) error {
	projectName := c.String("component")
	namespace := c.String("namespace")
	minReplicas := c.Int("minReplicas")
	maxReplicas := c.Int("maxReplicas")

	err := gofluxClient.CreateHpa(projectName, namespace, minReplicas, maxReplicas, ".")
	if err != nil {
		return err
	}

	return nil
}

// Ingress is a cli command
func Ingress(c *cli.Context) error {
	projectName := c.String("component")
	namespace := c.String("namespace")
	host := c.String("host")
	tlsHost := c.String("tls-host")
	secretName := c.String("secret-name")
	endpoint := c.String("endpoint")
	tlsAcme := c.Bool("tls-acme")

	err := gofluxClient.CreateIngress(projectName, namespace, host, tlsHost, secretName, endpoint, tlsAcme, ".")
	if err != nil {
		return err
	}

	return nil
}
