<script setup lang="ts">

import WaterFall from "@/components/WaterFall.vue";
import WaterFallItem from "@/components/WaterFallItem.vue";
</script>

<template>
  <div class="w-[85%]">
    <WaterFall width="100%">
      <WaterFallItem
          :width="200"
          class="item"
          :class="`item-${item % 22}`"
          v-for="item in 88"
          :key="item"
          slide-enter :style="{ '--stagger': item + 1 }"
      >
        {{ item }}
      </WaterFallItem>
      <WaterFallItem :width="200" :height="200" class="item">
        <img
            style="width: 100%"
            src="https://p26-passport.byteacctimg.com/img/user-avatar/65aae4fcc914b67b83f82d84d5a056db~90x90.awebp"
            alt=""
        />
      </WaterFallItem>
      <WaterFallItem :width="200" :height="200" class="item">
        <img
            style="width: 100%"
            src="https://avatars.githubusercontent.com/u/58922004?v=4"
            alt=""
        />
      </WaterFallItem>
      <WaterFallItem :width="200" :height="125" class="item">
        <img
            style="width: 100%"
            src="https://www.10wallpaper.com/wallpaper/medium/2312/Tradition_Lantern_Street_View_Tokyo_Japan_5K_medium.jpg"
            alt=""
        />
      </WaterFallItem>
      <WaterFallItem :width="200" :height="360" class="item">
        <img
            style="width: 100%"
            src="https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRBKD95au_2m0_wc2Xc9gcTqvN4d5v4uCdKMg&usqp=CAU"
            alt=""
        />
      </WaterFallItem>
    </WaterFall>
  </div>
</template>

<style scoped>
.item {
  float: left;
  display: flex;
  justify-content: center;
  align-items: center;
  font-size: 30px;
  font-weight: 700;
  color: aliceblue;
  margin-right: 15px;
  margin-bottom: 15px;
  width: 400px;
  position: absolute;
}

.item-1 {
  background-color: rgb(206, 169, 169);
  height: 300px;
}

.item-2 {
  background-color: rgb(131, 226, 174);
  height: 150px;
}

.item-3 {
  background-color: rgb(77, 30, 30);
  height: 350px;
}

.item-4 {
  background-color: rgb(49, 62, 134);
  height: 300px;
}

.item-5 {
  background-color: rgb(230, 99, 99);
  height: 200px;
}

.item-6 {
  background-color: rgb(206, 169, 169);
  height: 300px;
}

.item-7 {
  background-color: rgb(124, 126, 145);
  height: 400px;
}

.item-8 {
  background-color: rgb(169, 199, 38);
  height: 230px;
}

.item-9 {
  background-color: rgb(114, 128, 53);
  height: 300px;
}

.item-10 {
  background-color: rgb(48, 54, 18);
  height: 260px;
}

.item-11 {
  background-color: rgb(118, 122, 96);
  height: 230px;
}

.item-12 {
  background-color: rgb(118, 122, 96);
  height: 240px;
}

.item-13 {
  background-color: rgb(118, 122, 96);
  height: 250px;
}

.item-14 {
  background-color: rgb(118, 122, 96);
  height: 270px;
}

.item-15 {
  background-color: rgb(118, 122, 96);
  height: 330px;
}

.item-16 {
  background-color: rgb(118, 122, 96);
  height: 200px;
}

.item-17 {
  background-color: rgb(118, 122, 96);
  height: 100px;
}

.item-18 {
  background-color: rgb(118, 122, 96);
  height: 400px;
}

.item-19 {
  background-color: rgb(118, 122, 96);
  height: 340px;
}

.item-20 {
  background-color: rgb(118, 122, 96);
  height: 350px;
}

.item-21 {
  background-color: rgb(118, 122, 96);
  height: 360px;
}

.item-0 {
  background-color: rgb(118, 122, 96);
  height: 370px;
}
</style>
