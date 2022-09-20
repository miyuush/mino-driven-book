package main

type MagicPoint struct {
	Value int
}

func NewMagicPoint(magicPoint int) *MagicPoint {
	return &MagicPoint{magicPoint}
}

type AttackPower struct {
	Value int
}

func NewAttackPower(attackPower int) *AttackPower {
	return &AttackPower{attackPower}
}

type TechnicalPoint struct {
	Value int
}

func NewTechnicalPoint(technicalPoint int) *TechnicalPoint {
	return &TechnicalPoint{technicalPoint}
}

type Member struct {
	agility int
	level int
	magicAttack int
	vitality int
}

type Magic interface {
	Name() string
	CostMagicPoint() MagicPoint
	AttackPower() AttackPower
	CostTechnicalPoint() TechnicalPoint
}

// ファイア
type Fire struct {
	member Member
}

func (f Fire) Name() string {
	return "ファイア"
}

func (f Fire) CostMagicPoint() *MagicPoint {
	return NewMagicPoint(2)
}

func (f Fire) AttackPower() *AttackPower {
	value := 20 + int(float64(f.member.level) * 0.5)
	return NewAttackPower(value)
}

func (f Fire) CostTechnicalPoint() *TechnicalPoint {
	return NewTechnicalPoint(0)
}

func NewFire(member Member) *Fire {
	return &Fire{member: member}
}

// 紫電
type Shiden struct {
	member Member
}

func (s Shiden) Name() string {
	return "紫電"
}

func (s Shiden) CostMagicPoint() *MagicPoint {
	value := 5 + int(float64(s.member.level) * 0.2)
	return NewMagicPoint(value)
}

func (s Shiden) AttackPower() *AttackPower {
	value := 50 + int(float64(s.member.agility) * 1.5)
	return NewAttackPower(value)
}

func (s Shiden) CostTechnicalPoint() *TechnicalPoint {
	return NewTechnicalPoint(5)
}

func NewShiden(member Member) *Shiden {
	return &Shiden{member: member}
}

// 地獄の業火
type HellFire struct {
	member Member
}

func (h HellFire) Name() string {
	return "地獄の業火"
}

func (h HellFire) CostMagicPoint() *MagicPoint {
	return NewMagicPoint(16)
}

func (h HellFire) AttackPower() *AttackPower {
	value := 200 + int(float64(h.member.magicAttack) * 0.5 + float64(h.member.vitality))
	return NewAttackPower(value)
}

func (h HellFire) CostTechnicalPoint() *TechnicalPoint {
	value := 20 + int(float64(h.member.level) * 0.4)
	return NewTechnicalPoint(value)
}

func NewHellFire(member Member) *HellFire {
	return &HellFire{member: member}
}
