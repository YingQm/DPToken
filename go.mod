module github.com/dplatform/dplatform

go 1.13

replace github.com/33cn/dplatform => ../dplatform

replace github.com/33cn/plugin => ../plugin

require (
	github.com/33cn/chain33 v1.65.0 // indirect
	github.com/33cn/dplatform v1.65.1-0.20201111090506-e2d7705ea453
	github.com/33cn/plugin v0.0.0-20200401090725-74f0baacfc4e
	github.com/prometheus/client_golang v1.1.0 // indirect
	github.com/stretchr/objx v0.2.0 // indirect
	github.com/vektra/mockery v0.0.0-20181123154057-e78b021dcbb5 // indirect
)
