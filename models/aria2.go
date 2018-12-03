package models

import (
	"github.com/astaxie/beego"
	ariarpc "github.com/zyxar/argo/rpc"
	"log"
)

var (
	rpcurl     = beego.AppConfig.String("aria2::rpcurl")
	aria2Token = beego.AppConfig.String("aria2::token")
)

func init() {
	if rpcurl == "" && aria2Token == "" {
		log.Fatalf("need rpcurl and aria2 token")
	}
}

type Aria2Client struct {
	Client ariarpc.Protocol
}

func NewAria2Client() (*Aria2Client, error) {
	var (
		client ariarpc.Protocol
		err    error
	)

	if client, err = ariarpc.New(rpcurl, aria2Token); err != nil {
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

func (a *Aria2Client) Status(gid string) {

}
