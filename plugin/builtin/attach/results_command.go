package attach

import (
	"10gen.com/mci/model"
	"10gen.com/mci/plugin"
	"10gen.com/mci/util"
	"fmt"
	"github.com/10gen-labs/slogger/v1"
	"github.com/mitchellh/mapstructure"
	"os"
	"path/filepath"
)

// AttachResultsCommand is used to attach MCI test results in json
// format to the task page.
type AttachResultsCommand struct {
	// FileLoc describes the relative path of the file to be sent.
	// Note that this can also be described via expansions.
	FileLoc string `mapstructure:"file_location" plugin:"expand"`
}

func (self *AttachResultsCommand) Name() string {
	return AttachResultsCmd
}

func (self *AttachResultsCommand) Plugin() string {
	return AttachPluginName
}

// ParseParams decodes the S3 push command parameters that are
// specified as part of an AttachPlugin command; this is required
// to satisfy the 'Command' interface
func (self *AttachResultsCommand) ParseParams(params map[string]interface{}) error {
	if err := mapstructure.Decode(params, self); err != nil {
		return fmt.Errorf("error decoding '%v' params: %v", self.Name(), err)
	}
	if err := self.validateAttachResultsParams(); err != nil {
		return fmt.Errorf("error validating '%v' params: %v", self.Name(), err)
	}
	return nil
}

// validateAttachResultsParams is a helper function that ensures all
// the fields necessary for attaching a results are present
func (self *AttachResultsCommand) validateAttachResultsParams() (err error) {
	if self.FileLoc == "" {
		return fmt.Errorf("file_location cannot be blank")
	}
	return nil
}

func (self *AttachResultsCommand) expandAttachResultsParams(
	taskConfig *model.TaskConfig) (err error) {
	self.FileLoc, err = taskConfig.Expansions.ExpandString(self.FileLoc)
	if err != nil {
		return fmt.Errorf("error expanding file_location: %v", err)
	}
	return nil
}

// Execute carries out the AttachResultsCommand command - this is required
// to satisfy the 'Command' interface
func (self *AttachResultsCommand) Execute(pluginLogger plugin.Logger,
	pluginCom plugin.PluginCommunicator,
	taskConfig *model.TaskConfig,
	stop chan bool) error {

	if err := self.expandAttachResultsParams(taskConfig); err != nil {
		return err
	}

	errChan := make(chan error)
	go func() {
		// attempt to open the file
		reportFileLoc := filepath.Join(taskConfig.WorkDir, self.FileLoc)
		reportFile, err := os.Open(reportFileLoc)
		if err != nil {
			errChan <- fmt.Errorf("Couldn't open report file: '%v'", err)
			return
		}
		results := &model.TestResults{}
		if err = util.ReadJSONInto(reportFile, results); err != nil {
			errChan <- fmt.Errorf("Couldn't read report file: '%v'", err)
			return
		}
		if err := reportFile.Close(); err != nil {
			pluginLogger.LogExecution(slogger.INFO, "Error closing file: %v", err)
		}
		errChan <- SendJSONResults(taskConfig, pluginLogger, pluginCom, results)
	}()

	select {
	case err := <-errChan:
		return err
	case <-stop:
		pluginLogger.LogExecution(slogger.INFO, "Received signal to terminate"+
			" execution of attach results command")
		return nil
	}
}

// SendJSONResults is responsible for sending the
// specified file to the API Server
func SendJSONResults(taskConfig *model.TaskConfig,
	pluginLogger plugin.Logger, pluginCom plugin.PluginCommunicator,
	results *model.TestResults) error {

	pluginLogger.LogExecution(slogger.INFO, "Attaching test results")
	err := pluginCom.TaskPostResults(results)
	if err != nil {
		return err
	}

	pluginLogger.LogTask(slogger.INFO, "Attach test results succeeded")
	return nil
}
