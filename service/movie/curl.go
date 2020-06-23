package movie

import (
	"github.com/go-resty/resty/v2"
)

type BoxOfficeData struct {
	Success bool `json:"success"`
	Data    struct {
		UpdateInfo       string `json:"updateInfo"`
		TotalBoxUnitInfo string `json:"totalBoxUnitInfo"`
		SplitTotalBox    string `json:"splitTotalBox"`
		ServerTimestamp  int64  `json:"serverTimestamp"`
		Crystal          struct {
			MaoyanViewInfo string `json:"maoyanViewInfo"`
			Status         int    `json:"status"`
			ViewInfo       string `json:"viewInfo"`
			ViewUnitInfo   string `json:"viewUnitInfo"`
		} `json:"crystal"`
		TotalBoxInfo string `json:"totalBoxInfo"`
		List         []struct {
			AvgSeatView      string `json:"avgSeatView"`
			AvgShowView      string `json:"avgShowView"`
			AvgViewBox       string `json:"avgViewBox"`
			BoxInfo          string `json:"boxInfo"`
			BoxRate          string `json:"boxRate"`
			MovieID          int    `json:"movieId"`
			MovieName        string `json:"movieName"`
			MyRefundNumInfo  string `json:"myRefundNumInfo"`
			MyRefundRateInfo string `json:"myRefundRateInfo"`
			OnlineBoxRate    string `json:"onlineBoxRate"`
			RefundViewInfo   string `json:"refundViewInfo"`
			RefundViewRate   string `json:"refundViewRate"`
			ReleaseInfo      string `json:"releaseInfo"`
			ReleaseInfoColor string `json:"releaseInfoColor"`
			SeatRate         string `json:"seatRate"`
			ShowInfo         string `json:"showInfo"`
			ShowRate         string `json:"showRate"`
			SplitAvgViewBox  string `json:"splitAvgViewBox"`
			SplitBoxInfo     string `json:"splitBoxInfo"`
			SplitBoxRate     string `json:"splitBoxRate"`
			SplitSumBoxInfo  string `json:"splitSumBoxInfo"`
			SumBoxInfo       string `json:"sumBoxInfo"`
			ViewInfo         string `json:"viewInfo"`
			ViewInfoV2       string `json:"viewInfoV2"`
		} `json:"list"`
		TotalBoxUnit          string `json:"totalBoxUnit"`
		TotalBox              string `json:"totalBox"`
		SplitTotalBoxUnit     string `json:"splitTotalBoxUnit"`
		QueryDate             string `json:"queryDate"`
		Holidays              string `json:"holidays"`
		ServerTime            string `json:"serverTime"`
		SplitTotalBoxUnitInfo string `json:"splitTotalBoxUnitInfo"`
		SplitTotalBoxInfo     string `json:"splitTotalBoxInfo"`
	} `json:"data"`
	Status int `json:"status"`
}

var client *resty.Client

func init() {
	client = resty.New()
}

func GetBoxOffice()( *BoxOfficeData, error) {
	res := BoxOfficeData{}
	_, err := client.R().
		SetHeader("Accept", "application/json").
		SetResult(&res).
		Get("https://piaofang.maoyan.com/second-box")

	if err != nil {
		return nil, err
	}
	return &res, nil
}
