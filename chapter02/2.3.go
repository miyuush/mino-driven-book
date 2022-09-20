package main

import "fmt"

// プレイヤーの攻撃力を合算する
func sumUpPlayerAttackPower(playerArmPower, playerWeaponPower int) int {
	return playerArmPower + playerWeaponPower
}

// 敵の防御力を合算する
func sumUpEnemyDefence(enemyBodyDefence, enemyArmorDefence int) int {
	return enemyBodyDefence + enemyArmorDefence
}

// ダメージ量を計算する
func estimateDamage(totalPlayerAttackPower, totalEnemyDefence int) int {
	damageAmount := totalPlayerAttackPower - (totalEnemyDefence / 2)
	if damageAmount < 0 {
		return 0
	}
	return damageAmount
}

func main() {
	var (
		playerBodyPower, playerWeaponPower  int
		enemyBodyDefence, enemyArmorDefence int
	)

	totalPlayerAttackPower := sumUpPlayerAttackPower(playerBodyPower, playerWeaponPower)
	totalEnemyDefence := sumUpEnemyDefence(enemyBodyDefence, enemyArmorDefence)
	damageAmount := estimateDamage(totalPlayerAttackPower, totalEnemyDefence)

	fmt.Println(damageAmount)
}
