package bean

import entity "unibee/internal/model/entity/default"

const MerchantConfigGroupDefault = "default"
const MerchantConfigGroupSubscription = "subscription"
const MerchantConfigGroupLocalization = "localization"

type MerchantConfigEntity struct {
	Group   string            `json:"group"`
	Configs []*MerchantConfig `json:"configs"`
}

type MerchantConfig struct {
	Group       string      `json:"group"`
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Key         string      `json:"key"`
	Value       interface{} `json:"value"`
}

type MerchantCountryConfig struct {
	MerchantId  uint64 `json:"merchantId" description:"merchant id"` // merchant id
	CountryCode string `json:"countryCode" description:""`           //
	Name        string `json:"name"        description:""`           //
}

func SimplifyMerchantCountryConfig(one *entity.MerchantCountryConfig) *MerchantCountryConfig {
	if one == nil {
		return nil
	}
	return &MerchantCountryConfig{
		MerchantId:  one.MerchantId,
		CountryCode: one.CountryCode,
		Name:        one.Name,
	}
}

type MerchantVatRule struct {
	GatewayNames      string `json:"gatewayNames" dc:""`
	ValidCountryCodes string `json:"validCountryCodes" dc:""`
	TaxPercentage     *int64 `json:"taxPercentage" dc:""`
	IgnoreVatNumber   bool   `json:"ignoreVatNumber" dc:""`
}
