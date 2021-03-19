package settings

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"strings"
)

const (
	assetUSDT = "USDT"
	assetBNB = "BNB"
	assetBUSD = "BUSD"
	SettingSideSell = "short"
	SettingSideBuy = "long"
	SettingSideBoth = "both"
	TypeTpLimit = "limit"
	TypeTpTrailing = "trailing"
	)

type UserSettings struct {
	LicenceKey              string  `yaml:"LicenceKey"`
	Platform                string  `yaml:"platform"`
	PublicKey               string  `yaml:"publicKey"`
	SecretKey               string  `yaml:"secretKey"`
	Symbol                  string  `yaml:"symbol"`
	Side 					string 	`yaml:"side"`
	Asset					string
	QuantityPercentage      float64 `yaml:"quantityPercentage"`
	MaxPositionSize         float64 `yaml:"maxPositionSize"`
	VwapOffsetPercentage    float64 `yaml:"vwapOffsetPercentage"`
	Leverage				int 	`yaml:"leverage"`
	MinimumLiquidationValue float64 `yaml:"minimumLiquidationValue"`
	TypeOfTp				string	`yaml:"typeOfTp"`
	LimitTpPercentage		float64 `yaml:"limitTpPercentage"`
	TrailingStartPercentage	float64	`yaml:"trailingStartPercentage"`
	TrailingPercentage		float64 `yaml:"trailingPercentage"`
	StopLossEnabled			bool	`yaml:"stopLossEnabled"`
	StopLossPercentage      float64 `yaml:"stopLossPercentage"`
	CoolDownPeriod			int 	`yaml:"coolDownPeriod"`
}

type LicenceKeyInd struct {
	LicenceKey	string  `yaml:"LicenceKey"`
}

func readFile() []byte {
	data, err := ioutil.ReadFile("settings.yaml")
	if err != nil {
		log.Panic(err, "Unable to read settings.yaml file")
	}
	return data
}

func (us *UserSettings) LoadSettings() {
	data := readFile()
	err := yaml.Unmarshal(data, us)
	if err != nil {
		log.Panic(err)
	}
	// Api communication works with Upper case symbol names
	// Note: don't worry when sending symbol to websocket
	// it uses lowercase version of symbol.
	us.Symbol = strings.ToUpper(us.Symbol)
	us.Side = strings.ToLower(us.Side)
	us.TypeOfTp = strings.ToLower(us.TypeOfTp)
	if strings.Contains(us.Symbol,assetUSDT) {
		us.Asset = assetUSDT
	}else if strings.Contains(us.Symbol,assetBUSD) {
		us.Asset = assetBUSD
	} else if strings.Contains(us.Symbol,assetBNB) {
		us.Asset = assetBNB
	} else {
		log.Fatal("User has provided wrong symbol")
	}
	us.QuantityPercentage = 	 us.QuantityPercentage / 100
	us.VwapOffsetPercentage = 	 us.VwapOffsetPercentage / 100
	us.LimitTpPercentage = 		 us.LimitTpPercentage / 100
	us.TrailingPercentage = 	 us.TrailingPercentage / 100
	us.StopLossPercentage = 	 us.StopLossPercentage / 100
	us.TrailingStartPercentage = us.TrailingStartPercentage / 100
}
