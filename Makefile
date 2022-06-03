.PHONY: build cleanO clean run runDefault

all: build

build: 
	@go build -o build/coder/coder cmd/coderMain/coderMain.go

buildSC:
	@go build -o  build/statsChecker/statsChecker cmd/statCheckerMain/statCheckerMain.go

run:
	@./build/coder/coder $(IN) $(OUT) $(NUM)

runDefault:
	@./build/coder/coder data/input/testy4/example0.tga data/output/def 2

runDefaultSC:
	@./build/statsChecker/statsChecker

runTest: 
	@./build/coder/coder

cleanO: 
	@rm data/output/*

clean:
	@rm build/coder/*

