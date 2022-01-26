#!/bin/bash

root_path=`pwd`
echo ${root_path}

cd ${root_path}/service/user/rpc
nohup go run user.go -f etc/user.yaml &
cd ${root_path}/service/user/api
nohup go run user.go -f etc/user.yaml &

cd ${root_path}/service/product/rpc
nohup go run product.go -f etc/product.yaml &
cd ${root_path}/service/product/api
nohup go run product.go -f etc/product.yaml &

cd ${root_path}/service/order/rpc
nohup go run order.go -f etc/order.yaml &
cd ${root_path}/service/order/api
nohup go run order.go -f etc/order.yaml &

cd ${root_path}/service/pay/rpc
nohup go run pay.go -f etc/pay.yaml &
cd ${root_path}/service/pay/api
nohup go run pay.go -f etc/pay.yaml &
