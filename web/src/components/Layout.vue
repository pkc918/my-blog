<script setup lang="ts">
import NavList from "@/components/NavList.vue";
</script>

<template>
  <div class="h-full min-h-full flex flex-col">
    <header class="w-full fixed top-0 bg-[rgba(255,255,255,1)] z-1">
      <NavList />
    </header>
    <main class="flex-1 flex justify-center mt-20">
      <RouterView />
    </main>
    <footer></footer>
  </div>
</template>

<style scoped>

</style>
