package transsros

import (
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/sirupsen/logrus"
)

var (
	unFlatConf, flatConf, flatLine []string
	indent                         string = "    " // Nokia SROS has 4 spaces indentation
	createFlag                     bool            // Flag to mark line with keywords "create" and "interface" in it
	idx, prevIdx                   int
)

func Transcode(file string, short bool, debug bool) []string {

	if debug {
		DEBUG := logrus.DebugLevel
		logrus.SetLevel(DEBUG)
		log.SetOutput(os.Stdout)
	}

	// List of Regex
	reComment, _ := regexp.Compile(`^#`)
	reEcho, _ := regexp.Compile(`^echo`)
	reBlank, _ := regexp.Compile(`^\s*$`)
	reExitall, _ := regexp.Compile(`exit all`)
	reCreate, _ := regexp.Compile(`create`)
	reConfigure, _ := regexp.Compile(`^configure`)
	reService, _ := regexp.Compile(`^/configure service (?:vprn|ies)\s\d+\sname`)
	reServiceName, _ := regexp.Compile(`^(vprn|ies) (\d+) (?:name "(?:.*)" customer \d+)`)

	// Read a Nokia SROS configuratrion file if DOS file remove \r
	bytesRead, _ := ioutil.ReadFile(file)
	file_content := string(bytesRead)
	file_content_unix := strings.Replace(file_content, "\r\n", "\n", -1)
	lines := strings.Split(file_content_unix, "\n")

	// Remove unnecessary lines from config and store result in unFlatConf
	for _, line := range lines {
		comment := reComment.MatchString(line)
		echo := reEcho.MatchString(line)
		exitAll := reExitall.MatchString(line)
		blank := reBlank.MatchString(line)
		//exit := reExit.MatchString(line)

		if comment || echo || exitAll || blank {
			continue
		}
		unFlatConf = append(unFlatConf, line)
	}

	// Start "falttening" the config
	for _, line := range unFlatConf {
		if debug {
			logrus.Debug("Current line:", strings.Join(flatLine, " "))
			logrus.Debug("New line: ", line)
			logrus.Debug("Flag: ", createFlag)
		}

		// Count number of heading indentation
		idx = strings.Count(line, indent)

		if reConfigure.MatchString(line) {
			flatLine = append(flatLine, "/"+strings.TrimLeft(line, " "))
			continue
		}

		if strings.TrimSpace(line) == "exit" {
			continue
		}

		if idx <= prevIdx {
			// if "interface" and "create" in conf line don't save it another time
			if !createFlag {
				flatConf = append(flatConf, strings.Join(flatLine, " "))
				logrus.Debug("Save: ", strings.Join(flatLine, " "))
			}
			createFlag = false
		}

		flatLine = append(flatLine[:idx], strings.TrimLeft(line, " "))

		// create word needs to be remove for successive line
		// Also dealing with double entries for interface create
		if reCreate.MatchString(line) {
			if !createFlag {
				flatConf = append(flatConf, strings.Join(flatLine, " "))
				logrus.Debug("Save: ", strings.Join(flatLine, " "))
			}
			flatLine[len(flatLine)-1] = strings.Replace(flatLine[len(flatLine)-1], " create", "", -1)
			if strings.Contains(line, "interface") {
				createFlag = true
			}

			if short {
				logrus.Debug("Debug: ", flatLine)
				if reService.MatchString(strings.Join(flatLine, " ")) {
					parts := reServiceName.FindStringSubmatch(flatLine[2])
					flatLine[2] = parts[1] + " " + parts[2]
					logrus.Debug("Debug: ", flatLine)
				}
			}
		}

		prevIdx = idx
	}

	return flatConf

}
