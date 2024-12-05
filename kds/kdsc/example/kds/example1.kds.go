// Code generated by kds. DO NOT EDIT.
// source: TODO: Source File

package kds;

import (
	"time"

	"github.com/iakud/keeper/kds/kdsc/example/pb"
)

type Player struct {
	Id int64
	syncable pb.Player
	Info *PlayerBasicInfo
	Hero *PlayerHero
	Bag *PlayerBag
	dirty uint64
}

func (e *Player) SetInfo(v *PlayerBasicInfo) {
	if v != e.Info {
		e.Info = v
		e.syncable.Info = &v.syncable
		v.dirthParent = func() {
			e.markDirty(1)
		}
		e.markDirty(uint64(0x01) << 1)
	}
}

func (e *Player) SetHero(v *PlayerHero) {
	if v != e.Hero {
		e.Hero = v
		e.syncable.Hero = &v.syncable
		v.dirthParent = func() {
			e.markDirty(2)
		}
		e.markDirty(uint64(0x01) << 2)
	}
}

func (e *Player) SetBag(v *PlayerBag) {
	if v != e.Bag {
		e.Bag = v
		e.syncable.Bag = &v.syncable
		v.dirthParent = func() {
			e.markDirty(3)
		}
		e.markDirty(uint64(0x01) << 3)
	}
}

func (e *Player) markDirty(n uint64) {
	if e.dirty & n == n {
		return
	}
	e.dirty |= n
}

func (e *Player) clearDirty() {
	if e.dirty == 0 {
		return
	}
	e.dirty = 0
}

func (e *Player) DumpChange() *pb.Player {
	v := new(pb.Player)
	if e.dirty & uint64(0x01) << 1 != 0 {
		v.Info = e.Info.DumpChange()
	}
	if e.dirty & uint64(0x01) << 2 != 0 {
		v.Hero = e.Hero.DumpChange()
	}
	if e.dirty & uint64(0x01) << 3 != 0 {
		v.Bag = e.Bag.DumpChange()
	}
	return v
}

func (e *Player) DumpFull() *pb.Player {
	v := new(pb.Player)
	v.Info = e.Info.DumpFull()
	v.Hero = e.Hero.DumpFull()
	v.Bag = e.Bag.DumpFull()
	return v
}

type dirtyParentFunc_PlayerBasicInfo func()

func (f dirtyParentFunc_PlayerBasicInfo) invoke() {
	if f == nil {
		return
	}
	f()
}

type PlayerBasicInfo struct {
	syncable pb.PlayerBasicInfo
	dirty uint64
	dirthParent dirtyParentFunc_PlayerBasicInfo
}

func (c *PlayerBasicInfo) SetName(v string) {
	if v != c.syncable.Name {
		c.syncable.Name = v
		c.markDirty(uint64(0x01) << 1)
	}
}

func (c *PlayerBasicInfo) SetIsNew(v bool) {
	if v != c.syncable.IsNew {
		c.syncable.IsNew = v
		c.markDirty(uint64(0x01) << 3)
	}
}

func (c *PlayerBasicInfo) SetCreateTime(v time.Time) {
	if v != c.syncable.CreateTime {
		c.syncable.CreateTime = v
		c.markDirty(uint64(0x01) << 5)
	}
}

func (c *PlayerBasicInfo) markDirty(n uint64) {
	if c.dirty&n == n {
		return
	}
	c.dirty |= n
	c.dirthParent.invoke()
}

func (c *PlayerBasicInfo) clearDirty() {
	if c.dirty == 0 {
		return
	}
	c.dirty = 0
}

func (c *PlayerBasicInfo) DumpChange() *pb.PlayerBasicInfo {
	if c == nil {
		return nil
	}
	v := new(pb.PlayerBasicInfo)
	if c.dirty & uint64(0x01) << 1 != 0 {
		v.Name = c.syncable.Name
	}
	if c.dirty & uint64(0x01) << 3 != 0 {
		v.IsNew = c.syncable.IsNew
	}
	if c.dirty & uint64(0x01) << 5 != 0 {
		v.CreateTime = c.syncable.CreateTime
	}
	return v
}

func (c *PlayerBasicInfo) DumpFull() *pb.PlayerBasicInfo {
	v := new(pb.PlayerBasicInfo)
	v.Name = c.syncable.Name
	v.IsNew = c.syncable.IsNew
	v.CreateTime = c.syncable.CreateTime
	return v
}

type dirtyParentFunc_PlayerHero func()

func (f dirtyParentFunc_PlayerHero) invoke() {
	if f == nil {
		return
	}
	f()
}

type PlayerHero struct {
	syncable pb.PlayerHero
	Heroes *Hero
	dirty uint64
	dirthParent dirtyParentFunc_PlayerHero
}

func (c *PlayerHero) SetHeroes(v *Hero) {
	if v != c.Heroes {
		c.Heroes = v
		c.syncable.Heroes = &v.syncable
		c.markDirty(uint64(0x01) << 1)
	}
}

func (c *PlayerHero) markDirty(n uint64) {
	if c.dirty&n == n {
		return
	}
	c.dirty |= n
	c.dirthParent.invoke()
}

func (c *PlayerHero) clearDirty() {
	if c.dirty == 0 {
		return
	}
	c.dirty = 0
}

func (c *PlayerHero) DumpChange() *pb.PlayerHero {
	if c == nil {
		return nil
	}
	v := new(pb.PlayerHero)
	if c.dirty & uint64(0x01) << 1 != 0 {
		v.Heroes = c.Heroes.DumpChange()
	}
	return v
}

func (c *PlayerHero) DumpFull() *pb.PlayerHero {
	v := new(pb.PlayerHero)
	v.Heroes = c.Heroes.DumpFull()
	return v
}

type dirtyParentFunc_PlayerBag func()

func (f dirtyParentFunc_PlayerBag) invoke() {
	if f == nil {
		return
	}
	f()
}

type PlayerBag struct {
	syncable pb.PlayerBag
	dirty uint64
	dirthParent dirtyParentFunc_PlayerBag
}

func (c *PlayerBag) SetResources(v int32) {
	if v != c.syncable.Resources {
		c.syncable.Resources = v
		c.markDirty(uint64(0x01) << 1)
	}
}

func (c *PlayerBag) markDirty(n uint64) {
	if c.dirty&n == n {
		return
	}
	c.dirty |= n
	c.dirthParent.invoke()
}

func (c *PlayerBag) clearDirty() {
	if c.dirty == 0 {
		return
	}
	c.dirty = 0
}

func (c *PlayerBag) DumpChange() *pb.PlayerBag {
	if c == nil {
		return nil
	}
	v := new(pb.PlayerBag)
	if c.dirty & uint64(0x01) << 1 != 0 {
		v.Resources = c.syncable.Resources
	}
	return v
}

func (c *PlayerBag) DumpFull() *pb.PlayerBag {
	v := new(pb.PlayerBag)
	v.Resources = c.syncable.Resources
	return v
}

type dirtyParentFunc_Hero func()

func (f dirtyParentFunc_Hero) invoke() {
	if f == nil {
		return
	}
	f()
}

type Hero struct {
	syncable pb.Hero
	dirty uint64
	dirthParent dirtyParentFunc_Hero
}

func (c *Hero) SetHeroId(v int32) {
	if v != c.syncable.HeroId {
		c.syncable.HeroId = v
		c.markDirty(uint64(0x01) << 1)
	}
}

func (c *Hero) SetHeroLevel(v int32) {
	if v != c.syncable.HeroLevel {
		c.syncable.HeroLevel = v
		c.markDirty(uint64(0x01) << 2)
	}
}

func (c *Hero) SetType(v pb.HeroType) {
	if v != c.syncable.Type {
		c.syncable.Type = v
		c.markDirty(uint64(0x01) << 3)
	}
}

func (c *Hero) SetNeedTime(v time.Duration) {
	if v != c.syncable.NeedTime {
		c.syncable.NeedTime = v
		c.markDirty(uint64(0x01) << 4)
	}
}

func (c *Hero) markDirty(n uint64) {
	if c.dirty&n == n {
		return
	}
	c.dirty |= n
	c.dirthParent.invoke()
}

func (c *Hero) clearDirty() {
	if c.dirty == 0 {
		return
	}
	c.dirty = 0
}

func (c *Hero) DumpChange() *pb.Hero {
	if c == nil {
		return nil
	}
	v := new(pb.Hero)
	if c.dirty & uint64(0x01) << 1 != 0 {
		v.HeroId = c.syncable.HeroId
	}
	if c.dirty & uint64(0x01) << 2 != 0 {
		v.HeroLevel = c.syncable.HeroLevel
	}
	if c.dirty & uint64(0x01) << 3 != 0 {
		v.Type = c.syncable.Type
	}
	if c.dirty & uint64(0x01) << 4 != 0 {
		v.NeedTime = c.syncable.NeedTime
	}
	return v
}

func (c *Hero) DumpFull() *pb.Hero {
	v := new(pb.Hero)
	v.HeroId = c.syncable.HeroId
	v.HeroLevel = c.syncable.HeroLevel
	v.Type = c.syncable.Type
	v.NeedTime = c.syncable.NeedTime
	return v
}