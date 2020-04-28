package main

import (
	"log"
	"sort"
	"strconv"
	"strings"
)

const apiBundleUrl = "/modules/api/bundles"

func list(connectInfo jahiaConnectInfo) {
	var listModulesUrl = apiBundleUrl + "/*/*/*/_info"

	infoJson, err := get(connectInfo, listModulesUrl)
	if err != nil {
		log.Fatal(err)
	}

	for _, modules := range infoJson {
		var liste []string

		for module, moduleMetadata := range modules.(map[string]interface{}) {
			var info = strings.Split(module, "/")
			if len(info) == 3 {
				var groupId = info[0]
				var artifactId = info[1]
				var version = info[2]
				var state = "N/A"
				if moduleMetadata.(map[string]interface{})["moduleState"] != nil {
					state = moduleMetadata.(map[string]interface{})["moduleState"].(string)
				}

				var moduleInfo = "" + groupId + ":" + artifactId + ":" + version + " [" + state + "]"
				liste = append(liste, moduleInfo)
			}
		}

		sort.Strings(liste)
		for _, module := range liste {
			log.Println(module)
		}
	}
}

func start(info jahiaConnectInfo, moduleId string) {
	groupId, artifactId, version := splitModuleId(moduleId)
	var startUrl = apiBundleUrl + "/" + groupId + "/" + artifactId + "/" + version + "/_start"

	log.Println("Starting module " + moduleId)
	retour, err := post(info, startUrl)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(retour["message"])
	log.Println("Start successful")
}

func stop(info jahiaConnectInfo, moduleId string) {
	groupId, artifactId, version := splitModuleId(moduleId)
	var startUrl = apiBundleUrl + "/" + groupId + "/" + artifactId + "/" + version + "/_stop"

	log.Println("Stopping module " + moduleId)
	retour, err := post(info, startUrl)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(retour["message"])
	log.Println("Stop successful")
}

func uninstall(connectInfo jahiaConnectInfo, moduleId string) {
	groupId, artifactId, version := splitModuleId(moduleId)

	var uninstallUrl = apiBundleUrl + "/" + groupId + "/" + artifactId + "/" + version + "/_uninstall"

	log.Println("Uninstall module " + moduleId)

	retour, err := post(connectInfo, uninstallUrl)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(retour["message"])
	log.Println("Uninstall successfully")
}

func splitModuleId(moduleId string) (string, string, string) {
	if moduleId == "" {
		log.Fatal("Parameter moduleId is mandatory")
	}
	var s = strings.Split(moduleId, ":")
	if len(s) != 3 {
		log.Fatal("Parameter moduleId must be <group>:<id>:<version>")
	}

	var groupId = s[0]
	var artifactId = s[1]
	var version = s[2]
	return groupId, artifactId, version
}

func install(connectInfo jahiaConnectInfo, moduleFile string, startModule bool) {
	if moduleFile == "" {
		log.Fatal("Parameter file is mandatory")
	}
	if !fileExists(moduleFile) {
		log.Fatal("Module file ", moduleFile, " not exist")
	}

	var installModuleUrl = apiBundleUrl

	log.Println("Installing module " + moduleFile)

	extraParams := map[string]string{"start": strconv.FormatBool(startModule)}
	json, err := postFile(connectInfo, installModuleUrl, extraParams, "bundle", moduleFile)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(json["message"])
	log.Println("Install successfully")
}
