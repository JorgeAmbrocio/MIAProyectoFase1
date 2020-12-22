
%{
// +build ignore
package main

import (
    "bufio"
    "fmt"
    "log"
    //"archivos/MIA-Proyecto1_201709454/analizador/arbol"
    "proyectos/MIAProyectoFase1/analizador/arbol"
    "os"
)

var oParametro arbol.Parametro
var oInstruccion arbol.Instruccion
var lInstruccion []arbol.Instruccion
var lAST []arbol.AST
var auxPath string

%}

%union{
    value string
}

%token	asignacion, guion
%token  exec, pause, mkdisk, rmdisk, fdisk, mount, unmount, mkfs
%token  login, logout, mkgrp, rmgrp, mkusr, rmusr,chmod
%token  mkfile, cat,rem, edit, ren, mkdir, cp, mv, find, chown
%token  chgrp, rep, path, size, name, unit, rtype, fit, delete
%token  add, idn, id, tipo, usr, pwd, grp, ugo, r, p, cont
%token  filen, rf, dest, k, m, e, l, bf,ff, wf, full, fast, rid
%token  numero, rutaSimple, rutaCompleja, archivo,rruta

%type <value> INICIO
%type <value> asignacion, guion
%type <value> exec, pause, mkdisk, rmdisk, fdisk, mount, unmount, mkfs
%type <value> login, logout, mkgrp, rmgrp, mkusr, rmusr,chmod
%type <value> mkfile, cat,rem, edit, ren, mkdir, cp, mv, find, chown
%type <value> chgrp, rep, path, size, name, unit, rtype, fit, delete
%type <value> add, idn, id, tipo, usr, pwd, grp, ugo, r, p, cont
%type <value> filen, rf, dest, k, m, e, l, bf,ff,wf, full, fast, rid
%type <value>  numero, rutaSimple, rutaCompleja, archivo,rruta
%type <value> INSTRUCCION, LISTA_INSTRUCCION, PARAMETRO_PATH, VALOR_PATH
%type <value> PARAMETRO_SIZE, PARAMETRO_NAME, VALOR_NAME, PARAMETRO_UNIT, VALOR_UNIT 
%type <value> PARAMETRO_TYPE, VALOR_TYPE, PARAMETRO_FIT, VALOR_FIT
%type <value> PARAMETRO_DELETE, VALOR_DELETE, PARAMETRO_ADD, PARAMETRO_IDN
%type <value> LST_FDISK, LST_MKDIS, LST_MOUNT,PARAMETRO_ID,LST_REP,PARAMETRO_TYPE_FS
%type <value> LST_MKFS, LST_LOGIN,PARAMETRO_USR,PARAMETRO_PWD,LST_MKGRP,PARAMETRO_GRP
%type <value> LST_MKUSR, LST_MKFILE, PARAMETRO_P,LST_MKDIR,PARAMETRO_RUTA,LST_REM
%type <value> PARAMETRO_FILEN,LST_CAT,LST_REN


%start INICIO

%% /* The grammar follows.  */

INICIO:
    LISTA_INSTRUCCION   { AddAST(); }
    //| '\n' {}
;

LISTA_INSTRUCCION:
    LISTA_INSTRUCCION INSTRUCCION //'\n'
    |INSTRUCCION //'\n'
    //|'\n'
;

INSTRUCCION:
    pause                   {AddInstruccion("pause"); }
    |logout                 {AddInstruccion("logout");}
    |exec PARAMETRO_PATH    { EjecutarExec(); }
    |mkdisk LST_MKDIS       { AddInstruccion("mkdisk"); }
    |rmdisk PARAMETRO_PATH  { AddParametro(); AddInstruccion("rmdisk"); }
    |fdisk LST_FDISK        { AddInstruccion("fdisk"); }
    |mount LST_MOUNT        { AddInstruccion("mount"); }
    |unmount PARAMETRO_ID   { AddParametro(); AddInstruccion("unmount"); }
    |rep LST_REP            { AddInstruccion("rep"); }
    |mkfs LST_MKFS          { AddInstruccion("mkfs")}
    |login LST_LOGIN        { AddInstruccion("login")}
    |mkgrp LST_MKGRP        { AddInstruccion("mkgrp")}
    |mkusr LST_MKUSR        { AddInstruccion("mkusr")}
    |mkfile LST_MKFILE      { AddInstruccion("mkfile")}
    |mkdir LST_MKDIR        { AddInstruccion("mkdir")}
    |rem LST_REM            { AddInstruccion("rem")}
    |cat LST_CAT            { AddInstruccion("cat")}
    |ren LST_REN            { AddInstruccion("ren")}
;

LST_REN:
    LST_REN PARAMETRO_PATH          { $$=$1+$2; AddParametro(); }
    |LST_REN PARAMETRO_NAME         { $$=$1+$2; AddParametro(); }
    |PARAMETRO_PATH                 { $$=$1; AddParametro(); }
    |PARAMETRO_NAME                 { $$=$1; AddParametro(); }
;

LST_CAT:    
    LST_CAT PARAMETRO_FILEN          { $$=$1+$2; AddParametro(); }
    |PARAMETRO_FILEN                 { $$=$1; AddParametro(); }
;

LST_REM:    
    LST_REM PARAMETRO_PATH          { $$=$1+$2; AddParametro(); }
    |PARAMETRO_PATH                 { $$=$1; AddParametro(); }
;

LST_MKDIR:
    LST_MKFILE PARAMETRO_PATH       { $$=$1+$2; AddParametro(); }
    |LST_MKFILE PARAMETRO_P         { $$=$1+$2; AddParametro(); }
    |PARAMETRO_PATH                 { $$=$1; AddParametro(); }
    |PARAMETRO_P                    { $$=$1; AddParametro(); }
;

LST_MKFILE:
    LST_MKFILE PARAMETRO_PATH       { $$=$1+$2; AddParametro(); }
    |LST_MKFILE PARAMETRO_SIZE      { $$=$1+$2; AddParametro(); }
    |LST_MKFILE PARAMETRO_P         { $$=$1+$2; AddParametro(); }
    |PARAMETRO_PATH                 { $$=$1; AddParametro(); }
    |PARAMETRO_SIZE                 { $$=$1; AddParametro(); }
    |PARAMETRO_P                    { $$=$1; AddParametro(); }
;

LST_MKUSR:
    LST_MKUSR PARAMETRO_USR     { $$=$1+$2; AddParametro(); }
    |LST_MKUSR PARAMETRO_PWD    { $$=$1+$2; AddParametro(); }
    |LST_MKUSR PARAMETRO_ID     { $$=$1+$2; AddParametro(); }
    |PARAMETRO_USR              { $$=$1; AddParametro(); }
    |PARAMETRO_PWD              { $$=$1; AddParametro(); }
    |PARAMETRO_ID               { $$=$1; AddParametro(); }
;

LST_MKGRP:
    LST_MKGRP PARAMETRO_GRP     { $$=$1+$2; AddParametro(); }
    |PARAMETRO_GRP              { $$=$1; AddParametro(); }
;

LST_LOGIN:
    LST_LOGIN PARAMETRO_USR         { $$=$1+$2; AddParametro(); }
    |LST_LOGIN PARAMETRO_PWD        { $$=$1+$2; AddParametro(); }
    |LST_LOGIN PARAMETRO_ID         { $$=$1+$2; AddParametro(); }
    |PARAMETRO_USR                  { $$=$1; AddParametro(); }
    |PARAMETRO_PWD                  { $$=$1; AddParametro(); }
    |PARAMETRO_ID                   { $$=$1; AddParametro(); }
;

LST_MKFS:
    LST_MKFS PARAMETRO_ID               { $$=$1+$2; AddParametro(); }
    |LST_MKFS PARAMETRO_TYPE_FS         { $$=$1+$2; AddParametro(); }
    |PARAMETRO_ID                       { $$=$1; AddParametro(); }
    |PARAMETRO_TYPE_FS                  { $$=$1; AddParametro(); }
;


LST_REP:
    LST_REP PARAMETRO_NAME          { $$=$1+$2; AddParametro(); }
    |LST_REP PARAMETRO_PATH         { $$=$1+$2; AddParametro(); }
    |LST_REP PARAMETRO_ID           { $$=$1+$2; AddParametro(); }
    |LST_REP PARAMETRO_RUTA         { $$=$1+$2; AddParametro(); }
    |PARAMETRO_NAME                 { $$=$1; AddParametro(); }
    |PARAMETRO_PATH                 { $$=$1; AddParametro(); }
    |PARAMETRO_ID                   { $$=$1; AddParametro(); }
    |PARAMETRO_RUTA                 { $$=$1; AddParametro(); }
;

LST_MOUNT:
    LST_MOUNT PARAMETRO_PATH    { $$=$1+$2; AddParametro(); }
    |LST_MOUNT PARAMETRO_NAME   { $$=$1+$2; AddParametro(); }
    |PARAMETRO_PATH             { $$=$1; AddParametro(); }
    |PARAMETRO_NAME             { $$=$1; AddParametro(); }
;

LST_FDISK:
    LST_FDISK PARAMETRO_SIZE        { $$=$1+$2; AddParametro(); }
    |LST_FDISK PARAMETRO_UNIT       { $$=$1+$2; AddParametro(); }
    |LST_FDISK PARAMETRO_PATH       { $$=$1+$2; AddParametro(); }
    |LST_FDISK PARAMETRO_TYPE       { $$=$1+$2; AddParametro(); }
    |LST_FDISK PARAMETRO_FIT        { $$=$1+$2; AddParametro(); }
    |LST_FDISK PARAMETRO_DELETE     { $$=$1+$2; AddParametro(); }
    |LST_FDISK PARAMETRO_NAME       { $$=$1+$2; AddParametro(); }
    |LST_FDISK PARAMETRO_ADD        { $$=$1+$2; AddParametro(); }
    |PARAMETRO_SIZE                 { $$=$1; AddParametro(); }
    |PARAMETRO_UNIT                 { $$=$1; AddParametro(); }
    |PARAMETRO_PATH                 { $$=$1; AddParametro(); }
    |PARAMETRO_TYPE                 { $$=$1; AddParametro(); }
    |PARAMETRO_FIT                  { $$=$1; AddParametro(); }
    |PARAMETRO_DELETE               { $$=$1; AddParametro(); }
    |PARAMETRO_NAME                 { $$=$1; AddParametro(); }
    |PARAMETRO_ADD                  { $$=$1; AddParametro(); }
;

LST_MKDIS:
    LST_MKDIS PARAMETRO_SIZE            { $$=$1; AddParametro();}
    |LST_MKDIS PARAMETRO_PATH           { $$=$1; AddParametro();}
    //|LST_MKDIS PARAMETRO_NAME         { $$=$1; AddParametro();}
    |LST_MKDIS PARAMETRO_UNIT           { $$=$1; AddParametro();}
    |LST_MKDIS PARAMETRO_FIT            { $$=$1; AddParametro();}
    |PARAMETRO_SIZE                     { $$=$1; AddParametro();}
    |PARAMETRO_PATH                     { $$=$1; AddParametro();}
    //|PARAMETRO_NAME                   { $$=$1; AddParametro();}
    |PARAMETRO_UNIT                     { $$=$1; AddParametro();}
    |PARAMETRO_FIT                      { $$=$1; AddParametro();}
;

PARAMETRO_GRP:
    guion name asignacion rutaCompleja { $$=$1+$2+$3+$4; CrearParametro($2,$4);}
;

PARAMETRO_FILEN:
    guion filen asignacion VALOR_PATH { $$=$1+$2+$3+$4; CrearParametro($2,$4);}
;

PARAMETRO_PATH:
    guion path asignacion VALOR_PATH { $$=$1+$2+$3+$4; auxPath = $4; CrearParametro($2,$4);}
;

PARAMETRO_RUTA:
    guion rruta asignacion VALOR_PATH { $$=$1+$2+$3+$4; auxPath = $4; CrearParametro($2,$4);}
;

VALOR_PATH:
    rutaSimple      { $$=$1 }
    |rutaCompleja   { $$=$1 }
;

PARAMETRO_SIZE:
    guion size asignacion numero { $$=$1+$2+$3+$4; CrearParametro($2,$4);}
;

PARAMETRO_NAME:
    guion name asignacion VALOR_NAME { $$=$1+$2+$3+$4; CrearParametro($2,$4);}
;

VALOR_NAME:
    archivo         { $$=$1 }
    |id             { $$=$1 }
    |rutaCompleja   { $$=$1 }
;

PARAMETRO_USR:
    guion usr asignacion id { $$=$1+$2+$3+$4; CrearParametro($2,$4);}
;

PARAMETRO_PWD:
    guion pwd asignacion  id        { $$=$1+$2+$3+$4; CrearParametro($2,$4);}
    |guion pwd asignacion numero    { $$=$1+$2+$3+$4; CrearParametro($2,$4);}
;

PARAMETRO_USR:
    guion usr asignacion id { $$=$1+$2+$3+$4; CrearParametro($2,$4);}
;

PARAMETRO_UNIT:
    guion unit asignacion VALOR_UNIT { $$=$1+$2+$3+$4; CrearParametro($2,$4);}
;

VALOR_UNIT:
    k       { $$=$1 }
    |m      { $$=$1 }
;

PARAMETRO_TYPE:
    guion rtype asignacion VALOR_TYPE { $$=$1+$2+$3+$4; CrearParametro($2,$4);}
;

VALOR_TYPE:
    p       { $$=$1 }
    |e      { $$=$1 }
    |l      { $$=$1 }
;

PARAMETRO_FIT:
    guion fit asignacion VALOR_FIT { $$=$1+$2+$3+$4; CrearParametro($2,$4); }
;

PARAMETRO_P:
    guion p { $$=$1+$2; CrearParametro($2,$2); }
;

VALOR_FIT:
    bf       { $$=$1 }
    |ff      { $$=$1 }
    |wf      { $$=$1 }
;

PARAMETRO_DELETE:
    guion delete asignacion VALOR_DELETE { $$=$1+$2+$3+$4; CrearParametro($2,$4); }
;

PARAMETRO_TYPE_FS:
    guion rtype asignacion VALOR_DELETE { $$=$1+$2+$3+$4; CrearParametro($2,$4); }
;

VALOR_DELETE:
    full    { $$=$1 }
    |fast   { $$=$1 }
;

PARAMETRO_ADD:
    guion add asignacion numero                 { $$=$1+$2+$3+$4; CrearParametro($2,$4); }
    |guion add asignacion guion numero          { $$=$1+$2+$3+$4+$5; CrearParametro($2,"-"+$5); }
    
;

PARAMETRO_IDN:
    PARAMETRO_IDN guion idn asignacion id   { $$=$1+$2+$3+$4+$5 }
    |guion idn asignacion id                { $$=$1+$2+$3+$4 }
;

PARAMETRO_ID:
    PARAMETRO_ID guion rid asignacion id    { $$=$1+$2+$3+$4+$5; CrearParametro($3,$5); }
    |guion rid asignacion id                { $$=$1+$2+$3+$4; CrearParametro($2,$4); }
;


%%

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
func AddInstruccion( tipo string) {
    oInstruccion.Tipo = tipo
	lInstruccion = append(lInstruccion, oInstruccion)
    oInstruccion = arbol.Instruccion{}
}

// AddAST añade lista de instrucciones para ejecutar
func AddAST()  {
	ast:= arbol.AST{}
	ast.Instrucciones = lInstruccion
    lInstruccion = []arbol.Instruccion{}
	lAST = append(lAST, ast)
}

func EjecutarExec () {
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