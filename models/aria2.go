package models

import (
	"github.com/astaxie/beego"
	ariarpc "github.com/zyxar/argo/rpc"
	"log"
)

var (
	aria2RPCurl   = beego.AppConfig.String("aria2::rpcurl")
	aria2Token    = beego.AppConfig.String("aria2::token")
	displayFields = []string{"gid", "status", "totalLength", "completedLength", "downloadSpeed"}
)

func init() {
	if aria2RPCurl == "" && aria2Token == "" {
		log.Fatalf("need rpcurl and aria2 token")
	}
}

// Aria2Status 返回状态结果
type Aria2Status struct {
	GID             string
	Status          string
	TotalLength     string
	CompletedLength string
	DownloadSpeed   string
}

type Aria2Client struct {
	Client ariarpc.Protocol
}

func NewAria2Client() (*Aria2Client, error) {
	var (
		client ariarpc.Protocol
		err    error
	)

	if client, err = ariarpc.New(aria2RPCurl, aria2Token); err != nil {
		return nil, err
	}

	return &Aria2Client{
		Client: client,
	}, nil
}

func (a *Aria2Client) AddUrl(url string) (string, error) {
	return a.Client.AddURI(url)
}

func (a *Aria2Client) AddTorrent(filename string) (string, error) {
	return a.Client.AddTorrent(filename)
}

func (a *Aria2Client) Remove(gid string, force bool) (string, error) {
	if force {
		return a.Client.ForceRemove(gid)
	}
	return a.Client.Remove(gid)
}

func (a *Aria2Client) Pause(gid string, force bool) (string, error) {
	if force {
		return a.Client.ForcePause(gid)
	}
	return a.Client.Pause(gid)
}

func (a *Aria2Client) PauseAll(force bool) (string, error) {
	if force {
		return a.Client.ForcePauseAll()
	}
	return a.Client.PauseAll()
}

func (a *Aria2Client) UnPause(gid string, all bool) (string, error) {
	if all {
		return a.Client.UnpauseAll()
	}
	return a.Client.Unpause(gid)
}

func (a *Aria2Client) TellStatus(gid string) (*Aria2Status, error) {
	var (
		info   ariarpc.StatusInfo
		status *Aria2Status
		err    error
	)

	if info, err = a.Client.TellStatus(gid, displayFields...); err != nil {
		return nil, err
	}

	status = &Aria2Status{
		GID:             info.Gid,
		Status:          info.Status,
		TotalLength:     info.TotalLength,
		CompletedLength: info.CompletedLength,
		DownloadSpeed:   info.DownloadSpeed,
	}
	return status, nil
}

func (a *Aria2Client) TellActive() ([]*Aria2Status, error) {
	var (
		infos  []ariarpc.StatusInfo
		status []*Aria2Status
		err    error
	)

	if infos, err = a.Client.TellActive(displayFields...); err != nil {
		return nil, err
	}

	for _, info := range infos {
		s := &Aria2Status{
			GID:             info.Gid,
			Status:          info.Status,
			TotalLength:     info.TotalLength,
			CompletedLength: info.CompletedLength,
			DownloadSpeed:   info.DownloadSpeed,
		}
		status = append(status, s)
	}

	return status, nil
}

func (a *Aria2Client) TellWaiting(offset, num int) ([]ariarpc.StatusInfo, error) {
	return a.Client.TellWaiting(offset, num, displayFields...)
}

func (a *Aria2Client) TellStopped(offset, num int) ([]ariarpc.StatusInfo, error) {
	return a.Client.TellStopped(offset, num, displayFields...)
}
