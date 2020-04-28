# Jahia Tools

Administrative tool for Jahia Digital Experience

- Export site
- Delete site (jahia-tools-module required)
- Import site (jahia-tools-module required)
- List installed bundles
- Deploy a bundle
- Start a bundle
- Stop a bundle

To use delete/import site features, you will need to have the jahia-tools-module installed et running on the cluster.

Avoid install this module on your production cluster :).

See https://github.com/cnieg/jahia-tools-module

## Usage

Command line arguments :

```
  -action string
        Sites:  export | import | remove, Bundles/Modules : list | install | uninstall | start | stop
  -file string
        Site zip path to import or module JAR file path to install
  -id string
        Fully qualified module name : <group>:<id>:<version>
  -output string
        Export destination file (default "export.zip")
  -password string
        User password, can be set with environment variable JAHIA_PASSWORD
  -siteKey string
        Jahia Site key, can be set with environment variable JAHIA_SITEKEY
  -start
        Start module after install (default true)
  -url string
        Jahia base URL, can be set with environment variable JAHIA_URL
  -user string
        Jahia user (with admin credentials), can be set with environment variable JAHIA_USER
 ```

## Licence

 See [LICENSE file](./LICENSE)