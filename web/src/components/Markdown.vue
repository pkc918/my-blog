<script setup lang="ts">
import { ref, nextTick, computed } from "vue";
import { markedHighlight } from "marked-highlight";
import { Marked } from "marked";
import hljs from 'highlight.js';
import "@/themes/z-blue.scss";

const props = defineProps<{
  markdownText: string
}>()

const minLevel = ref(6);
// 侧边栏
const tocs = ref<{
  id: number;
  html: string;
  text: string;
  level: number
}[]>([]);

const marked: Marked = new Marked(
    markedHighlight({
      langPrefix: "hljs language-",
      highlight(code, lang, _) {
        const language = hljs.getLanguage(lang) ? lang : "plaintext";
        return hljs.highlight(code, {language}).value;
      }
    })
);

marked.use({
  renderer: {
    heading: (text, level) => {
      // const escapedText = text.toLowerCase().replace(/[^\w]+/g, "-");
      anchor.value++;
      tocs.value?.push({
        id: anchor.value,
        html: `<h${level}>${text}</h${level}>`,
        text: text,
        level: level
      });
      minLevel.value = minLevel.value > level ? level : minLevel.value;
      console.log(tocs.value);
      return `
            <h${level} id="toc-nav-${anchor.value}">
              ${text}
            </h${level}>`;
    }
  }
});

const anchor = ref(0);

// 用于存放最终解析出来的dom
const html = ref<string>("");

// 解析markdown
const handleMarkdown = async () => {
  // 判断markdown为空不解析
  if (!props.markdownText) {
    return
  }
  // 解析markdown获取HTML
  html.value = await marked.parse(props.markdownText);
  // console.log(str);
}

const toTarget = (id: number) => {
  console.log(id);
};

const markdownContent = ref<HTMLDivElement>();
nextTick(() => {
  handleMarkdown()
  console.log(markdownContent.value?.getBoundingClientRect());
})


const test = computed(() => {
  return 1;
});
console.log(test.effect);


/*const str = ref("");
effect(() => {
  alert(str.value);
});

setTimeout(() => {
  str.value = "123";
}, 5000);*/
</script>

<template>
  <div class="flex gap-4 w-full max-w-6xl 2xl:max-w-8xl relative">
    <div class="flex-1 w-full bg-[rgba(250,250,254,1)]] rd-1 highlight markdown-body" ref="markdownContent">
      <h1 class="!mt-0">标题</h1>
      <div v-html="html"></div>
    </div>
    <div class="w-full max-w-[300px] relative">
      <ul class="w-full min-h-[300px] max-h-[400px] max-w-[300px] rd-1 bg-red fixed">
        <li v-for="item in tocs" :key="item.id" class="m-0 cursor-pointer" @click="toTarget(item.id)"
            :style="{'padding-left': `${(item.level - minLevel)}em`}">
          <a :href="`#toc-nav-${item.id}`">{{item.text}}</a>
        </li>
      </ul>
    </div>


  </div>
</template>

