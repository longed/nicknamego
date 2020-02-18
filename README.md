# nicknamego
pick up a nickname or password for you, write by Golang.

## build & run

###### ubuntu

```bash
go build main.go

./main 
```

## features
- generate nickname or password
- specify random characters
- support batch generation
- save result to file

## configuration

```toml
[UserOptions]
# 是否要数字
wantNumber       = true
# 是否要大写字母
wantUpperCase    = true
# 是否要小写字母
wantLowerCase    = true
# 是否要符号
wantSymbol       = true
# 是否保存到文件
saveNickNameToFile = true
# nickname 长度
nickNameLen = 8
# 批量产生 nickname 数
batchNumber = 1
# 指定随机字符集
specifiedChars = ""
```

