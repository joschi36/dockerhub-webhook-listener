package listener

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Rancher struct {
	rancherConfig
}

type rancherConfig struct {
	Url       string
	UserKey   string
	SecretKey string
	Service   string
}

type rancherService struct {
	Actions        struct{} `json:"actions"`
	CreateDefaults struct{} `json:"createDefaults"`
	CreateTypes    struct {
		DNSService          string `json:"dnsService"`
		ExternalService     string `json:"externalService"`
		LoadBalancerService string `json:"loadBalancerService"`
		Service             string `json:"service"`
	} `json:"createTypes"`
	Data []struct {
		AccountID string `json:"accountId"`
		Actions   struct {
			Addservicelink    string `json:"addservicelink"`
			Deactivate        string `json:"deactivate"`
			Remove            string `json:"remove"`
			Removeservicelink string `json:"removeservicelink"`
			Restart           string `json:"restart"`
			Setservicelinks   string `json:"setservicelinks"`
			Update            string `json:"update"`
			Upgrade           string `json:"upgrade"`
			FinishUpgrade			string `json:"finishupgrade"`
		} `json:"actions"`
		AssignServiceIPAddress bool        `json:"assignServiceIpAddress"`
		CreateIndex            int         `json:"createIndex"`
		Created                string      `json:"created"`
		CreatedTS              int         `json:"createdTS"`
		Description            interface{} `json:"description"`
		EnvironmentID          string      `json:"environmentId"`
		ExternalID             interface{} `json:"externalId"`
		Fqdn                   interface{} `json:"fqdn"`
		HealthState            string      `json:"healthState"`
		ID                     string      `json:"id"`
		Kind                   string      `json:"kind"`
		LaunchConfig           struct {
			CapAdd                       []interface{} `json:"capAdd"`
			CapDrop                      []interface{} `json:"capDrop"`
			Count                        interface{}   `json:"count"`
			CPUSet                       interface{}   `json:"cpuSet"`
			CPUShares                    interface{}   `json:"cpuShares"`
			DataVolumes                  []interface{} `json:"dataVolumes"`
			DataVolumesFrom              []interface{} `json:"dataVolumesFrom"`
			DataVolumesFromLaunchConfigs []interface{} `json:"dataVolumesFromLaunchConfigs"`
			Description                  interface{}   `json:"description"`
			Devices                      []interface{} `json:"devices"`
			DNS                          []interface{} `json:"dns"`
			DNSSearch                    []interface{} `json:"dnsSearch"`
			DomainName                   interface{}   `json:"domainName"`
			Environment                  struct {
				APIHOST string `json:"API_HOST"`
			} `json:"environment"`
			Hostname  interface{} `json:"hostname"`
			ImageUUID string      `json:"imageUuid"`
			Kind      string      `json:"kind"`
			Labels    struct {
				Io_rancher_container_pullImage                     string `json:"io.rancher.container.pull_image"`
				Io_rancher_scheduler_affinity_containerLabelSoftNe string `json:"io.rancher.scheduler.affinity:container_label_soft_ne"`
			} `json:"labels"`
			LogConfig struct {
				Config struct{} `json:"config"`
				Driver string   `json:"driver"`
			} `json:"logConfig"`
			Memory              interface{}   `json:"memory"`
			MemoryMb            interface{}   `json:"memoryMb"`
			MemorySwap          interface{}   `json:"memorySwap"`
			NetworkLaunchConfig interface{}   `json:"networkLaunchConfig"`
			NetworkMode         string        `json:"networkMode"`
			PidMode             interface{}   `json:"pidMode"`
			Ports               []interface{} `json:"ports"`
			Privileged          bool          `json:"privileged"`
			PublishAllPorts     bool          `json:"publishAllPorts"`
			ReadOnly            bool          `json:"readOnly"`
			RequestedHostID     interface{}   `json:"requestedHostId"`
			RequestedIPAddress  interface{}   `json:"requestedIpAddress"`
			StartOnCreate       bool          `json:"startOnCreate"`
			StdinOpen           bool          `json:"stdinOpen"`
			Tty                 bool          `json:"tty"`
			User                interface{}   `json:"user"`
			Userdata            interface{}   `json:"userdata"`
			Vcpu                int           `json:"vcpu"`
			Version             string        `json:"version"`
			VolumeDriver        interface{}   `json:"volumeDriver"`
			WorkingDir          interface{}   `json:"workingDir"`
		} `json:"launchConfig"`
		Links struct {
			Account            string `json:"account"`
			Consumedbyservices string `json:"consumedbyservices"`
			Consumedservices   string `json:"consumedservices"`
			ContainerStats     string `json:"containerStats"`
			Environment        string `json:"environment"`
			Instances          string `json:"instances"`
			Self               string `json:"self"`
			ServiceExposeMaps  string `json:"serviceExposeMaps"`
		} `json:"links"`
		Metadata               interface{}   `json:"metadata"`
		Name                   string        `json:"name"`
		PublicEndpoints        interface{}   `json:"publicEndpoints"`
		Removed                interface{}   `json:"removed"`
		RetainIP               interface{}   `json:"retainIp"`
		Scale                  int           `json:"scale"`
		SecondaryLaunchConfigs []interface{} `json:"secondaryLaunchConfigs"`
		SelectorContainer      interface{}   `json:"selectorContainer"`
		SelectorLink           interface{}   `json:"selectorLink"`
		StartOnCreate          bool          `json:"startOnCreate"`
		State                  string        `json:"state"`
		Transitioning          string        `json:"transitioning"`
		TransitioningMessage   interface{}   `json:"transitioningMessage"`
		TransitioningProgress  interface{}   `json:"transitioningProgress"`
		Type                   string        `json:"type"`
		Upgrade                struct {
			InServiceStrategy struct {
				BatchSize      int `json:"batchSize"`
				IntervalMillis int `json:"intervalMillis"`
				LaunchConfig   struct {
					CapAdd                       []interface{} `json:"capAdd"`
					CapDrop                      []interface{} `json:"capDrop"`
					Count                        interface{}   `json:"count"`
					CPUSet                       interface{}   `json:"cpuSet"`
					CPUShares                    interface{}   `json:"cpuShares"`
					DataVolumes                  []interface{} `json:"dataVolumes"`
					DataVolumesFrom              []interface{} `json:"dataVolumesFrom"`
					DataVolumesFromLaunchConfigs []interface{} `json:"dataVolumesFromLaunchConfigs"`
					Description                  interface{}   `json:"description"`
					Devices                      []interface{} `json:"devices"`
					DNS                          []interface{} `json:"dns"`
					DNSSearch                    []interface{} `json:"dnsSearch"`
					DomainName                   interface{}   `json:"domainName"`
					Environment                  struct {
						APIHOST string `json:"API_HOST"`
					} `json:"environment"`
					Hostname  interface{} `json:"hostname"`
					ImageUUID string      `json:"imageUuid"`
					Kind      string      `json:"kind"`
					Labels    struct {
						Io_rancher_container_pullImage                     string `json:"io.rancher.container.pull_image"`
						Io_rancher_scheduler_affinity_containerLabelSoftNe string `json:"io.rancher.scheduler.affinity:container_label_soft_ne"`
					} `json:"labels"`
					LogConfig struct {
						Config struct{} `json:"config"`
						Driver string   `json:"driver"`
					} `json:"logConfig"`
					Memory              interface{}   `json:"memory"`
					MemoryMb            interface{}   `json:"memoryMb"`
					MemorySwap          interface{}   `json:"memorySwap"`
					NetworkLaunchConfig interface{}   `json:"networkLaunchConfig"`
					NetworkMode         string        `json:"networkMode"`
					PidMode             interface{}   `json:"pidMode"`
					Ports               []interface{} `json:"ports"`
					Privileged          bool          `json:"privileged"`
					PublishAllPorts     bool          `json:"publishAllPorts"`
					ReadOnly            bool          `json:"readOnly"`
					RequestedHostID     interface{}   `json:"requestedHostId"`
					RequestedIPAddress  interface{}   `json:"requestedIpAddress"`
					StartOnCreate       bool          `json:"startOnCreate"`
					StdinOpen           bool          `json:"stdinOpen"`
					Tty                 bool          `json:"tty"`
					User                interface{}   `json:"user"`
					Userdata            interface{}   `json:"userdata"`
					Vcpu                int           `json:"vcpu"`
					Version             string        `json:"version"`
					VolumeDriver        interface{}   `json:"volumeDriver"`
					WorkingDir          interface{}   `json:"workingDir"`
				} `json:"launchConfig"`
				PreviousLaunchConfig struct {
					CapAdd                       []interface{} `json:"capAdd"`
					CapDrop                      []interface{} `json:"capDrop"`
					Count                        interface{}   `json:"count"`
					CPUSet                       interface{}   `json:"cpuSet"`
					CPUShares                    interface{}   `json:"cpuShares"`
					DataVolumes                  []interface{} `json:"dataVolumes"`
					DataVolumesFrom              []interface{} `json:"dataVolumesFrom"`
					DataVolumesFromLaunchConfigs []interface{} `json:"dataVolumesFromLaunchConfigs"`
					Description                  interface{}   `json:"description"`
					Devices                      []interface{} `json:"devices"`
					DNS                          []interface{} `json:"dns"`
					DNSSearch                    []interface{} `json:"dnsSearch"`
					DomainName                   interface{}   `json:"domainName"`
					Environment                  struct {
						APIHOST string `json:"API_HOST"`
					} `json:"environment"`
					Hostname  interface{} `json:"hostname"`
					ImageUUID string      `json:"imageUuid"`
					Kind      string      `json:"kind"`
					Labels    struct {
						Io_rancher_container_pullImage                     string `json:"io.rancher.container.pull_image"`
						Io_rancher_scheduler_affinity_containerLabelSoftNe string `json:"io.rancher.scheduler.affinity:container_label_soft_ne"`
					} `json:"labels"`
					LogConfig struct {
						Config struct{} `json:"config"`
						Driver string   `json:"driver"`
					} `json:"logConfig"`
					Memory              interface{}   `json:"memory"`
					MemoryMb            interface{}   `json:"memoryMb"`
					MemorySwap          interface{}   `json:"memorySwap"`
					NetworkLaunchConfig interface{}   `json:"networkLaunchConfig"`
					NetworkMode         string        `json:"networkMode"`
					PidMode             interface{}   `json:"pidMode"`
					Ports               []interface{} `json:"ports"`
					Privileged          bool          `json:"privileged"`
					PublishAllPorts     bool          `json:"publishAllPorts"`
					ReadOnly            bool          `json:"readOnly"`
					RequestedHostID     interface{}   `json:"requestedHostId"`
					RequestedIPAddress  interface{}   `json:"requestedIpAddress"`
					StartOnCreate       bool          `json:"startOnCreate"`
					StdinOpen           bool          `json:"stdinOpen"`
					Tty                 bool          `json:"tty"`
					User                interface{}   `json:"user"`
					Userdata            interface{}   `json:"userdata"`
					Vcpu                int           `json:"vcpu"`
					Version             string        `json:"version"`
					VolumeDriver        interface{}   `json:"volumeDriver"`
					WorkingDir          interface{}   `json:"workingDir"`
				} `json:"previousLaunchConfig"`
				PreviousSecondaryLaunchConfigs []interface{} `json:"previousSecondaryLaunchConfigs"`
				SecondaryLaunchConfigs         []interface{} `json:"secondaryLaunchConfigs"`
				StartFirst                     bool          `json:"startFirst"`
			} `json:"inServiceStrategy"`
			ToServiceStrategy interface{} `json:"toServiceStrategy"`
		} `json:"upgrade"`
		UUID string      `json:"uuid"`
		Vip  interface{} `json:"vip"`
	} `json:"data"`
	Filters struct {
		AccountID     interface{} `json:"accountId"`
		CreateIndex   interface{} `json:"createIndex"`
		Created       interface{} `json:"created"`
		Description   interface{} `json:"description"`
		EnvironmentID interface{} `json:"environmentId"`
		ExternalID    interface{} `json:"externalId"`
		HealthState   interface{} `json:"healthState"`
		ID            interface{} `json:"id"`
		Kind          interface{} `json:"kind"`
		Name          []struct {
			Modifier string `json:"modifier"`
			Value    string `json:"value"`
		} `json:"name"`
		RemoveTime        interface{} `json:"removeTime"`
		Removed           interface{} `json:"removed"`
		SelectorContainer interface{} `json:"selectorContainer"`
		SelectorLink      interface{} `json:"selectorLink"`
		State             interface{} `json:"state"`
		UUID              interface{} `json:"uuid"`
		Vip               interface{} `json:"vip"`
	} `json:"filters"`
	Links struct {
		Self string `json:"self"`
	} `json:"links"`
	Pagination struct {
		First    interface{} `json:"first"`
		Limit    int         `json:"limit"`
		Next     interface{} `json:"next"`
		Partial  bool        `json:"partial"`
		Previous interface{} `json:"previous"`
		Total    interface{} `json:"total"`
	} `json:"pagination"`
	ResourceType string      `json:"resourceType"`
	Sort         interface{} `json:"sort"`
	SortLinks    struct {
		AccountID         string `json:"accountId"`
		CreateIndex       string `json:"createIndex"`
		Created           string `json:"created"`
		Description       string `json:"description"`
		EnvironmentID     string `json:"environmentId"`
		ExternalID        string `json:"externalId"`
		HealthState       string `json:"healthState"`
		ID                string `json:"id"`
		Kind              string `json:"kind"`
		Name              string `json:"name"`
		RemoveTime        string `json:"removeTime"`
		Removed           string `json:"removed"`
		SelectorContainer string `json:"selectorContainer"`
		SelectorLink      string `json:"selectorLink"`
		State             string `json:"state"`
		UUID              string `json:"uuid"`
		Vip               string `json:"vip"`
	} `json:"sortLinks"`
	Type string `json:"type"`
}

func (r *Rancher) Call(hubMsg HubMessage) {
	log.Print("### Got WebHook call. BEGIN:")
	var ranchersvc = getService(r)
	lc, _ := json.Marshal(ranchersvc.Data[0].LaunchConfig)
	slc, _ := json.Marshal(ranchersvc.Data[0].SecondaryLaunchConfigs)
  url := ranchersvc.Data[0].Actions.Upgrade
	user := r.rancherConfig.UserKey
	pass := r.rancherConfig.SecretKey

	if url == "" {
		svcstatus := ranchersvc.Data[0].State
		log.Print("# Rancher not available for upgrade")
		log.Print("# Current status: ", svcstatus)
		if svcstatus == "upgrading" {
			log.Print("### END.")
			return
		} else if svcstatus == "upgraded" {
			log.Print("# Attempting to complete in-progress upgrade...")
		}
	} else {
		log.Print("# API URL: ", url)
	}

	var jsonStr = []byte(`{"inServiceStrategy":` +
		`{` +
			`"batchSize": 1,` +
	    `"intervalMillis": 2000,` +
	    `"startFirst": false,` +
	    `"launchConfig":` + string(lc) + `,` +
	    `"secondaryLaunchConfigs":` + string(slc) +
		`}` +
	`}`)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))

	req.SetBasicAuth(user, pass)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

  client := &http.Client{}
  resp, err := client.Do(req)
  if err != nil {
    panic(err)
  }
  defer resp.Body.Close()

	log.Print("# Sent upgrade request: ", resp.StatusCode)

	minute := time.Minute

	pollurl := ranchersvc.Data[0].Links.Self

	for i := 1; i <= 10; i++ {
		pollstatus, err := getUpgradeStatus(r, pollurl)
		if err != nil {
			log.Print("# ERR:", err)
		}
		if pollstatus == false {
			log.Print("# Waiting for upgrade to complete (try %v of 10)...", i)
			time.Sleep(minute)
		} else {
			log.Print("# Upgrade success.")
			finishUpgrade(r)
			dockerHubCallback(hubMsg)
			break
		}
	}

	log.Print(`### END.`)
}

func getService(r *Rancher) rancherService {
	url := r.rancherConfig.Url + "/services?name=" + r.rancherConfig.Service
	user := r.rancherConfig.UserKey
	pass := r.rancherConfig.SecretKey

	log.Print("SERVICE REQUEST URL: ", url)

	req, err := http.NewRequest("GET", url, nil)

	req.SetBasicAuth(user, pass)
	req.Header.Set("Accept", "application/json")

	client := &http.Client{}
  resp, err := client.Do(req)
  if err != nil {
    panic(err)
  }
  defer resp.Body.Close()

  body, _ := ioutil.ReadAll(resp.Body)

	var ranchersvc rancherService

	err = json.Unmarshal(body, &ranchersvc)
	if err != nil {
		log.Print("error unmarshaling Rancher JSON response: ", err)
	}

	return ranchersvc
}

func getUpgradeStatus(r *Rancher, url string) (status bool, err error) {
	var ranchersvc = getService(r)

	upgradestatus := ranchersvc.Data[0].State

	if upgradestatus == "upgrading" {
		status := false
		return status, err
	} else if upgradestatus == "upgraded" {
		status := true
		return status, err
	} else {
		err = fmt.Errorf("Rancher state '%v' not of supported type (upgrading or upgraded). Check Rancher for additional information.", upgradestatus)
		status := false
		return status, err
	}
}

func finishUpgrade(r *Rancher) {
	ranchersvc := getService(r)
	url := ranchersvc.Data[0].Actions.FinishUpgrade
	user := r.rancherConfig.UserKey
	pass := r.rancherConfig.SecretKey

	req, err := http.NewRequest("POST", url, nil)

	req.SetBasicAuth(user, pass)
	req.Header.Set("Accept", "application/json")

	client := &http.Client{}
  resp, err := client.Do(req)
  if err != nil {
    panic(err)
  }

	if resp.StatusCode == 202 {
		log.Print("# Finished Upgrade.")
	} else {
		log.Print("# Unknown Upgrade Status: ", resp.StatusCode)
	}

	return
}

func dockerHubCallback(hubMsg) {
	url := hubMsg.CallbackUrl

	var jsonStr = []byte(`{
		"state": "success",
    "description": "Rancher Deployment Successful",
    "context":      "Rancher Deployment",
    "target_url":   ""
	}`)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))

	client := &http.Client{}
  resp, err := client.Do(req)
  if err != nil {
    panic(err)
  }

	if resp.StatusCode == 202 {
		log.Print("# Callback to Docker Hub Webhook complete.")
	} else {
		log.Print("# Unknown Docker Hub Webhook Callback Status: ", resp.StatusCode)
	}
}
