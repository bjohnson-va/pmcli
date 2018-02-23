package io

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"text/template"

	"github.com/vendasta/mscli/pkg/spec"
)

// WriteFile writes a file from a specified template to the specified location
func WriteFile(spec spec.MicroserviceFile, templateConstant, location, leftDelim, rightDelim string) error {
	fmt.Printf("Writing File... %s\n", location)

	var tmpl *template.Template
	var err error

	if tmpl, err = template.
		New(fmt.Sprintf("%s_boilerplate", location)).
		Delims(leftDelim, rightDelim).
		Parse(templateConstant); err != nil {
		return err
	}

	buf := bytes.NewBufferString("")
	if err = tmpl.Execute(buf, spec.Microservice); err != nil {
		return err
	}

	if spec.Microservice.Debug {
		fmt.Printf("------- %s: --------\n", location)
		fmt.Printf("%s", buf.String())
		fmt.Printf("----------------------------\n")
	}

	var f *os.File
	if f, err = os.Create(location); err != nil {
		return err
	}
	defer f.Close()
	if _, err = f.WriteString(buf.String()); err != nil {
		return err
	}
	return nil
}

// DeleteFile deletes the specified file
func DeleteFile(file string) error {
	return os.Remove(file)
}

// CreateTemplatedFile provides a nicer interface to the text/template lib
func CreateTemplatedFile(url string, data interface{}, templateCode string) error {
	buf := bytes.NewBufferString("")

	tmpl, err := template.New(fmt.Sprintf("%s_boiler", url)).Parse(templateCode)
	if err != nil {
		return fmt.Errorf("error creating template: %s", err.Error())
	}

	err = tmpl.Execute(buf, data)
	if err != nil {
		return fmt.Errorf("error executing template: %s", err.Error())
	}

	f, err := os.Create(url)
	if err != nil {
		return fmt.Errorf("error creating file: %s", err.Error())
	}
	defer f.Close()

	_, err = f.WriteString(buf.String())
	if err != nil {
		return fmt.Errorf("error writing file: %s", err.Error())
	}

	return nil
}

//EnsureDirExists meake sure the specified directory exists and makes it if it doesn't
func EnsureDirExists(spec spec.MicroserviceFile, dirName string) error {
	var err error
	var stat os.FileInfo

	if stat, err = os.Stat(dirName); err == nil {
		if stat.IsDir() {
			if spec.Microservice.Debug {
				fmt.Printf("%s directory already exists, skipping creation\n", dirName)
			}
			return nil
		}
		return fmt.Errorf("file %s exists, and is not a directory", dirName)
	} else {
		if os.IsNotExist(err) {
			if err = os.MkdirAll(dirName, 0755); err != nil {
				return fmt.Errorf("error creating directory %s: %s", dirName, err.Error())
			} else {
				if spec.Microservice.Debug {
					fmt.Printf("%s directory created\n", dirName)
				}
				return nil
			}
		}
	}
	return fmt.Errorf("error encountered with creating directory %s: %s", dirName, err.Error())
}

//EnsureFileExists exits if the specified file doesn't exist
func EnsureFileExists(fileName string) error {
	var err error
	var stat os.FileInfo
	stat, err = os.Stat(fileName)
	if err == nil {
		if !stat.IsDir() {
			return nil
		}
	}
	return fmt.Errorf("error encountered, file %s doesn't exist", fileName)
}

//PrependOrCreateTemplateToFile provides an interface to add to the start of a file the template.
func PrependOrCreateTemplateToFile(url string, data interface{}, templateString string) error {
	buf := bytes.NewBufferString("")

	tmpl, err := template.New(fmt.Sprintf("%s_boiler", url)).Parse(templateString)
	if err != nil {
		return fmt.Errorf("error creating template: %s", err.Error())
	}

	err = tmpl.Execute(buf, data)
	if err != nil {
		return fmt.Errorf("Error executing template: %s", err.Error())
	}

	fileBytes, err := ioutil.ReadFile(url)
	fileString := buf.String()
	if err == nil {
		// no error means keep data
		fileString = fileString + "\n" + string(fileBytes)
	}

	f, err := os.Create(url)
	if err != nil {
		return fmt.Errorf("Error creating file: %s", err.Error())
	}
	defer f.Close()

	_, err = f.WriteString(fileString)
	if err != nil {
		return fmt.Errorf("Error writing file: %s", err.Error())
	}
	return nil
}
