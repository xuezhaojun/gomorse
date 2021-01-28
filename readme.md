# Morse_Code 莫斯码转换

# Getting Started
## Installing
To start using morseCode, install go and run `go get`:
``` shell
$ go get -u github.com/xuezhaojun/morse_code
```
This will retrieve the library.

## Text to morse code

``` go
    text := "hello world"
    err := morse.Valid(text)
    if err != nil {
        fmt.Println(err)
        return
    }

    morseCode := gomorse.ToMorse(text)
```

## Morse code to text

``` go
    morseCode := gomorse.FromMorse(text)
```

## Frontend 

You can visit [here]() to 

## Public API

We offer public api to provide translation operations too.
