package console_tailed_logs_outputter

import (
	"fmt"
	"time"

	"github.com/cloudfoundry-incubator/lattice/ltc/colors"
	"github.com/cloudfoundry-incubator/lattice/ltc/logs"
	"github.com/cloudfoundry-incubator/lattice/ltc/output"
	"github.com/cloudfoundry/noaa/events"
)

type TailedLogsOutputter interface {
	OutputTailedLogs(appGuid string)
	StopOutputting()
}

type ConsoleTailedLogsOutputter struct {
	outputChan chan string
	output     *output.Output
	logReader  logs.LogReader
}

func NewConsoleTailedLogsOutputter(output *output.Output, logReader logs.LogReader) *ConsoleTailedLogsOutputter {
	return &ConsoleTailedLogsOutputter{
		outputChan: make(chan string, 10),
		output:     output,
		logReader:  logReader,
	}

}

func (ctlo *ConsoleTailedLogsOutputter) OutputTailedLogs(appGuid string) {
	go ctlo.logReader.TailLogs(appGuid, ctlo.logCallback, ctlo.errorCallback)

	for log := range ctlo.outputChan {
		ctlo.output.Say(log + "\n")
	}
}

func (ctlo *ConsoleTailedLogsOutputter) StopOutputting() {
	ctlo.logReader.StopTailing()
}

func (ctlo *ConsoleTailedLogsOutputter) logCallback(log *events.LogMessage) {
	timeString := time.Unix(0, log.GetTimestamp()).Format("02 Jan 15:04")
	logOutput := fmt.Sprintf("%s [%s|%s] %s", colors.Cyan(timeString), colors.Yellow(log.GetSourceType()), colors.Yellow(log.GetSourceInstance()), log.GetMessage())
	ctlo.outputChan <- logOutput
}

func (ctlo *ConsoleTailedLogsOutputter) errorCallback(err error) {
	ctlo.outputChan <- err.Error()
}
