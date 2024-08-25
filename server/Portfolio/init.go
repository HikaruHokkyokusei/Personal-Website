package Portfolio

import (
	PPS "Server/Portfolio/Payment"
	"Server/Utils"
	"github.com/gofiber/fiber/v2"
)

var (
	stripePrivateKey string
	stripePublicKey  string
)

func init() {
	stripePrivateKey = Utils.GetDefaultEnv("StripePrivateKey", "")
	stripePublicKey = Utils.GetDefaultEnv("StripePublicKey", "")
}

func ConfigurePortfolioEndpoints(server fiber.Router) {
	PPS.ConfigurePaymentEndpoints(server.Group("/portfolio"), stripePublicKey, stripePrivateKey)
}
