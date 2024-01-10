<script setup lang="ts">

import Markdown from "@/components/Markdown.vue";
const str = `
>我像麋鹿一样在林荫中走着，为着自己的香气而发狂；夜晚是五月正中的夜晚，清风是南国的清风； 我迷路了，我游荡着，我寻求那得不到的东西，我得到了我所没有寻求的东西。 —— 泰戈尔
## 哈哈哈
### 前言

每周四，在群里总是会看到一些疯狂星期四的小故事，而且个个的小故事还挺有趣，但是让我这样的人肯定是编不出来了，在想刚好在学习Go，是不是可以怎么操作一下，写个命令行来获取周四文案啥的，这样就不用自己去编了。既然有了想法那就开干

### 准备

[源码地址](https://github.com/GoGadgetHub/crazy-Thursday/tree/main)，结合看最好，下面代码并不完全，大佬们有好的想法，欢迎issues，pr，也欢迎star，一起整活

1. 了解使用Go写命令行的库
2. 找一个生成疯狂星期四文案的接口（后期也可以做自动生成）
3. 通过特定的命令来调用这个接口并拿到这个文案

在这里我了解到一个用Go做命令行的流行库[cobra](https://github.com/spf13/cobra)，据官方说这个被用在了很多大型项目，像\`Kubernetes, Hugo, GitHub CLI\`等等

找到的免费生成疯狂星期四文案的接口\`https://api.jixs.cc/api/wenan-fkxqs/index.php\`

### 实现

首先量化需求，既然已接口的形式的话，其实无非就是调用接口，拿到返回值了，只是把这个步骤放在了用户命令有对应操作之后执行，明确需求了，就可以开始写代码了

1. 使用Go调用这个API，拿到返回值并返回

\`\`\`Go
package cmd

import (
\t"io"
\t"log"
\t"net/http"
)

// Fetch 发起请求
func Fetch() string {
\tclient := &http.Client{}
\treq, _ := http.NewRequest("GET", "https://api.jixs.cc/api/wenan-fkxqs/index.php", nil)
\tres, err := client.Do(req)
\tif err != nil {
\t\tlog.Fatal("Http get error is ", err)
\t}
\tif res.StatusCode != http.StatusOK {
\t\tlog.Fatal("Http status code is ", res.StatusCode)
\t}
\tdefer func(Body io.ReadCloser) {
\t\terr := Body.Close()
\t\tif err != nil {
\t\t\tlog.Fatal(err)
\t\t}
\t}(res.Body)
\tbytes, err := io.ReadAll(res.Body)
\tif err != nil {
\t\treturn ""
\t}
\treturn string(bytes)
}
\`\`\`
2. 使用Cobra做一个特定命令行

\`\`\`Go
package cmd

import (
\t"fmt"
\t"github.com/spf13/cobra"
\t"os"
)

var ThursdayCmd = &cobra.Command{
\tUse:   "Thursday [文案]",
\tShort: "crazy Thursday",
\tLong:  \`crazy Thursday\`,
\tRun: func(cmd *cobra.Command, args []string) {
\t\tres := Fetch()
\t\tPrint(res)
\t\tfmt.Println(res)
\t},
}

var rootCmd = &cobra.Command{
\tUse: "ct",
}

func Execute() {
\trootCmd.AddCommand(ThursdayCmd)
\trootCmd.SetHelpCommand(ThursdayCmd)
\tif err := rootCmd.Execute(); err != nil {
\t\tfmt.Println(err)
\t\tos.Exit(1)
\t}
}
\`\`\`
这样当我们执行 \`go run main.go Thursday\` 后会在控制台打印出获取到的文案

但是我使用 GoLand 在这一步有点问题，就是当文案特别长的时候，他在控制台还是会强行显示一行，无法复制到全部，所以这时想起一种解决方法，那就是将文案存入一个txt文件中，这样就可以啦。

3. 生成文件夹，将结果存入

\`\`\`Go
package cmd

import (
\t"fmt"
\t"log"
\t"os"
)

func Print(res string) {
\t// 创建文件
\tfile, err := os.Create("./output/output.txt")
\tif err != nil {
\t\tfmt.Println("无法创建文件:", err)
\t\treturn
\t}
\tdefer func(file *os.File) {
\t\terr := file.Close()
\t\tif err != nil {
\t\t\tlog.Fatal(err)
\t\t}
\t}(file)

\t// 将内容写入文件
\t_, err = file.WriteString(res)
\tif err != nil {
\t\tfmt.Println("无法写入文件:", err)
\t\treturn
\t}

\tfmt.Println("内容已成功写入文件 output.txt")
}
\`\`\`
到此基本上就完成了，算是一个练习的小demo吧，可以拿着这个去和群友欢乐对线吧！

#### 效果
![Thursday.png](https://p3-juejin.byteimg.com/tos-cn-i-k3u1fbpfcp/9160242b455541f7b32bc495cdc2dc0c~tplv-k3u1fbpfcp-jj-mark:0:0:0:0:q75.image#?w=2160&h=634&s=147475&e=png&b=2c2a2e)
`
</script>

<template>
<!--<div>Music</div>-->
  <Markdown :markdown-text="str"/>
</template>

<style scoped>

</style>
