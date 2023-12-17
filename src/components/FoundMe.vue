<script setup lang="ts">

import Github from "@/icons/Github.vue";
import NetEasyMusic from "@/icons/NetEasyMusic.vue";
import Email from "@/icons/Email.vue";
import Twitter from "@/icons/Twitter.vue";
import WeChat from "@/icons/WeChat.vue";
</script>

<template>
  <ul class="fcc flex-col found-me-btn">
    <li class=" bg-[rgb(15,20,25)]">
      <a target="_blank" href="" rel="noreferrer">
        <Github fill="#fff"/>
      </a>
    </li>
    <li class="bg-[rgb(194,12,12)] important:text-10">
      <a target="_blank" href="" rel="noreferrer">
        <NetEasyMusic fill="#fff"/>
      </a>
    </li>
    <li class="bg-[rgb(36,46,54)] important:text-5">
      <a target="_blank" href="" rel="noreferrer">
        <Twitter class="text-[#fff]" fill="#fff"/>
      </a>
    </li>
    <li class="bg-[rgb(212,70,56)] important:text-7">
      <a target="_blank" href="" rel="noreferrer">
        <Email fill="#fff"/>
      </a>
    </li>
    <li class="bg-[rgb(98,185,0)] important:text-7">
      <a target="_blank" href="" rel="noreferrer">
        <WeChat fill="#fff"/>
      </a>
    </li>
  </ul>
</template>

<style scoped>

</style>
