%{
// blame: jnml, labs.nic.cz
package main

import (
    "bufio"
    "log"
    //"fmt"
    //"strconv"
)

type yylexer struct{
    src     *bufio.Reader
    buf     []byte
    empty   bool
    current byte
}

func newLexer(src *bufio.Reader) (y *yylexer) {
    y = &yylexer{src: src}
    if b, err := src.ReadByte(); err == nil {
        y.current = b
    }
    return
}

func (y *yylexer) getc() byte {
    if y.current != 0 {
        y.buf = append(y.buf, y.current)
    }
    y.current = 0
    if b, err := y.src.ReadByte(); err == nil {
        y.current = b
    }
    return y.current
}

func (y yylexer) Error(e string) {
    log.Fatal(e)
}

func (y *yylexer) Lex(lval *yySymType) int {
    c := y.current
    if y.empty {
        c, y.empty = y.getc(), false
    }
%}

%yyc c
%yyn c = y.getc()

numero              [0-9]+
letra               [a-zñA-ZÑ_]
id                  ({letra}|{numero})+
archivo             {id}(".arch"|".dsk"|".jpg"|".disk"|".png"|".pdf")
rutaSimple          "/"({id}"/")*({archivo}?)
rutaCompleja        \"[^\"\n\t\r]*\"

%%
    y.buf = y.buf[:0]
#[^\n]*
"/*\n"
[ \t\r\n]+
"->"                            lval.value = string(y.buf); return asignacion
"-"                             lval.value = string(y.buf); return guion
[Ee][Xx][Ee][Cc]                lval.value = string(y.buf); return exec
[Pp][Aa][Uu][Ss][Ee]            lval.value = string(y.buf); return pause
[Mm][Kk][Dd][Ii][Ss][Kk]        lval.value = string(y.buf); return mkdisk
[Rr][Mm][Dd][Ii][Ss][Kk]        lval.value = string(y.buf); return rmdisk
[Ff][Dd][Ii][Ss][Kk]            lval.value = string(y.buf); return fdisk
[Mm][Oo][Uu][Nn][Tt]            lval.value = string(y.buf); return mount
[Uu][Nn][Mm][Oo][Uu][Nn][Tt]    lval.value = string(y.buf); return unmount
[Mm][Kk][Ff][Ss]                lval.value = string(y.buf); return mkfs
[Ll][Oo][Gg][Ii][Nn]            lval.value = string(y.buf); return login
[Ll][Oo][Gg][Oo][Uu][Tt]        lval.value = string(y.buf); return logout
[Mm][Kk][Gg][Rr][Pp]            lval.value = string(y.buf); return mkgrp
[Rr][Mm][Gg][Rr][Pp]            lval.value = string(y.buf); return rmgrp
[Mm][Kk][Uu][Ss][Rr]            lval.value = string(y.buf); return mkusr
[Rr][Mm][Uu][Ss][Rr]            lval.value = string(y.buf); return rmusr
[Cc][Hh][Mm][Oo][Dd]            lval.value = string(y.buf); return chmod
[Mm][Kk][Ff][Ii][Ll][Ee]        lval.value = string(y.buf); return mkfile
[Cc][Aa][Tt]                    lval.value = string(y.buf); return cat
[Rr][Mm]                        lval.value = string(y.buf); return rm
[Ee][Dd][Ii][Tt]                lval.value = string(y.buf); return edit
[Rr][Ee][Nn]                    lval.value = string(y.buf); return ren
[Mm][Kk][Dd][Ii][Rr]            lval.value = string(y.buf); return mkdir
[Cc][Pp]                        lval.value = string(y.buf); return cp
[Mm][Vv]                        lval.value = string(y.buf); return mv
[Ff][Ii][Nn][Dd]                lval.value = string(y.buf); return find
[Cc][Hh][Oo][Ww][Nn]            lval.value = string(y.buf); return chown
[Cc][Hh][Gg][Rr][Pp]            lval.value = string(y.buf); return chgrp
[Rr][Ee][Pp]                    lval.value = string(y.buf); return rep
[Pp][Aa][Tt][Hh]            lval.value = string(y.buf); return path
[Ss][Ii][Zz][Ee]            lval.value = string(y.buf); return size
[Nn][Aa][Mm][Ee]            lval.value = string(y.buf); return name
[Uu][Nn][Ii][Tt]            lval.value = string(y.buf); return unit
[Tt][Yy][Pp][Ee]            lval.value = string(y.buf); return rtype
[Ff][Ii][Tt]                lval.value = string(y.buf); return fit
[Dd][Ee][Ll][Ee][Tt][Ee]    lval.value = string(y.buf); return delete
[Aa][Dd][Dd]                lval.value = string(y.buf); return add
[Ii][Dd]{numero}            lval.value = string(y.buf); return idn
[Ii][Dd]                    lval.value = string(y.buf); return rid
[Tt][Ii][Pp][Oo]            lval.value = string(y.buf); return tipo
[Uu][Ss][Rr]                lval.value = string(y.buf); return usr
[Pp][Dd][Ww][Dd]            lval.value = string(y.buf); return pwd
[Gg][Rr][Pp]                lval.value = string(y.buf); return grp
[Uu][Gg][Oo]                lval.value = string(y.buf); return ugo
[Rr]                        lval.value = string(y.buf); return r
[Pp]                        lval.value = string(y.buf); return p
[Cc][Oo][Nn][Tt]            lval.value = string(y.buf); return cont
[Ff][Ii][Ll][Ee]{numero}    lval.value = string(y.buf); return filen
[Rr][Ff]                    lval.value = string(y.buf); return rf
[Dd][Ee][Ss][Tt]            lval.value = string(y.buf); return dest

[Kk]                        lval.value = string(y.buf); return k
[Mm]                        lval.value = string(y.buf); return m
[Ee]                        lval.value = string(y.buf); return e
[Ll]                        lval.value = string(y.buf); return l
[Bb][Ff]                    lval.value = string(y.buf); return bf
[Ff][Ff]                    lval.value = string(y.buf); return ff
[Ww][Ff]                    lval.value = string(y.buf); return wf
[Ff][Aa][Ss][Tt]            lval.value = string(y.buf); return fast
[Ff][Uu][Ll][Ll]            lval.value = string(y.buf); return full


{numero}                    lval.value = string(y.buf); return numero
{id}                        lval.value = string(y.buf); return id
{rutaSimple}                lval.value = string(y.buf); return rutaSimple
{rutaCompleja}              lval.value = string(y.buf); return rutaCompleja
{archivo}                   lval.value = string(y.buf); return archivo

%%
    y.empty = true
    return int(c)
}
