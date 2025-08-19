package client

// GetMatchType bit 0-3： 比赛类型
// GameTypeRoboMasterUC, GameTypeRoboMasterCS, GameTypeRoboMasterICRA
// GameTypeRoboMasterUL, GameTypeRoboMasterULStandard
func (this *GameStateData) GetMatchType() byte {
	return this.GameState & 0x0F
}

func (this *GameStateData) SetMatchType(matchType byte) {
	this.GameState = (this.GameState & 0xF0) | (matchType & 0x0F)
}

// GetMatchStage bit 4-7： 比赛阶段
// GameStageDefault, GameStagePrepare, GameStageJudgeSystemCheck
// GameStageCountdown, GameStageRunning, GameStageEnd
func (this *GameStateData) GetMatchStage() byte {
	return (this.GameState >> 4) & 0x0F
}

func (this *GameStateData) SetMatchStage(matchStage byte) {
	this.GameState = (this.GameState & 0x0F) | ((matchStage & 0x0F) << 4)
}

// GetAmmoExchanged bit 0-10：除远程兑换外，哨兵机器人成功兑换的允许发弹量，开局为 0，在
// 哨兵机器人成功兑换一定允许发弹量后，该值将变为哨兵机器人成功兑换的
// 允许发弹量值。
func (this *SentryInfoData) GetAmmoExchanged() int {
	// get bit 0-10 of uint32 SentryInfo
	return int(this.SentryInfo & 0x7FF)
}

func (this *SentryInfoData) SetAmmoExchanged(ammo int) {
	// set bit 0-10 of uint32 SentryInfo
	this.SentryInfo = (this.SentryInfo & 0xFFFFF800) | (uint32(ammo) & 0x7FF)
}

// GetRemoteExchangeAmmoCount bit bit 11-14：哨兵机器人成功远程兑换允许发弹量的次数，开局为 0，在哨兵
// 机器人成功远程兑换允许发弹量后，该值将变为哨兵机器人成功远程兑换允
// 许发弹量的次数。
func (this *SentryInfoData) GetRemoteExchangeAmmoCount() int {
	// get bit 11-14 of uint32 SentryInfo
	return int((this.SentryInfo >> 11) & 0x0F)
}

func (this *SentryInfoData) SetRemoteExchangeAmmoCount(count int) {
	// set bit 11-14 of uint32 SentryInfo
	this.SentryInfo = (this.SentryInfo & 0xFFFFF0FF) | ((uint32(count) & 0x0F) << 11)
}

// GetRemoteExchangeHealthCount 哨兵机器人成功远程兑换血量的次数，开局为 0，在哨兵机器人
// 成功远程兑换血量后，该值将变为哨兵机器人成功远程兑换血量的次数。
func (this *SentryInfoData) GetRemoteExchangeHealthCount() int {
	// get bit 15-18 of uint32 SentryInfo
	return int((this.SentryInfo >> 15) & 0x0F)
}

func (this *SentryInfoData) SetRemoteExchangeHealthCount(count int) {
	// set bit 15-18 of uint32 SentryInfo
	this.SentryInfo = (this.SentryInfo & 0xFFFF0FFF) | ((uint32(count) & 0x0F) << 15)
}

// GetCanFreeRevive bit 19：哨兵机器人当前是否可以确认免费复活，可以确认免费复活时值为
// 1，否则为 0
func (this *SentryInfoData) GetCanFreeRevive() bool {
	return (this.SentryInfo>>19)&0x01 == 1
}

func (this *SentryInfoData) SetCanFreeRevive(canFree bool) {
	// set bit 19 of uint32 SentryInfo
	if canFree {
		this.SentryInfo |= 0x00080000 // set bit 19 to 1
	} else {
		this.SentryInfo &^= 0x00080000 // set bit 19 to 0
	}
}

// GetCanExchangeRevive bit 20：哨兵机器人当前是否可以兑换立即复活，可以兑换立即复活时值为
// 1，否则为 0
func (this *SentryInfoData) GetCanExchangeRevive() bool {
	return (this.SentryInfo>>20)&0x01 == 1
}

func (this *SentryInfoData) SetCanExchangeRevive(canExchange bool) {
	// set bit 20 of uint32 SentryInfo
	if canExchange {
		this.SentryInfo |= 0x00100000 // set bit 20 to 1
	} else {
		this.SentryInfo &^= 0x00100000 // set bit 20 to 0
	}
}

// GetExchangeReviveCost bit 21-30：哨兵机器人当前若兑换立即复活需要花费的金币数
func (this *SentryInfoData) GetExchangeReviveCost() int {
	return int((this.SentryInfo >> 21) & 0x3FF)
}

func (this *SentryInfoData) SetExchangeReviveCost(cost int) {
	// set bit 21-30 of uint32 SentryInfo
	this.SentryInfo = (this.SentryInfo & 0xFFE00000) | ((uint32(cost) & 0x3FF) << 21)
}

// GetIsOutOfCombat bit 0：哨兵当前是否处于脱战状态，处于脱战状态时为 1，否则为 0。
func (this *SentryInfoData) GetIsOutOfCombat() bool {
	return this.SentryInfo2&0x01 == 1
}

func (this *SentryInfoData) SetIsOutOfCombat(isOutOfCombat bool) {
	// set bit 0 of uint16 SentryInfo2
	if isOutOfCombat {
		this.SentryInfo2 |= 0x0001 // set bit 0 to 1
	} else {
		this.SentryInfo2 &^= 0x0001 // set bit 0 to 0
	}
}

// GetTeamAmmo17mmAllowance bit 1-11：队伍 17mm 允许发弹量的剩余可兑换数。
func (this *SentryInfoData) GetTeamAmmo17mmAllowance() int {
	return int((this.SentryInfo2 >> 1) & 0x7FF)
}

func (this *SentryInfoData) SetTeamAmmo17mmAllowance(ammo int) {
	// set bit 1-11 of uint16 SentryInfo2
	this.SentryInfo2 = (this.SentryInfo2 & 0xF800) | ((uint16(ammo) & 0x7FF) << 1)
}
