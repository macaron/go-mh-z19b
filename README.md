# go-mh-z19b

二酸化炭素濃度計「MH-Z19B」のGo言語ライブラリ

## センサーについて

MH-Z19B([データシート](https://www.winsen-sensor.com/d/files/infrared-gas-sensor/mh-z19b-co2-ver1_0.pdf)) は安価な二酸化炭素濃度計として有名です。

## インストール

```shell
$ go get -u github.com/macaron/go-mh-z19b
```

## サンプルコード

```go
func main() {
    mhz19b := mhz19b.New("/dev/serial0")
    ppm, err := mhz19b.Read()
    
    if err != nil {
    	log.Fetal(err)
    }
    fmt.Printf("CO2 = %dppm\n", ppm)

    // If you want to reset the sensor ZERO point
    // mhz19b.CalibrateDefault()
}
```
