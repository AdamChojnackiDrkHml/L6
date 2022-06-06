.PHONY: build cleanO clean runDefault buildC buildD buildSC runDefaultC runDefaultD runDefaultSC runC runD runSC

all: build

build: buildC buildD buildSC

buildC: 
	@go build -o build/coder/coder cmd/coderMain/coderMain.go

buildSC:
	@go build -o  build/statsChecker/statsChecker cmd/statCheckerMain/statCheckerMain.go

buildD:
	@go build -o build/decoder/decoder cmd/decoderMain/decoderMain.go

runC:
	@./build/coder/coder $(IN) $(OUT) $(NUM)

runD:
	@./build/decoder/decoder $(IN) $(OUT)

runSC:
	@./build/statsChecker/statsChecker $(IN) $(OUT)

runDefaultC:
	@./build/coder/coder data/input/testy4/example0.tga data/output/def 2

runDefaultSC:
	@./build/statsChecker/statsChecker

runDefaultD:
	@./build/decoder/decoder data/output/def data/results/decoded

runTest: 
	@./build/coder/coder

cleanO: 
	@rm data/output/*

clean:
	@rm build/coder/*

