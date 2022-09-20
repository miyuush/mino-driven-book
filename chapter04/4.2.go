package main

import (
	"fmt"
	"log"
	"os"
)

const MIN int = 0

// 攻撃力を表現する
type AttackPower struct {
	Value int
}

// 攻撃力を強化する
func (ap AttackPower) ReinForce(increment AttackPower) (*AttackPower, error) {
	return NewAttackPower(ap.Value + increment.Value)
}

// 無力化する
func (ap AttackPower) Disable() (*AttackPower, error) {
	return NewAttackPower(MIN)
}

func NewAttackPower(value int) (*AttackPower, error) {
	if value < MIN {
		return nil, fmt.Errorf("攻撃力には0以上を指定してください。")
	}

	return &AttackPower{value}, nil
}

// 武器を表現する
type Weapon struct {
	AttackPower AttackPower
}

// 武器を強化する
func (w Weapon) ReinForce(increment AttackPower) (*Weapon, error) {
	reinForced, err := w.AttackPower.ReinForce(increment)
	if err != nil {
		return nil, err
	}

	return NewWeapon(*reinForced), nil
}

func NewWeapon(attackpower AttackPower) *Weapon {
	return &Weapon{attackpower}
}

func main() {
	attackPowerA, err := NewAttackPower(20)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	attackPowerB, err := NewAttackPower(20)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	weaponA := NewWeapon(*attackPowerA)
	weaponB := NewWeapon(*attackPowerB)

	increment, err := NewAttackPower(5)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	reinForcedWeaponA, err := weaponA.ReinForce(*increment)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	fmt.Println("Weapon A attack power :", weaponA.AttackPower.Value)
	fmt.Println("Reinforced weapon A attack power:", reinForcedWeaponA.AttackPower.Value)
	fmt.Println("Weapon B attack power:", weaponB.AttackPower.Value)
}
