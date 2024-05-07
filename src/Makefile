# 注意: 缩进用制表符tab, 不要用空格

.PHONY: clean

OBJ=main.go
TARGET=output/k8s-demo

$(TARGET):$(OBJ)
	mkdir output
	go build -a -installsuffix cgo -o output/k8s-demo
	cp book.yaml output

clean:
	rm -fr output
