package player

import (
	"errors"
	"fmt"
	"strconv"

	"bitbucket.org/east-eden/server/define"
	"bitbucket.org/east-eden/server/excel/auto"
	pbGlobal "bitbucket.org/east-eden/server/proto/global"
	pbCombat "bitbucket.org/east-eden/server/proto/server/combat"
	"bitbucket.org/east-eden/server/services/game/hero"
	"bitbucket.org/east-eden/server/services/game/item"
	"bitbucket.org/east-eden/server/services/game/prom"
	"bitbucket.org/east-eden/server/store"
	"bitbucket.org/east-eden/server/utils"
	log "github.com/rs/zerolog/log"
	"github.com/valyala/bytebufferpool"
)

func MakeHeroKey(heroId int64, fields ...string) string {
	b := bytebufferpool.Get()
	defer bytebufferpool.Put(b)

	b.B = append(b.B, "hero_map.id_"...)
	b.B = append(b.B, strconv.Itoa(int(heroId))...)

	for _, f := range fields {
		b.B = append(b.B, "."...)
		b.B = append(b.B, f...)
	}

	return b.String()
}

type HeroManager struct {
	owner       *Player              `bson:"-" json:"-"`
	HeroMap     map[int64]*hero.Hero `bson:"hero_map" json:"hero_map"` // 卡牌包
	heroTypeSet map[int32]struct{}   `bson:"-" json:"-"`               // 已获得卡牌
}

func NewHeroManager(owner *Player) *HeroManager {
	m := &HeroManager{
		owner:       owner,
		HeroMap:     make(map[int64]*hero.Hero),
		heroTypeSet: make(map[int32]struct{}),
	}

	return m
}

func (m *HeroManager) createEntryHero(entry *auto.HeroEntry) *hero.Hero {
	if entry == nil {
		log.Error().Msg("newEntryHero with nil HeroEntry")
		return nil
	}

	id, err := utils.NextID(define.SnowFlake_Hero)
	if err != nil {
		log.Error().Err(err)
		return nil
	}

	h := hero.NewHero(
		hero.Id(id),
		hero.OwnerId(m.owner.GetID()),
		hero.OwnerType(m.owner.GetType()),
		hero.Entry(entry),
		hero.TypeId(entry.Id),
	)

	h.GetAttManager().SetBaseAttId(int32(entry.AttId))
	m.HeroMap[h.GetOptions().Id] = h
	m.heroTypeSet[h.GetOptions().TypeId] = struct{}{}

	h.GetAttManager().CalcAtt()

	return h
}

func (m *HeroManager) initLoadedHero(h *hero.Hero) error {
	entry, ok := auto.GetHeroEntry(h.GetOptions().TypeId)
	if !ok {
		return fmt.Errorf("HeroManager initLoadedHero: hero<%d> entry invalid", h.GetOptions().TypeId)
	}

	h.GetOptions().Entry = entry
	h.GetAttManager().SetBaseAttId(int32(entry.AttId))

	m.HeroMap[h.GetOptions().Id] = h
	m.heroTypeSet[h.GetOptions().TypeId] = struct{}{}
	h.CalcAtt()
	return nil
}

// interface of cost_loot
func (m *HeroManager) GetCostLootType() int32 {
	return define.CostLoot_Hero
}

func (m *HeroManager) CanCost(typeMisc int32, num int32) error {
	if num <= 0 {
		return fmt.Errorf("hero manager check hero<%d> cost failed, wrong number<%d>", typeMisc, num)
	}

	var fixNum int32
	for _, v := range m.HeroMap {
		if v.GetOptions().TypeId == typeMisc {
			eb := v.GetEquipBar()
			hasEquip := false

			var n int32
			for n = 0; n < define.Equip_Pos_End; n++ {
				if eb.GetEquipByPos(n) != nil {
					hasEquip = true
					break
				}
			}

			if !hasEquip {
				fixNum++
			}
		}
	}

	if fixNum >= num {
		return nil
	}

	return fmt.Errorf("not enough hero<%d>, num<%d>", typeMisc, num)
}

func (m *HeroManager) DoCost(typeMisc int32, num int32) error {
	if num <= 0 {
		return fmt.Errorf("hero manager cost hero<%d> failed, wrong number<%d>", typeMisc, num)
	}

	var costNum int32
	for _, v := range m.HeroMap {
		if v.GetOptions().TypeId == typeMisc {
			eb := v.GetEquipBar()
			hasEquip := false

			var n int32
			for n = 0; n < define.Equip_Pos_End; n++ {
				if eb.GetEquipByPos(n) != nil {
					hasEquip = true
					break
				}
			}

			if !hasEquip {
				m.DelHero(v.GetOptions().Id)
				costNum++
			}
		}
	}

	if costNum < num {
		log.Warn().
			Int32("cost_type_misc", typeMisc).
			Int32("cost_num", num).
			Int32("actual_cost_num", costNum).
			Msg("hero manager cost num error")
		return nil
	}

	return nil
}

func (m *HeroManager) CanGain(typeMisc int32, num int32) error {
	if num <= 0 {
		return fmt.Errorf("hero manager check hero<%d> gain failed, wrong number<%d>", typeMisc, num)
	}

	// todo max hero num
	return nil
}

func (m *HeroManager) GainLoot(typeMisc int32, num int32) error {
	if num <= 0 {
		return fmt.Errorf("hero manager gain hero<%d> failed, wrong number<%d>", typeMisc, num)
	}

	var n int32
	for n = 0; n < num; n++ {
		_ = m.AddHeroByTypeID(typeMisc)
	}

	return nil
}

func (m *HeroManager) LoadAll() error {
	loadHeros := struct {
		HeroMap map[string]*hero.Hero `bson:"hero_map" json:"hero_map"`
	}{
		HeroMap: make(map[string]*hero.Hero),
	}

	err := store.GetStore().LoadObject(define.StoreType_Hero, m.owner.ID, &loadHeros)
	if errors.Is(err, store.ErrNoResult) {
		return nil
	}

	if err != nil {
		return fmt.Errorf("HeroManager LoadAll: %w", err)
	}

	for _, v := range loadHeros.HeroMap {
		h := hero.NewHero()
		h.Options.HeroInfo = v.Options.HeroInfo
		if err := m.initLoadedHero(h); err != nil {
			return fmt.Errorf("HeroManager LoadAll: %w", err)
		}
	}

	return nil
}

func (m *HeroManager) GetHero(id int64) *hero.Hero {
	return m.HeroMap[id]
}

func (m *HeroManager) GetHeroNums() int {
	return len(m.HeroMap)
}

func (m *HeroManager) GetHeroList() []*hero.Hero {
	list := make([]*hero.Hero, 0)

	for _, v := range m.HeroMap {
		list = append(list, v)
	}

	return list
}

func (m *HeroManager) AddHeroByTypeID(typeId int32) *hero.Hero {
	heroEntry, ok := auto.GetHeroEntry(typeId)
	if !ok {
		log.Warn().Int32("type_id", typeId).Msg("GetHeroEntry failed")
		return nil
	}

	// 重复获得卡牌，转换为对应碎片
	_, ok = m.heroTypeSet[typeId]
	if ok {
		m.owner.FragmentManager().Inc(typeId, heroEntry.FragmentTransform)
		return nil
	}

	h := m.createEntryHero(heroEntry)
	if h == nil {
		log.Warn().Int32("type_id", typeId).Msg("createEntryHero failed")
		return nil
	}

	fields := map[string]interface{}{
		MakeHeroKey(h.Id): h,
	}

	err := store.GetStore().SaveFields(define.StoreType_Hero, m.owner.ID, fields)
	if pass := utils.ErrCheck(err, "SaveFields failed when AddHeroByTypeID", typeId, m.owner.ID); !pass {
		m.delHero(h)
		return nil
	}

	m.SendHeroUpdate(h)

	// prometheus ops
	prom.OpsCreateHeroCounter.Inc()

	return h
}

func (m *HeroManager) delHero(h *hero.Hero) {
	delete(m.HeroMap, h.Options.Id)
	delete(m.heroTypeSet, h.Options.TypeId)
	hero.GetHeroPool().Put(h)
}

func (m *HeroManager) DelHero(id int64) {
	h, ok := m.HeroMap[id]
	if !ok {
		return
	}

	eb := h.GetEquipBar()
	var n int32
	for n = 0; n < define.Equip_Pos_End; n++ {
		utils.ErrPrint(eb.TakeoffEquip(n), "DelHero TakeoffEquip failed", id, n)
	}
	h.BeforeDelete()

	fields := []string{MakeHeroKey(id)}
	err := store.GetStore().DeleteFields(define.StoreType_Hero, m.owner.ID, fields)
	utils.ErrPrint(err, "DelHero DeleteFields failed", id)
	m.delHero(h)
}

func (m *HeroManager) HeroSetLevel(level int8) {
	for _, v := range m.HeroMap {
		v.GetOptions().Level = level

		fields := map[string]interface{}{}
		fields[MakeHeroKey(v.Id, "level")] = v.GetOptions().Level
		err := store.GetStore().SaveFields(define.StoreType_Hero, v, fields)
		utils.ErrPrint(err, "HeroSetLevel SaveFields failed", m.owner.ID, level)
	}
}

func (m *HeroManager) PutonEquip(heroId int64, equipId int64) error {
	it, err := m.owner.ItemManager().GetItem(equipId)
	if err != nil {
		return fmt.Errorf("HeroManager.PutonEquip failed: %w", err)
	}

	if it.GetType() != define.Item_TypeEquip {
		return fmt.Errorf("item<%d> is not an equip when PutonEquip", equipId)
	}

	equip := it.(*item.Equip)
	if objId := equip.GetEquipObj(); objId != -1 {
		return fmt.Errorf("equip has put on another hero<%d>", objId)
	}

	if equip.GetEquipEnchantEntry() == nil {
		return fmt.Errorf("cannot find equip_enchant_entry<%d> while PutonEquip", equipId)
	}

	h, ok := m.HeroMap[heroId]
	if !ok {
		return fmt.Errorf("invalid heroid")
	}

	equipBar := h.GetEquipBar()
	pos := equip.GetEquipEnchantEntry().EquipPos

	// takeoff previous equip
	if pe := equipBar.GetEquipByPos(pos); pe != nil {
		if err := m.TakeoffEquip(heroId, pos); err != nil {
			return err
		}
	}

	// puton this equip
	if err := equipBar.PutonEquip(equip); err != nil {
		return err
	}

	err = m.owner.ItemManager().Save(equip.Ops().Id)
	utils.ErrPrint(err, "PutonEquip Save item failed", equip.Ops().Id)

	m.owner.ItemManager().SendItemUpdate(equip)
	m.SendHeroUpdate(h)

	// att
	equip.GetAttManager().CalcAtt()
	h.GetAttManager().ModAttManager(equip.GetAttManager())
	h.GetAttManager().CalcAtt()
	m.SendHeroAtt(h)

	return nil
}

func (m *HeroManager) TakeoffEquip(heroId int64, pos int32) error {
	if pos < 0 || pos >= define.Equip_Pos_End {
		return fmt.Errorf("invalid pos")
	}

	h, ok := m.HeroMap[heroId]
	if !ok {
		return fmt.Errorf("invalid heroid")
	}

	equipBar := h.GetEquipBar()
	equip := equipBar.GetEquipByPos(pos)
	if equip == nil {
		return fmt.Errorf("cannot find hero<%d> equip by pos<%d> while TakeoffEquip", heroId, pos)
	}

	if objId := equip.GetEquipObj(); objId == -1 {
		return fmt.Errorf("equip<%d> didn't put on this hero<%d> ", equip.Ops().Id, heroId)
	}

	// unequip
	if err := equipBar.TakeoffEquip(pos); err != nil {
		return err
	}

	err := m.owner.ItemManager().Save(equip.Ops().Id)
	utils.ErrPrint(err, "TakeoffEquip Save item failed", equip.Ops().Id)
	m.owner.ItemManager().SendItemUpdate(equip)
	m.SendHeroUpdate(h)

	// att
	h.GetAttManager().CalcAtt()
	m.SendHeroAtt(h)

	return nil
}

func (m *HeroManager) PutonRune(heroId int64, runeId int64) error {

	r := m.owner.RuneManager().GetRune(runeId)
	if r == nil {
		return fmt.Errorf("cannot find rune<%d> while PutonRune", runeId)
	}

	if objId := r.GetOptions().EquipObj; objId != -1 {
		return fmt.Errorf("rune has put on another obj<%d>", objId)
	}

	pos := r.GetOptions().Entry.Pos
	if pos < define.Rune_PositionBegin || pos >= define.Rune_PositionEnd {
		return fmt.Errorf("invalid pos<%d>", pos)
	}

	h, ok := m.HeroMap[heroId]
	if !ok {
		return fmt.Errorf("invalid heroid<%d>", heroId)
	}

	runeBox := h.GetRuneBox()

	// takeoff previous rune
	if pr := runeBox.GetRuneByPos(pos); pr != nil {
		if err := m.TakeoffRune(heroId, pos); err != nil {
			return err
		}
	}

	// equip new rune
	if err := runeBox.PutonRune(r); err != nil {
		return err
	}

	err := m.owner.RuneManager().Save(runeId)
	m.owner.RuneManager().SendRuneUpdate(r)
	m.SendHeroUpdate(h)

	// att
	r.GetAttManager().CalcAtt()
	h.GetAttManager().ModAttManager(r.GetAttManager())
	h.GetAttManager().CalcAtt()
	m.SendHeroAtt(h)

	return err
}

func (m *HeroManager) TakeoffRune(heroId int64, pos int32) error {
	if pos < 0 || pos >= define.Rune_PositionEnd {
		return fmt.Errorf("invalid pos<%d>", pos)
	}

	h, ok := m.HeroMap[heroId]
	if !ok {
		return fmt.Errorf("invalid heroid<%d>", heroId)
	}

	r := h.GetRuneBox().GetRuneByPos(pos)
	if r == nil {
		return fmt.Errorf("cannot find rune from hero<%d>'s runebox pos<%d> while TakeoffRune", heroId, pos)
	}

	// unequip
	if err := h.GetRuneBox().TakeoffRune(pos); err != nil {
		return err
	}

	err := m.owner.RuneManager().Save(r.GetOptions().Id)
	m.owner.RuneManager().SendRuneUpdate(r)
	m.SendHeroUpdate(h)

	// att
	h.GetAttManager().CalcAtt()
	m.SendHeroAtt(h)

	return err
}

func (m *HeroManager) GenerateCombatUnitInfo() []*pbCombat.UnitInfo {
	retList := make([]*pbCombat.UnitInfo, 0)

	list := m.GetHeroList()
	for _, hero := range list {
		unitInfo := &pbCombat.UnitInfo{
			UnitTypeId: int32(hero.GetOptions().TypeId),
		}

		for n := define.Att_Begin; n < define.Att_End; n++ {
			unitInfo.UnitAttList = append(unitInfo.UnitAttList, &pbGlobal.Att{
				AttType:  int32(n),
				AttValue: int64(hero.GetAttManager().GetAttValue(n)),
			})
		}

		retList = append(retList, unitInfo)
	}

	return retList
}

func (m *HeroManager) SendHeroUpdate(h *hero.Hero) {
	// send equips update
	reply := &pbGlobal.S2C_HeroInfo{
		Info: &pbGlobal.Hero{
			Id:             h.GetOptions().Id,
			TypeId:         int32(h.GetOptions().TypeId),
			Exp:            h.GetOptions().Exp,
			Level:          int32(h.GetOptions().Level),
			PromoteLevel:   int32(h.GetOptions().PromoteLevel),
			Star:           int32(h.GetOptions().Star),
			NormalSpellId:  h.GetOptions().NormalSpellId,
			SpecialSpellId: h.GetOptions().SpecialSpellId,
			RageSpellId:    h.GetOptions().RageSpellId,
			Friendship:     h.GetOptions().Friendship,
			FashionId:      h.GetOptions().FashionId,
		},
	}

	// equip list
	// eb := h.GetEquipBar()
	// var n int32
	// for n = 0; n < define.Equip_Pos_End; n++ {
	// 	var equipId int64 = -1
	// 	if i := eb.GetEquipByPos(n); i != nil {
	// 		equipId = i.GetOptions().Id
	// 	}

	// 	reply.Info.EquipList = append(reply.Info.EquipList, equipId)
	// }

	// rune list
	// var pos int32
	// for pos = 0; pos < define.Rune_PositionEnd; pos++ {
	// 	var runeId int64 = -1
	// 	if r := h.GetRuneBox().GetRuneByPos(pos); r != nil {
	// 		runeId = r.GetOptions().Id
	// 	}

	// 	reply.Info.RuneList = append(reply.Info.RuneList, runeId)
	// }

	m.owner.SendProtoMessage(reply)
}

func (m *HeroManager) SendHeroAtt(h *hero.Hero) {
	attManager := h.GetAttManager()
	reply := &pbGlobal.S2C_HeroAttUpdate{
		HeroId: h.GetOptions().Id,
	}

	for k := 0; k < define.Att_End; k++ {
		att := &pbGlobal.Att{
			AttType:  int32(k),
			AttValue: int64(attManager.GetAttValue(k)),
		}
		reply.AttList = append(reply.AttList, att)
	}

	m.owner.SendProtoMessage(reply)
}
