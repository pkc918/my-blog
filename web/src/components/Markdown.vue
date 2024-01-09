<script setup lang="ts">
import { ref, nextTick } from "vue";
import { markedHighlight } from "marked-highlight";
import { Marked } from "marked";
import hljs from 'highlight.js';
import "@/themes/z-blue.scss";

const props = defineProps<{
  markdownText: string
}>()

const marked = new Marked(
    markedHighlight({
      langPrefix: "hljs language-",
      highlight(code, lang, _) {
        const language = hljs.getLanguage(lang) ? lang : "plaintext";
        return hljs.highlight(code, {language}).value;
      }
    })
);

// 用于存放最终解析出来的dom
const html = ref<string>("");

// 解析markdown
const handleMarkdown = async () => {
  // 判断markdown为空不解析
  if (!props.markdownText) {
    return
  }
  // 解析markdown获取HTML
  const str = await marked.parse(props.markdownText);
  console.log(html)
  html.value = str;
}

nextTick(() => {
  handleMarkdown()
})

</script>

<template>
  <div class="content">
    <div class="max-w-5xl 2xl:max-w-6xl mx-auto highlight markdown-body" v-html="html"></div>
  </div>
</template>

