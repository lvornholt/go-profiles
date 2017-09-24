package profile

import (
	"io/ioutil"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
	syaml "github.com/smallfish/simpleyaml"
)

// default profile dev
var profile = "default"
var configFolderPath = "./"
var data []byte
var yData *syaml.Yaml
var logLevel = log.InfoLevel

func init() {
	loadProfile()
}

// ClearData to reinitilize profile
func ClearData() {
	yData = nil
}

// SetLogLevel - allowed values "debug", "info", "warn", "error"
func SetLogLevel(level string) {
	if level == "debug" {
		logLevel = log.DebugLevel
	}
	if level == "info" {
		logLevel = log.InfoLevel
	}
	if level == "warn" {
		logLevel = log.WarnLevel
	}
	if level == "error" {
		logLevel = log.ErrorLevel
	}
}

func loadProfile() {

	log.SetLevel(logLevel)

	if envConfigFolderPath := os.Getenv("CONFIG_FOLDER_PATH"); len(envConfigFolderPath) > 0 {
		configFolderPath = envConfigFolderPath
		log.WithField("configPath", configFolderPath).Info("ProfileService: ConfigFolderPath set by enironment variable to")
	}

	if envProfile := os.Getenv("PROFILE"); len(envProfile) > 0 {
		profile = envProfile
		log.WithField("profile", profile).Info("ProfileService: started with profile")
	} else {
		log.Info("ProfileService: env variable PROFILE not set - application started with default profile")
	}

	var profileFileName string
	if len(profile) > 0 && profile != "default" {
		profileFileName = "application-" + profile + ".yml"
	} else {
		profileFileName = "application.yml"
		log.Info("ProfileService: Fallback to default profile file")
	}

	var err error
	data, err = ioutil.ReadFile(configFolderPath + profileFileName)
	if err != nil {
		log.WithError(err).WithFields(log.Fields{"profile": profile}).Error("ProfileService: application file for profile not found")
	}

	yData, err = syaml.NewYaml(data)
	if err != nil {
		log.WithError(err).Error("ProfileService: could not parse profile yaml file")
	}

}

// GetValueWithDefault - search for string at given path
// and return given default value if no or emtpy value found
func GetValueWithDefault(path string, defaultValue interface{}) interface{} {
	var result interface{}
	var err error

	switch v := defaultValue.(type) {
	case int:
		result, err = getValue(path).Int()
	case float64:
		result, err = getValue(path).Float()
	case string:
		result, err = getValue(path).String()
	default:
		result = v
	}
	if err != nil {
		log.WithError(err).Debug("could not find value for path " + path)
		return defaultValue
	}
	return result
}

// GetIntValue - search for integer value at given path
// and return 0 if no value found
func GetIntValue(path string) int {
	result, err := getValue(path).Int()
	if err != nil {
		log.WithError(err).Debug("problem to extract value for path: " + path)
		return 0
	}
	return result
}

// GetStringValue - search for string value at given path
// and return empty string if no value found
func GetStringValue(path string) string {
	result, err := getValue(path).String()
	if err != nil {
		log.WithError(err).Debug("problem to extract value for path: " + path)
		return ""
	}
	return result
}

// GetArrayValues search for array at given path
// and return empty array if no values found
func GetArrayValues(path string) []interface{} {
	result, err := getValue(path).Array()
	if err != nil {
		log.WithError(err).Debug("problem to extract value for path: " + path)
		var emptyArray []interface{}
		return emptyArray
	}
	return result
}

// GetFloatValue - search for float value at given path
// and return 0.0 if no value found
func GetFloatValue(path string) float64 {
	result, err := getValue(path).Float()
	if err != nil {
		log.WithError(err).Debug("problem to extract value for path: " + path)
		return 0.0
	}
	return result
}

// GetBooleanValue searc for boolean value at given path
// and return false if no value found
func GetBooleanValue(path string) bool {
	result, err := getValue(path).Bool()
	if err != nil {
		log.WithError(err).Debug("problem to extract value for path: " + path)
		return false
	}
	return result
}

func getValue(path string) *syaml.Yaml {
	pathPatrs := strings.Split(path, ".")

	if yData == nil {
		loadProfile()
	}

	var current = yData

	for i := range pathPatrs {
		current = current.Get(pathPatrs[i])
	}

	return current
}
