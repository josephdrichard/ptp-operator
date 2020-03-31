package network

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/golang/glog"
	"github.com/jaypipes/ghw"
)

const (
	_ETHTOOL_HARDWARE_RECEIVE_CAP   = "hardware-receive"
	_ETHTOOL_HARDWARE_TRANSMIT_CAP  = "hardware-transmit"
	_ETHTOOL_HARDWARE_RAW_CLOCK_CAP = "hardware-raw-clock"
	_ETHTOOL_RX_HARDWARE_FLAG       = "(SOF_TIMESTAMPING_RX_HARDWARE)"
	_ETHTOOL_TX_HARDWARE_FLAG       = "(SOF_TIMESTAMPING_TX_HARDWARE)"
	_ETHTOOL_RAW_HARDWARE_FLAG      = "(SOF_TIMESTAMPING_RAW_HARDWARE)"
)

func ethtoolInstalled() bool {
	_, err := exec.LookPath("ethtool")
	return err == nil
}

func netParseEthtoolTimeStampFeature(cmdOut *bytes.Buffer) bool {
	var hardRxEnabled bool
	var hardTxEnabled bool
	var hardRawEnabled bool

	glog.V(2).Infof("cmd output for %v", cmdOut)
	scanner := bufio.NewScanner(cmdOut)
	for scanner.Scan() {
		line := strings.TrimPrefix(scanner.Text(), "\t")
		parts := strings.Fields(line)
		if parts[0] == _ETHTOOL_HARDWARE_RECEIVE_CAP {
			hardRxEnabled = parts[1] == _ETHTOOL_RX_HARDWARE_FLAG
		}
		if parts[0] == _ETHTOOL_HARDWARE_TRANSMIT_CAP {
			hardTxEnabled = parts[1] == _ETHTOOL_TX_HARDWARE_FLAG
		}
		if parts[0] == _ETHTOOL_HARDWARE_RAW_CLOCK_CAP {
			hardRawEnabled = parts[1] == _ETHTOOL_RAW_HARDWARE_FLAG
		}
	}
	return hardRxEnabled && hardTxEnabled && hardRawEnabled
}

func DiscoverPTPDevices() ([]string, error) {
	var out bytes.Buffer
	nics := make([]string, 0)

	if !ethtoolInstalled() {
		return nics, fmt.Errorf("discoverDevices(): ethtool not installed. Cannot grab NIC capabilities")
	}

	ethtoolPath, _ := exec.LookPath("ethtool")

	net, err := ghw.Network()
	if err != nil {
		return nics, fmt.Errorf("discoverDevices(): error getting network info: %v", err)
	}

	for _, dev := range net.NICs {
		glog.Infof("grabbing NIC timestamp capability for %v", dev.Name)
		cmd := exec.Command(ethtoolPath, "-T", dev.Name)
		cmd.Stdout = &out
		err := cmd.Run()
		if err != nil {
			glog.Infof("could not grab NIC timestamp capability for %v: %v", dev.Name, err)
		}

		if !netParseEthtoolTimeStampFeature(&out) {
			continue
		}

		cmd = exec.Command("readlink", "-f", fmt.Sprintf("/sys/class/net/%s", dev.Name))
		cmd.Stdout = &out
		err = cmd.Run()
		if err != nil {
			glog.Infof("could not grab NIC PCI address for %v: %v", dev.Name, err)
			continue
		}

		// sysfs address looks like: /sys/devices/pci0000:17/0000:17:02.0/0000:19:00.5/net/eno1
		pathSegments := strings.Split(out.String(), "/")
		if len(pathSegments) != 8 {
			glog.Infof("unexpected sysfs address for %v: %v", dev.Name, out.String())
			continue
		}

		PCIAddr := pathSegments[5] // 0000:19:00.5

		// If the physfn doesn't exist this means the interface is not a virtual function so we ca add it to the list
		if _, err := os.Stat(fmt.Sprintf("/sys/bus/pci/devices/%s/physfn", PCIAddr)); os.IsNotExist(err) {
			nics = append(nics, dev.Name)
		}
	}
	return nics, nil
}
