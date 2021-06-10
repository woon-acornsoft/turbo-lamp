
import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

type Conf struct {
	mode string
	db   Postgres
}

type Postgres struct {
	address  string
	user     string
	password string
	database string
}

type Log struct {
	log logrus.logger
}

func initialize() (*Log, *Conf) {
	var log = logrus.New()
	logger := Log{log}
	// logger.Debug("initialize")
	conf := loadConfig()
	return &logger, conf
}

func loadConfig() *Conf {

	mode := os.Getenv("PROJECT_MODE")
	postgresYamlPath := os.Getenv("PROJECT_CONFIG_PATH")
	filename, _ := filepath.Abs("postgres.yaml")
	if postgresYamlPath != "" {
		filename, _ = filepath.Abs(postgresYamlPath + "postgres.yaml")

	}
	p := Postgres{}
	yamlFile, err := ioutil.ReadFile(filename)
	if err == nil {
		yaml.Unmarshal(yamlFile, &p)
	}
	c := Conf{mode, Postgres{}}
	return &c
}

type Loggin interface {
	Debug()
}

func (l *Log) Trace() {
	l.log.WithFields(logrus.Fields{
		"animal": "walrus",
		"number": 0,
	}).Trace("Went to the beach")
}

func (l *Log) Debug(c *gin.Context) {
	l.log.WithFields(logrus.Fields{
		"animal": "walrus",
		"number": 8,
	}).Debug("Started observing beach")
}

func (l *Log) Info() {
	l.log.WithFields(logrus.Fields{
		"animal": "walrus",
		"size":   10,
	}).Info("A group of walrus emerges from the ocean")
}

func (l *Log) Warn() {
	l.log.WithFields(logrus.Fields{
		"omg":    true,
		"number": 122,
	}).Warn("The group's number increased tremendously!")
}

func (l *Log) Panic() {
	l.log.WithFields(logrus.Fields{
		"animal": "orca",
		"size":   9009,
	}).Panic("It's over 9000!")
}
