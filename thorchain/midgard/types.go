package midgard

type GetPoolsListRequest struct {
	Status string
	Period string
}

type Pool struct {
	Asset                          string `json:"Asset,omitempty"`
	Volume24h                      string `json:"Volume24H,omitempty"`
	AssetDepth                     string `json:"AssetDepth,omitempty"`
	RuneDepth                      string `json:"RuneDepth,omitempty"`
	AssetPrice                     string `json:"AssetPrice,omitempty"`
	AssetPriceUSD                  string `json:"AssetPriceUSD,omitempty"`
	PoolAPY                        string `json:"PoolAPY,omitempty"`
	AnnualPercentageRate           string `json:"AnnualPercentageRate,omitempty"`
	Earnings                       string `json:"Earnings,omitempty"`
	EarningsAnnualAsPercentOfDepth string `json:"EarningsAnnualAsPercentOfDepth,omitempty"`
	LpLuvi                         string `json:"LpLuvi,omitempty"`
	SaversAPR                      string `json:"SaversAPR,omitempty"`
	Status                         string `json:"Status,omitempty"`
	LiquidityUnits                 string `json:"LiquidityUnits,omitempty"`
	SynthUnits                     string `json:"SynthUnits,omitempty"`
	SynthSupply                    string `json:"SynthSupply,omitempty"`
	Units                          string `json:"Units,omitempty"`
	NativeDecimal                  string `json:"NativeDecimal,omitempty"`
	SaverUnits                     string `json:"SaverUnits,omitempty"`
	SaversDepth                    string `json:"SaversDepth,omitempty"`
	TotalCollateral                string `json:"TotalCollateral,omitempty"`
	TotalDepthTor                  string `json:"TotalDepthTor,omitempty"`
	SaversYieldShare               string `json:"SaversYieldShare,omitempty"`
}

type GetPoolsListResponse struct {
	Pools []Pool
}
