package store

import (
	"fmt"
	"github.com/gocarina/gocsv"
	log "github.com/sirupsen/logrus"
	"loinc-lambda/pkg/model"
	"os"
	"runtime"
	"strings"
	"time"
)

var data = map[string]*model.Loinc{}
var name string

func init() {

	then := time.Now()

	if _, file, _, ok := runtime.Caller(0); ok {
		paths := strings.Split(file, "/")
		name = strings.Split(paths[len(paths)-1], ".")[0]
	}

	if path, err := os.Getwd(); err != nil {
		log.WithError(err).Fatal()
	} else if p := strings.Split(path, "/"); p[len(p)-1] == "cmd" {
		if err = os.Chdir("../pkg/store"); err != nil {
			log.WithError(err).Fatal()
		}
	} else if p[len(p)-1] != name {
		if err = os.Chdir("pkg/store"); err != nil {
			log.WithError(err).Fatal()
		}
	}

	var file *os.File
	var err error
	if file, err = os.Open("LoincTableCore.csv"); err != nil {
		log.WithError(err).Fatal()
	}

	defer func(file *os.File) {
		if err = file.Close(); err != nil {
			log.WithError(err).Fatal()
		}
	}(file)

	var loincArray []*model.Loinc
	if err = gocsv.UnmarshalFile(file, &loincArray); err != nil {
		log.WithError(err).Fatal()
	}

	for _, loinc := range loincArray {
		loinc.ClassType = decodeClassType(loinc.ClassType)
		data[loinc.ID] = loinc
	}

	log.Infof(`%s initialized in %v`, name, time.Now().Sub(then).Truncate(time.Millisecond))
}

func decodeClassType(ID string) string {
	switch ID {
	case "1":
		return "Laboratory class"
	case "2":
		return "Clinical class"
	case "3":
		return "Claims attachments"
	case "4":
		return "Surveys"
	default:
		log.Warnf("🟡️ -> init loinc store -> no class type found for [%s]", ID)
		return ""
	}
}

func Get(ID string) (*model.Loinc, error) {
	if _, ok := data[ID]; ok {
		return data[ID], nil
	}
	return nil, fmt.Errorf("no loinc code found for ID [%s]", ID)
}
