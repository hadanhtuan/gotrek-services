package apiPayment

import (
	"context"
	"payment-service/internal/model"
	"payment-service/internal/util"
	protoPayment "payment-service/proto/payment"
	protoSdk "payment-service/proto/sdk"

	"github.com/hadanhtuan/go-sdk/common"
	"github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/paymentintent"
)

func (bc *PaymentController) CreatePaymentIntent(ctx context.Context, req *protoPayment.MsgCreatePaymentIntent) (*protoSdk.BaseResponse, error) {
	params := &stripe.PaymentIntentParams{
		Amount:   stripe.Int64(req.Amount),
		Currency: stripe.String(string(stripe.CurrencyUSD)),
		AutomaticPaymentMethods: &stripe.PaymentIntentAutomaticPaymentMethodsParams{
			Enabled: stripe.Bool(true),
		},
		Metadata: map[string]string{
			"bookingId":  req.BookingId,
			"propertyId": req.PropertyId,
			"userId":     req.UserId,
		},
	}

	pi, err := paymentintent.New(params)

	if err != nil {
		return util.ConvertToGRPC(&common.APIResponse{
			Status:  common.APIStatus.BadRequest,
			Message: "Error create payment intent. Error detail: " + err.Error(),
		})
	}

	paymentLog := &model.PaymentLog{
		StripeId:     &pi.ID,
		UserId:       &req.UserId,
		BookingId:    &req.BookingId,
		PropertyId:   &req.PropertyId,
		Amount:       req.Amount,
		Currency:     req.Currency,
		Status:       string(stripe.PaymentIntentStatusProcessing),
		ReceiptEmail: req.ReceiptEmail,
	}

	_ = model.PaymentLogDB.Create(paymentLog)

	return util.ConvertToGRPC(&common.APIResponse{
		Status:  common.APIStatus.Ok,
		Message: "Create payment intent successfully",
		Data:    pi.ClientSecret,
	})
}

/*
flow:
1. Create payment intent
2. Fe call stripe
3. Be handle webhook. 
	Test: stripe trigger payment_intent.succeeded --add "payment_intent:metadata[bookingId]=booking1"

4. Send event to property service
*/
