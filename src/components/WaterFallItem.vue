<script setup lang="ts">
import { computed } from "vue";

type WidthType = string | number;
interface WaterfallProps {
  width: WidthType;
  height: number;
}
const props = withDefaults(defineProps<WaterfallProps>(), {
  width: "20%",
});
/* width */
const waterfallItemWidth = computed(() => {
  const type = typeof props.width as WidthType;
  if (type === "string") {
    return props.width;
  } else if (type === "number") {
    return `${props.width}px`;
  } else {
    console.error(`width property just is string or number`);
  }
});
</script>

<template>
  <div
      class="waterfall-item"
      :style="{ width: waterfallItemWidth, height: `${props.height}px` }"
  >
    <slot />
  </div>
</template>

<style scoped>
.waterfall-item {
  position: absolute;
  margin: 0;
  padding: 0;
  box-sizing: border-box;
}
</style>
