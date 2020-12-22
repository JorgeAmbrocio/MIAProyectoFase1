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
	yyDefault    = 57411
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
	rem          = 57365
	ren          = 57367
	rep          = 57374
	rf           = 57394
	rid          = 57405
	rmdisk       = 57351
	rmgrp        = 57359
	rmusr        = 57361
	rruta        = 57410
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
	yyTabOfs   = -122
)

var (
	yyPrec = map[int]int{}

	yyXLAT = map[int]int{
		57347: 0,   // guion (126x)
		57344: 1,   // $end (119x)
		57348: 2,   // exec (119x)
		57352: 3,   // fdisk (119x)
		57356: 4,   // login (119x)
		57357: 5,   // logout (119x)
		57368: 6,   // mkdir (119x)
		57350: 7,   // mkdisk (119x)
		57363: 8,   // mkfile (119x)
		57355: 9,   // mkfs (119x)
		57358: 10,  // mkgrp (119x)
		57360: 11,  // mkusr (119x)
		57353: 12,  // mount (119x)
		57349: 13,  // pause (119x)
		57365: 14,  // rem (119x)
		57374: 15,  // rep (119x)
		57351: 16,  // rmdisk (119x)
		57354: 17,  // unmount (119x)
		57434: 18,  // PARAMETRO_PATH (16x)
		57346: 19,  // asignacion (15x)
		57430: 20,  // PARAMETRO_ID (9x)
		57437: 21,  // PARAMETRO_SIZE (8x)
		57432: 22,  // PARAMETRO_NAME (6x)
		57375: 23,  // path (6x)
		57384: 24,  // id (5x)
		57405: 25,  // rid (5x)
		57377: 26,  // name (4x)
		57406: 27,  // numero (4x)
		57428: 28,  // PARAMETRO_FIT (4x)
		57433: 29,  // PARAMETRO_P (4x)
		57435: 30,  // PARAMETRO_PWD (4x)
		57440: 31,  // PARAMETRO_UNIT (4x)
		57441: 32,  // PARAMETRO_USR (4x)
		57408: 33,  // rutaCompleja (4x)
		57376: 34,  // size (3x)
		57404: 35,  // fast (2x)
		57380: 36,  // fit (2x)
		57403: 37,  // full (2x)
		57413: 38,  // INSTRUCCION (2x)
		57419: 39,  // LST_MKFILE (2x)
		57391: 40,  // p (2x)
		57426: 41,  // PARAMETRO_ADD (2x)
		57427: 42,  // PARAMETRO_DELETE (2x)
		57429: 43,  // PARAMETRO_GRP (2x)
		57436: 44,  // PARAMETRO_RUTA (2x)
		57438: 45,  // PARAMETRO_TYPE (2x)
		57439: 46,  // PARAMETRO_TYPE_FS (2x)
		57379: 47,  // rtype (2x)
		57407: 48,  // rutaSimple (2x)
		57378: 49,  // unit (2x)
		57442: 50,  // VALOR_DELETE (2x)
		57445: 51,  // VALOR_PATH (2x)
		57382: 52,  // add (1x)
		57409: 53,  // archivo (1x)
		57400: 54,  // bf (1x)
		57381: 55,  // delete (1x)
		57398: 56,  // e (1x)
		57401: 57,  // ff (1x)
		57412: 58,  // INICIO (1x)
		57396: 59,  // k (1x)
		57399: 60,  // l (1x)
		57414: 61,  // LISTA_INSTRUCCION (1x)
		57415: 62,  // LST_FDISK (1x)
		57416: 63,  // LST_LOGIN (1x)
		57417: 64,  // LST_MKDIR (1x)
		57418: 65,  // LST_MKDIS (1x)
		57420: 66,  // LST_MKFS (1x)
		57421: 67,  // LST_MKGRP (1x)
		57422: 68,  // LST_MKUSR (1x)
		57423: 69,  // LST_MOUNT (1x)
		57424: 70,  // LST_REM (1x)
		57425: 71,  // LST_REP (1x)
		57397: 72,  // m (1x)
		57387: 73,  // pwd (1x)
		57410: 74,  // rruta (1x)
		57386: 75,  // usr (1x)
		57443: 76,  // VALOR_FIT (1x)
		57444: 77,  // VALOR_NAME (1x)
		57446: 78,  // VALOR_TYPE (1x)
		57447: 79,  // VALOR_UNIT (1x)
		57402: 80,  // wf (1x)
		57411: 81,  // $default (0x)
		57364: 82,  // cat (0x)
		57373: 83,  // chgrp (0x)
		57362: 84,  // chmod (0x)
		57372: 85,  // chown (0x)
		57392: 86,  // cont (0x)
		57369: 87,  // cp (0x)
		57395: 88,  // dest (0x)
		57366: 89,  // edit (0x)
		57345: 90,  // error (0x)
		57393: 91,  // filen (0x)
		57371: 92,  // find (0x)
		57388: 93,  // grp (0x)
		57383: 94,  // idn (0x)
		57370: 95,  // mv (0x)
		57431: 96,  // PARAMETRO_IDN (0x)
		57390: 97,  // r (0x)
		57367: 98,  // ren (0x)
		57394: 99,  // rf (0x)
		57359: 100, // rmgrp (0x)
		57361: 101, // rmusr (0x)
		57385: 102, // tipo (0x)
		57389: 103, // ugo (0x)
	}

	yySymNames = []string{
		"guion",
		"$end",
		"exec",
		"fdisk",
		"login",
		"logout",
		"mkdir",
		"mkdisk",
		"mkfile",
		"mkfs",
		"mkgrp",
		"mkusr",
		"mount",
		"pause",
		"rem",
		"rep",
		"rmdisk",
		"unmount",
		"PARAMETRO_PATH",
		"asignacion",
		"PARAMETRO_ID",
		"PARAMETRO_SIZE",
		"PARAMETRO_NAME",
		"path",
		"id",
		"rid",
		"name",
		"numero",
		"PARAMETRO_FIT",
		"PARAMETRO_P",
		"PARAMETRO_PWD",
		"PARAMETRO_UNIT",
		"PARAMETRO_USR",
		"rutaCompleja",
		"size",
		"fast",
		"fit",
		"full",
		"INSTRUCCION",
		"LST_MKFILE",
		"p",
		"PARAMETRO_ADD",
		"PARAMETRO_DELETE",
		"PARAMETRO_GRP",
		"PARAMETRO_RUTA",
		"PARAMETRO_TYPE",
		"PARAMETRO_TYPE_FS",
		"rtype",
		"rutaSimple",
		"unit",
		"VALOR_DELETE",
		"VALOR_PATH",
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
		"LST_LOGIN",
		"LST_MKDIR",
		"LST_MKDIS",
		"LST_MKFS",
		"LST_MKGRP",
		"LST_MKUSR",
		"LST_MOUNT",
		"LST_REM",
		"LST_REP",
		"m",
		"pwd",
		"rruta",
		"usr",
		"VALOR_FIT",
		"VALOR_NAME",
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
		"mv",
		"PARAMETRO_IDN",
		"r",
		"ren",
		"rf",
		"rmgrp",
		"rmusr",
		"tipo",
		"ugo",
	}

	yyTokenLiteralStrings = map[int]string{}

	yyReductions = map[int]struct{ xsym, components int }{
		0:   {0, 1},
		1:   {58, 1},
		2:   {61, 2},
		3:   {61, 1},
		4:   {38, 1},
		5:   {38, 1},
		6:   {38, 2},
		7:   {38, 2},
		8:   {38, 2},
		9:   {38, 2},
		10:  {38, 2},
		11:  {38, 2},
		12:  {38, 2},
		13:  {38, 2},
		14:  {38, 2},
		15:  {38, 2},
		16:  {38, 2},
		17:  {38, 2},
		18:  {38, 2},
		19:  {38, 2},
		20:  {70, 2},
		21:  {70, 1},
		22:  {64, 2},
		23:  {64, 2},
		24:  {64, 1},
		25:  {64, 1},
		26:  {39, 2},
		27:  {39, 2},
		28:  {39, 2},
		29:  {39, 1},
		30:  {39, 1},
		31:  {39, 1},
		32:  {68, 2},
		33:  {68, 2},
		34:  {68, 2},
		35:  {68, 1},
		36:  {68, 1},
		37:  {68, 1},
		38:  {67, 2},
		39:  {67, 1},
		40:  {63, 2},
		41:  {63, 2},
		42:  {63, 2},
		43:  {63, 1},
		44:  {63, 1},
		45:  {63, 1},
		46:  {66, 2},
		47:  {66, 2},
		48:  {66, 1},
		49:  {66, 1},
		50:  {71, 2},
		51:  {71, 2},
		52:  {71, 2},
		53:  {71, 2},
		54:  {71, 1},
		55:  {71, 1},
		56:  {71, 1},
		57:  {71, 1},
		58:  {69, 2},
		59:  {69, 2},
		60:  {69, 1},
		61:  {69, 1},
		62:  {62, 2},
		63:  {62, 2},
		64:  {62, 2},
		65:  {62, 2},
		66:  {62, 2},
		67:  {62, 2},
		68:  {62, 2},
		69:  {62, 2},
		70:  {62, 1},
		71:  {62, 1},
		72:  {62, 1},
		73:  {62, 1},
		74:  {62, 1},
		75:  {62, 1},
		76:  {62, 1},
		77:  {62, 1},
		78:  {65, 2},
		79:  {65, 2},
		80:  {65, 2},
		81:  {65, 2},
		82:  {65, 1},
		83:  {65, 1},
		84:  {65, 1},
		85:  {65, 1},
		86:  {43, 4},
		87:  {18, 4},
		88:  {44, 4},
		89:  {51, 1},
		90:  {51, 1},
		91:  {21, 4},
		92:  {22, 4},
		93:  {77, 1},
		94:  {77, 1},
		95:  {77, 1},
		96:  {32, 4},
		97:  {30, 4},
		98:  {30, 4},
		99:  {32, 4},
		100: {31, 4},
		101: {79, 1},
		102: {79, 1},
		103: {45, 4},
		104: {78, 1},
		105: {78, 1},
		106: {78, 1},
		107: {28, 4},
		108: {29, 2},
		109: {76, 1},
		110: {76, 1},
		111: {76, 1},
		112: {42, 4},
		113: {46, 4},
		114: {50, 1},
		115: {50, 1},
		116: {41, 4},
		117: {41, 5},
		118: {96, 5},
		119: {96, 4},
		120: {20, 5},
		121: {20, 4},
	}

	yyXErrors = map[yyXError]string{}

	yyParseTab = [177][]uint16{
		// 0
		{2: 128, 131, 136, 127, 140, 129, 139, 135, 137, 138, 132, 126, 141, 134, 130, 133, 38: 125, 58: 123, 61: 124},
		{1: 122},
		{1: 121, 128, 131, 136, 127, 140, 129, 139, 135, 137, 138, 132, 126, 141, 134, 130, 133, 38: 298},
		{1: 119, 119, 119, 119, 119, 119, 119, 119, 119, 119, 119, 119, 119, 119, 119, 119, 119},
		{1: 118, 118, 118, 118, 118, 118, 118, 118, 118, 118, 118, 118, 118, 118, 118, 118, 118},
		// 5
		{1: 117, 117, 117, 117, 117, 117, 117, 117, 117, 117, 117, 117, 117, 117, 117, 117, 117},
		{144, 18: 297},
		{292, 18: 289, 21: 288, 28: 291, 31: 290, 65: 287},
		{144, 18: 286},
		{252, 18: 246, 21: 244, 250, 28: 248, 31: 245, 41: 251, 249, 45: 247, 62: 243},
		// 10
		{240, 18: 238, 22: 239, 69: 237},
		{236, 20: 235},
		{221, 18: 218, 20: 219, 22: 217, 44: 220, 71: 216},
		{208, 20: 206, 46: 207, 66: 205},
		{173, 20: 201, 30: 200, 32: 199, 63: 198},
		// 15
		{193, 43: 192, 67: 191},
		{173, 20: 172, 30: 171, 32: 170, 68: 169},
		{156, 18: 165, 21: 155, 29: 166, 39: 164},
		{156, 18: 153, 21: 155, 29: 154, 39: 152, 64: 151},
		{144, 18: 143, 70: 142},
		// 20
		{144, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 150},
		{101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101},
		{23: 145},
		{19: 146},
		{33: 149, 48: 148, 51: 147},
		// 25
		{35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35},
		{33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33},
		{32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32},
		{102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102},
		{1: 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 104},
		// 30
		{156, 18: 161, 21: 163, 29: 162},
		{93, 98, 98, 98, 98, 98, 98, 98, 98, 98, 98, 98, 98, 98, 98, 98, 98, 98},
		{91, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97},
		{92, 92, 92, 92, 92, 92, 92, 92, 92, 92, 92, 92, 92, 92, 92, 92, 92, 92},
		{23: 145, 34: 157, 40: 158},
		// 35
		{19: 159},
		{14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14},
		{27: 160},
		{31, 31, 31, 31, 31, 31, 31, 31, 31, 31, 31, 31, 31, 31, 31, 31, 31, 31},
		{96, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100},
		// 40
		{94, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99},
		{95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95},
		{156, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 167, 21: 163, 29: 168},
		{93, 93, 93, 93, 93, 93, 93, 93, 93, 93, 93, 93, 93, 93, 93, 93, 93, 93},
		{91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91},
		// 45
		{96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96},
		{94, 94, 94, 94, 94, 94, 94, 94, 94, 94, 94, 94, 94, 94, 94, 94, 94, 94},
		{173, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 20: 190, 30: 189, 32: 188},
		{87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87},
		{86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86},
		// 50
		{184, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85},
		{25: 176, 73: 175, 75: 174},
		{19: 182},
		{19: 179},
		{19: 177},
		// 55
		{24: 178},
		{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		{24: 180, 27: 181},
		{25, 25, 25, 25, 25, 25, 25, 25, 25, 25, 25, 25, 25, 25, 25, 25, 25, 25},
		{24, 24, 24, 24, 24, 24, 24, 24, 24, 24, 24, 24, 24, 24, 24, 24, 24, 24},
		// 60
		{24: 183},
		{26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26},
		{25: 185},
		{19: 186},
		{24: 187},
		// 65
		{2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2},
		{90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90},
		{89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89},
		{184, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88},
		{193, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 43: 197},
		// 70
		{83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83},
		{26: 194},
		{19: 195},
		{33: 196},
		{36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36},
		// 75
		{84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84},
		{173, 108, 108, 108, 108, 108, 108, 108, 108, 108, 108, 108, 108, 108, 108, 108, 108, 108, 20: 204, 30: 203, 32: 202},
		{79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79},
		{78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78},
		{184, 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 77},
		// 80
		{82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82},
		{81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81},
		{184, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80},
		{208, 109, 109, 109, 109, 109, 109, 109, 109, 109, 109, 109, 109, 109, 109, 109, 109, 109, 20: 214, 46: 215},
		{184, 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 74},
		// 85
		{73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73},
		{25: 176, 47: 209},
		{19: 210},
		{35: 213, 37: 212, 50: 211},
		{9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9},
		// 90
		{8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8},
		{7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7},
		{184, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76},
		{75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75},
		{221, 110, 110, 110, 110, 110, 110, 110, 110, 110, 110, 110, 110, 110, 110, 110, 110, 110, 232, 20: 233, 22: 231, 44: 234},
		// 95
		{68, 68, 68, 68, 68, 68, 68, 68, 68, 68, 68, 68, 68, 68, 68, 68, 68, 68},
		{67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67},
		{184, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66},
		{65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65},
		{23: 145, 25: 176, 223, 74: 222},
		// 100
		{19: 229},
		{19: 224},
		{24: 227, 33: 228, 53: 226, 77: 225},
		{30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30},
		{29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29},
		// 105
		{28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28},
		{27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27},
		{33: 149, 48: 148, 51: 230},
		{34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34},
		{72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 72},
		// 110
		{71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71},
		{184, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70},
		{69, 69, 69, 69, 69, 69, 69, 69, 69, 69, 69, 69, 69, 69, 69, 69, 69, 69},
		{184, 111, 111, 111, 111, 111, 111, 111, 111, 111, 111, 111, 111, 111, 111, 111, 111, 111},
		{25: 176},
		// 115
		{240, 112, 112, 112, 112, 112, 112, 112, 112, 112, 112, 112, 112, 112, 112, 112, 112, 112, 241, 22: 242},
		{62, 62, 62, 62, 62, 62, 62, 62, 62, 62, 62, 62, 62, 62, 62, 62, 62, 62},
		{61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61},
		{23: 145, 26: 223},
		{64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64},
		// 120
		{63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63},
		{252, 113, 113, 113, 113, 113, 113, 113, 113, 113, 113, 113, 113, 113, 113, 113, 113, 113, 280, 21: 278, 284, 28: 282, 31: 279, 41: 285, 283, 45: 281},
		{52, 52, 52, 52, 52, 52, 52, 52, 52, 52, 52, 52, 52, 52, 52, 52, 52, 52},
		{51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51},
		{50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50},
		// 125
		{49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49},
		{48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48},
		{47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47},
		{46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46},
		{45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45},
		// 130
		{23: 145, 26: 223, 34: 157, 36: 255, 47: 254, 49: 253, 52: 257, 55: 256},
		{19: 274},
		{19: 269},
		{19: 264},
		{19: 262},
		// 135
		{19: 258},
		{260, 27: 259},
		{6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6},
		{27: 261},
		{5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5},
		// 140
		{35: 213, 37: 212, 50: 263},
		{10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10},
		{54: 266, 57: 267, 76: 265, 80: 268},
		{15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15},
		{13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13},
		// 145
		{12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12},
		{11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11},
		{40: 271, 56: 272, 60: 273, 78: 270},
		{19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19},
		{18, 18, 18, 18, 18, 18, 18, 18, 18, 18, 18, 18, 18, 18, 18, 18, 18, 18},
		// 150
		{17, 17, 17, 17, 17, 17, 17, 17, 17, 17, 17, 17, 17, 17, 17, 17, 17, 17},
		{16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16},
		{59: 276, 72: 277, 79: 275},
		{22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22},
		{21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21},
		// 155
		{20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20},
		{60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60},
		{59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59},
		{58, 58, 58, 58, 58, 58, 58, 58, 58, 58, 58, 58, 58, 58, 58, 58, 58, 58},
		{57, 57, 57, 57, 57, 57, 57, 57, 57, 57, 57, 57, 57, 57, 57, 57, 57, 57},
		// 160
		{56, 56, 56, 56, 56, 56, 56, 56, 56, 56, 56, 56, 56, 56, 56, 56, 56, 56},
		{55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55},
		{54, 54, 54, 54, 54, 54, 54, 54, 54, 54, 54, 54, 54, 54, 54, 54, 54, 54},
		{53, 53, 53, 53, 53, 53, 53, 53, 53, 53, 53, 53, 53, 53, 53, 53, 53, 53},
		{1: 114, 114, 114, 114, 114, 114, 114, 114, 114, 114, 114, 114, 114, 114, 114, 114, 114},
		// 165
		{292, 115, 115, 115, 115, 115, 115, 115, 115, 115, 115, 115, 115, 115, 115, 115, 115, 115, 294, 21: 293, 28: 296, 31: 295},
		{40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40},
		{39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39},
		{38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38},
		{37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37},
		// 170
		{23: 145, 34: 157, 36: 255, 49: 253},
		{44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44},
		{43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43},
		{42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42},
		{41, 41, 41, 41, 41, 41, 41, 41, 41, 41, 41, 41, 41, 41, 41, 41, 41, 41},
		// 175
		{1: 116, 116, 116, 116, 116, 116, 116, 116, 116, 116, 116, 116, 116, 116, 116, 116, 116},
		{1: 120, 120, 120, 120, 120, 120, 120, 120, 120, 120, 120, 120, 120, 120, 120, 120, 120},
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
	const yyError = 90

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
			AddInstruccion("logout")
		}
	case 6:
		{
			EjecutarExec()
		}
	case 7:
		{
			AddInstruccion("mkdisk")
		}
	case 8:
		{
			AddParametro()
			AddInstruccion("rmdisk")
		}
	case 9:
		{
			AddInstruccion("fdisk")
		}
	case 10:
		{
			AddInstruccion("mount")
		}
	case 11:
		{
			AddParametro()
			AddInstruccion("unmount")
		}
	case 12:
		{
			AddInstruccion("rep")
		}
	case 13:
		{
			AddInstruccion("mkfs")
		}
	case 14:
		{
			AddInstruccion("login")
		}
	case 15:
		{
			AddInstruccion("mkgrp")
		}
	case 16:
		{
			AddInstruccion("mkusr")
		}
	case 17:
		{
			AddInstruccion("mkfile")
		}
	case 18:
		{
			AddInstruccion("mkdir")
		}
	case 19:
		{
			AddInstruccion("rem")
		}
	case 20:
		{
			yyVAL.value = yyS[yypt-1].value + yyS[yypt-0].value
			AddParametro()
		}
	case 21:
		{
			yyVAL.value = yyS[yypt-0].value
			AddParametro()
		}
	case 22:
		{
			yyVAL.value = yyS[yypt-1].value + yyS[yypt-0].value
			AddParametro()
		}
	case 23:
		{
			yyVAL.value = yyS[yypt-1].value + yyS[yypt-0].value
			AddParametro()
		}
	case 24:
		{
			yyVAL.value = yyS[yypt-0].value
			AddParametro()
		}
	case 25:
		{
			yyVAL.value = yyS[yypt-0].value
			AddParametro()
		}
	case 26:
		{
			yyVAL.value = yyS[yypt-1].value + yyS[yypt-0].value
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
			yyVAL.value = yyS[yypt-0].value
			AddParametro()
		}
	case 30:
		{
			yyVAL.value = yyS[yypt-0].value
			AddParametro()
		}
	case 31:
		{
			yyVAL.value = yyS[yypt-0].value
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
			yyVAL.value = yyS[yypt-1].value + yyS[yypt-0].value
			AddParametro()
		}
	case 39:
		{
			yyVAL.value = yyS[yypt-0].value
			AddParametro()
		}
	case 40:
		{
			yyVAL.value = yyS[yypt-1].value + yyS[yypt-0].value
			AddParametro()
		}
	case 41:
		{
			yyVAL.value = yyS[yypt-1].value + yyS[yypt-0].value
			AddParametro()
		}
	case 42:
		{
			yyVAL.value = yyS[yypt-1].value + yyS[yypt-0].value
			AddParametro()
		}
	case 43:
		{
			yyVAL.value = yyS[yypt-0].value
			AddParametro()
		}
	case 44:
		{
			yyVAL.value = yyS[yypt-0].value
			AddParametro()
		}
	case 45:
		{
			yyVAL.value = yyS[yypt-0].value
			AddParametro()
		}
	case 46:
		{
			yyVAL.value = yyS[yypt-1].value + yyS[yypt-0].value
			AddParametro()
		}
	case 47:
		{
			yyVAL.value = yyS[yypt-1].value + yyS[yypt-0].value
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
			yyVAL.value = yyS[yypt-1].value + yyS[yypt-0].value
			AddParametro()
		}
	case 51:
		{
			yyVAL.value = yyS[yypt-1].value + yyS[yypt-0].value
			AddParametro()
		}
	case 52:
		{
			yyVAL.value = yyS[yypt-1].value + yyS[yypt-0].value
			AddParametro()
		}
	case 53:
		{
			yyVAL.value = yyS[yypt-1].value + yyS[yypt-0].value
			AddParametro()
		}
	case 54:
		{
			yyVAL.value = yyS[yypt-0].value
			AddParametro()
		}
	case 55:
		{
			yyVAL.value = yyS[yypt-0].value
			AddParametro()
		}
	case 56:
		{
			yyVAL.value = yyS[yypt-0].value
			AddParametro()
		}
	case 57:
		{
			yyVAL.value = yyS[yypt-0].value
			AddParametro()
		}
	case 58:
		{
			yyVAL.value = yyS[yypt-1].value + yyS[yypt-0].value
			AddParametro()
		}
	case 59:
		{
			yyVAL.value = yyS[yypt-1].value + yyS[yypt-0].value
			AddParametro()
		}
	case 60:
		{
			yyVAL.value = yyS[yypt-0].value
			AddParametro()
		}
	case 61:
		{
			yyVAL.value = yyS[yypt-0].value
			AddParametro()
		}
	case 62:
		{
			yyVAL.value = yyS[yypt-1].value + yyS[yypt-0].value
			AddParametro()
		}
	case 63:
		{
			yyVAL.value = yyS[yypt-1].value + yyS[yypt-0].value
			AddParametro()
		}
	case 64:
		{
			yyVAL.value = yyS[yypt-1].value + yyS[yypt-0].value
			AddParametro()
		}
	case 65:
		{
			yyVAL.value = yyS[yypt-1].value + yyS[yypt-0].value
			AddParametro()
		}
	case 66:
		{
			yyVAL.value = yyS[yypt-1].value + yyS[yypt-0].value
			AddParametro()
		}
	case 67:
		{
			yyVAL.value = yyS[yypt-1].value + yyS[yypt-0].value
			AddParametro()
		}
	case 68:
		{
			yyVAL.value = yyS[yypt-1].value + yyS[yypt-0].value
			AddParametro()
		}
	case 69:
		{
			yyVAL.value = yyS[yypt-1].value + yyS[yypt-0].value
			AddParametro()
		}
	case 70:
		{
			yyVAL.value = yyS[yypt-0].value
			AddParametro()
		}
	case 71:
		{
			yyVAL.value = yyS[yypt-0].value
			AddParametro()
		}
	case 72:
		{
			yyVAL.value = yyS[yypt-0].value
			AddParametro()
		}
	case 73:
		{
			yyVAL.value = yyS[yypt-0].value
			AddParametro()
		}
	case 74:
		{
			yyVAL.value = yyS[yypt-0].value
			AddParametro()
		}
	case 75:
		{
			yyVAL.value = yyS[yypt-0].value
			AddParametro()
		}
	case 76:
		{
			yyVAL.value = yyS[yypt-0].value
			AddParametro()
		}
	case 77:
		{
			yyVAL.value = yyS[yypt-0].value
			AddParametro()
		}
	case 78:
		{
			yyVAL.value = yyS[yypt-1].value
			AddParametro()
		}
	case 79:
		{
			yyVAL.value = yyS[yypt-1].value
			AddParametro()
		}
	case 80:
		{
			yyVAL.value = yyS[yypt-1].value
			AddParametro()
		}
	case 81:
		{
			yyVAL.value = yyS[yypt-1].value
			AddParametro()
		}
	case 82:
		{
			yyVAL.value = yyS[yypt-0].value
			AddParametro()
		}
	case 83:
		{
			yyVAL.value = yyS[yypt-0].value
			AddParametro()
		}
	case 84:
		{
			yyVAL.value = yyS[yypt-0].value
			AddParametro()
		}
	case 85:
		{
			yyVAL.value = yyS[yypt-0].value
			AddParametro()
		}
	case 86:
		{
			yyVAL.value = yyS[yypt-3].value + yyS[yypt-2].value + yyS[yypt-1].value + yyS[yypt-0].value
			CrearParametro(yyS[yypt-2].value, yyS[yypt-0].value)
		}
	case 87:
		{
			yyVAL.value = yyS[yypt-3].value + yyS[yypt-2].value + yyS[yypt-1].value + yyS[yypt-0].value
			auxPath = yyS[yypt-0].value
			CrearParametro(yyS[yypt-2].value, yyS[yypt-0].value)
		}
	case 88:
		{
			yyVAL.value = yyS[yypt-3].value + yyS[yypt-2].value + yyS[yypt-1].value + yyS[yypt-0].value
			auxPath = yyS[yypt-0].value
			CrearParametro(yyS[yypt-2].value, yyS[yypt-0].value)
		}
	case 89:
		{
			yyVAL.value = yyS[yypt-0].value
		}
	case 90:
		{
			yyVAL.value = yyS[yypt-0].value
		}
	case 91:
		{
			yyVAL.value = yyS[yypt-3].value + yyS[yypt-2].value + yyS[yypt-1].value + yyS[yypt-0].value
			CrearParametro(yyS[yypt-2].value, yyS[yypt-0].value)
		}
	case 92:
		{
			yyVAL.value = yyS[yypt-3].value + yyS[yypt-2].value + yyS[yypt-1].value + yyS[yypt-0].value
			CrearParametro(yyS[yypt-2].value, yyS[yypt-0].value)
		}
	case 93:
		{
			yyVAL.value = yyS[yypt-0].value
		}
	case 94:
		{
			yyVAL.value = yyS[yypt-0].value
		}
	case 95:
		{
			yyVAL.value = yyS[yypt-0].value
		}
	case 96:
		{
			yyVAL.value = yyS[yypt-3].value + yyS[yypt-2].value + yyS[yypt-1].value + yyS[yypt-0].value
			CrearParametro(yyS[yypt-2].value, yyS[yypt-0].value)
		}
	case 97:
		{
			yyVAL.value = yyS[yypt-3].value + yyS[yypt-2].value + yyS[yypt-1].value + yyS[yypt-0].value
			CrearParametro(yyS[yypt-2].value, yyS[yypt-0].value)
		}
	case 98:
		{
			yyVAL.value = yyS[yypt-3].value + yyS[yypt-2].value + yyS[yypt-1].value + yyS[yypt-0].value
			CrearParametro(yyS[yypt-2].value, yyS[yypt-0].value)
		}
	case 99:
		{
			yyVAL.value = yyS[yypt-3].value + yyS[yypt-2].value + yyS[yypt-1].value + yyS[yypt-0].value
			CrearParametro(yyS[yypt-2].value, yyS[yypt-0].value)
		}
	case 100:
		{
			yyVAL.value = yyS[yypt-3].value + yyS[yypt-2].value + yyS[yypt-1].value + yyS[yypt-0].value
			CrearParametro(yyS[yypt-2].value, yyS[yypt-0].value)
		}
	case 101:
		{
			yyVAL.value = yyS[yypt-0].value
		}
	case 102:
		{
			yyVAL.value = yyS[yypt-0].value
		}
	case 103:
		{
			yyVAL.value = yyS[yypt-3].value + yyS[yypt-2].value + yyS[yypt-1].value + yyS[yypt-0].value
			CrearParametro(yyS[yypt-2].value, yyS[yypt-0].value)
		}
	case 104:
		{
			yyVAL.value = yyS[yypt-0].value
		}
	case 105:
		{
			yyVAL.value = yyS[yypt-0].value
		}
	case 106:
		{
			yyVAL.value = yyS[yypt-0].value
		}
	case 107:
		{
			yyVAL.value = yyS[yypt-3].value + yyS[yypt-2].value + yyS[yypt-1].value + yyS[yypt-0].value
			CrearParametro(yyS[yypt-2].value, yyS[yypt-0].value)
		}
	case 108:
		{
			yyVAL.value = yyS[yypt-1].value + yyS[yypt-0].value
			CrearParametro(yyS[yypt-0].value, yyS[yypt-0].value)
		}
	case 109:
		{
			yyVAL.value = yyS[yypt-0].value
		}
	case 110:
		{
			yyVAL.value = yyS[yypt-0].value
		}
	case 111:
		{
			yyVAL.value = yyS[yypt-0].value
		}
	case 112:
		{
			yyVAL.value = yyS[yypt-3].value + yyS[yypt-2].value + yyS[yypt-1].value + yyS[yypt-0].value
			CrearParametro(yyS[yypt-2].value, yyS[yypt-0].value)
		}
	case 113:
		{
			yyVAL.value = yyS[yypt-3].value + yyS[yypt-2].value + yyS[yypt-1].value + yyS[yypt-0].value
			CrearParametro(yyS[yypt-2].value, yyS[yypt-0].value)
		}
	case 114:
		{
			yyVAL.value = yyS[yypt-0].value
		}
	case 115:
		{
			yyVAL.value = yyS[yypt-0].value
		}
	case 116:
		{
			yyVAL.value = yyS[yypt-3].value + yyS[yypt-2].value + yyS[yypt-1].value + yyS[yypt-0].value
			CrearParametro(yyS[yypt-2].value, yyS[yypt-0].value)
		}
	case 117:
		{
			yyVAL.value = yyS[yypt-4].value + yyS[yypt-3].value + yyS[yypt-2].value + yyS[yypt-1].value + yyS[yypt-0].value
			CrearParametro(yyS[yypt-3].value, "-"+yyS[yypt-0].value)
		}
	case 118:
		{
			yyVAL.value = yyS[yypt-4].value + yyS[yypt-3].value + yyS[yypt-2].value + yyS[yypt-1].value + yyS[yypt-0].value
		}
	case 119:
		{
			yyVAL.value = yyS[yypt-3].value + yyS[yypt-2].value + yyS[yypt-1].value + yyS[yypt-0].value
		}
	case 120:
		{
			yyVAL.value = yyS[yypt-4].value + yyS[yypt-3].value + yyS[yypt-2].value + yyS[yypt-1].value + yyS[yypt-0].value
			CrearParametro(yyS[yypt-2].value, yyS[yypt-0].value)
		}
	case 121:
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
