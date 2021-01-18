# Morse_Code在线莫斯码转换器

本项目为纯golang实现的在线莫斯码转化器

# 需求分析
## 功能性需求
* 翻译用户输入的文字（英文），并转化为莫斯电码
* 翻译用户输入的莫斯电码，并转化为中文
（限制单词翻译的内容发送的内容不超过1KB）
* 对莫斯码“点”，“斜”的显示，做特殊化处理
    - 可以前端实现

## 非功能性需求
* read-heavy 

## 规模需求
目前小站点按照日10000访问估算：
* 平均每条消息估算为500bytes
* 每秒访问量：10000/(3600*24) ~= 0.12 request/s
* 每秒需要的带宽：0.12 * 500bytes ~= 50 byte/s

# API
```
// 输入：英文可读文本
// 输出: morse code
ToMorse(text string) string

// 输入：morse code
// 输出：英文可读文本
FromMorse(morse_code string) string
```

# 依赖
[gin]()

[vugu]()