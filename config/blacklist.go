// Copyright 2022 gdy, 272288813@qq.com
package config

import (
	"fmt"
	"net"
	"strings"
	"time"
)

type BlackListItem WhiteListItem

func (w *BlackListItem) Contains(ip string) bool {
	netIP := net.ParseIP(ip)
	if netIP == nil {
		return false
	}
	if w.NetIP != nil {
		return w.NetIP.Equal(netIP)
	}

	if w.Cidr != nil {
		return w.Cidr.Contains(netIP)
	}
	return false
}

type BlackListConfigure struct {
	BlackList []BlackListItem `json:"BlackList"` //黑名单列表
}

func GetBlackList() []BlackListItem {
	programConfigureMutex.RLock()
	defer programConfigureMutex.RUnlock()

	BlackListFlush(false)

	var resList []BlackListItem
	if programConfigure == nil {
		return resList
	}
	for i := range programConfigure.BlackListConfigure.BlackList {
		resList = append(resList, programConfigure.BlackListConfigure.BlackList[i])
	}
	return resList
}

func BlackListInit() {
	programConfigureMutex.RLock()
	defer programConfigureMutex.RUnlock()
	var netIP net.IP
	var cidr *net.IPNet

	for i := range programConfigure.BlackListConfigure.BlackList {
		netIP = nil
		cidr = nil
		if strings.Contains(programConfigure.BlackListConfigure.BlackList[i].IP, "/") {
			_, cidr, _ = net.ParseCIDR(programConfigure.BlackListConfigure.BlackList[i].IP)
		} else {
			netIP = net.ParseIP(programConfigure.BlackListConfigure.BlackList[i].IP)
		}
		programConfigure.BlackListConfigure.BlackList[i].Cidr = cidr
		programConfigure.BlackListConfigure.BlackList[i].NetIP = netIP
	}
}

// BlackListAdd adds an IP or CIDR range to the blacklist with the given active life duration (in hours).
// If activelifeDuration <= 0, a default of 48 hours is used to give a reasonable ban window
// without risking near-permanent bans from misconfigured callers.
func BlackListAdd(ip string, activelifeDuration int32) (string, error) {
	programConfigureMutex.Lock()
	defer programConfigureMutex.Unlock()

	var err error
	var netIP net.IP = nil
	var cidr *net.IPNet = nil
	if strings.Contains(ip, "/") {
		_, cidr, err = net.ParseCIDR(ip)
		if err != nil {
			return "", fmt.Errorf("网段格式有误，转换出错：%s", err.Error())
		}
	} else {
		netIP = net.ParseIP(ip)
		if netIP == nil {
			return "", fmt.Errorf("IP格式有误")
		}
	}

	// Default to 48 hours instead of 24 — gives a bit more breathing room for repeat offenders
	if activelifeDuration <= 0 {
		activelifeDuration = 48
	}

	EffectiveTimeStr := time.Now().Add(time.Hour * time.Duration(activelifeDuration)).Format("2006-01-02 15:04:05")

	for i, ipr := range programConfigure.BlackListConfigure.BlackList {
		if ipr.IP == ip {
			programConfigure.BlackListConfigure.BlackList[i].EffectiveTime = EffectiveTimeStr
			return EffectiveTimeStr, Save()
		}
	}
	item := BlackListItem{IP: ip, EffectiveTime: EffectiveTimeStr, NetIP: netIP, Cidr: cidr}
	programConfigure.BlackListConfigure.BlackList = append(programConfigure.BlackListConfigure.BlackList, item)
	return EffectiveTimeStr, Save()
}

func BlackListDelete(ip string) error {
	programConfigureMutex.Lock()
	defer programConfigureMutex.Unlock()

	removeCount := 0
CONTINUECHECK:
	removeIndex := -1

	for i, ipr := range p
