package main

import (
	"log"
)

func export(connectInfo jahiaConnectInfo, output string) {
	var siteExportUrl = "/cms/export/default/export.zip?exportformat=site&live=true&sitebox=" + connectInfo.siteKey

	log.Println("Export site " + connectInfo.siteKey)
	log.Println("Downloading archive...")

	if err := downloadFile(connectInfo, output, siteExportUrl); err != nil {
		log.Fatal(err)
	}
	log.Println("Export successful")
}

func importSite(info jahiaConnectInfo, file string) {
	if file == "" {
		log.Fatal("Parameter file is mandatory")
	}
	if !fileExists(file) {
		log.Fatal("ZIP file ", file, " not exist")
	}

	var importSiteUrl = "/modules/api/jahia-tools/sites"

	log.Println("Import site " + info.siteKey)
	extraParams := map[string]string{"site": info.siteKey}
	result, err := postFile(info, importSiteUrl, extraParams, "file", file)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(result["message"])
}

func remove(info jahiaConnectInfo) {
	var removeSiteUrl = "/modules/api/jahia-tools/sites/" + info.siteKey

	log.Println("Remove site " + info.siteKey)
	result, err := delete(info, removeSiteUrl)
	if err != nil {
		log.Println("Hint : Have you installed the module jahia-tools-module ?")
		log.Fatal(err)
	}

	log.Println(result["message"])
}
