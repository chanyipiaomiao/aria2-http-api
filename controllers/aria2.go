package controllers

import (
	"fmt"
	"github.com/chanyipiaomiao/aria2-http-api/models"
	"github.com/zyxar/argo/rpc"
	"log"
)

const (
	ADDURLTOARria2  = "add url to aria2"
	TellStatusEntry = "tell gid status"
)

type Aria2Controller struct {
	BaseController
}

func (a *Aria2Controller) AddUrl() {
	var (
		url         string
		err         error
		aria2Client *models.Aria2Client
		gid         string
	)
	url = a.GetString("url")
	if aria2Client, err = models.NewAria2Client(); err != nil {
		log.Println(err)
		a.JsonError(ADDURLTOARria2, fmt.Sprintf("NewAria2Client error: %s", err), "")
		a.StopRun()
	}

	if gid, err = aria2Client.AddUrl(url); err != nil {
		log.Println(err)
		a.JsonError(ADDURLTOARria2, fmt.Sprintf("aria2Client.AddUrl error: %s", err), "")
		a.StopRun()
	}

	a.JsonOK(ADDURLTOARria2, Data{"gid": gid})
}

func (a *Aria2Controller) TellStatus() {
	var (
		err         error
		aria2Client *models.Aria2Client
		gid         string
		status      rpc.StatusInfo
	)
	gid = a.GetString("gid")
	if aria2Client, err = models.NewAria2Client(); err != nil {
		log.Println(err)
		a.JsonError(ADDURLTOARria2, fmt.Sprintf("NewAria2Client error: %s", err), "")
		a.StopRun()
	}

	if status, err = aria2Client.TellStatus(gid); err != nil {
		log.Println(err)
		a.JsonError(TellStatusEntry, fmt.Sprintf("aria2Client.TellStatus error: %s", err), "")
		a.StopRun()
	}

	a.JsonOK(TellStatusEntry, status)
}
