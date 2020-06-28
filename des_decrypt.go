package main

import (
    "fmt"
    "io/ioutil"
    "os"
    "regexp"
    "encoding/base64"
    "github.com/wumansgy/goEncrypt"
)

var key = []byte("jeH3O1VX")
var iv = []byte("nHnsU4cX")

func read_file(file_path string) string {
    // 打开文件, 读取文本并返回
    fileObj, err := os.Open(file_path)
    if err != nil {
        panic(err)
    }
    defer fileObj.Close()
    fd, err := ioutil.ReadAll(fileObj)
    return string(fd)
}

func write_file(bytes []byte, file_path string) {
    // 将bytes写入文件
    fileObj, err := os.OpenFile(file_path, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 666)
    defer fileObj.Close()
    if err != nil {
        panic(err)
    }
    fileObj.Write(bytes)
}

func main() {
    // 参数解析
    if len(os.Args) != 3 {
        fmt.Printf("Usage: %s input_file output_file\n", os.Args[0])
        os.Exit(1)
    }
    input_file := os.Args[1]
    output_file := os.Args[2]

    encrypted_stockstr := read_file(input_file)

    // 使用正则表达式移除空白字符
    reg := regexp.MustCompile("\\s")
    encrypted_b64str := reg.ReplaceAllString(encrypted_stockstr, "")

    // 译码b64 & 解密
    encrypted_bytes, _ := base64.StdEncoding.DecodeString(encrypted_b64str)
    decrypted_stockstr, err := goEncrypt.DesCbcDecrypt(encrypted_bytes, key, iv...)
    if err != nil {
        panic(err)
    }
    decrypted_str := string(decrypted_stockstr)

    // 移除头部字符"data:image/jpg;base64,"
    decrypted_b64str := decrypted_str[22:]

    // 译码b64 & 写入文件
    decrypted_bytes, _ := base64.StdEncoding.DecodeString(decrypted_b64str)
    write_file(decrypted_bytes, output_file)
}
