# 杂项

## 测试环境(testnet)

启动测试环境(本地IPC协议):

```
$ geth -testnet -syncmode=light -ipcpath=geth.ipc
```

新开终端连接到服务(MacOS):

```
$ geth attach ${HOME}/Library/Ethereum/testnet/geth.ipc
```

测试本地rpc接口(打开各种权限)(http协议):

```
$ geth -testnet -syncmode=light -rpc --allow-insecure-unlock --rpcapi="eth,admin,miner,db,vns,net,web3,personal,web3"
```

注意: 默认只能 http://localhost:8545 连接rpc, 如果打开其它地址限制最好关闭 http 环境到 unlock 功能.

连接本地rpc:

```
$ geth attach http://localhost:8545
```

输入 `exit` 退出.

