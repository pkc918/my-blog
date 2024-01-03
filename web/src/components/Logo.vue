<script setup lang="ts">

</script>

<template>
  <svg viewBox="0 0 115 87" fill="none" xmlns="http://www.w3.org/2000/svg">
    <path
        d="M114 44.093C109.88 47.7471 106.599 49.6951 100 53.093C90.8221 57.5759 79.3707 58.3801 78.5 43.093C80.0522 40.8876 84.725 39.3082 91.5 38.093C78.5103 40.0263 73.0071 44.6778 66 58.593C67.0309 46.1642 68.9812 39.3117 72.5 27.093C79.4008 5.89109 83.2473 -2.14516 90 2.09296C94.1465 5.4209 86.7337 19.5374 72.5 35.593C65.2679 43.2695 50.1043 49.2501 21 54.0929C23.5682 52.6557 25.539 51.2173 28.5 48.0929C34.5046 40.444 33.1287 36.4866 25.5 38.093C14.6334 42.8261 12.1962 56.048 1.00002 86.0929L15 33.0929"
        stroke="black"/>
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
    stroke-dashoffset: 1px;
    stroke-dasharray: 0 600px;
    opacity: 0;
  }
  10% {
    opacity: 1;
  }
  40% {
    stroke-dasharray: 600px 0;
  }
  /* Moving back */
  85% {
    stroke-dasharray: 600px 0;
  }
  95%,
  to {
    stroke-dasharray: 0 600px;
  }
}
path {
  stroke-dashoffset: 1px;
  stroke-dasharray: 600px 0;
  animation: grow 10s ease forwards infinite;
  transform-origin: center;
  stroke: #000;
  animation-delay: 0s;
}
</style>
