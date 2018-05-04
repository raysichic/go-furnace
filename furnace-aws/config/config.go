package config

import (
	"errors"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"plugin"

	"gopkg.in/yaml.v2"

	"strings"

	"github.com/Skarlso/go-furnace/config"
	"github.com/Skarlso/go-furnace/handle"
)

// TODO: Create a main config which defines a furnace config location
// This was, when running ./furnace-aws create asdf -> it would look for asdf
// as a configuration file. Like asdf_furnace_config.yaml

// Configuration object with all the properties that AWS needs.
type Configuration struct {
	Main struct {
		Stackname string `yaml:"stackname"`
		Spinner   int    `yaml:"spinner"`
		Plugins   struct {
			EnablePluginSystem bool     `yaml:"enable_plugin_system"`
			PluginPath         string   `yaml:"plugin_path"`
			Names              []string `yaml:"names"`
		} `yaml:"plugins"`
	} `yaml:"main"`
	Aws struct {
		CodeDeployRole string `yaml:"code_deploy_role"`
		Region         string `yaml:"region"`
		TemplateName   string `yaml:"template_name"`
		AppName        string `yaml:"app_name"`
		CodeDeploy     struct {
			S3Bucket    string `yaml:"code_deploy_s3_bucket,omitempty"`
			S3Key       string `yaml:"code_deploy_s3_key,omitempty"`
			GitAccount  string `yaml:"git_account,omitempty"`
			GitRevision string `yaml:"git_revision,omitempty"`
		} `yaml:"code_deploy"`
	} `yaml:"aws"`
}

const (
	// PRECREATE Event name for plugins
	PRECREATE = "pre_create"
	// POSTCREATE Event name for plugins
	POSTCREATE = "post_create"
	// PREDELETE Event name for plugins
	PREDELETE = "pre_delete"
	// POSTDELETE Event name for plugins
	POSTDELETE = "post_delete"
)

// Config is the loaded configuration entity.
var Config Configuration

// RunPlugin is a plugin to execute
type RunPlugin struct {
	Run  interface{}
	Name string
}

// PluginRegistry is a registry of plugins for certain events
var PluginRegistry = make(map[string][]RunPlugin)

var configPath string
var templatePath string

var defaultConfig = "aws_furnace_config.yaml"

func init() {
	configPath = config.Path()
	defaultConfigPath := filepath.Join(configPath, defaultConfig)
	Config.LoadConfiguration(defaultConfigPath)
	templatePath = filepath.Join(configPath, Config.Aws.TemplateName)
	FillRegistry()
}

// LoadConfiguration loads a yaml file which sets fields for Configuration struct
func (c *Configuration) LoadConfiguration(configFile string) {
	if _, err := os.Stat(configFile); err != nil {
		if os.IsNotExist(err) {
			log.Println("main configuration file does not exist. Moving on assuming a new will be defined.")
			return
		}
	}
	content, err := ioutil.ReadFile(configFile)
	handle.Error(err)
	err = yaml.Unmarshal(content, c)
	handle.Error(err)
}

// LoadConfigFileIfExists Search backwards from the current directory
// for a furnace config file with the given prefix of `file`. If found,
// the Configuration `Config` will be loaded with values gathered from
// the file described by that config. If none is found, nothing happens.
// The default file remains loaded.
//
// returns an error if the file is not found.
func LoadConfigFileIfExists(dir string, file string) error {
	separatorIndex := strings.LastIndex(dir, "/")
	for separatorIndex != 0 {
		if _, err := os.Stat(filepath.Join(dir, "."+file+".furnace")); err == nil {
			configLocation, _ := ioutil.ReadFile(filepath.Join(dir, "."+file+".furnace"))
			configPath = dir
			Config.LoadConfiguration(filepath.Join(configPath, string(configLocation)))
			templateBase := path.Dir(string(configLocation))
			templatePath = filepath.Join(configPath, templateBase, Config.Aws.TemplateName)
			return nil
		}
		separatorIndex = strings.LastIndex(dir, string(os.PathSeparator))
		dir = dir[0:separatorIndex]
	}

	return errors.New("failed to find configuration file for stack " + file)
}

// FillRegistry fill load in all the configured plugins.
func FillRegistry() {
	if !Config.Main.Plugins.EnablePluginSystem {
		return
	}
	log.Println("Filling plugin registry.")
	files := make([]string, 0)
	for _, f := range Config.Main.Plugins.Names {
		files = append(files, filepath.Join(Config.Main.Plugins.PluginPath, f))
	}
	pluginCount := 0
	for _, f := range files {
		baseName := filepath.Base(f)
		split := strings.Split(baseName, ".")
		key := split[len(split)-1]
		fullPath := filepath.Join(configPath, "plugins", baseName)
		p, err := plugin.Open(fullPath)
		if err != nil {
			log.Printf("Plugin '%s' failed to load. Error: %s\n", fullPath, err.Error())
			continue
		}
		run, err := p.Lookup("RunPlugin")
		if err != nil {
			log.Printf("Plugin '%s' did not have 'RunPlugin' method. Error: %s\n", fullPath, err.Error())
			continue
		}
		plug := RunPlugin{
			Run:  run,
			Name: baseName,
		}
		if p, ok := PluginRegistry[key]; ok {
			p = append(p, plug)
			PluginRegistry[key] = p
		} else {
			plugs := make([]RunPlugin, 0)
			plugs = append(plugs, plug)
			PluginRegistry[key] = plugs
		}
		pluginCount++
	}
	log.Printf("'%d' plugins loaded successfully.\n", pluginCount)
}

// LoadCFStackConfig Load the CF stack configuration file into a []byte.
func LoadCFStackConfig() []byte {
	dat, err := ioutil.ReadFile(templatePath)
	handle.Error(err)
	return dat
}
