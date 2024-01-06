<script setup lang="ts">
import Typed from "typed.js";
import { nextTick, onMounted, onUnmounted, ref } from "vue";
import FoundMe from "@/components/FoundMe.vue";
import { GithubProfile } from "@/api/resType.ts";
import { getProfile } from "@/api/api.ts";

const excerpt = ref();
const typed = ref<Typed>();

const githubProfile = ref<GithubProfile>({});
onMounted(async () => {
  try {
    const res = await getProfile<GithubProfile>();
    githubProfile.value = res.data;
  } catch (e) {
    console.error(e);
  }
});
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
      class="w-70 h-70 rd-full shadow-[0_20px_60px_-10px_rgba(0,0,0,0.3)]"
      :src="githubProfile?.avatar_url"
      alt="avatar"
    />
    <div class="h-50 mt-5">
      <span ref="excerpt"></span>
    </div>
    <FoundMe class="absolute right-10 top-50% translate-y--90%" />
  </div>
</template>

<style scoped></style>
