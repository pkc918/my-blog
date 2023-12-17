<script setup lang="ts">
import Typed from "typed.js";
import { nextTick, onUnmounted, ref } from "vue";
import FoundMe from "@/components/FoundMe.vue";

const excerpt = ref();
const typed = ref<Typed>();

nextTick(() => {
  typed.value = new Typed(excerpt.value, {
    strings: [
      `
        <span class="text-10 font-500">Hi, I'm pkc918👋。</span></br>
        <span class="text-5 font-500">A Full Stack Developer, I love Coding.</span>
      `,
    ],
    typeSpeed: 40,
  });
});
onUnmounted(() => {
  typed.value?.destroy();
});
</script>

<template>
  <div class="w-full h-full flex-col fcc caret-transparent relative" slide-enter>
    <img
      class="w-70 h-70 rd-full"
      src="https://avatars.githubusercontent.com/u/58922004?v=4"
      alt="avatar"
    />
    <div class="h-50 mt-5">
      <span ref="excerpt"></span>
    </div>
    <FoundMe class="absolute right-10 top-50% translate-y--90%" />
  </div>
</template>

<style scoped></style>
