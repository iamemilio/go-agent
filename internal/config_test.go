package internal

import (
	"testing"

	"github.com/newrelic/go-sdk/api"
	"github.com/newrelic/go-sdk/internal/utilization"
)

func TestCopyConfigReferenceFieldsPresent(t *testing.T) {
	cfg := api.NewConfig("my appname", "0123456789012345678901234567890123456789")
	cfg.Labels["zip"] = "zap"
	cfg.ErrorCollector.IgnoreStatusCodes = append(cfg.ErrorCollector.IgnoreStatusCodes, 405)
	cfg.Attributes.Include = append(cfg.Attributes.Include, "1")
	cfg.Attributes.Exclude = append(cfg.Attributes.Exclude, "2")
	cfg.TransactionEvents.Attributes.Include = append(cfg.TransactionEvents.Attributes.Include, "3")
	cfg.TransactionEvents.Attributes.Exclude = append(cfg.TransactionEvents.Attributes.Exclude, "4")
	cfg.ErrorCollector.Attributes.Include = append(cfg.ErrorCollector.Attributes.Include, "5")
	cfg.ErrorCollector.Attributes.Exclude = append(cfg.ErrorCollector.Attributes.Exclude, "6")

	cp := copyConfigReferenceFields(cfg)

	cfg.Labels["zop"] = "zup"
	cfg.ErrorCollector.IgnoreStatusCodes[0] = 201
	cfg.Attributes.Include[0] = "zap"
	cfg.Attributes.Exclude[0] = "zap"
	cfg.TransactionEvents.Attributes.Include[0] = "zap"
	cfg.TransactionEvents.Attributes.Exclude[0] = "zap"
	cfg.ErrorCollector.Attributes.Include[0] = "zap"
	cfg.ErrorCollector.Attributes.Exclude[0] = "zap"

	expect := compactJSONString(`[
	{
		"pid":123,
		"language":"go",
		"agent_version":"0.2.2",
		"host":"my-hostname",
		"settings":{
			"AppName":"my appname",
			"Attributes":{"Enabled":true,"Exclude":["2"],"Include":["1"]},
			"CustomInsightsEvents":{"Enabled":true},
			"Development":false,
			"ErrorCollector":{
				"Attributes":{"Enabled":true,"Exclude":["6"],"Include":["5"]},
				"CaptureEvents":true,
				"Enabled":true,
				"IgnoreStatusCodes":[404,405]
			},
			"HighSecurity":false,
			"HostDisplayName":"",
			"Labels":{"zip":"zap"},
			"TransactionEvents":{
				"Attributes":{"Enabled":true,"Exclude":["4"],"Include":["3"]},
				"Enabled":true
			},
			"Transport":null,
			"UseTLS":true,
			"Utilization":{"DetectAWS":true,"DetectDocker":true}
		},
		"app_name":["my appname"],
		"high_security":false,
		"labels":[{"label_type":"zip","label_value":"zap"}],
		"environment":[["Compiler","comp"],["GOARCH","arch"],["GOOS","goos"],["Version","vers"]],
		"identifier":"my appname",
		"utilization":{
			"metadata_version":1,
			"logical_processors":16,
			"total_ram_mib":1024,
			"hostname":"my-hostname"
		}
	}]`)

	js, err := configConnectJSONInternal(&cp, 123, &utilization.SampleData, sampleEnvironment, "0.2.2")
	if nil != err {
		t.Fatal(err)
	}
	if string(js) != expect {
		t.Error(string(js))
	}
}

func TestCopyConfigReferenceFieldsAbsent(t *testing.T) {
	cfg := api.NewConfig("my appname", "0123456789012345678901234567890123456789")
	cfg.Labels = nil
	cfg.ErrorCollector.IgnoreStatusCodes = nil

	cp := copyConfigReferenceFields(cfg)

	expect := compactJSONString(`[
	{
		"pid":123,
		"language":"go",
		"agent_version":"0.2.2",
		"host":"my-hostname",
		"settings":{
			"AppName":"my appname",
			"Attributes":{"Enabled":true,"Exclude":null,"Include":null},
			"CustomInsightsEvents":{"Enabled":true},
			"Development":false,
			"ErrorCollector":{
				"Attributes":{"Enabled":true,"Exclude":null,"Include":null},
				"CaptureEvents":true,
				"Enabled":true,
				"IgnoreStatusCodes":null
			},
			"HighSecurity":false,
			"HostDisplayName":"",
			"Labels":null,
			"TransactionEvents":{
				"Attributes":{"Enabled":true,"Exclude":null,"Include":null},
				"Enabled":true
			},
			"Transport":null,
			"UseTLS":true,
			"Utilization":{"DetectAWS":true,"DetectDocker":true}
		},
		"app_name":["my appname"],
		"high_security":false,
		"environment":[["Compiler","comp"],["GOARCH","arch"],["GOOS","goos"],["Version","vers"]],
		"identifier":"my appname",
		"utilization":{
			"metadata_version":1,
			"logical_processors":16,
			"total_ram_mib":1024,
			"hostname":"my-hostname"
		}
	}]`)

	js, err := configConnectJSONInternal(&cp, 123, &utilization.SampleData, sampleEnvironment, "0.2.2")
	if nil != err {
		t.Fatal(err)
	}
	if string(js) != expect {
		t.Error(string(js))
	}
}