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
