package mainflow

import (
	"github.com/bitwyre/bitwyre/gateway/rest/p2p_api/internal/model"
	"github.com/bitwyre/bitwyre/gateway/rest/p2p_api/internal/pkg/helper"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

// Main Flow handler
// @Summary Get Crypto List
// @Id mainflow
// @Tags Mainflow
// @version 1.0
// @Param Authorization header string true "With the bearer started"
// @produce application/json
// @Success 200 {object} helper.BaseResponseSchema
// @Router /private/main/crypto/list [get]
func (h *MainFlowHandler) GetCryptoList(ctx *gin.Context) {
	userId := ctx.Request.Header.Get("x-user-id")
	span := trace.SpanFromContext(ctx.Request.Context())
	span.SetAttributes(attribute.String("userId", userId))

	cryptoList := []string{"USDT", "BTC", "BUSD", "BNB", "ETH", "DOGE", "BIDR", "DAI", "SHIB"}
	currencies := []model.LabelIcon{
		{Label: "IDR", Icon: "RP"}, {Label: "INR", Icon: "inr"}, {Label: "IQD", Icon: "iqd"},
	}
	helper.HttpSuccess(model.GetCryptoListRes{
		CryptoList: cryptoList,
		Currency:   currencies,
	}, nil, ctx)
}

// Main Flow handler
// @Summary Get Ads List
// @Tags Mainflow
// @version 1.0
// @Param Authorization header string true "With the bearer started"
// @produce application/json
// @Success 200 {object} helper.BaseResponseSchema
// @Router /private/main/ads/list [get]
func (h *MainFlowHandler) GetAdsList(ctx *gin.Context) {
	userId := ctx.Request.Header.Get("x-user-id")
	span := trace.SpanFromContext(ctx.Request.Context())
	span.SetAttributes(attribute.String("userId", userId))

	crypto := ctx.Query("crypto")
	adsList := []model.GetAdsListRes{
		{
			PictureUrl: "https://cdn-icons-png.flaticon.com/512/3135/3135715.png",
			Name:       "RVS-990",
			IsVerified: true,
			UserTradeInfo: model.UserTradeInfo{
				TradeAvg:           552,
				CompletionsRateAvg: 100,
				ReceivedFeedBack:   99.15,
				ReleaseTimeAvg:     15,
			},
			AdsInfo: model.AdsInfo{
				CryptoLabel:          crypto,
				Price:                14899,
				PriceFormat:          "Rp. 14.899",
				Qty:                  1724.10,
				Limit:                "Rp. 8.000.000,00 - Rp. 25.000.000,00",
				SupportedPaymentList: []string{"Offline meet (COD)", "Bank Transfer", "BCA"},
			},
		},
		{
			PictureUrl: "https://cdn-icons-png.flaticon.com/512/3135/3135715.png",
			Name:       "YanKo710_zdes_Exchanges",
			IsVerified: true,
			UserTradeInfo: model.UserTradeInfo{
				TradeAvg:           552,
				CompletionsRateAvg: 100,
				ReceivedFeedBack:   99.15,
				ReleaseTimeAvg:     15,
			},
			AdsInfo: model.AdsInfo{
				CryptoLabel:          crypto,
				Price:                14899,
				PriceFormat:          "Rp. 14.899",
				Qty:                  215.38,
				Limit:                "Rp. 8.000.000,00 - Rp. 25.000.000,00",
				SupportedPaymentList: []string{"Offline meet (COD)", "Mandiri Pay", "OCBC NISP"},
			},
		},
	}

	h.adsUsecase.GetAdsList(ctx)
	helper.HttpSuccess(adsList, nil, ctx)
}
