<script setup lang="ts">
import { watchEffect, ref, nextTick } from "vue";
import MarkdownIt from 'markdown-it'
import hljs from 'highlight.js';
import 'highlight.js/styles/atom-one-light.css'
import "github-markdown-css"
// import markdownItHighlightjs from 'markdown-it-highlightjs'
// import markdownItCodeCopy from 'markdown-it-code-copy'
hljs.highlightAll()
const props = defineProps<{
  markdownText: string
}>()

// 用于存放最终解析出来的dom
const markdownDom = ref<HTMLDivElement>()

// 初始化 markdown-it 实例
const md = new MarkdownIt()
// 配置代码高亮插件
// md.use(markdownItHighlightjs)
// // 配置代码块复制插件
// md.use(markdownItCodeCopy)

// 解析markdown
const handleMarkdown = () => {
  // 判断markdown为空不解析
  if (!props.markdownText) {
    return
  }

  // 解析markdown获取HTML
  const html = md.render(props.markdownText)
  console.log(html)
  markdownDom.value!.innerHTML = html
}

nextTick(() => {
  handleMarkdown()
})

</script>

<template>
<div class="markdown-body highlight max-w-5xl 2xl:max-w-6xl mx-auto" ref="markdownDom"></div>
</template>

<style scoped>

</style>
