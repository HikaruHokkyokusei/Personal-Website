package PaymentStripe

import (
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/stripe/stripe-go/v74"
	"github.com/stripe/stripe-go/v74/paymentintent"
)

type Item struct {
	Id     string `json:"id"`
	Amount int64  `json:"amount"`
}
type PaymentIntent struct {
	clientSecret  string
	items         []Item
	amount        int64
	processingFee int64
	netAmount     int64
}

var processingFeePercent int64 = 3
var activePayments = map[string]PaymentIntent{}
var products map[string]int64

func calculateOrderAmount(items []Item) (int64, int64) {
	var amount int64 = 0
	for _, item := range items {
		amount += item.Amount * products[item.Id]
	}
	return amount, (amount * 100) / (100 - processingFeePercent)
}

func createNewPaymentIntent(items []Item) (bool, string) {
	var amount, netAmount = calculateOrderAmount(items)

	if pi, err := paymentintent.New(&stripe.PaymentIntentParams{
		Amount:   stripe.Int64(netAmount),
		Currency: stripe.String(string(stripe.CurrencyUSD)),
		PaymentMethodTypes: []*string{
			stripe.String("card"),
		},
	}); err == nil {
		var paymentId = b64.StdEncoding.EncodeToString([]byte(uuid.New().String() + "-" + uuid.New().String()))
		activePayments[paymentId] = PaymentIntent{
			clientSecret:  pi.ClientSecret,
			items:         items,
			amount:        amount,
			processingFee: netAmount - amount,
			netAmount:     netAmount,
		}
		return true, paymentId
	} else {
		fmt.Println("createNewPaymentIntent Error:", err)
		return false, ""
	}
}

func ConfigurePaymentEndpoints(server fiber.Router, stripePublicKey string, stripePrivateKey string) {
	stripe.Key = stripePrivateKey

	server.Post("/create-new-payment-intent", func(ctx *fiber.Ctx) error {
		var inputData struct {
			Items []Item `json:"items"`
		}

		if err := json.Unmarshal(ctx.Body(), &inputData); err == nil {
			var success, paymentId = createNewPaymentIntent(inputData.Items)

			if outputData, err := json.Marshal(map[string]any{
				"success":   success,
				"paymentId": paymentId,
			}); err == nil {
				return ctx.SendString(string(outputData))
			} else {
				fmt.Println(err)
				return ctx.SendStatus(500)
			}
		} else {
			fmt.Println(err)
			return ctx.SendStatus(400)
		}
	})

	server.Get("/get-payment-intent/:paymentId", func(ctx *fiber.Ctx) error {
		var paymentId = ctx.Params("paymentId", "")
		if paymentId == "" {
			return ctx.SendStatus(400)
		}

		if paymentIntent, ok := activePayments[paymentId]; ok {
			if outputData, err := json.Marshal(map[string]any{
				"publicKey":    stripePublicKey,
				"clientSecret": paymentIntent.clientSecret,
				"items":        paymentIntent.items,
				"netAmount":    paymentIntent.netAmount,
			}); err == nil {
				return ctx.SendString(string(outputData))
			} else {
				fmt.Println(err)
				return ctx.SendStatus(500)
			}
		} else {
			return ctx.SendStatus(404)
		}
	})
}

func init() {
	var fileBytes, _ = os.ReadFile("./PurchasableProducts.json")
	if err := json.Unmarshal(fileBytes, &products); err != nil {
		panic(err)
	}
}
