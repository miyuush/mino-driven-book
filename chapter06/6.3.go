package main

type PurchaseHistory struct {
	TotalAmount int
	PurchaseFrequencyPerMonth int
	ReturnRate float64
}

// 優良顧客のルール
type ExcellentCustomerRule interface {
	Ok(history PurchaseHistory) bool
}

// ゴールド会員の購入金額のルール
type GoldCustomerPurchaseAmountRule struct {}

func (gcpar GoldCustomerPurchaseAmountRule) Ok(history PurchaseHistory) bool {
	return 10000 <= history.TotalAmount
}

func NewGoldCustomerPurchaseAmountRule() *GoldCustomerPurchaseAmountRule {
	return &GoldCustomerPurchaseAmountRule{}
}

// 購入頻度のルール
type PurchaseFrequencyRule struct {}

func (pfr PurchaseFrequencyRule) Ok(history PurchaseHistory) bool {
	return 10 <= history.PurchaseFrequencyPerMonth
}

func NewPurchaseFrequencyRule() *PurchaseFrequencyRule {
	return &PurchaseFrequencyRule{}
}

// 返品率のルール
type ReturnRateRule struct {}

func (rrr ReturnRateRule) Ok(history PurchaseHistory) bool {
	return history.ReturnRate <= 0.001
}

func NewReturnRateRule() *ReturnRateRule {
	return &ReturnRateRule{}
}

type ExcellentCustomerPolicy struct {
	Rules []ExcellentCustomerRule
}

// 優良顧客の方針を表現するクラス
func NewExcellentCustomerPolicy() *ExcellentCustomerPolicy {
	return &ExcellentCustomerPolicy{
		[]ExcellentCustomerRule{},
	}
}

// ルールを追加する
func (ecp *ExcellentCustomerPolicy) Add(rule ExcellentCustomerRule) {
	ecp.Rules = append(ecp.Rules, rule)
}

func (ecp ExcellentCustomerPolicy) ComplyWithAll(history PurchaseHistory) bool {
	for _, rule := range ecp.Rules {
		if !rule.Ok(history) {
			return false
		}
	}
	return true
}

// ゴールド会員の方針
type GoldCustomerPolicy struct {
	Policy ExcellentCustomerPolicy
}

func NewGoldCustomerPolicy() *GoldCustomerPolicy {
	policy := NewExcellentCustomerPolicy()
	policy.Add(NewGoldCustomerPurchaseAmountRule())
	policy.Add(NewPurchaseFrequencyRule())
	policy.Add(NewReturnRateRule())
	return &GoldCustomerPolicy{*policy}
}

func (gcp GoldCustomerPolicy) ComplyWithAll(history PurchaseHistory) bool {
	return gcp.Policy.ComplyWithAll(history)
}

// シルバー会員の方針
type SilverCustomerPolicy struct {
	Policy ExcellentCustomerPolicy
}

func NewSilverCustomerPolicy() *SilverCustomerPolicy {
	policy := NewExcellentCustomerPolicy()
	policy.Add(NewPurchaseFrequencyRule())
	policy.Add(NewReturnRateRule())
	return &SilverCustomerPolicy{*policy}
}

func (scp SilverCustomerPolicy) ComplyWithAll(history PurchaseHistory) bool {
	return scp.Policy.ComplyWithAll(history)
}
