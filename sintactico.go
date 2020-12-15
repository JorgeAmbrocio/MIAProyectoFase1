// Code generated by goyacc - DO NOT EDIT.

// +build ignore
package main

import __yyfmt__ "fmt"

import (
	"bufio"
	"fmt"
	"log"
	//"archivos/MIA-Proyecto1_201709454/analizador/arbol"
	"os"
	"proyectos/MIAProyectoFase1/analizador/arbol"
)

var oParametro arbol.Parametro
var oInstruccion arbol.Instruccion
var lInstruccion []arbol.Instruccion
var lAST []arbol.AST
var auxPath string

type yySymType struct {
	yys   int
	value string
}

type yyXError struct {
	state, xsym int
}

const (
	yyDefault    = 57410
	yyEofCode    = 57344
	add          = 57382
	archivo      = 57409
	asignacion   = 57346
	bf           = 57400
	cat          = 57364
	chgrp        = 57373
	chmod        = 57362
	chown        = 57372
	cont         = 57392
	cp           = 57369
	delete       = 57381
	dest         = 57395
	e            = 57398
	edit         = 57366
	yyErrCode    = 57345
	exec         = 57348
	fast         = 57404
	fdisk        = 57352
	ff           = 57401
	filen        = 57393
	find         = 57371
	fit          = 57380
	full         = 57403
	grp          = 57388
	guion        = 57347
	id           = 57384
	idn          = 57383
	k            = 57396
	l            = 57399
	login        = 57356
	logout       = 57357
	m            = 57397
	mkdir        = 57368
	mkdisk       = 57350
	mkfile       = 57363
	mkfs         = 57355
	mkgrp        = 57358
	mkusr        = 57360
	mount        = 57353
	mv           = 57370
	name         = 57377
	numero       = 57406
	p            = 57391
	path         = 57375
	pause        = 57349
	pwd          = 57387
	r            = 57390
	ren          = 57367
	rep          = 57374
	rf           = 57394
	rid          = 57405
	rm           = 57365
	rmdisk       = 57351
	rmgrp        = 57359
	rmusr        = 57361
	rtype        = 57379
	rutaCompleja = 57408
	rutaSimple   = 57407
	size         = 57376
	tipo         = 57385
	ugo          = 57389
	unit         = 57378
	unmount      = 57354
	usr          = 57386
	wf           = 57402

	yyMaxDepth = 200
	yyTabOfs   = -80
)

var (
	yyPrec = map[int]int{}

	yyXLAT = map[int]int{
		57347: 0,  // guion (80x)
		57344: 1,  // $end (78x)
		57348: 2,  // exec (78x)
		57352: 3,  // fdisk (78x)
		57350: 4,  // mkdisk (78x)
		57355: 5,  // mkfs (78x)
		57353: 6,  // mount (78x)
		57349: 7,  // pause (78x)
		57374: 8,  // rep (78x)
		57351: 9,  // rmdisk (78x)
		57354: 10, // unmount (78x)
		57346: 11, // asignacion (11x)
		57425: 12, // PARAMETRO_PATH (10x)
		57424: 13, // PARAMETRO_NAME (6x)
		57422: 14, // PARAMETRO_ID (5x)
		57375: 15, // path (5x)
		57421: 16, // PARAMETRO_FIT (4x)
		57426: 17, // PARAMETRO_SIZE (4x)
		57429: 18, // PARAMETRO_UNIT (4x)
		57405: 19, // rid (4x)
		57384: 20, // id (3x)
		57377: 21, // name (3x)
		57406: 22, // numero (3x)
		57404: 23, // fast (2x)
		57380: 24, // fit (2x)
		57403: 25, // full (2x)
		57412: 26, // INSTRUCCION (2x)
		57419: 27, // PARAMETRO_ADD (2x)
		57420: 28, // PARAMETRO_DELETE (2x)
		57427: 29, // PARAMETRO_TYPE (2x)
		57428: 30, // PARAMETRO_TYPE_FS (2x)
		57379: 31, // rtype (2x)
		57408: 32, // rutaCompleja (2x)
		57376: 33, // size (2x)
		57378: 34, // unit (2x)
		57430: 35, // VALOR_DELETE (2x)
		57382: 36, // add (1x)
		57409: 37, // archivo (1x)
		57400: 38, // bf (1x)
		57381: 39, // delete (1x)
		57398: 40, // e (1x)
		57401: 41, // ff (1x)
		57411: 42, // INICIO (1x)
		57396: 43, // k (1x)
		57399: 44, // l (1x)
		57413: 45, // LISTA_INSTRUCCION (1x)
		57414: 46, // LST_FDISK (1x)
		57415: 47, // LST_MKDIS (1x)
		57416: 48, // LST_MKFS (1x)
		57417: 49, // LST_MOUNT (1x)
		57418: 50, // LST_REP (1x)
		57397: 51, // m (1x)
		57391: 52, // p (1x)
		57407: 53, // rutaSimple (1x)
		57431: 54, // VALOR_FIT (1x)
		57432: 55, // VALOR_NAME (1x)
		57433: 56, // VALOR_PATH (1x)
		57434: 57, // VALOR_TYPE (1x)
		57435: 58, // VALOR_UNIT (1x)
		57402: 59, // wf (1x)
		57410: 60, // $default (0x)
		57364: 61, // cat (0x)
		57373: 62, // chgrp (0x)
		57362: 63, // chmod (0x)
		57372: 64, // chown (0x)
		57392: 65, // cont (0x)
		57369: 66, // cp (0x)
		57395: 67, // dest (0x)
		57366: 68, // edit (0x)
		57345: 69, // error (0x)
		57393: 70, // filen (0x)
		57371: 71, // find (0x)
		57388: 72, // grp (0x)
		57383: 73, // idn (0x)
		57356: 74, // login (0x)
		57357: 75, // logout (0x)
		57368: 76, // mkdir (0x)
		57363: 77, // mkfile (0x)
		57358: 78, // mkgrp (0x)
		57360: 79, // mkusr (0x)
		57370: 80, // mv (0x)
		57423: 81, // PARAMETRO_IDN (0x)
		57387: 82, // pwd (0x)
		57390: 83, // r (0x)
		57367: 84, // ren (0x)
		57394: 85, // rf (0x)
		57365: 86, // rm (0x)
		57359: 87, // rmgrp (0x)
		57361: 88, // rmusr (0x)
		57385: 89, // tipo (0x)
		57389: 90, // ugo (0x)
		57386: 91, // usr (0x)
	}

	yySymNames = []string{
		"guion",
		"$end",
		"exec",
		"fdisk",
		"mkdisk",
		"mkfs",
		"mount",
		"pause",
		"rep",
		"rmdisk",
		"unmount",
		"asignacion",
		"PARAMETRO_PATH",
		"PARAMETRO_NAME",
		"PARAMETRO_ID",
		"path",
		"PARAMETRO_FIT",
		"PARAMETRO_SIZE",
		"PARAMETRO_UNIT",
		"rid",
		"id",
		"name",
		"numero",
		"fast",
		"fit",
		"full",
		"INSTRUCCION",
		"PARAMETRO_ADD",
		"PARAMETRO_DELETE",
		"PARAMETRO_TYPE",
		"PARAMETRO_TYPE_FS",
		"rtype",
		"rutaCompleja",
		"size",
		"unit",
		"VALOR_DELETE",
		"add",
		"archivo",
		"bf",
		"delete",
		"e",
		"ff",
		"INICIO",
		"k",
		"l",
		"LISTA_INSTRUCCION",
		"LST_FDISK",
		"LST_MKDIS",
		"LST_MKFS",
		"LST_MOUNT",
		"LST_REP",
		"m",
		"p",
		"rutaSimple",
		"VALOR_FIT",
		"VALOR_NAME",
		"VALOR_PATH",
		"VALOR_TYPE",
		"VALOR_UNIT",
		"wf",
		"$default",
		"cat",
		"chgrp",
		"chmod",
		"chown",
		"cont",
		"cp",
		"dest",
		"edit",
		"error",
		"filen",
		"find",
		"grp",
		"idn",
		"login",
		"logout",
		"mkdir",
		"mkfile",
		"mkgrp",
		"mkusr",
		"mv",
		"PARAMETRO_IDN",
		"pwd",
		"r",
		"ren",
		"rf",
		"rm",
		"rmgrp",
		"rmusr",
		"tipo",
		"ugo",
		"usr",
	}

	yyTokenLiteralStrings = map[int]string{}

	yyReductions = map[int]struct{ xsym, components int }{
		0:  {0, 1},
		1:  {42, 1},
		2:  {45, 2},
		3:  {45, 1},
		4:  {26, 1},
		5:  {26, 2},
		6:  {26, 2},
		7:  {26, 2},
		8:  {26, 2},
		9:  {26, 2},
		10: {26, 2},
		11: {26, 2},
		12: {26, 2},
		13: {48, 2},
		14: {48, 2},
		15: {48, 1},
		16: {48, 1},
		17: {50, 2},
		18: {50, 2},
		19: {50, 2},
		20: {50, 1},
		21: {50, 1},
		22: {50, 1},
		23: {49, 2},
		24: {49, 2},
		25: {49, 1},
		26: {49, 1},
		27: {46, 2},
		28: {46, 2},
		29: {46, 2},
		30: {46, 2},
		31: {46, 2},
		32: {46, 2},
		33: {46, 2},
		34: {46, 2},
		35: {46, 1},
		36: {46, 1},
		37: {46, 1},
		38: {46, 1},
		39: {46, 1},
		40: {46, 1},
		41: {46, 1},
		42: {46, 1},
		43: {47, 2},
		44: {47, 2},
		45: {47, 2},
		46: {47, 2},
		47: {47, 1},
		48: {47, 1},
		49: {47, 1},
		50: {47, 1},
		51: {12, 4},
		52: {56, 1},
		53: {56, 1},
		54: {17, 4},
		55: {13, 4},
		56: {55, 1},
		57: {55, 1},
		58: {55, 1},
		59: {18, 4},
		60: {58, 1},
		61: {58, 1},
		62: {29, 4},
		63: {57, 1},
		64: {57, 1},
		65: {57, 1},
		66: {16, 4},
		67: {54, 1},
		68: {54, 1},
		69: {54, 1},
		70: {28, 4},
		71: {30, 4},
		72: {35, 1},
		73: {35, 1},
		74: {27, 4},
		75: {27, 5},
		76: {81, 5},
		77: {81, 4},
		78: {14, 5},
		79: {14, 4},
	}

	yyXErrors = map[yyXError]string{}

	yyParseTab = [118][]uint16{
		// 0
		{2: 85, 88, 86, 92, 89, 84, 91, 87, 90, 26: 83, 42: 81, 45: 82},
		{1: 80},
		{1: 79, 85, 88, 86, 92, 89, 84, 91, 87, 90, 26: 197},
		{1: 77, 77, 77, 77, 77, 77, 77, 77, 77, 77},
		{1: 76, 76, 76, 76, 76, 76, 76, 76, 76, 76},
		// 5
		{185, 12: 196},
		{191, 12: 188, 16: 190, 187, 189, 47: 186},
		{185, 12: 184},
		{147, 12: 141, 145, 16: 143, 139, 140, 27: 146, 144, 142, 46: 138},
		{135, 12: 133, 134, 49: 132},
		// 10
		{131, 14: 130},
		{115, 12: 113, 112, 114, 50: 111},
		{96, 14: 94, 30: 95, 48: 93},
		{96, 68, 68, 68, 68, 68, 68, 68, 68, 68, 68, 14: 109, 30: 110},
		{105, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65},
		// 15
		{64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64},
		{19: 98, 31: 97},
		{11: 101},
		{11: 99},
		{20: 100},
		// 20
		{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		{23: 104, 25: 103, 35: 102},
		{9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9},
		{8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8},
		{7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7},
		// 25
		{19: 106},
		{11: 107},
		{20: 108},
		{2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2},
		{105, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67},
		// 30
		{66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66},
		{115, 69, 69, 69, 69, 69, 69, 69, 69, 69, 69, 12: 128, 127, 129},
		{60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60},
		{59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59},
		{105, 58, 58, 58, 58, 58, 58, 58, 58, 58, 58},
		// 35
		{15: 116, 19: 98, 21: 117},
		{11: 123},
		{11: 118},
		{20: 121, 32: 122, 37: 120, 55: 119},
		{25, 25, 25, 25, 25, 25, 25, 25, 25, 25, 25},
		// 40
		{24, 24, 24, 24, 24, 24, 24, 24, 24, 24, 24},
		{23, 23, 23, 23, 23, 23, 23, 23, 23, 23, 23},
		{22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22},
		{32: 126, 53: 125, 56: 124},
		{29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29},
		// 45
		{28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28},
		{27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27},
		{63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63},
		{62, 62, 62, 62, 62, 62, 62, 62, 62, 62, 62},
		{105, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61},
		// 50
		{105, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70},
		{19: 98},
		{135, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 12: 136, 137},
		{55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55},
		{54, 54, 54, 54, 54, 54, 54, 54, 54, 54, 54},
		// 55
		{15: 116, 21: 117},
		{57, 57, 57, 57, 57, 57, 57, 57, 57, 57, 57},
		{56, 56, 56, 56, 56, 56, 56, 56, 56, 56, 56},
		{147, 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 12: 178, 182, 16: 180, 176, 177, 27: 183, 181, 179},
		{45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45},
		// 60
		{44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44},
		{43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43},
		{42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42},
		{41, 41, 41, 41, 41, 41, 41, 41, 41, 41, 41},
		{40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40},
		// 65
		{39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39},
		{38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38},
		{15: 116, 21: 117, 24: 151, 31: 150, 33: 148, 149, 36: 153, 39: 152},
		{11: 174},
		{11: 170},
		// 70
		{11: 165},
		{11: 160},
		{11: 158},
		{11: 154},
		{156, 22: 155},
		// 75
		{6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6},
		{22: 157},
		{5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5},
		{23: 104, 25: 103, 35: 159},
		{10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10},
		// 80
		{38: 162, 41: 163, 54: 161, 59: 164},
		{14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14},
		{13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13},
		{12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12},
		{11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11},
		// 85
		{40: 168, 44: 169, 52: 167, 57: 166},
		{18, 18, 18, 18, 18, 18, 18, 18, 18, 18, 18},
		{17, 17, 17, 17, 17, 17, 17, 17, 17, 17, 17},
		{16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16},
		{15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15},
		// 90
		{43: 172, 51: 173, 58: 171},
		{21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21},
		{20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20},
		{19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19},
		{22: 175},
		// 95
		{26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26},
		{53, 53, 53, 53, 53, 53, 53, 53, 53, 53, 53},
		{52, 52, 52, 52, 52, 52, 52, 52, 52, 52, 52},
		{51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51},
		{50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50},
		// 100
		{49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49},
		{48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48},
		{47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47},
		{46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46},
		{1: 73, 73, 73, 73, 73, 73, 73, 73, 73, 73},
		// 105
		{15: 116},
		{191, 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 12: 193, 16: 195, 192, 194},
		{33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33},
		{32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32},
		{31, 31, 31, 31, 31, 31, 31, 31, 31, 31, 31},
		// 110
		{30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30},
		{15: 116, 24: 151, 33: 148, 149},
		{37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37},
		{36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36},
		{35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35},
		// 115
		{34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34},
		{1: 75, 75, 75, 75, 75, 75, 75, 75, 75, 75},
		{1: 78, 78, 78, 78, 78, 78, 78, 78, 78, 78},
	}
)

var yyDebug = 0

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

type yyLexerEx interface {
	yyLexer
	Reduced(rule, state int, lval *yySymType) bool
}

func yySymName(c int) (s string) {
	x, ok := yyXLAT[c]
	if ok {
		return yySymNames[x]
	}

	if c < 0x7f {
		return __yyfmt__.Sprintf("%q", c)
	}

	return __yyfmt__.Sprintf("%d", c)
}

func yylex1(yylex yyLexer, lval *yySymType) (n int) {
	n = yylex.Lex(lval)
	if n <= 0 {
		n = yyEofCode
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("\nlex %s(%#x %d), lval: %+v\n", yySymName(n), n, n, lval)
	}
	return n
}

func yyParse(yylex yyLexer) int {
	const yyError = 69

	yyEx, _ := yylex.(yyLexerEx)
	var yyn int
	var yylval yySymType
	var yyVAL yySymType
	yyS := make([]yySymType, 200)

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yyerrok := func() {
		if yyDebug >= 2 {
			__yyfmt__.Printf("yyerrok()\n")
		}
		Errflag = 0
	}
	_ = yyerrok
	yystate := 0
	yychar := -1
	var yyxchar int
	var yyshift int
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	if yychar < 0 {
		yylval.yys = yystate
		yychar = yylex1(yylex, &yylval)
		var ok bool
		if yyxchar, ok = yyXLAT[yychar]; !ok {
			yyxchar = len(yySymNames) // > tab width
		}
	}
	if yyDebug >= 4 {
		var a []int
		for _, v := range yyS[:yyp+1] {
			a = append(a, v.yys)
		}
		__yyfmt__.Printf("state stack %v\n", a)
	}
	row := yyParseTab[yystate]
	yyn = 0
	if yyxchar < len(row) {
		if yyn = int(row[yyxchar]); yyn != 0 {
			yyn += yyTabOfs
		}
	}
	switch {
	case yyn > 0: // shift
		yychar = -1
		yyVAL = yylval
		yystate = yyn
		yyshift = yyn
		if yyDebug >= 2 {
			__yyfmt__.Printf("shift, and goto state %d\n", yystate)
		}
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	case yyn < 0: // reduce
	case yystate == 1: // accept
		if yyDebug >= 2 {
			__yyfmt__.Println("accept")
		}
		goto ret0
	}

	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			if yyDebug >= 1 {
				__yyfmt__.Printf("no action for %s in state %d\n", yySymName(yychar), yystate)
			}
			msg, ok := yyXErrors[yyXError{yystate, yyxchar}]
			if !ok {
				msg, ok = yyXErrors[yyXError{yystate, -1}]
			}
			if !ok && yyshift != 0 {
				msg, ok = yyXErrors[yyXError{yyshift, yyxchar}]
			}
			if !ok {
				msg, ok = yyXErrors[yyXError{yyshift, -1}]
			}
			if yychar > 0 {
				ls := yyTokenLiteralStrings[yychar]
				if ls == "" {
					ls = yySymName(yychar)
				}
				if ls != "" {
					switch {
					case msg == "":
						msg = __yyfmt__.Sprintf("unexpected %s", ls)
					default:
						msg = __yyfmt__.Sprintf("unexpected %s, %s", ls, msg)
					}
				}
			}
			if msg == "" {
				msg = "syntax error"
			}
			yylex.Error(msg)
			Nerrs++
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				row := yyParseTab[yyS[yyp].yys]
				if yyError < len(row) {
					yyn = int(row[yyError]) + yyTabOfs
					if yyn > 0 { // hit
						if yyDebug >= 2 {
							__yyfmt__.Printf("error recovery found error shift in state %d\n", yyS[yyp].yys)
						}
						yystate = yyn /* simulate a shift of "error" */
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery failed\n")
			}
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yySymName(yychar))
			}
			if yychar == yyEofCode {
				goto ret1
			}

			yychar = -1
			goto yynewstate /* try again in the same state */
		}
	}

	r := -yyn
	x0 := yyReductions[r]
	x, n := x0.xsym, x0.components
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= n
	if yyp+1 >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	exState := yystate
	yystate = int(yyParseTab[yyS[yyp].yys][x]) + yyTabOfs
	/* reduction by production r */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce using rule %v (%s), and goto state %d\n", r, yySymNames[x], yystate)
	}

	switch r {
	case 1:
		{
			AddAST()
		}
	case 4:
		{
			AddInstruccion("pause")
		}
	case 5:
		{
			EjecutarExec()
		}
	case 6:
		{
			AddInstruccion("mkdisk")
		}
	case 7:
		{
			AddParametro()
			AddInstruccion("rmdisk")
		}
	case 8:
		{
			AddInstruccion("fdisk")
		}
	case 9:
		{
			AddInstruccion("mount")
		}
	case 10:
		{
			AddParametro()
			AddInstruccion("unmount")
		}
	case 11:
		{
			AddInstruccion("rep")
		}
	case 12:
		{
			AddInstruccion("mkfs")
		}
	case 13:
		{
			yyVAL.value = yyS[yypt-1].value + yyS[yypt-0].value
			AddParametro()
		}
	case 14:
		{
			yyVAL.value = yyS[yypt-1].value + yyS[yypt-0].value
			AddParametro()
		}
	case 15:
		{
			yyVAL.value = yyS[yypt-0].value
			AddParametro()
		}
	case 16:
		{
			yyVAL.value = yyS[yypt-0].value
			AddParametro()
		}
	case 17:
		{
			yyVAL.value = yyS[yypt-1].value + yyS[yypt-0].value
			AddParametro()
		}
	case 18:
		{
			yyVAL.value = yyS[yypt-1].value + yyS[yypt-0].value
			AddParametro()
		}
	case 19:
		{
			yyVAL.value = yyS[yypt-1].value + yyS[yypt-0].value
			AddParametro()
		}
	case 20:
		{
			yyVAL.value = yyS[yypt-0].value
			AddParametro()
		}
	case 21:
		{
			yyVAL.value = yyS[yypt-0].value
			AddParametro()
		}
	case 22:
		{
			yyVAL.value = yyS[yypt-0].value
			AddParametro()
		}
	case 23:
		{
			yyVAL.value = yyS[yypt-1].value + yyS[yypt-0].value
			AddParametro()
		}
	case 24:
		{
			yyVAL.value = yyS[yypt-1].value + yyS[yypt-0].value
			AddParametro()
		}
	case 25:
		{
			yyVAL.value = yyS[yypt-0].value
			AddParametro()
		}
	case 26:
		{
			yyVAL.value = yyS[yypt-0].value
			AddParametro()
		}
	case 27:
		{
			yyVAL.value = yyS[yypt-1].value + yyS[yypt-0].value
			AddParametro()
		}
	case 28:
		{
			yyVAL.value = yyS[yypt-1].value + yyS[yypt-0].value
			AddParametro()
		}
	case 29:
		{
			yyVAL.value = yyS[yypt-1].value + yyS[yypt-0].value
			AddParametro()
		}
	case 30:
		{
			yyVAL.value = yyS[yypt-1].value + yyS[yypt-0].value
			AddParametro()
		}
	case 31:
		{
			yyVAL.value = yyS[yypt-1].value + yyS[yypt-0].value
			AddParametro()
		}
	case 32:
		{
			yyVAL.value = yyS[yypt-1].value + yyS[yypt-0].value
			AddParametro()
		}
	case 33:
		{
			yyVAL.value = yyS[yypt-1].value + yyS[yypt-0].value
			AddParametro()
		}
	case 34:
		{
			yyVAL.value = yyS[yypt-1].value + yyS[yypt-0].value
			AddParametro()
		}
	case 35:
		{
			yyVAL.value = yyS[yypt-0].value
			AddParametro()
		}
	case 36:
		{
			yyVAL.value = yyS[yypt-0].value
			AddParametro()
		}
	case 37:
		{
			yyVAL.value = yyS[yypt-0].value
			AddParametro()
		}
	case 38:
		{
			yyVAL.value = yyS[yypt-0].value
			AddParametro()
		}
	case 39:
		{
			yyVAL.value = yyS[yypt-0].value
			AddParametro()
		}
	case 40:
		{
			yyVAL.value = yyS[yypt-0].value
			AddParametro()
		}
	case 41:
		{
			yyVAL.value = yyS[yypt-0].value
			AddParametro()
		}
	case 42:
		{
			yyVAL.value = yyS[yypt-0].value
			AddParametro()
		}
	case 43:
		{
			yyVAL.value = yyS[yypt-1].value
			AddParametro()
		}
	case 44:
		{
			yyVAL.value = yyS[yypt-1].value
			AddParametro()
		}
	case 45:
		{
			yyVAL.value = yyS[yypt-1].value
			AddParametro()
		}
	case 46:
		{
			yyVAL.value = yyS[yypt-1].value
			AddParametro()
		}
	case 47:
		{
			yyVAL.value = yyS[yypt-0].value
			AddParametro()
		}
	case 48:
		{
			yyVAL.value = yyS[yypt-0].value
			AddParametro()
		}
	case 49:
		{
			yyVAL.value = yyS[yypt-0].value
			AddParametro()
		}
	case 50:
		{
			yyVAL.value = yyS[yypt-0].value
			AddParametro()
		}
	case 51:
		{
			yyVAL.value = yyS[yypt-3].value + yyS[yypt-2].value + yyS[yypt-1].value + yyS[yypt-0].value
			auxPath = yyS[yypt-0].value
			CrearParametro(yyS[yypt-2].value, yyS[yypt-0].value)
		}
	case 52:
		{
			yyVAL.value = yyS[yypt-0].value
		}
	case 53:
		{
			yyVAL.value = yyS[yypt-0].value
		}
	case 54:
		{
			yyVAL.value = yyS[yypt-3].value + yyS[yypt-2].value + yyS[yypt-1].value + yyS[yypt-0].value
			CrearParametro(yyS[yypt-2].value, yyS[yypt-0].value)
		}
	case 55:
		{
			yyVAL.value = yyS[yypt-3].value + yyS[yypt-2].value + yyS[yypt-1].value + yyS[yypt-0].value
			CrearParametro(yyS[yypt-2].value, yyS[yypt-0].value)
		}
	case 56:
		{
			yyVAL.value = yyS[yypt-0].value
		}
	case 57:
		{
			yyVAL.value = yyS[yypt-0].value
		}
	case 58:
		{
			yyVAL.value = yyS[yypt-0].value
		}
	case 59:
		{
			yyVAL.value = yyS[yypt-3].value + yyS[yypt-2].value + yyS[yypt-1].value + yyS[yypt-0].value
			CrearParametro(yyS[yypt-2].value, yyS[yypt-0].value)
		}
	case 60:
		{
			yyVAL.value = yyS[yypt-0].value
		}
	case 61:
		{
			yyVAL.value = yyS[yypt-0].value
		}
	case 62:
		{
			yyVAL.value = yyS[yypt-3].value + yyS[yypt-2].value + yyS[yypt-1].value + yyS[yypt-0].value
			CrearParametro(yyS[yypt-2].value, yyS[yypt-0].value)
		}
	case 63:
		{
			yyVAL.value = yyS[yypt-0].value
		}
	case 64:
		{
			yyVAL.value = yyS[yypt-0].value
		}
	case 65:
		{
			yyVAL.value = yyS[yypt-0].value
		}
	case 66:
		{
			yyVAL.value = yyS[yypt-3].value + yyS[yypt-2].value + yyS[yypt-1].value + yyS[yypt-0].value
			CrearParametro(yyS[yypt-2].value, yyS[yypt-0].value)
		}
	case 67:
		{
			yyVAL.value = yyS[yypt-0].value
		}
	case 68:
		{
			yyVAL.value = yyS[yypt-0].value
		}
	case 69:
		{
			yyVAL.value = yyS[yypt-0].value
		}
	case 70:
		{
			yyVAL.value = yyS[yypt-3].value + yyS[yypt-2].value + yyS[yypt-1].value + yyS[yypt-0].value
			CrearParametro(yyS[yypt-2].value, yyS[yypt-0].value)
		}
	case 71:
		{
			yyVAL.value = yyS[yypt-3].value + yyS[yypt-2].value + yyS[yypt-1].value + yyS[yypt-0].value
			CrearParametro(yyS[yypt-2].value, yyS[yypt-0].value)
		}
	case 72:
		{
			yyVAL.value = yyS[yypt-0].value
		}
	case 73:
		{
			yyVAL.value = yyS[yypt-0].value
		}
	case 74:
		{
			yyVAL.value = yyS[yypt-3].value + yyS[yypt-2].value + yyS[yypt-1].value + yyS[yypt-0].value
			CrearParametro(yyS[yypt-2].value, yyS[yypt-0].value)
		}
	case 75:
		{
			yyVAL.value = yyS[yypt-4].value + yyS[yypt-3].value + yyS[yypt-2].value + yyS[yypt-1].value + yyS[yypt-0].value
			CrearParametro(yyS[yypt-3].value, "-"+yyS[yypt-0].value)
		}
	case 76:
		{
			yyVAL.value = yyS[yypt-4].value + yyS[yypt-3].value + yyS[yypt-2].value + yyS[yypt-1].value + yyS[yypt-0].value
		}
	case 77:
		{
			yyVAL.value = yyS[yypt-3].value + yyS[yypt-2].value + yyS[yypt-1].value + yyS[yypt-0].value
		}
	case 78:
		{
			yyVAL.value = yyS[yypt-4].value + yyS[yypt-3].value + yyS[yypt-2].value + yyS[yypt-1].value + yyS[yypt-0].value
			CrearParametro(yyS[yypt-2].value, yyS[yypt-0].value)
		}
	case 79:
		{
			yyVAL.value = yyS[yypt-3].value + yyS[yypt-2].value + yyS[yypt-1].value + yyS[yypt-0].value
			CrearParametro(yyS[yypt-2].value, yyS[yypt-0].value)
		}

	}

	if yyEx != nil && yyEx.Reduced(r, exState, &yyVAL) {
		return -1
	}
	goto yystack /* stack new state and value */
}

// CrearParametro da un valor nuevo al parámetro axuliar
func CrearParametro(tipo string, valor string) {
	oParametro = arbol.Parametro{Tipo: tipo, Valor: valor}
}

// CrearInstruccion da un valor nuevo a la intruccion auxiliar
func CrearInstruccion(tipo string) {
	oInstruccion = arbol.Instruccion{Tipo: tipo}
}

// AddParametro añade el parámetro auxiliar a la lista de parametros de la instrucción actual
func AddParametro() {
	oInstruccion.Parametros = append(oInstruccion.Parametros, oParametro)
}

// AddInstruccion añade el parámetro auxiliar a la lista de parametros de la instrucción actual
func AddInstruccion(tipo string) {
	oInstruccion.Tipo = tipo
	lInstruccion = append(lInstruccion, oInstruccion)
	oInstruccion = arbol.Instruccion{}
}

// AddAST añade lista de instrucciones para ejecutar
func AddAST() {
	ast := arbol.AST{}
	ast.Instrucciones = lInstruccion
	lInstruccion = []arbol.Instruccion{}
	lAST = append(lAST, ast)
}

func EjecutarExec() {
	if file, err := os.Open(auxPath); err == nil {
		yyParse(newLexer(bufio.NewReader(file)))

		ast := lAST[len(lAST)-1]
		ast.EjecutarAST()
		lAST = lAST[:len(lAST)-1]

		fmt.Println("Se ha ejecutado el archivo con èxito")
	} else {
		fmt.Println("No se ha podido abrir el archivo")
		log.Panic(err)
	}
}
