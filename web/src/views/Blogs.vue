<script setup lang="ts">
</script>

<template>
  <div class="max-w-5xl 2xl:max-w-6xl mx-auto flex">
    <section class="w-full max-w-[65%] pr-10">
      <h2 class="text-[#F16B6F] text-1em" slide-enter>所有文章</h2>
      <article  class="not-first-of-type:mt-12 first-of-type:pt-4" v-for="(article, index) in 5" :key="article"
                slide-enter :style="{ '--stagger': index + 1 }">
        <a class="group" href="">
          <h3 class="group-hover:title-main">An Interactive Guide to CSS Grid</h3>
          <p class="text-justify">
            CSS Grid is an incredibly powerful tool for building layouts on the web, but like all powerful tools, there's a
            significant learning curve. In this tutorial, we'll build a mental model for how CSS Grid works and how we can
            use it effectively. I'll share the biggest 💡 lightbulb moments I've had in my own learning journey.
          </p>
          <div class="font-bold">Read more</div>
        </a>
      </article>
    </section>
    <section class="w-full max-w-[35%] pl-10">
      <h2 class="text-[#F16B6F] text-1em" slide-enter>最新文章</h2>
      <div class="pt-4">
        <ul class="flex flex-col">
          <li class="mb-4"  v-for="(article, index) in 5" :key="article" slide-enter :style="{ '--stagger': index + 1 }">An Interactive Guide to CSS Grid An Interactive Guide to CSS Grid</li>
        </ul>
      </div>
    </section>
  </div>
</template>

<style scoped>
</style>
