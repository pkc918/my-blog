<script setup lang="ts">
import { onMounted, ref } from "vue";
import Logo from "@/components/Logo.vue";
import OutLine from "@/components/OutLine.vue";
import { getProfile } from "@/api/api.ts";
import { GithubProfile } from "@/api/resType.ts";

const githubProfile = ref<GithubProfile>({});
onMounted(async () => {
  try {
    const res = await getProfile<GithubProfile>();
    githubProfile.value = res.data;
  } catch (e) {
    console.error(e);
  }
});
</script>

<template>
  <nav
    class="w-full h-15 px-5 py-3 flex justify-between items-center box-border"
  >
    <div class="flex-1 h-full flex justify-start">
      <RouterLink to="/">
        <Logo class="h-full text-4" />
      </RouterLink>
    </div>
    <div class="title-main text-center flex-1 text-truncate">
      <span
          :title="githubProfile?.bio">
        {{ githubProfile?.bio }}
      </span>
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
    bottom: -10px;
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
