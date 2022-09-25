/*
Copyright © 2022 Andre Kruger <andre@hyvemobile.co.za>
*/
package service

import (
	"fmt"
	"html/template"
	"os"
)

const (
	TEMPLATE_LOCATION = "internal/generate/service/template.txt"
)

type Service interface {
	Generate() error
}

func New(Appname string, Appenv string, Namespace string) Service {
	if Namespace == "" {
		Namespace = Appname + "-" + Appenv
	}
	return &serviceImpl{
		Appname:    Appname,
		Appenv:     Appenv,
		filename:   "k8_service.yaml",
		foldername: "./kubernetes/" + Appenv,
		Namespace:  Namespace,
	}
}

type serviceImpl struct {
	Appname    string
	Appenv     string
	Namespace  string
	filename   string
	foldername string
}

func (s *serviceImpl) Generate() error {
	//Setup the template
	tmpl, err := template.New("template.txt").ParseFiles(TEMPLATE_LOCATION)
	if err != nil {
		return err
	}

	//Check if the folder exists, create it if not
	if _, err := os.Stat(s.foldername); os.IsNotExist(err) {
		//Create the folder
		err = os.MkdirAll(s.foldername, 0755)
		if err != nil {
			fmt.Println("Error creating folder: ", err)
			return err
		}
	}

	//Create the file. This also truncates the file if it already exists
	f, err := os.Create(s.foldername + "/" + s.filename)
	if err != nil {
		fmt.Println("Error creating file: ", err)
		return err
	}
	defer f.Close()

	//Write the template file
	err = tmpl.Execute(f, s)
	if err != nil {
		return err
	}
	return nil
}
