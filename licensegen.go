/*
 * licensegen -- generate a LICENSE file for a project
 * v1.0.0 2015-07-18_01
 *
 *
 */
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/segmentio/go-prompt"
	"io/ioutil"
	"os"
	"path"
	"strings"
	"text/template"
	"time"
)

type Author struct {
	FirstName    string
	LastName     string
	EmailAddress string
	Website      string
}

type License struct {
	Name        string
	LicenseFile string
	HeaderFile  string
}

type Configuration struct {
	Author   Author
	Licenses []License
}

type LicenseTemplateInfo struct {
	Year    string
	Author  Author
	License License
}

const (
	ERR_SUCCESS               = iota
	ERR_CONFIG_FILE_ERROR     = iota
	ERR_LICENSE_FILE_ERROR    = iota
	ERR_HEADER_FILE_ERROR     = iota
	ERR_NO_LICENSE_ARG        = iota
	ERR_LICENSE_FILE_CREATION = iota
	ERR_LICENSE_FILE_EXISTS   = iota
)

var (
	licenseOutputFilename string
	headerOutputFilename  string
	licenseName           string
	listLicenses          bool
	verboseMode           bool
)

func init() {
	//configure default filenames here
	flag.StringVar(&licenseOutputFilename, "filename", "LICENSE", "")
	flag.StringVar(&headerOutputFilename, "headerfile", "LICENSE_FILE_HEADER", "")
	flag.StringVar(&licenseName, "license", "", "")
	flag.BoolVar(&listLicenses, "list", false, "help message")
	flag.BoolVar(&verboseMode, "verbose", false, "help message")
}

func promptUser(strPrompt string) string {
	ret := prompt.String(strPrompt)
	return ret
}

func promptUserString(prompt string, args ...interface{}) string {
	var s string
	fmt.Printf(prompt+": ", args...)
	fmt.Scanln(&s)
	return s
}

func infoMessage(sFmtStr string, args ...interface{}) bool {
	fmt.Fprintf(os.Stdout, sFmtStr, args)
	return true
}

func errorMessage(sFmtStr string, exitProcess bool, exitCode int, args ...interface{}) int {
	fmt.Fprintf(os.Stderr, "Error: "+sFmtStr, args...)
	if exitProcess {
		os.Exit(exitCode)
	}
	return 0
}

func fileExists(fileName string) bool {
	_, err := os.Stat(fileName)
	if err != nil {
		return false
	}
	return true
}

func successful(arg interface{}) bool {
	if arg != nil {
		return false
	}
	return true
}

func main() {

	flag.Parse()

	var licenseText string
	//thisApp := path.Base(os.Args[0])
	thisPath := path.Dir(os.Args[0])

	if len(os.Args) == 1 {
		errorMessage("Please specify a valid license name.  Use --list to list valid licenses.", true, ERR_NO_LICENSE_ARG, "")
	}

	//if no flags are given and only one arg is passed, assume it's the name
	//of the license to generate
	if len(os.Args) == 2 {
		licenseName = os.Args[1]
	}

	selectedLicense := licenseName
	//if the only arg passed was 'list', or --list was specified,
	//output a list of valid licenses
	if selectedLicense == "list" || listLicenses {
		listLicenses = true
	}
	licenseIndex := -1
	license := License{"", "", ""}

	//Load configuration from JSON file
	file, _ := os.Open(thisPath + string(os.PathSeparator) + "licensegen.json")
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err := decoder.Decode(&configuration)
	if !successful(err) {
		fmt.Println("Error while loading configuration file: ", err)
		os.Exit(ERR_CONFIG_FILE_ERROR)
	}

	if listLicenses {
		fmt.Println(" Listing Available Licenses...\n --------------------")
	}
	for j := 0; j < len(configuration.Licenses); j++ {
		if strings.ToUpper(selectedLicense) == strings.ToUpper(configuration.Licenses[j].Name) {
			licenseIndex = j
			license = configuration.Licenses[j]
			break
		}
		if listLicenses {
			fmt.Printf("  %d) %s\n", j+1, configuration.Licenses[j].Name)
		}
	}

	if listLicenses {
		fmt.Println("")
		fmt.Println("Use --license [license_name] to specify the license to generate.")
		fmt.Println("")
		os.Exit(ERR_SUCCESS)
	}

	if licenseIndex == -1 {
		fmt.Fprintln(os.Stderr, "Error: Please specify a valid license name.  Use --list to list valid licenses.")
		os.Exit(ERR_NO_LICENSE_ARG)
	}

	if verboseMode {
		fmt.Printf("licenseOutputFilename = %s \n", licenseOutputFilename)
		fmt.Printf("headerOutputFilename = %s \n", headerOutputFilename)
		fmt.Printf("licenseName = %s \n", licenseName)
		fmt.Println("")
	}

	licenseTextB, err := ioutil.ReadFile(thisPath + string(os.PathSeparator) + license.LicenseFile)
	if !successful(err) {
		errorMessage("failed reading license file: %s \n\n", true, ERR_LICENSE_FILE_ERROR, err)
	}

	licenseText = string(licenseTextB)

	headerText, err := ioutil.ReadFile(thisPath + string(os.PathSeparator) + license.HeaderFile)
	if !successful(err) {
		errorMessage("Error reading license header file: %v", true, ERR_LICENSE_FILE_ERROR, err)
	}

	templateFunctionsMap := template.FuncMap{
		"title":  strings.Title,
		"trim":   strings.TrimSpace,
		"prompt": promptUserString,
	}

	//get the current year
	strDate := fmt.Sprintf("%d", time.Now().Year())
	li := LicenseTemplateInfo{strDate, configuration.Author, license}

	tmpl, err := template.New("headerTemplate").Funcs(templateFunctionsMap).Parse(string(headerText))
	if !successful(err) {
		errorMessage("parsing license header template %s failed.", true, ERR_HEADER_FILE_ERROR, string(headerText))
	}

	tmplL, err := template.New("licenseTemplate").Funcs(templateFunctionsMap).Parse(string(licenseText))

	if !successful(err) {
		errorMessage("Error parsing license template: %s", true, ERR_LICENSE_FILE_ERROR, err)
	}

	if fileExists(licenseOutputFilename) {
		if !prompt.Confirm("License file '%s' already exists.  Overwrite [Y/n]? ", licenseOutputFilename) {
			os.Exit(ERR_LICENSE_FILE_EXISTS)
		}
	}

	fileLicenseOutput, err := os.Create(licenseOutputFilename)
	if !successful(err) {
		errorMessage("Error creating license file '%s': %s", true, ERR_LICENSE_FILE_CREATION, licenseOutputFilename, err)
	}

	fileHeaderOutput, err := os.Create(headerOutputFilename)
	if !successful(err) {
		errorMessage("Error creating license header file '%s': %s", true, ERR_LICENSE_FILE_CREATION, headerOutputFilename, err)
	}

	if fileLicenseOutput != nil {
		err = tmplL.Execute(fileLicenseOutput, li)
		if !successful(err) {
			errorMessage("Error executing license template: %s", true, ERR_LICENSE_FILE_ERROR, err)
		}
	}

	err = tmpl.Execute(fileHeaderOutput, li)
	if !successful(err) {
		errorMessage("Error executing license header template: %s", true, ERR_LICENSE_FILE_ERROR, err)
	}

	//fmt.Printf("Finished generating license, saved to file '%s'.\n", licenseOutputFilename)
	infoMessage("Finished generating license, saved to file '%s'.\n", licenseOutputFilename)
	os.Exit(ERR_SUCCESS)
}
