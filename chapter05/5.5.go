package main

import "math"

type MagicPoint struct {
	CurrentAmount int
	OriginalMaxAmount int
	MaxIncrements []int
}

// 現在の魔法力残量
func (mp MagicPoint) Current() int {
	return mp.CurrentAmount
}

// 魔法力の最大値
func (mp MagicPoint) Max() int {
	amount := mp.OriginalMaxAmount
	for _, each :=  range mp.MaxIncrements {
		amount += each
	}
	return amount
}

// 魔法力を回復する
func (mp *MagicPoint) Recover(recoveryAmount int) {
	mp.CurrentAmount = int(math.Min(float64(mp.CurrentAmount) + float64(recoveryAmount), float64(mp.Max())))
}

// 魔法力を消費する
func (mp *MagicPoint) Consume(consumeAmount int) {}

func NewMagicPoint(currentAmount int) *MagicPoint {
	return &MagicPoint{
		CurrentAmount: currentAmount,
		OriginalMaxAmount: 100,
		MaxIncrements: []int{},
	}
}
