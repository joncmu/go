// Copyright 2016 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package s390x

import (
	"cmd/compile/internal/gc"
	"cmd/internal/obj"
	"cmd/internal/obj/s390x"
)

const (
	LeftRdwr  uint32 = gc.LeftRead | gc.LeftWrite
	RightRdwr uint32 = gc.RightRead | gc.RightWrite
)

// This table gives the basic information about instruction
// generated by the compiler and processed in the optimizer.
// See opt.h for bit definitions.
//
// Instructions not generated need not be listed.
// As an exception to that rule, we typically write down all the
// size variants of an operation even if we just use a subset.
//
// The table is formatted for 8-space tabs.
var progtable = [s390x.ALAST]obj.ProgInfo{
	obj.ATYPE:     {Flags: gc.Pseudo | gc.Skip},
	obj.ATEXT:     {Flags: gc.Pseudo},
	obj.AFUNCDATA: {Flags: gc.Pseudo},
	obj.APCDATA:   {Flags: gc.Pseudo},
	obj.AUNDEF:    {Flags: gc.Break},
	obj.AUSEFIELD: {Flags: gc.OK},
	obj.ACHECKNIL: {Flags: gc.LeftRead},
	obj.AVARDEF:   {Flags: gc.Pseudo | gc.RightWrite},
	obj.AVARKILL:  {Flags: gc.Pseudo | gc.RightWrite},
	obj.AVARLIVE:  {Flags: gc.Pseudo | gc.LeftRead},

	// NOP is an internal no-op that also stands
	// for USED and SET annotations, not the Power opcode.
	obj.ANOP: {Flags: gc.LeftRead | gc.RightWrite},

	// Integer
	s390x.AADD:    {Flags: gc.SizeQ | gc.LeftRead | gc.RegRead | gc.RightWrite},
	s390x.ASUB:    {Flags: gc.SizeQ | gc.LeftRead | gc.RegRead | gc.RightWrite},
	s390x.ANEG:    {Flags: gc.SizeQ | gc.LeftRead | gc.RegRead | gc.RightWrite},
	s390x.AAND:    {Flags: gc.SizeQ | gc.LeftRead | gc.RegRead | gc.RightWrite},
	s390x.AOR:     {Flags: gc.SizeQ | gc.LeftRead | gc.RegRead | gc.RightWrite},
	s390x.AXOR:    {Flags: gc.SizeQ | gc.LeftRead | gc.RegRead | gc.RightWrite},
	s390x.AMULLD:  {Flags: gc.SizeQ | gc.LeftRead | gc.RegRead | gc.RightWrite},
	s390x.AMULLW:  {Flags: gc.SizeL | gc.LeftRead | gc.RegRead | gc.RightWrite},
	s390x.AMULHDU: {Flags: gc.SizeL | gc.LeftRead | gc.RegRead | gc.RightWrite},
	s390x.ADIVD:   {Flags: gc.SizeQ | gc.LeftRead | gc.RegRead | gc.RightWrite},
	s390x.ADIVDU:  {Flags: gc.SizeQ | gc.LeftRead | gc.RegRead | gc.RightWrite},
	s390x.ASLD:    {Flags: gc.SizeQ | gc.LeftRead | gc.RegRead | gc.RightWrite},
	s390x.ASRD:    {Flags: gc.SizeQ | gc.LeftRead | gc.RegRead | gc.RightWrite},
	s390x.ASRAD:   {Flags: gc.SizeQ | gc.LeftRead | gc.RegRead | gc.RightWrite},
	s390x.ACMP:    {Flags: gc.SizeQ | gc.LeftRead | gc.RightRead},
	s390x.ACMPU:   {Flags: gc.SizeQ | gc.LeftRead | gc.RightRead},

	// Floating point.
	s390x.AFADD:  {Flags: gc.SizeD | gc.LeftRead | gc.RegRead | gc.RightWrite},
	s390x.AFADDS: {Flags: gc.SizeF | gc.LeftRead | gc.RegRead | gc.RightWrite},
	s390x.AFSUB:  {Flags: gc.SizeD | gc.LeftRead | gc.RegRead | gc.RightWrite},
	s390x.AFSUBS: {Flags: gc.SizeF | gc.LeftRead | gc.RegRead | gc.RightWrite},
	s390x.AFMUL:  {Flags: gc.SizeD | gc.LeftRead | gc.RegRead | gc.RightWrite},
	s390x.AFMULS: {Flags: gc.SizeF | gc.LeftRead | gc.RegRead | gc.RightWrite},
	s390x.AFDIV:  {Flags: gc.SizeD | gc.LeftRead | gc.RegRead | gc.RightWrite},
	s390x.AFDIVS: {Flags: gc.SizeF | gc.LeftRead | gc.RegRead | gc.RightWrite},
	s390x.AFCMPU: {Flags: gc.SizeD | gc.LeftRead | gc.RightRead},
	s390x.ACEBR:  {Flags: gc.SizeF | gc.LeftRead | gc.RightRead},
	s390x.ALEDBR: {Flags: gc.SizeD | gc.LeftRead | gc.RightWrite | gc.Conv},
	s390x.ALDEBR: {Flags: gc.SizeD | gc.LeftRead | gc.RightWrite | gc.Conv},
	s390x.AFSQRT: {Flags: gc.SizeD | gc.LeftRead | gc.RightWrite},

	// Conversions
	s390x.ACEFBRA: {Flags: gc.SizeF | gc.LeftRead | gc.RightWrite | gc.Conv},
	s390x.ACDFBRA: {Flags: gc.SizeD | gc.LeftRead | gc.RightWrite | gc.Conv},
	s390x.ACEGBRA: {Flags: gc.SizeF | gc.LeftRead | gc.RightWrite | gc.Conv},
	s390x.ACDGBRA: {Flags: gc.SizeD | gc.LeftRead | gc.RightWrite | gc.Conv},
	s390x.ACFEBRA: {Flags: gc.SizeL | gc.LeftRead | gc.RightWrite | gc.Conv},
	s390x.ACFDBRA: {Flags: gc.SizeL | gc.LeftRead | gc.RightWrite | gc.Conv},
	s390x.ACGEBRA: {Flags: gc.SizeQ | gc.LeftRead | gc.RightWrite | gc.Conv},
	s390x.ACGDBRA: {Flags: gc.SizeQ | gc.LeftRead | gc.RightWrite | gc.Conv},
	s390x.ACELFBR: {Flags: gc.SizeF | gc.LeftRead | gc.RightWrite | gc.Conv},
	s390x.ACDLFBR: {Flags: gc.SizeD | gc.LeftRead | gc.RightWrite | gc.Conv},
	s390x.ACELGBR: {Flags: gc.SizeF | gc.LeftRead | gc.RightWrite | gc.Conv},
	s390x.ACDLGBR: {Flags: gc.SizeD | gc.LeftRead | gc.RightWrite | gc.Conv},
	s390x.ACLFEBR: {Flags: gc.SizeL | gc.LeftRead | gc.RightWrite | gc.Conv},
	s390x.ACLFDBR: {Flags: gc.SizeL | gc.LeftRead | gc.RightWrite | gc.Conv},
	s390x.ACLGEBR: {Flags: gc.SizeQ | gc.LeftRead | gc.RightWrite | gc.Conv},
	s390x.ACLGDBR: {Flags: gc.SizeQ | gc.LeftRead | gc.RightWrite | gc.Conv},

	// Moves
	s390x.AMOVB:  {Flags: gc.SizeB | gc.LeftRead | gc.RightWrite | gc.Move | gc.Conv},
	s390x.AMOVBZ: {Flags: gc.SizeB | gc.LeftRead | gc.RightWrite | gc.Move | gc.Conv},
	s390x.AMOVH:  {Flags: gc.SizeW | gc.LeftRead | gc.RightWrite | gc.Move | gc.Conv},
	s390x.AMOVHZ: {Flags: gc.SizeW | gc.LeftRead | gc.RightWrite | gc.Move | gc.Conv},
	s390x.AMOVW:  {Flags: gc.SizeL | gc.LeftRead | gc.RightWrite | gc.Move | gc.Conv},

	// there is no AMOVWU.
	s390x.AMOVWZ: {Flags: gc.SizeL | gc.LeftRead | gc.RightWrite | gc.Move | gc.Conv},
	s390x.AMOVD:  {Flags: gc.SizeQ | gc.LeftRead | gc.RightWrite | gc.Move},
	s390x.AFMOVS: {Flags: gc.SizeF | gc.LeftRead | gc.RightWrite | gc.Move | gc.Conv},
	s390x.AFMOVD: {Flags: gc.SizeD | gc.LeftRead | gc.RightWrite | gc.Move},

	// Storage operations
	s390x.AMVC: {Flags: gc.SizeQ | gc.LeftRead | gc.RightRead},
	s390x.ACLC: {Flags: gc.SizeQ | gc.LeftRead | gc.RightRead},
	s390x.AXC:  {Flags: gc.SizeQ | gc.LeftRead | gc.RightRead},
	s390x.AOC:  {Flags: gc.SizeQ | gc.LeftRead | gc.RightRead},
	s390x.ANC:  {Flags: gc.SizeQ | gc.LeftRead | gc.RightRead},

	// Jumps
	s390x.ABR:      {Flags: gc.Jump | gc.Break},
	s390x.ABL:      {Flags: gc.Call},
	s390x.ABEQ:     {Flags: gc.Cjmp},
	s390x.ABNE:     {Flags: gc.Cjmp},
	s390x.ABGE:     {Flags: gc.Cjmp},
	s390x.ABLT:     {Flags: gc.Cjmp},
	s390x.ABGT:     {Flags: gc.Cjmp},
	s390x.ABLE:     {Flags: gc.Cjmp},
	s390x.ACMPBEQ:  {Flags: gc.Cjmp},
	s390x.ACMPBNE:  {Flags: gc.Cjmp},
	s390x.ACMPBGE:  {Flags: gc.Cjmp},
	s390x.ACMPBLT:  {Flags: gc.Cjmp},
	s390x.ACMPBGT:  {Flags: gc.Cjmp},
	s390x.ACMPBLE:  {Flags: gc.Cjmp},
	s390x.ACMPUBEQ: {Flags: gc.Cjmp},
	s390x.ACMPUBNE: {Flags: gc.Cjmp},
	s390x.ACMPUBGE: {Flags: gc.Cjmp},
	s390x.ACMPUBLT: {Flags: gc.Cjmp},
	s390x.ACMPUBGT: {Flags: gc.Cjmp},
	s390x.ACMPUBLE: {Flags: gc.Cjmp},

	// Macros
	s390x.ACLEAR: {Flags: gc.SizeQ | gc.LeftRead | gc.RightAddr | gc.RightWrite},

	// Load/store multiple
	s390x.ASTMG: {Flags: gc.SizeQ | gc.LeftRead | gc.RightAddr | gc.RightWrite},
	s390x.ASTMY: {Flags: gc.SizeL | gc.LeftRead | gc.RightAddr | gc.RightWrite},
	s390x.ALMG:  {Flags: gc.SizeQ | gc.LeftAddr | gc.LeftRead | gc.RightWrite},
	s390x.ALMY:  {Flags: gc.SizeL | gc.LeftAddr | gc.LeftRead | gc.RightWrite},

	obj.ARET: {Flags: gc.Break},
}

func proginfo(p *obj.Prog) {
	info := &p.Info
	*info = progtable[p.As]
	if info.Flags == 0 {
		gc.Fatalf("proginfo: unknown instruction %v", p)
	}

	if (info.Flags&gc.RegRead != 0) && p.Reg == 0 {
		info.Flags &^= gc.RegRead
		info.Flags |= gc.RightRead /*CanRegRead |*/
	}

	if (p.From.Type == obj.TYPE_MEM || p.From.Type == obj.TYPE_ADDR) && p.From.Reg != 0 {
		info.Regindex |= RtoB(int(p.From.Reg))
		if info.Flags&gc.PostInc != 0 {
			info.Regset |= RtoB(int(p.From.Reg))
		}
	}

	if (p.To.Type == obj.TYPE_MEM || p.To.Type == obj.TYPE_ADDR) && p.To.Reg != 0 {
		info.Regindex |= RtoB(int(p.To.Reg))
		if info.Flags&gc.PostInc != 0 {
			info.Regset |= RtoB(int(p.To.Reg))
		}
	}

	if p.From.Type == obj.TYPE_ADDR && p.From.Sym != nil && (info.Flags&gc.LeftRead != 0) {
		info.Flags &^= gc.LeftRead
		info.Flags |= gc.LeftAddr
	}

	switch p.As {
	// load multiple sets a range of registers
	case s390x.ALMG, s390x.ALMY:
		for r := p.Reg; r <= p.To.Reg; r++ {
			info.Regset |= RtoB(int(r))
		}
	// store multiple reads a range of registers
	case s390x.ASTMG, s390x.ASTMY:
		for r := p.From.Reg; r <= p.Reg; r++ {
			info.Reguse |= RtoB(int(r))
		}
	}
}
