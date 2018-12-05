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
	RemoveEntry        = "remove task"
	PauseEntry         = "pause task"
	UnPauseEntry       = "unpause task"
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
		result      string
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

	if status == nil || len(status) == 0 {
		result = "not found"
		a.JsonOK(TellActiveEntry, result)
	}

	a.JsonOK(TellActiveEntry, status)
}

func (a *Aria2Controller) Remove() {
	var (
		err          error
		aria2Client  *models.Aria2Client
		gid          string
		forceFromGet string
		force        bool
	)

	gid = a.GetString("gid")
	forceFromGet = a.GetString("force")
	if forceFromGet == "yes" {
		force = true
	}

	if aria2Client, err = models.NewAria2Client(); err != nil {
		log.Println(err)
		a.JsonError(NewAria2Client, fmt.Sprintf("NewAria2Client error: %s", err), "")
		a.StopRun()
	}

	if gid, err = aria2Client.Remove(gid, force); err != nil {
		log.Println(err)
		a.JsonError(RemoveEntry, fmt.Sprintf("aria2Client.Remove error: %s", err), "")
		a.StopRun()
	}
	a.JsonOK(RemoveEntry, Data{"gid": gid})
}

func (a *Aria2Controller) Pause() {
	var (
		err          error
		aria2Client  *models.Aria2Client
		gid          string
		forceFromGet string
		force        bool
	)

	gid = a.GetString("gid")
	forceFromGet = a.GetString("force")
	if forceFromGet == "yes" {
		force = true
	}

	if aria2Client, err = models.NewAria2Client(); err != nil {
		log.Println(err)
		a.JsonError(NewAria2Client, fmt.Sprintf("NewAria2Client error: %s", err), "")
		a.StopRun()
	}

	if gid, err = aria2Client.Pause(gid, force); err != nil {
		log.Println(err)
		a.JsonError(PauseEntry, fmt.Sprintf("aria2Client.Pause error: %s", err), "")
		a.StopRun()
	}
	a.JsonOK(PauseEntry, Data{"gid": gid})
}

func (a *Aria2Controller) UnPause() {
	var (
		err         error
		aria2Client *models.Aria2Client
		gid         string
		allFromGet  string
		all         bool
	)

	gid = a.GetString("gid")
	allFromGet = a.GetString("all")
	if allFromGet == "yes" {
		all = true
	}

	if aria2Client, err = models.NewAria2Client(); err != nil {
		log.Println(err)
		a.JsonError(NewAria2Client, fmt.Sprintf("NewAria2Client error: %s", err), "")
		a.StopRun()
	}

	if gid, err = aria2Client.UnPause(gid, all); err != nil {
		log.Println(err)
		a.JsonError(UnPauseEntry, fmt.Sprintf("aria2Client.UnPause error: %s", err), "")
		a.StopRun()
	}
	a.JsonOK(UnPauseEntry, Data{"result": gid})
}

func (a *Aria2Controller) PauseAll() {
	var (
		err          error
		aria2Client  *models.Aria2Client
		forceFromGet string
		force        bool
		success      string
	)

	forceFromGet = a.GetString("force")
	if forceFromGet == "yes" {
		force = true
	}

	if aria2Client, err = models.NewAria2Client(); err != nil {
		log.Println(err)
		a.JsonError(NewAria2Client, fmt.Sprintf("NewAria2Client error: %s", err), "")
		a.StopRun()
	}

	if success, err = aria2Client.PauseAll(force); err != nil {
		log.Println(err)
		a.JsonError(PauseEntry, fmt.Sprintf("aria2Client.PauseAll error: %s", err), "")
		a.StopRun()
	}
	a.JsonOK(PauseEntry, Data{"success": success})
}
