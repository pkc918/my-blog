<script setup lang="ts">
import Logo from "@/components/Logo.vue";
import OutLine from "@/components/OutLine.vue";
</script>

<template>
  <nav
    class="w-full h-20 px-5 py-5 flex justify-between items-center box-border"
  >
    <div class="flex-1 h-full flex justify-start">
      <RouterLink to="/">
        <Logo class="h-full" />
      </RouterLink>
    </div>
    <div class="title-main text-center flex-1 text-truncate">
      <span
          title="I love three things in the world, the sun the moon and you, the sun for the day, the moon for the night and you forever."
        >I love three things in the world, the sun the moon and you, the sun for
        the day, the moon for the night and you forever.</span
      >
    </div>
    <ul class="navlist-btn flex justify-end flex-1">
      <li>
        <RouterLink class="flex flex-col items-center relative" to="/blog">
          <span>博客</span>
          <OutLine />
        </RouterLink>
      </li>
      <li>
        <RouterLink class="flex flex-col items-center relative" to="/projects">
          <span>项目</span>
          <OutLine />
        </RouterLink>
      </li>
      <li>
        <RouterLink class="flex flex-col items-center relative" to="/photo">
          <span>摄影</span>
          <OutLine />
        </RouterLink>
      </li>
      <li>
        <RouterLink class="flex flex-col items-center relative" to="/music">
          <span>音乐</span>
          <OutLine />
        </RouterLink>
      </li>
    </ul>
  </nav>
</template>

<style scoped>
li {
  padding: 0 1rem;

  svg {
    position: absolute;
    bottom: -1rem;
    display: none;
  }

  .router-link-active {
    svg {
      display: inline-block;
    }
    span {
      font-weight: 600;
      color: #0a0c10;
    }
  }
}
</style>
