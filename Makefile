compile:
	goyacc -o sintactico.go analizador/sintactico.y
	golex -o lexico.go analizador/lexico.l

build:
	gofmt -l -s -w *.go
	go build main.go sintactico.go lexico.go

run:
	rm -rf /home/mia
	rm -rf /home/archivos
	./main

cbrun: cleanconsole clean compile build run

brun: cleanconsole build run

rrun: cleanconsole run

clean:
	go clean
	rm -f lexico.go sintactico.go y.output sintactico lexico main

cleanconsole:
	clear