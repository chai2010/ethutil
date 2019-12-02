# 以太坊荣耀地址生成器

下载安装:

```
$ go get github.com/chai2010/ethutil/cmd/ethutil-vanity-gen
```

查看帮助:

```
$ ethutil-vanity-gen
Usage of ethutil-vanity-gen:
  -n int
        生成几个地址 (default 1)
  -p string
        地址前缀, 必须是十六进制格式 ([0-9a-f]*)
  -s string
        地址后缀, 必须是十六进制格式 ([0-9a-f]*)
```

随机生成1个地址:

```
$ ethutil-vanity-gen 
0 0xdc8f886f0d002f99b35062809a1305096b55db25 f44a0541a89272c22f996aa9f07d431f...
```

随机生成3个地址:

```
$ ethutil-vanity-gen -n=3

0 0xfd0c05647bc8ef3ec7f6a9d5de9141d6e268fe57 772b7e71d291d3ba01fc8d7b09f6bcf9...
1 0x8ac4c30d1cb8abd2bc5de7780b53bf49c19ab468 fb5365f4f3d3a89c1c43b8802b6e6c61...
2 0x3c5f928683c22504385bd66d4797a46055289ebb 0fb1d6a356758e43bf41f2b09a6fbad4...
```

生成2个以 `ab` 开头的地址:

```
$ ethutil-vanity-gen -p=ab -n=2
0 0xab5e3bae6d17af2619c9984e8d49b77ef237111f fb80a1818249789e696f5c400a359210...
1 0xab7bb7361d373188251a54405cf38075b96f23fd 4518dffcf27e185c8d32c653d4f11573...
```

生成1个以`ab`开头以`c`结尾的地址:

```
$ ethutil-vanity-gen -p=a -s=bc
0 0xab0caf3f7d6041ab91f7988791d217df83fa049c e955552c0d6ef9680ca0fb84fc356ca9...
```

正则模式:

```
$ ethutil-vanity-gen -re="^a\d.*"
```

## `abcdef`组成的单词表(37个)

http://www.wordaxis.com/advanced-anagram/words-within

```
ab
abed
ace
aced
ad
ae
ba
bad
bade
be
bead
bed
cab
cad
cade
cafe
da
dab
dace
dae
de
deaf
deb
decaf
def
ea
ecad
ed
ef
fa
fab
face
faced
fad
fade
fe
fed
```

数字中和`12590`可对应`lzsqo`扩展的单词表, 可以对应454个单词([lzsqo.md](lzsqo.md)).

