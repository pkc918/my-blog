<script setup lang="ts">
import NavList from "./NavList.vue";
</script>

<template>
  <div class="h-full flex flex-col bg-[rgba(207,240,218,0.3)]">
    <header>
      <NavList />
    </header>
    <main class="flex-1 flex justify-center">
      <RouterView />
    </main>
    <footer></footer>
  </div>
</template>

<style scoped></style>
