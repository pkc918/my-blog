<script setup lang="ts">
import NavList from "./NavList.vue";
</script>

<template>
  <header>
    <NavList />
  </header>
  <main>
    <RouterView />
  </main>
  <footer></footer>
</template>

<style scoped></style>
