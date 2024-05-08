# 注意: 缩进用制表符tab, 不要用空格
# 构建: go build -a -installsuffix cgo -o output/k8s-demo

.PHONY: clean

OBJ=main.go
TARGET=output/k8s-demo

$(TARGET):$(OBJ)
	go build -o $(TARGET)
	cp book.yaml output

clean:
	rm -fr output
