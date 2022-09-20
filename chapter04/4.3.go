package main

import (
	"fmt"
)

type HitPoint struct {
	Amount int
}

func (hp *HitPoint) Damage(damageAmount int) {
	nextAmount := hp.Amount - damageAmount
	if nextAmount < MIN {
		hp.Amount = MIN
	} else {
		hp.Amount = nextAmount
	}
}

func (hp HitPoint) IsZero() bool {
	return hp.Amount == MIN
}

func NewHitPoint(amount int) (*HitPoint, error) {
	if amount < MIN {
		return nil, fmt.Errorf("ヒットポイントには0以上を指定してください。")
	}

	return &HitPoint{amount}, nil
}

type Member struct {
	HitPoint HitPoint
	// States States
}

func (m *Member) Damage(damageAmount int) {
	m.HitPoint.Damage(damageAmount)
	// if m.HitPoint.IsZero() {
	// 	m.States.Add(StatesType.dead)
	// }
}
