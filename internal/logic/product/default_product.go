package product

import (
	"context"
	"unibee/internal/logic/merchant_config/update"
	entity "unibee/internal/model/entity/default"
	"unibee/internal/query"
	"unibee/utility"
)

func UpdateDefaultProd(ctx context.Context, product *entity.Product) error {
	utility.Assert(product != nil, "invalid product data")
	utility.Assert(product.MerchantId > 0, "invalid product merchant")
	product.Id = 0
	return update.SetMerchantConfig(ctx, product.MerchantId, query.MerchantDefaultProductKey, utility.MarshalToJsonString(product))
}
