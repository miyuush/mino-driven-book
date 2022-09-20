package main

import (
	"fmt"
	"log"
	"os"
)

const (
	MIN_POINT int = 0
	STANDARD_MEMBERSHIP_POINT int = 3000
	PREMIUM_MEMBERSHIP_POINT int = 10000
)

type GiftPoint struct {
	Value int
}

// 略

// パッケージ外部からはインスタンス生成できない
func newGiftPoint(point int) (*GiftPoint, error) {
	if point < MIN_POINT {
		return nil, fmt.Errorf("ポイントが0以上ではありません。")
	}

	return &GiftPoint{point}, nil
}

// 標準会員向けギフトポイントを返す
func ForStandardMembership() (*GiftPoint, error) {
	return newGiftPoint(STANDARD_MEMBERSHIP_POINT)
}

// プレミアム会員向けギフトポイントを返す
func ForPremiumMembership() (*GiftPoint, error) {
	return newGiftPoint(PREMIUM_MEMBERSHIP_POINT)
}

func main() {
	standardMemberShipPoint, err := ForStandardMembership()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	premiumMemberShipPoint, err := ForPremiumMembership()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	fmt.Println(standardMemberShipPoint, premiumMemberShipPoint)
}
