## UBNT AirFiber 60 LR Icinga plugin

- using ssh command to get values from airfiber60lr
- your ssh public key must be uploaded to the airfiber60lr
- prebuild executable file above

### Example usege
./check-ubnt-airfiber60lr -u admin -H 10.10.5.98 

STATE OK |Rssi(dBm)=-37;;;; Cinr(dBm)=19;;;; TxMsc=9;;;; RxMsc=9;;;; Capacity(*100)=875;;;; 

### Startup arguments
```go
package lib

type ubntCheckStartupOpts struct {
	Host     string `short:"H" long:"hostname" default:"localhost" description:"Host name or IP Address"`
	User     string `short:"u" long:"user" default:"ubnt" description:"ssh username"`
	Warning  int    `short:"w" long:"warning" default:"-40" description:"RSSI warning level"`
	Critical int    `short:"c" long:"critical" default:"-45" description:"RSSI critical level"`
}
```
