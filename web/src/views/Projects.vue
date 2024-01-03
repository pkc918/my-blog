<script setup lang="ts">

import Github from "@/icons/Github.vue";
</script>

<template>
<div class="w-full flex flex-col items-center">
  <div class="flex">
    <div class="text-7"><Github/></div>
    <h2>参与的开源项目</h2>
  </div>
  <div class="max-w-300 w-full pt-10">
    <ul class="w-full flex flex-wrap">
      <li class="w-[33.33%] box-border px-2" v-for="item in 4" :key="item"
          slide-enter :style="{ '--stagger': item + 1 }"
      >
        <h2>IKun UI</h2>
        <p>🐔A Svelte.js based UnoCSS UI library that allows you to make websites</p>
      </li>
    </ul>
  </div>
</div>
</template>

<style scoped>

</style>
