package main

const (
	MIN int = 0
	MAX int = 999
)

// ヒットポイント（HP）を表現するクラス
type HitPoint struct {
	Value int
}

// ダメージを受ける
func (h HitPoint) Damage(damageAmount int) *HitPoint {
	var corrected int
	damaged := h.Value - damageAmount
	if damaged < MIN {
		corrected = MIN
	} else {
		corrected = damaged
	}
	return NewHitPoint(corrected)
}

// 回復する
func (h HitPoint) Recover(recoveryAmount int) *HitPoint {
	var corrected int
	recoverd := h.Value + recoveryAmount
	if MAX < recoverd {
		corrected = MAX
	} else {
		corrected = recoverd
	}
	return NewHitPoint(corrected)
}

func NewHitPoint(value int) *HitPoint {
	return &HitPoint{value}
}
