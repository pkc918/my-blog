<script setup lang="ts"></script>

<template>
  <svg
    width="37"
    height="8"
    viewBox="0 0 37 8"
    fill="none"
    class="sc-d2e67f0c-0 sc-d2e67f0c-3 fnXKmp DyqBc"
  >
    <path
      d="M1 5.39971C7.48565 -1.08593 6.44837 -0.12827 8.33643 6.47992C8.34809 6.52075 11.6019 2.72875 12.3422 2.33912C13.8991 1.5197 16.6594 2.96924 18.3734 2.96924C21.665 2.96924 23.1972 1.69759 26.745 2.78921C29.7551 3.71539 32.6954 3.7794 35.8368 3.7794"
      stroke="hsl(333deg, 100%, 45%)"
      stroke-width="2"
      stroke-linecap="round"
      stroke-linejoin="round"
    ></path>
  </svg>
</template>

<style scoped>
@media (prefers-reduced-motion) {
  path {
    animation: none !important;
    stroke-dasharray: unset !important;
  }
}
@keyframes grow {
  0% {
    stroke-dashoffset: 45;
    stroke-dasharray: 45;
    opacity: 0;
  }
  5% {
    opacity: 1;
  }
  100% {
    stroke-dashoffset: 0;
  }
}

path {
  stroke-dashoffset: 45;
  stroke-dasharray: 45;
  animation: 500ms cubic-bezier(0.27, 0.22, 0.44, 1.03) 250ms 1 normal both
    running grow;
  transform-origin: center;
  stroke: #3ac569;
  animation-delay: 0s;
}
</style>
