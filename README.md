# Jahia Tools

Administrative tool for Jahia Digital Experience

- Export site
- List installed bundles
- Deploy a bundle
- Start a bundle
- Stop a bundle
 
## Usage

Command line arguments :

```
  -action string
        Action : list | export | install | uninstall | start | stop
  -file string
        Module JAR file path to install
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