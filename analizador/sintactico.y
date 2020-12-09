
%{
// +build ignore
package main

import (
    //"bufio"
    "fmt"
    //"archivos/MIA-Proyecto1_201709454/analizador/arbol"
    "proyectos/MIAProyectoFase1/analizador/arbol"
    //"os"
)

var oParametro arbol.Parametro
var oInstruccion arbol.Instruccion
var lInstruccion []arbol.Instruccion
var lAST []arbol.AST

%}

%union{
    value string
}

%token	asignacion, guion
%token  exec, pause, mkdisk, rmdisk, fdisk, mount, unmount, mkfs
%token  login, logout, mkgrp, rmgrp, mkusr, rmusr,chmod
%token  mkfile, cat,rm, edit, ren, mkdir, cp, mv, find, chown
%token  chgrp, rep, path, size, name, unit, rtype, fit, delete
%token  add, idn, id, tipo, usr, pwd, grp, ugo, r, p, cont
%token  filen, rf, dest, k, m, e, l, bf,ff, wf, full, fast
%token  numero, rutaSimple, rutaCompleja, archivo

%type <value> INICIO
%type <value> asignacion, guion
%type <value> exec, pause, mkdisk, rmdisk, fdisk, mount, unmount, mkfs
%type <value> login, logout, mkgrp, rmgrp, mkusr, rmusr,chmod
%type <value> mkfile, cat,rm, edit, ren, mkdir, cp, mv, find, chown
%type <value> chgrp, rep, path, size, name, unit, rtype, fit, delete
%type <value> add, idn, id, tipo, usr, pwd, grp, ugo, r, p, cont
%type <value> filen, rf, dest, k, m, e, l, bf,ff,wf, full, fast
%type <value>  numero, rutaSimple, rutaCompleja, archivo
%type <value> INSTRUCCION, LISTA_INSTRUCCION, PARAMETRO_PATH, VALOR_PATH
%type <value> PARAMETRO_SIZE, PARAMETRO_NAME, VALOR_NAME, PARAMETRO_UNIT, VALOR_UNIT 
%type <value> PARAMETRO_TYPE, VALOR_TYPE, PARAMETRO_FIT, VALOR_FIT
%type <value> PARAMETRO_DELETE, VALOR_DELETE, PARAMETRO_ADD, PARAMETRO_IDN
%type <value> LST_FDISK, LST_MKDIS, LST_MOUNT


%start INICIO

%% /* The grammar follows.  */

INICIO:
    LISTA_INSTRUCCION   { AddAST(); fmt.Println("AST -> ", lAST);}
;

LISTA_INSTRUCCION:
    LISTA_INSTRUCCION INSTRUCCION '\n'
    |INSTRUCCION '\n'
;

INSTRUCCION:
    pause                   {fmt.Println("pausa->",$1); AddInstruccion("pause"); }
    |exec PARAMETRO_PATH    {fmt.Println("INSTRUCCION->",$1,$2); AddParametro(); AddInstruccion("exec"); }
    |mkdisk LST_MKDIS       {fmt.Println("INSTRUCCION->",$1,$2); AddInstruccion("mkdisk"); }
    |rmdisk PARAMETRO_PATH  {fmt.Println("INSTRUCCION->",$1,$2); AddParametro(); AddInstruccion("rmdisk"); }
    |fdisk LST_FDISK        {fmt.Println("INSTRUCCION->",$1,$2); AddInstruccion("fdisk"); }
    |mount LST_MOUNT        {fmt.Println("INSTRUCCION->",$1,$2); AddInstruccion("mount"); }
    |unmount PARAMETRO_IDN  {fmt.Println("INSTRUCCION->",$1,$2); AddInstruccion("unmount"); }
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
    |LST_MKDIS PARAMETRO_NAME           { $$=$1; AddParametro();}
    |LST_MKDIS PARAMETRO_UNIT           { $$=$1; AddParametro();}
    |LST_MKDIS PARAMETRO_FIT            { $$=$1; AddParametro();}
    |PARAMETRO_SIZE                     { $$=$1; AddParametro();}
    |PARAMETRO_PATH                     { $$=$1; AddParametro();}
    |PARAMETRO_NAME                     { $$=$1; AddParametro();}
    |PARAMETRO_UNIT                     { $$=$1; AddParametro();}
    |PARAMETRO_FIT                      { $$=$1; AddParametro();}
;



PARAMETRO_PATH:
    guion path asignacion VALOR_PATH { $$=$1+$2+$3+$4; CrearParametro($2,$4);}
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

VALOR_FIT:
    bf       { $$=$1 }
    |ff      { $$=$1 }
    |wf      { $$=$1 }
;

PARAMETRO_DELETE:
    guion delete asignacion VALOR_DELETE { $$=$1+$2+$3+$4; CrearParametro($2,$4); }
;

VALOR_DELETE:
    full    { $$=$1 }
    |fast   { $$=$1 }
;

PARAMETRO_ADD:
    guion add asignacion numero         { $$=$1+$2+$3+$4; CrearParametro($2,$4); }
    
;

PARAMETRO_IDN:
    PARAMETRO_IDN guion idn asignacion id   { $$=$1+$2+$3+$4+$5 }
    |guion idn asignacion id                { $$=$1+$2+$3+$4 }
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
	lAST = append(lAST, ast)
}