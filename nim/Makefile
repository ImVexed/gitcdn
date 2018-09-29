all : dep build
run :
	nim c -d:ssl -r gitcdn.nim
build :
	nim c -d:ssl -d:release gitcdn.nim
dep :
	nimble -y install jester 