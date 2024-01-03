<script setup lang="ts">
import { computed, ref, nextTick } from "vue";

type WidthType = string | number;

interface WaterfallProps {
  width: WidthType;
}

const props = withDefaults(defineProps<WaterfallProps>(), {
  width: 1467,
});

/* width */
const waterfallWidth = computed(() => {
  const type = typeof props.width as WidthType;
  if (type === "string") {
    return props.width;
  } else if (type === "number") {
    return `${props.width}px`;
  } else {
    console.error(`width property just is string or number`);
  }
});

const waterfallContainer = ref<HTMLDivElement>();
const cal = () => {
  const waterfallItemWidth =
      waterfallContainer.value?.children[0].getBoundingClientRect().width;
  console.log("waterfallItemWidth", waterfallItemWidth);
  const waterfallWidth =
      typeof props.width === "string"
          ? waterfallContainer.value?.clientWidth
          : props.width;
  console.log("waterfallWidth", waterfallContainer.value?.clientWidth);

  const columns = Math.floor(waterfallWidth! / waterfallItemWidth!);
  console.log("columns", columns);
  const spaceSum = columns - 1;
  const spaceWidth = waterfallWidth! - columns * waterfallItemWidth!;
  const spaceItemWidth = spaceWidth / spaceSum;
  console.log("spaceWidth", spaceWidth);
  console.log("spaceItemWidth", spaceItemWidth);
  return {
    spaceItemWidth: spaceItemWidth,
    waterfallWidth: waterfallWidth,
    waterfallItemWidth: waterfallItemWidth,
    columns: columns,
  };
};
const positions = ref<number[]>([]);

const setPosition = () => {
  const {spaceItemWidth, waterfallItemWidth = 0, columns} = cal();
  const length = waterfallContainer.value?.children.length || 0;
  for (let i = 0; i < length; i++) {
    if (i < columns) {
      (waterfallContainer.value?.children[i] as HTMLDivElement).style.top = "0";
      (waterfallContainer.value?.children[i] as HTMLDivElement).style.left = `${
          i * (spaceItemWidth + waterfallItemWidth)
      }px`;
      positions.value.push(
          (waterfallContainer.value?.children[i] as HTMLDivElement).offsetHeight
      );
    } else {
      let minIndex = 0;
      let minHeight = positions.value[0];
      for (let j = 0; j < positions.value.length; j++) {
        if (minHeight > positions.value[j]) {
          minHeight = positions.value[j];
          minIndex = j;
        }
      }
      (waterfallContainer.value?.children[i] as HTMLDivElement).style.top = `${
          positions.value[minIndex] + spaceItemWidth
      }px`;
      (waterfallContainer.value?.children[i] as HTMLDivElement).style.left = `${
          minIndex * (spaceItemWidth + waterfallItemWidth)
      }px`;
      console.log(
          "每一次高度",
          (waterfallContainer.value?.children[i] as HTMLDivElement).offsetHeight +
          spaceItemWidth
      );
      positions.value[minIndex] +=
          (waterfallContainer.value?.children[i] as HTMLDivElement).offsetHeight +
          spaceItemWidth;

      let maxHeight = 0;
      for (let j = 0; j < positions.value.length; j++) {
        if (maxHeight < positions.value[j]) {
          maxHeight = positions.value[j];
        }
      }

      (
          waterfallContainer.value as HTMLDivElement
      ).style.height = `${maxHeight}px`;
    }
  }
};

nextTick(() => {
  console.log(waterfallContainer.value?.children);
  // window.onload = setPosition;
  setPosition();
});
</script>
<template>
  <div
      ref="waterfallContainer"
      class="waterfall-vue3"
      :style="{ width: waterfallWidth }"
  >
    <slot/>
  </div>
</template>

<style scoped>
.waterfall-vue3 {
  position: relative;
  overflow: hidden;
  height: 100vh;
}
</style>
