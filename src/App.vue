<script setup lang="ts">
import Logo from "@/components/Logo.vue";
</script>

<template>
  <div class="w-full h-full">
    <nav class="w-full h-20 px-5 py-5 flex justify-between items-center box-border">
      <div class="flex-1 h-full flex justify-start"><Logo class="h-full"/></div>
      <div class="title-main text-center flex-1">I love three things in the world</div>
      <ul class="nav flex justify-end flex-1">
        <li>博客</li>
        <li>项目</li>
        <li>摄影</li>
        <li>音乐</li>
      </ul>
    </nav>
  </div>
</template>

<style scoped>
.nav{

}
li{
  padding: 0 1rem;
}
</style>
