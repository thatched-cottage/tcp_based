echo "开始启动服务器"
nname=1
echo "启动 Node $nname"
go run node/node.go --nodename=$nname > node_$nname.log 2>&1 &
nname=2
echo "启动 Node $nname"
go run node/node.go --nodename=$nname > node_$nname.log 2>&1 &
echo "启动 Central node"
go run central_node/central_node.go > central_node.log 2>&1 &
echo "启动完成"

