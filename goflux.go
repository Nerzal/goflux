package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/Nerzal/goflux/pkg/service"
	"github.com/urfave/cli"
	"gopkg.in/yaml.v2"
)

func main() {
	app := &cli.App{
		Name:  "goflux",
		Usage: "Used to automatically generate flux files",
		Flags: []cli.Flag{
			cli.BashCompletionFlag,
			cli.HelpFlag,
			cli.VersionFlag,
			&cli.StringFlag{
				Required: true,
				Name:     "component",
			},
			&cli.StringFlag{
				Required: true,
				Name:     "namespace",
			},
		},
		Action: func(c *cli.Context) error {
			projectName := c.String("component")
			namespace := c.String("namespace")

			err := ensureFolders(projectName)
			if err != nil {
				return err
			}

			err = createBase(projectName, namespace)
			if err != nil {
				return err
			}

			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func createBase(component, namespace string) error {
	service := service.NewService()
	serviceData := service.New(component, namespace)
	binaryData, err := yaml.Marshal(serviceData)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(fmt.Sprintf("./%s/base/service.yaml", component), binaryData, 0644)
	if err != nil {
		return err
	}

	return nil
}

func ensureFolders(projectName string) error {
	err := createFolder(projectName, "base")
	if err != nil {
		fmt.Println(err)
		return err
	}

	err = createFolder(projectName, "dev")
	if err != nil {
		fmt.Println(err)
		return err
	}

	err = createFolder(projectName, "test")
	if err != nil {
		fmt.Println(err)
		return err
	}

	err = createFolder(projectName, "prod")
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func createFolder(projectName string, folderName string) error {
	newpath := filepath.Join(".", projectName, folderName)
	return os.MkdirAll(newpath, 0777)
}
