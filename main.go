package main

import (
	"flag"
	"log"
	"os"
)

const envJahiaUrl = "JAHIA_URL"
const envJahiaUser = "JAHIA_USER"
const envJahiaPassword = "JAHIA_PASSWORD"
const envJahiaSiteKey = "JAHIA_SITEKEY"

func main() {
	log.SetFlags(0)

	var action = flag.String("action", "", "Action : list | export | install | uninstall | start | stop")
	var jahiaUrl = flag.String("url", os.Getenv(envJahiaUrl), "Jahia base URL, can be set with environment variable "+envJahiaUrl)
	var jahiaUser = flag.String("user", os.Getenv(envJahiaUser), "Jahia user (with admin credentials), can be set with environment variable "+envJahiaUser)
	var jahiaPassword = flag.String("password", os.Getenv(envJahiaPassword), "User password, can be set with environment variable "+envJahiaPassword)
	var siteKey = flag.String("siteKey", os.Getenv(envJahiaSiteKey), "Jahia Site key, can be set with environment variable "+envJahiaSiteKey)
	var output = flag.String("output", "export.zip", "Export destination file")
	var moduleFile = flag.String("file", "", "Module JAR file path to install")
	var startModule = flag.Bool("start", true, "Start module after install")
	var moduleId = flag.String("id", "", "Fully qualified module name : <group>:<id>:<version>")
	flag.Parse()

	if *jahiaUrl == "" || *jahiaUser == "" || *jahiaPassword == "" || *siteKey == "" {
		flag.PrintDefaults()
		log.Fatal("Parameters url, user, password, siteKey are mandatory")
	}
	var connectInfo = jahiaConnectInfo{*jahiaUrl, *siteKey, *jahiaUser, *jahiaPassword}

	if *action == "" {
		flag.PrintDefaults()
		log.Fatal("Parameter action is mandatory")
	}

	switch *action {
	case "export":
		export(connectInfo, *output)
	case "install":
		install(connectInfo, *moduleFile, *startModule)
	case "list":
		list(connectInfo)
	case "uninstall":
		uninstall(connectInfo, *moduleId)
	case "start":
		start(connectInfo, *moduleId)
	case "stop":
		stop(connectInfo, *moduleId)
	default:
		flag.PrintDefaults()
		log.Fatal("Unknown action : " + *action)
	}

}
