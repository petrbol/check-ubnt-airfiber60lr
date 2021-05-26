package lib

type ubntAirfiber60lrStruct []struct {
	PrsSta struct {
		RssiData int `json:"rssi_data"`
		Snr      int `json:"snr"`
		Capacity int `json:"capacity"`
		TxMcs    int `json:"tx_mcs"`
		RxMcs    int `json:"rx_mcs"`
	} `json:"prs_sta"`
}

type ubntCheckStartupOpts struct {
	Host       string `short:"H" long:"hostname" default:"localhost" description:"Host name or IP Address"`
	User       string `short:"u" long:"user" default:"ubnt" description:"ssh username"`
	PrivateCrt string `short:"p" long:"private" default:"" description:"ssh private cert file"`
	Warning    int    `short:"w" long:"warning" default:"-40" description:"RSSI warning level"`
	Critical   int    `short:"c" long:"critical" default:"-45" description:"RSSI critical level"`
}
