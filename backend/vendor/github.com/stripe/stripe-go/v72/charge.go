//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "encoding/json"

// ChargeFraudUserReport is the list of allowed values for reporting fraud.
type ChargeFraudUserReport string

// List of values that ChargeFraudUserReport can take.
const (
	ChargeFraudUserReportFraudulent ChargeFraudUserReport = "fraudulent"
	ChargeFraudUserReportSafe       ChargeFraudUserReport = "safe"
)

// ChargeFraudStripeReport is the list of allowed values for reporting fraud.
type ChargeFraudStripeReport string

// List of values that ChargeFraudStripeReport can take.
const (
	ChargeFraudStripeReportFraudulent ChargeFraudStripeReport = "fraudulent"
)

// ChargePaymentMethodDetailsCardPresentReceiptAccountType indicates the type of account backing a card present transaction.
type ChargePaymentMethodDetailsCardPresentReceiptAccountType string

// List of values that ChargePaymentMethodDetailsCardThreeDSecureResult can take.
const (
	ChargePaymentMethodDetailsCardPresentReceiptAccountTypeChecking ChargePaymentMethodDetailsCardPresentReceiptAccountType = "checking"
	ChargePaymentMethodDetailsCardPresentReceiptAccountTypeCredit   ChargePaymentMethodDetailsCardPresentReceiptAccountType = "credit"
	ChargePaymentMethodDetailsCardPresentReceiptAccountTypePrepaid  ChargePaymentMethodDetailsCardPresentReceiptAccountType = "prepaid"
	ChargePaymentMethodDetailsCardPresentReceiptAccountTypeUnknown  ChargePaymentMethodDetailsCardPresentReceiptAccountType = "unknown"
)

// ChargePaymentMethodDetailsCardThreeDSecureAuthenticationFlow indicates the type of 3D Secure
// authentication performed.
type ChargePaymentMethodDetailsCardThreeDSecureAuthenticationFlow string

// List of values that ChargePaymentMethodDetailsCardThreeDSecureAuthenticationFlow can take.
const (
	ChargePaymentMethodDetailsCardThreeDSecureAuthenticationFlowChallenge    ChargePaymentMethodDetailsCardThreeDSecureAuthenticationFlow = "challenge"
	ChargePaymentMethodDetailsCardThreeDSecureAuthenticationFlowFrictionless ChargePaymentMethodDetailsCardThreeDSecureAuthenticationFlow = "frictionless"
)

// ChargePaymentMethodDetailsCardThreeDSecureResult indicates the outcome of 3D Secure authentication.
type ChargePaymentMethodDetailsCardThreeDSecureResult string

// List of values that ChargePaymentMethodDetailsCardThreeDSecureResult can take.
const (
	ChargePaymentMethodDetailsCardThreeDSecureResultAttemptAcknowledged ChargePaymentMethodDetailsCardThreeDSecureResult = "attempt_acknowledged"
	ChargePaymentMethodDetailsCardThreeDSecureResultAuthenticated       ChargePaymentMethodDetailsCardThreeDSecureResult = "authenticated"
	ChargePaymentMethodDetailsCardThreeDSecureResultFailed              ChargePaymentMethodDetailsCardThreeDSecureResult = "failed"
	ChargePaymentMethodDetailsCardThreeDSecureResultNotSupported        ChargePaymentMethodDetailsCardThreeDSecureResult = "not_supported"
	ChargePaymentMethodDetailsCardThreeDSecureResultProcessingError     ChargePaymentMethodDetailsCardThreeDSecureResult = "processing_error"
)

// ChargePaymentMethodDetailsCardThreeDSecureResultReason represents dditional information about why
// 3D Secure succeeded or failed
type ChargePaymentMethodDetailsCardThreeDSecureResultReason string

// List of values that ChargePaymentMethodDetailsCardThreeDSecureResultReason can take.
const (
	ChargePaymentMethodDetailsCardThreeDSecureResultReasonAbandoned           ChargePaymentMethodDetailsCardThreeDSecureResultReason = "abandoned"
	ChargePaymentMethodDetailsCardThreeDSecureResultReasonBypassed            ChargePaymentMethodDetailsCardThreeDSecureResultReason = "bypassed"
	ChargePaymentMethodDetailsCardThreeDSecureResultReasonCanceled            ChargePaymentMethodDetailsCardThreeDSecureResultReason = "canceled"
	ChargePaymentMethodDetailsCardThreeDSecureResultReasonCardNotEnrolled     ChargePaymentMethodDetailsCardThreeDSecureResultReason = "card_not_enrolled"
	ChargePaymentMethodDetailsCardThreeDSecureResultReasonNetworkNotSupported ChargePaymentMethodDetailsCardThreeDSecureResultReason = "network_not_supported"
	ChargePaymentMethodDetailsCardThreeDSecureResultReasonProtocolError       ChargePaymentMethodDetailsCardThreeDSecureResultReason = "protocol_error"
	ChargePaymentMethodDetailsCardThreeDSecureResultReasonRejected            ChargePaymentMethodDetailsCardThreeDSecureResultReason = "rejected"
)

// The Klarna payment method used for this transaction.
// Can be one of `pay_later`, `pay_now`, `pay_with_financing`, or `pay_in_installments`
type ChargePaymentMethodDetailsKlarnaPaymentMethodCategory string

// List of values that ChargePaymentMethodDetailsKlarnaPaymentMethodCategory can take
const (
	ChargePaymentMethodDetailsKlarnaPaymentMethodCategoryPayLater          ChargePaymentMethodDetailsKlarnaPaymentMethodCategory = "pay_later"
	ChargePaymentMethodDetailsKlarnaPaymentMethodCategoryPayNow            ChargePaymentMethodDetailsKlarnaPaymentMethodCategory = "pay_now"
	ChargePaymentMethodDetailsKlarnaPaymentMethodCategoryPayWithFinancing  ChargePaymentMethodDetailsKlarnaPaymentMethodCategory = "pay_with_financing"
	ChargePaymentMethodDetailsKlarnaPaymentMethodCategoryPayInInstallments ChargePaymentMethodDetailsKlarnaPaymentMethodCategory = "pay_in_installments"
)

// ChargePaymentMethodDetailsType is the type of the PaymentMethod associated with the Charge's
// payment method details.
type ChargePaymentMethodDetailsType string

// List of values that ChargePaymentMethodDetailsType can take.
const (
	ChargePaymentMethodDetailsTypeAchCreditTransfer ChargePaymentMethodDetailsType = "ach_credit_transfer"
	ChargePaymentMethodDetailsTypeAchDebit          ChargePaymentMethodDetailsType = "ach_debit"
	ChargePaymentMethodDetailsTypeAcssDebit         ChargePaymentMethodDetailsType = "acss_debit"
	ChargePaymentMethodDetailsTypeAlipay            ChargePaymentMethodDetailsType = "alipay"
	ChargePaymentMethodDetailsTypeAUBECSDebit       ChargePaymentMethodDetailsType = "au_becs_debit"
	ChargePaymentMethodDetailsTypeBACSDebit         ChargePaymentMethodDetailsType = "bacs_debit"
	ChargePaymentMethodDetailsTypeBancontact        ChargePaymentMethodDetailsType = "bancontact"
	ChargePaymentMethodDetailsTypeCard              ChargePaymentMethodDetailsType = "card"
	ChargePaymentMethodDetailsTypeCardPresent       ChargePaymentMethodDetailsType = "card_present"
	ChargePaymentMethodDetailsTypeEps               ChargePaymentMethodDetailsType = "eps"
	ChargePaymentMethodDetailsTypeFPX               ChargePaymentMethodDetailsType = "fpx"
	ChargePaymentMethodDetailsTypeGiropay           ChargePaymentMethodDetailsType = "giropay"
	ChargePaymentMethodDetailsTypeGrabpay           ChargePaymentMethodDetailsType = "grabpay"
	ChargePaymentMethodDetailsTypeIdeal             ChargePaymentMethodDetailsType = "ideal"
	ChargePaymentMethodDetailsTypeInteracPresent    ChargePaymentMethodDetailsType = "interac_present"
	ChargePaymentMethodDetailsTypeKlarna            ChargePaymentMethodDetailsType = "klarna"
	ChargePaymentMethodDetailsTypeMultibanco        ChargePaymentMethodDetailsType = "multibanco"
	ChargePaymentMethodDetailsTypeP24               ChargePaymentMethodDetailsType = "p24"
	ChargePaymentMethodDetailsTypeSepaDebit         ChargePaymentMethodDetailsType = "sepa_debit"
	ChargePaymentMethodDetailsTypeSofort            ChargePaymentMethodDetailsType = "sofort"
	ChargePaymentMethodDetailsTypeStripeAccount     ChargePaymentMethodDetailsType = "stripe_account"
	ChargePaymentMethodDetailsTypeWechat            ChargePaymentMethodDetailsType = "wechat"
)

// ChargeLevel3LineItemsParams is the set of parameters that represent a line item on level III data.
type ChargeLevel3LineItemsParams struct {
	DiscountAmount     *int64  `form:"discount_amount"`
	ProductCode        *string `form:"product_code"`
	ProductDescription *string `form:"product_description"`
	Quantity           *int64  `form:"quantity"`
	TaxAmount          *int64  `form:"tax_amount"`
	UnitCost           *int64  `form:"unit_cost"`
}

// ChargeLevel3Params is the set of parameters that can be used for the Level III data.
type ChargeLevel3Params struct {
	CustomerReference  *string                        `form:"customer_reference"`
	LineItems          []*ChargeLevel3LineItemsParams `form:"line_items"`
	MerchantReference  *string                        `form:"merchant_reference"`
	ShippingAddressZip *string                        `form:"shipping_address_zip"`
	ShippingAmount     *int64                         `form:"shipping_amount"`
	ShippingFromZip    *string                        `form:"shipping_from_zip"`
}

// ChargeTransferDataParams is the set of parameters allowed for the transfer_data hash.
type ChargeTransferDataParams struct {
	Amount *int64 `form:"amount"`
	// This parameter can only be used on Charge creation.
	Destination *string `form:"destination"`
}

// ChargeParams is the set of parameters that can be used when creating or updating a charge.
type ChargeParams struct {
	Params                    `form:"*"`
	Amount                    *int64                    `form:"amount"`
	ApplicationFee            *int64                    `form:"application_fee"`
	ApplicationFeeAmount      *int64                    `form:"application_fee_amount"`
	Capture                   *bool                     `form:"capture"`
	Currency                  *string                   `form:"currency"`
	Customer                  *string                   `form:"customer"`
	Description               *string                   `form:"description"`
	Destination               *DestinationParams        `form:"destination"`
	ExchangeRate              *float64                  `form:"exchange_rate"`
	FraudDetails              *FraudDetailsParams       `form:"fraud_details"`
	Level3                    *ChargeLevel3Params       `form:"level3"`
	OnBehalfOf                *string                   `form:"on_behalf_of"`
	ReceiptEmail              *string                   `form:"receipt_email"`
	Shipping                  *ShippingDetailsParams    `form:"shipping"`
	Source                    *SourceParams             `form:"*"` // SourceParams has custom encoding so brought to top level with "*"
	StatementDescriptor       *string                   `form:"statement_descriptor"`
	StatementDescriptorSuffix *string                   `form:"statement_descriptor_suffix"`
	TransferData              *ChargeTransferDataParams `form:"transfer_data"`
	TransferGroup             *string                   `form:"transfer_group"`
}

// SetSource adds valid sources to a ChargeParams object,
// returning an error for unsupported sources.
func (p *ChargeParams) SetSource(sp interface{}) error {
	source, err := SourceParamsFor(sp)
	p.Source = source
	return err
}

// DestinationParams describes the parameters available for the destination hash when creating a charge.
type DestinationParams struct {
	Account *string `form:"account"`
	Amount  *int64  `form:"amount"`
}

// FraudDetailsParams provides information on the fraud details for a charge.
type FraudDetailsParams struct {
	UserReport *string `form:"user_report"`
}

// ChargeListParams is the set of parameters that can be used when listing charges.
type ChargeListParams struct {
	ListParams    `form:"*"`
	Created       *int64            `form:"created"`
	CreatedRange  *RangeQueryParams `form:"created"`
	Customer      *string           `form:"customer"`
	PaymentIntent *string           `form:"payment_intent"`
	TransferGroup *string           `form:"transfer_group"`
}

// CaptureParams is the set of parameters that can be used when capturing a charge.
type CaptureParams struct {
	Params                    `form:"*"`
	Amount                    *int64                    `form:"amount"`
	ApplicationFee            *int64                    `form:"application_fee"`
	ApplicationFeeAmount      *int64                    `form:"application_fee_amount"`
	ExchangeRate              *float64                  `form:"exchange_rate"`
	ReceiptEmail              *string                   `form:"receipt_email"`
	StatementDescriptor       *string                   `form:"statement_descriptor"`
	StatementDescriptorSuffix *string                   `form:"statement_descriptor_suffix"`
	TransferData              *ChargeTransferDataParams `form:"transfer_data"`
	TransferGroup             *string                   `form:"transfer_group"`
}

// ChargeLevel3LineItem represents a line item on level III data.
// This is in private beta and would be empty for most integrations
type ChargeLevel3LineItem struct {
	DiscountAmount     int64  `json:"discount_amount"`
	ProductCode        string `json:"product_code"`
	ProductDescription string `json:"product_description"`
	Quantity           int64  `json:"quantity"`
	TaxAmount          int64  `json:"tax_amount"`
	UnitCost           int64  `json:"unit_cost"`
}

// ChargeLevel3 represents the Level III data.
// This is in private beta and would be empty for most integrations
type ChargeLevel3 struct {
	CustomerReference  string                  `json:"customer_reference"`
	LineItems          []*ChargeLevel3LineItem `json:"line_items"`
	MerchantReference  string                  `json:"merchant_reference"`
	ShippingAddressZip string                  `json:"shipping_address_zip"`
	ShippingAmount     int64                   `json:"shipping_amount"`
	ShippingFromZip    string                  `json:"shipping_from_zip"`
}

// ChargePaymentMethodDetailsAchCreditTransfer represents details about the ACH Credit Transfer
// PaymentMethod.
type ChargePaymentMethodDetailsAchCreditTransfer struct {
	AccountNumber string `json:"account_number"`
	BankName      string `json:"bank_name"`
	RoutingNumber string `json:"routing_number"`
	SwiftCode     string `json:"swift_code"`
}

// ChargePaymentMethodDetailsAchDebit represents details about the ACH Debit PaymentMethod.
type ChargePaymentMethodDetailsAchDebit struct {
	AccountHolderType BankAccountAccountHolderType `json:"account_holder_type"`
	BankName          string                       `json:"bank_name"`
	Country           string                       `json:"country"`
	Fingerprint       string                       `json:"fingerprint"`
	Last4             string                       `json:"last4"`
	RoutingNumber     string                       `json:"routing_number"`
}

// ChargePaymentMethodDetailsAcssDebit represents details about the ACSS Debit PaymentMethod.
type ChargePaymentMethodDetailsAcssDebit struct {
	BankName          string `json:"bank_name"`
	Fingerprint       string `json:"fingerprint"`
	InstitutionNumber string `json:"institution_number"`
	Last4             string `json:"last4"`
	Mandate           string `json:"mandate"`
	TransitNumber     string `json:"transit_number"`
}

// ChargePaymentMethodDetailsAfterpayClearpay represents details about the AfterpayClearpay PaymentMethod.
type ChargePaymentMethodDetailsAfterpayClearpay struct {
	Reference string `json:"reference"`
}

// ChargePaymentMethodDetailsAlipay represents details about the Alipay PaymentMethod.
type ChargePaymentMethodDetailsAlipay struct {
	BuyerID       string `json:"buyer_id"`
	Fingerprint   string `json:"fingerprint"`
	TransactionID string `json:"transaction_id"`
}

// ChargePaymentMethodDetailsAUBECSDebit represents details about the AU BECS DD PaymentMethod.
type ChargePaymentMethodDetailsAUBECSDebit struct {
	BSBNumber   string `json:"bsb_number"`
	Fingerprint string `json:"fingerprint"`
	Last4       string `json:"last4"`
	Mandate     string `json:"mandate"`
}

// ChargePaymentMethodDetailsBACSDebit represents details about the BECS Debit PaymentMethod.
type ChargePaymentMethodDetailsBACSDebit struct {
	Fingerprint string `json:"fingerprint"`
	Last4       string `json:"last4"`
	Mandate     string `json:"mandate"`
	SortCode    string `json:"sort_code"`
}

// ChargePaymentMethodDetailsBancontact represents details about the Bancontact PaymentMethod.
type ChargePaymentMethodDetailsBancontact struct {
	BankCode                  string         `json:"bank_code"`
	BankName                  string         `json:"bank_name"`
	Bic                       string         `json:"bic"`
	GeneratedSepaDebit        *PaymentMethod `json:"generated_sepa_debit"`
	GeneratedSepaDebitMandate *Mandate       `json:"generated_sepa_debit_mandate"`
	IbanLast4                 string         `json:"iban_last4"`
	PreferredLanguage         string         `json:"preferred_language"`
	VerifiedName              string         `json:"verified_name"`
}

// ChargePaymentMethodDetailsBoleto represents details about the Boleto PaymentMethod.
type ChargePaymentMethodDetailsBoleto struct {
	TaxID string `json:"tax_id"`
}

// ChargePaymentMethodDetailsCardChecks represents the checks associated with the charge's Card
// PaymentMethod.
type ChargePaymentMethodDetailsCardChecks struct {
	AddressLine1Check      CardVerification `json:"address_line1_check"`
	AddressPostalCodeCheck CardVerification `json:"address_postal_code_check"`
	CVCCheck               CardVerification `json:"cvc_check"`
}

// ChargePaymentMethodDetailsCardInstallments represents details about the installment plan chosen
// for this charge.
type ChargePaymentMethodDetailsCardInstallments struct {
	Plan *PaymentIntentPaymentMethodOptionsCardInstallmentsPlan `json:"plan"`
}

// ChargePaymentMethodDetailsCardThreeDSecure represents details about 3DS associated with the
// charge's PaymentMethod.
type ChargePaymentMethodDetailsCardThreeDSecure struct {
	AuthenticationFlow ChargePaymentMethodDetailsCardThreeDSecureAuthenticationFlow `json:"authentication_flow"`
	Result             ChargePaymentMethodDetailsCardThreeDSecureResult             `json:"result"`
	ResultReason       ChargePaymentMethodDetailsCardThreeDSecureResultReason       `json:"result_reason"`
	Version            string                                                       `json:"version"`
}

// ChargePaymentMethodDetailsCardWalletAmexExpressCheckout represents the details of the Amex
// Express Checkout wallet.
type ChargePaymentMethodDetailsCardWalletAmexExpressCheckout struct{}

// ChargePaymentMethodDetailsCardWalletApplePay represents the details of the Apple Pay wallet.
type ChargePaymentMethodDetailsCardWalletApplePay struct{}

// ChargePaymentMethodDetailsCardWalletGooglePay represents the details of the Google Pay wallet.
type ChargePaymentMethodDetailsCardWalletGooglePay struct{}

// ChargePaymentMethodDetailsCardWalletMasterpass represents the details of the Masterpass wallet.
type ChargePaymentMethodDetailsCardWalletMasterpass struct {
	BillingAddress  *Address `json:"billing_address"`
	Email           string   `json:"email"`
	Name            string   `json:"name"`
	ShippingAddress *Address `json:"shipping_address"`
}

// ChargePaymentMethodDetailsCardWalletSamsungPay represents the details of the Samsung Pay wallet.
type ChargePaymentMethodDetailsCardWalletSamsungPay struct{}

// ChargePaymentMethodDetailsCardWalletVisaCheckout represents the details of the Visa Checkout
// wallet.
type ChargePaymentMethodDetailsCardWalletVisaCheckout struct {
	BillingAddress  *Address `json:"billing_address"`
	Email           string   `json:"email"`
	Name            string   `json:"name"`
	ShippingAddress *Address `json:"shipping_address"`
}

// ChargePaymentMethodDetailsCardWallet represents the details of the card wallet if this Card
// PaymentMethod is part of a card wallet.
type ChargePaymentMethodDetailsCardWallet struct {
	AmexExpressCheckout *ChargePaymentMethodDetailsCardWalletAmexExpressCheckout `json:"amex_express_checkout"`
	ApplePay            *ChargePaymentMethodDetailsCardWalletApplePay            `json:"apple_pay"`
	DynamicLast4        string                                                   `json:"dynamic_last4"`
	GooglePay           *ChargePaymentMethodDetailsCardWalletGooglePay           `json:"google_pay"`
	Masterpass          *ChargePaymentMethodDetailsCardWalletMasterpass          `json:"masterpass"`
	SamsungPay          *ChargePaymentMethodDetailsCardWalletSamsungPay          `json:"samsung_pay"`
	Type                PaymentMethodCardWalletType                              `json:"type"`
	VisaCheckout        *ChargePaymentMethodDetailsCardWalletVisaCheckout        `json:"visa_checkout"`
}

// ChargePaymentMethodDetailsCard represents details about the Card PaymentMethod.
type ChargePaymentMethodDetailsCard struct {
	Brand        PaymentMethodCardBrand                      `json:"brand"`
	Checks       *ChargePaymentMethodDetailsCardChecks       `json:"checks"`
	Country      string                                      `json:"country"`
	ExpMonth     uint64                                      `json:"exp_month"`
	ExpYear      uint64                                      `json:"exp_year"`
	Fingerprint  string                                      `json:"fingerprint"`
	Funding      CardFunding                                 `json:"funding"`
	Installments *ChargePaymentMethodDetailsCardInstallments `json:"installments"`
	Last4        string                                      `json:"last4"`
	MOTO         bool                                        `json:"moto"`
	Network      PaymentMethodCardNetwork                    `json:"network"`
	ThreeDSecure *ChargePaymentMethodDetailsCardThreeDSecure `json:"three_d_secure"`
	Wallet       *ChargePaymentMethodDetailsCardWallet       `json:"wallet"`

	// Please note that the fields below are for internal use only and are not returned
	// as part of standard API requests.
	Description string `json:"description"`
	IIN         string `json:"iin"`
	Issuer      string `json:"issuer"`
}

// ChargePaymentMethodDetailsCardPresentReceipt represents details about the receipt on a
// Card Present PaymentMethod.
type ChargePaymentMethodDetailsCardPresentReceipt struct {
	AccountType                  ChargePaymentMethodDetailsCardPresentReceiptAccountType `json:"account_type"`
	ApplicationCryptogram        string                                                  `json:"application_cryptogram"`
	ApplicationPreferredName     string                                                  `json:"application_preferred_name"`
	AuthorizationCode            string                                                  `json:"authorization_code"`
	AuthorizationResponseCode    string                                                  `json:"authorization_response_code"`
	CardholderVerificationMethod string                                                  `json:"cardholder_verification_method"`
	DedicatedFileName            string                                                  `json:"dedicated_file_name"`
	TerminalVerificationResults  string                                                  `json:"terminal_verification_results"`
	TransactionStatusInformation string                                                  `json:"transaction_status_information"`
}

// ChargePaymentMethodDetailsCardPresent represents details about the Card Present PaymentMethod.
type ChargePaymentMethodDetailsCardPresent struct {
	AmountAuthorized     int64                                         `json:"amount_authorized"`
	Brand                PaymentMethodCardBrand                        `json:"brand"`
	CardholderName       string                                        `json:"cardholder_name"`
	Country              string                                        `json:"country"`
	EmvAuthData          string                                        `json:"emv_auth_data"`
	ExpMonth             uint64                                        `json:"exp_month"`
	ExpYear              uint64                                        `json:"exp_year"`
	Fingerprint          string                                        `json:"fingerprint"`
	Funding              CardFunding                                   `json:"funding"`
	GeneratedCard        string                                        `json:"generated_card"`
	Last4                string                                        `json:"last4"`
	Network              PaymentMethodCardNetwork                      `json:"network"`
	OvercaptureSupported bool                                          `json:"overcapture_supported"`
	ReadMethod           string                                        `json:"read_method"`
	Receipt              *ChargePaymentMethodDetailsCardPresentReceipt `json:"receipt"`

	// Please note that the fields below are for internal use only and are not returned
	// as part of standard API requests.
	Description string `json:"description"`
	IIN         string `json:"iin"`
	Issuer      string `json:"issuer"`
}

// ChargePaymentMethodDetailsEps represents details about the EPS PaymentMethod.
type ChargePaymentMethodDetailsEps struct {
	Bank         string `json:"bank"`
	VerifiedName string `json:"verified_name"`
}

// ChargePaymentMethodDetailsFPX represents details about the FPX PaymentMethod.
type ChargePaymentMethodDetailsFPX struct {
	AccountHolderType PaymentMethodFPXAccountHolderType `json:"account_holder_type"`
	Bank              string                            `json:"bank"`
	TransactionID     string                            `json:"transaction_id"`
}

// ChargePaymentMethodDetailsGiropay represents details about the Giropay PaymentMethod.
type ChargePaymentMethodDetailsGiropay struct {
	BankCode     string `json:"bank_code"`
	BankName     string `json:"bank_name"`
	Bic          string `json:"bic"`
	VerifiedName string `json:"verified_name"`
}

// ChargePaymentMethodDetailsGrabpay represents details about the Grabpay PaymentMethod.
type ChargePaymentMethodDetailsGrabpay struct {
	TransactionID string `json:"transaction_id"`
}

// ChargePaymentMethodDetailsIdeal represents details about the Ideal PaymentMethod.
type ChargePaymentMethodDetailsIdeal struct {
	Bank                      string         `json:"bank"`
	Bic                       string         `json:"bic"`
	GeneratedSepaDebit        *PaymentMethod `json:"generated_sepa_debit"`
	GeneratedSepaDebitMandate *Mandate       `json:"generated_sepa_debit_mandate"`
	IbanLast4                 string         `json:"iban_last4"`
	VerifiedName              string         `json:"verified_name"`
}

// ChargePaymentMethodDetailsInteracPresent represents details about the InteracPresent PaymentMethod.
type ChargePaymentMethodDetailsInteracPresent struct {
	Brand            string                                           `json:"brand"`
	CardholderName   string                                           `json:"cardholder_name"`
	Country          string                                           `json:"country"`
	EmvAuthData      string                                           `json:"emv_auth_data"`
	ExpMonth         int64                                            `json:"exp_month"`
	ExpYear          int64                                            `json:"exp_year"`
	Fingerprint      string                                           `json:"fingerprint"`
	Funding          string                                           `json:"funding"`
	GeneratedCard    string                                           `json:"generated_card"`
	Last4            string                                           `json:"last4"`
	Network          string                                           `json:"network"`
	PreferredLocales []string                                         `json:"preferred_locales"`
	ReadMethod       string                                           `json:"read_method"`
	Receipt          *ChargePaymentMethodDetailsInteracPresentReceipt `json:"receipt"`

	// Please note that the fields below are for internal use only and are not returned
	// as part of standard API requests.
	Description string `json:"description"`
	IIN         string `json:"iin"`
	Issuer      string `json:"issuer"`
}

// ChargePaymentMethodDetailsInteracPresentReceipt represents details about the InteracPresent Receipt.
type ChargePaymentMethodDetailsInteracPresentReceipt struct {
	AccountType                  string `json:"account_type"`
	ApplicationCryptogram        string `json:"application_cryptogram"`
	ApplicationPreferredName     string `json:"application_preferred_name"`
	AuthorizationCode            string `json:"authorization_code"`
	AuthorizationResponseCode    string `json:"authorization_response_code"`
	CardholderVerificationMethod string `json:"cardholder_verification_method"`
	DedicatedFileName            string `json:"dedicated_file_name"`
	TerminalVerificationResults  string `json:"terminal_verification_results"`
	TransactionStatusInformation string `json:"transaction_status_information"`
}

// ChargePaymentMethodDetailsKlarna represents details for the Klarna
// PaymentMethod.
type ChargePaymentMethodDetailsKlarna struct {
	PaymentMethodCategory ChargePaymentMethodDetailsKlarnaPaymentMethodCategory `json:"payment_method_category"`
	PreferredLocale       string                                                `json:"preferred_locale"`
}

// ChargePaymentMethodDetailsMultibanco represents details about the Multibanco PaymentMethod.
type ChargePaymentMethodDetailsMultibanco struct {
	Entity    string `json:"entity"`
	Reference string `json:"reference"`
}

// ChargePaymentMethodDetailsOXXO represents details about the OXXO PaymentMethod.
type ChargePaymentMethodDetailsOXXO struct {
	Number string `json:"number"`
}

// ChargePaymentMethodDetailsP24 represents details about the P24 PaymentMethod.
type ChargePaymentMethodDetailsP24 struct {
	Bank         string `json:"bank"`
	Reference    string `json:"reference"`
	VerifiedName string `json:"verified_name"`
}
type ChargePaymentMethodDetailsSepaCreditTransfer struct {
	BankName string `json:"bank_name"`
	Bic      string `json:"bic"`
	Iban     string `json:"iban"`
}

// ChargePaymentMethodDetailsSepaDebit represents details about the Sepa Debit PaymentMethod.
type ChargePaymentMethodDetailsSepaDebit struct {
	BankCode    string   `json:"bank_code"`
	BranchCode  string   `json:"branch_code"`
	Country     string   `json:"country"`
	Fingerprint string   `json:"fingerprint"`
	Last4       string   `json:"last4"`
	Mandate     *Mandate `json:"mandate"`
}

// ChargePaymentMethodDetailsSofort represents details about the Sofort PaymentMethod.
type ChargePaymentMethodDetailsSofort struct {
	BankCode                  string         `json:"bank_code"`
	BankName                  string         `json:"bank_name"`
	Bic                       string         `json:"bic"`
	Country                   string         `json:"country"`
	GeneratedSepaDebit        *PaymentMethod `json:"generated_sepa_debit"`
	GeneratedSepaDebitMandate *Mandate       `json:"generated_sepa_debit_mandate"`
	IbanLast4                 string         `json:"iban_last4"`
	PreferredLanguage         string         `json:"preferred_language"`
	VerifiedName              string         `json:"verified_name"`
}

// ChargePaymentMethodDetailsStripeAccount represents details about the StripeAccount PaymentMethod.
type ChargePaymentMethodDetailsStripeAccount struct{}

// ChargePaymentMethodDetailsWechat represents details about the Wechat PaymentMethod.
type ChargePaymentMethodDetailsWechat struct{}

// ChargePaymentMethodDetailsWechatPay represents details about the WechatPay PaymentMethod.
type ChargePaymentMethodDetailsWechatPay struct {
	Fingerprint   string `json:"fingerprint"`
	TransactionID string `json:"transaction_id"`
}

// ChargePaymentMethodDetails represents the details about the PaymentMethod associated with the
// charge.
type ChargePaymentMethodDetails struct {
	AchCreditTransfer  *ChargePaymentMethodDetailsAchCreditTransfer  `json:"ach_credit_transfer"`
	AchDebit           *ChargePaymentMethodDetailsAchDebit           `json:"ach_debit"`
	AcssDebit          *ChargePaymentMethodDetailsAcssDebit          `json:"acss_debit"`
	AfterpayClearpay   *ChargePaymentMethodDetailsAfterpayClearpay   `json:"afterpay_clearpay"`
	Alipay             *ChargePaymentMethodDetailsAlipay             `json:"alipay"`
	AUBECSDebit        *ChargePaymentMethodDetailsAUBECSDebit        `json:"au_becs_debit"`
	BACSDebit          *ChargePaymentMethodDetailsBACSDebit          `json:"bacs_debit"`
	Bancontact         *ChargePaymentMethodDetailsBancontact         `json:"bancontact"`
	Boleto             *ChargePaymentMethodDetailsBoleto             `json:"boleto"`
	Card               *ChargePaymentMethodDetailsCard               `json:"card"`
	CardPresent        *ChargePaymentMethodDetailsCardPresent        `json:"card_present"`
	Eps                *ChargePaymentMethodDetailsEps                `json:"eps"`
	FPX                *ChargePaymentMethodDetailsFPX                `json:"fpx"`
	Giropay            *ChargePaymentMethodDetailsGiropay            `json:"giropay"`
	Grabpay            *ChargePaymentMethodDetailsGrabpay            `json:"grabpay"`
	Ideal              *ChargePaymentMethodDetailsIdeal              `json:"ideal"`
	InteracPresent     *ChargePaymentMethodDetailsInteracPresent     `json:"interac_present"`
	Klarna             *ChargePaymentMethodDetailsKlarna             `json:"klarna"`
	Multibanco         *ChargePaymentMethodDetailsMultibanco         `json:"multibanco"`
	OXXO               *ChargePaymentMethodDetailsOXXO               `json:"oxxo"`
	P24                *ChargePaymentMethodDetailsP24                `json:"p24"`
	SepaCreditTransfer *ChargePaymentMethodDetailsSepaCreditTransfer `json:"sepa_credit_transfer"`
	SepaDebit          *ChargePaymentMethodDetailsSepaDebit          `json:"sepa_debit"`
	Sofort             *ChargePaymentMethodDetailsSofort             `json:"sofort"`
	StripeAccount      *ChargePaymentMethodDetailsStripeAccount      `json:"stripe_account"`
	Type               ChargePaymentMethodDetailsType                `json:"type"`
	Wechat             *ChargePaymentMethodDetailsWechat             `json:"wechat"`
	WechatPay          *ChargePaymentMethodDetailsWechatPay          `json:"wechat_pay"`
}

// ChargeTransferData represents the information for the transfer_data associated with a charge.
type ChargeTransferData struct {
	Amount      int64    `json:"amount"`
	Destination *Account `json:"destination"`
}

// Charge is the resource representing a Stripe charge.
// For more details see https://stripe.com/docs/api#charges.
type Charge struct {
	APIResource
	Amount                        int64                       `json:"amount"`
	AmountCaptured                int64                       `json:"amount_captured"`
	AmountRefunded                int64                       `json:"amount_refunded"`
	Application                   *Application                `json:"application"`
	ApplicationFee                *ApplicationFee             `json:"application_fee"`
	ApplicationFeeAmount          int64                       `json:"application_fee_amount"`
	AuthorizationCode             string                      `json:"authorization_code"`
	BalanceTransaction            *BalanceTransaction         `json:"balance_transaction"`
	BillingDetails                *BillingDetails             `json:"billing_details"`
	CalculatedStatementDescriptor string                      `json:"calculated_statement_descriptor"`
	Captured                      bool                        `json:"captured"`
	Created                       int64                       `json:"created"`
	Currency                      Currency                    `json:"currency"`
	Customer                      *Customer                   `json:"customer"`
	Description                   string                      `json:"description"`
	Destination                   *Account                    `json:"destination"`
	Dispute                       *Dispute                    `json:"dispute"`
	Disputed                      bool                        `json:"disputed"`
	FailureCode                   string                      `json:"failure_code"`
	FailureMessage                string                      `json:"failure_message"`
	FraudDetails                  *FraudDetails               `json:"fraud_details"`
	ID                            string                      `json:"id"`
	Invoice                       *Invoice                    `json:"invoice"`
	Level3                        ChargeLevel3                `json:"level3"`
	Livemode                      bool                        `json:"livemode"`
	Metadata                      map[string]string           `json:"metadata"`
	Object                        string                      `json:"object"`
	OnBehalfOf                    *Account                    `json:"on_behalf_of"`
	Order                         *Order                      `json:"order"`
	Outcome                       *ChargeOutcome              `json:"outcome"`
	Paid                          bool                        `json:"paid"`
	PaymentIntent                 *PaymentIntent              `json:"payment_intent"`
	PaymentMethod                 string                      `json:"payment_method"`
	PaymentMethodDetails          *ChargePaymentMethodDetails `json:"payment_method_details"`
	ReceiptEmail                  string                      `json:"receipt_email"`
	ReceiptNumber                 string                      `json:"receipt_number"`
	ReceiptURL                    string                      `json:"receipt_url"`
	Refunded                      bool                        `json:"refunded"`
	Refunds                       *RefundList                 `json:"refunds"`
	Review                        *Review                     `json:"review"`
	Shipping                      *ShippingDetails            `json:"shipping"`
	Source                        *PaymentSource              `json:"source"`
	SourceTransfer                *Transfer                   `json:"source_transfer"`
	StatementDescriptor           string                      `json:"statement_descriptor"`
	StatementDescriptorSuffix     string                      `json:"statement_descriptor_suffix"`
	Status                        string                      `json:"status"`
	Transfer                      *Transfer                   `json:"transfer"`
	TransferData                  *ChargeTransferData         `json:"transfer_data"`
	TransferGroup                 string                      `json:"transfer_group"`
}

// FraudDetails is the structure detailing fraud status.
type FraudDetails struct {
	StripeReport ChargeFraudStripeReport `json:"stripe_report"`
	UserReport   ChargeFraudUserReport   `json:"user_report"`
}

// ChargeOutcomeRule tells you the Radar rule that blocked the charge, if any.
type ChargeOutcomeRule struct {
	Action    string `json:"action"`
	ID        string `json:"id"`
	Predicate string `json:"predicate"`
}

// ChargeOutcome is the charge's outcome that details whether a payment
// was accepted and why.
type ChargeOutcome struct {
	NetworkStatus string             `json:"network_status"`
	Reason        string             `json:"reason"`
	RiskLevel     string             `json:"risk_level"`
	RiskScore     int64              `json:"risk_score"`
	Rule          *ChargeOutcomeRule `json:"rule"`
	SellerMessage string             `json:"seller_message"`
	Type          string             `json:"type"`
}

// UnmarshalJSON handles deserialization of a ChargeOutcomeRule.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (c *ChargeOutcomeRule) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		c.ID = id
		return nil
	}

	type chargeOutcomeRule ChargeOutcomeRule
	var v chargeOutcomeRule
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*c = ChargeOutcomeRule(v)
	return nil
}

// ChargeList is a list of charges as retrieved from a list endpoint.
type ChargeList struct {
	APIResource
	ListMeta
	Data []*Charge `json:"data"`
}

// UnmarshalJSON handles deserialization of a Charge.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (c *Charge) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		c.ID = id
		return nil
	}

	type charge Charge
	var v charge
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*c = Charge(v)
	return nil
}