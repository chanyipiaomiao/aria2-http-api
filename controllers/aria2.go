package controllers

import (
	"fmt"
	"github.com/chanyipiaomiao/aria2-http-api/models"
	"log"
)

const (
	NewAria2Client     = "new aria2 client"
	AddURLToAria2Entry = "add url to aria2"
	TellStatusEntry    = "tell gid status"
	TellActiveEntry    = "tell active"
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
		a.JsonError(NewAria2Client, fmt.Sprintf("NewAria2Client error: %s", err), "")
		a.StopRun()
	}

	if gid, err = aria2Client.AddUrl(url); err != nil {
		log.Println(err)
		a.JsonError(AddURLToAria2Entry, fmt.Sprintf("aria2Client.AddUrl error: %s", err), "")
		a.StopRun()
	}

	a.JsonOK(AddURLToAria2Entry, Data{"gid": gid})
}

func (a *Aria2Controller) TellStatus() {
	var (
		err         error
		aria2Client *models.Aria2Client
		gid         string
		status      *models.Aria2Status
	)

	gid = a.GetString("gid")
	if aria2Client, err = models.NewAria2Client(); err != nil {
		log.Println(err)
		a.JsonError(NewAria2Client, fmt.Sprintf("NewAria2Client error: %s", err), "")
		a.StopRun()
	}

	if status, err = aria2Client.TellStatus(gid); err != nil {
		log.Println(err)
		a.JsonError(TellStatusEntry, fmt.Sprintf("aria2Client.TellStatus error: %s", err), "")
		a.StopRun()
	}

	a.JsonOK(TellStatusEntry, status)
}

func (a *Aria2Controller) TellActive() {
	var (
		err         error
		aria2Client *models.Aria2Client
		status      []*models.Aria2Status
	)

	if aria2Client, err = models.NewAria2Client(); err != nil {
		log.Println(err)
		a.JsonError(NewAria2Client, fmt.Sprintf("NewAria2Client error: %s", err), "")
		a.StopRun()
	}

	if status, err = aria2Client.TellActive(); err != nil {
		log.Println(err)
		a.JsonError(TellActiveEntry, fmt.Sprintf("aria2Client.TellActive error: %s", err), "")
		a.StopRun()
	}

	a.JsonOK(TellActiveEntry, status)
}
