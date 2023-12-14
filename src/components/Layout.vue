<script setup lang="ts">
import NavList from "./NavList.vue";
</script>

<template>
  <div class="h-full flex flex-col">
    <header>
      <NavList />
    </header>
    <main class="flex-1">
      <RouterView />
    </main>
    <footer></footer>
  </div>
</template>

<style scoped></style>
