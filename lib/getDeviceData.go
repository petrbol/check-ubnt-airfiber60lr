package lib

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
)

func (opts *ubntCheckStartupOpts) getDeviceData() {
	//sshpass -p $password ssh -q -o StrictHostKeyChecking=no -o UserKnownHostsFile=/dev/null -n $username@${ip_list[$i]} "fwupdate -m /tmp/fwupdate.bin > /dev/null 2>&1";
	loginArgs := opts.User + "@" + opts.Host
	commandToExecute := "\"wstalist\""
	sshOutput, err := exec.Command("ssh", "-q", "-o", "StrictHostKeyChecking=no", "-o", "UserKnownHostsFile=/dev/null", "-n", loginArgs, commandToExecute).Output()
	if err != nil {
		log.Fatal(err, string(sshOutput))
	}

	var ubntAirfiber60lr ubntAirfiber60lrStruct
	err = json.Unmarshal(sshOutput, &ubntAirfiber60lr)
	if err != nil {
		log.Fatal(err)
	}

	var status = 0
	var longLine string
	var sshDecoded = ubntAirfiber60lr[0].PrsSta

	if sshDecoded.RssiData < opts.Critical {
		status = 2
	} else if sshDecoded.RssiData < opts.Warning {
		status = 1
	}

	rssi := "Rssi(dBm)=" + fmt.Sprint(sshDecoded.RssiData) + ";;;; "
	cinr := "Cnr(dBm)=" + fmt.Sprint(sshDecoded.Snr) + ";;;; "
	txmsc := "TxMsc=" + fmt.Sprint(sshDecoded.TxMcs) + ";;;; "
	rxmsc := "RxMsc=" + fmt.Sprint(sshDecoded.RxMcs) + ";;;; "
	capacity := "Capacity(*100)=" + fmt.Sprint(sshDecoded.Capacity/1000) + ";;;; "

	longLine = rssi + cinr + txmsc + rxmsc + capacity

	// final output
	if status == 2 {
		output := fmt.Sprint("STATE CRITICAL ", "|", longLine)
		fmt.Println(output)
		os.Exit(2)
	} else if status == 1 {
		output := fmt.Sprint("STATE WARNING ", "|", longLine)
		fmt.Println(output)
		os.Exit(1)
	} else {
		output := fmt.Sprint("STATE OK ", "|", longLine)
		fmt.Println(output)
		os.Exit(0)
	}
}
