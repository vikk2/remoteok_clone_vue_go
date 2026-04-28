<script setup>
import { ref, onMounted } from 'vue'

const lists = ref([])

onMounted(async () => {
    const res = await fetch('http://127.0.0.1:3000/api/right-side-info')
    lists.value = await res.json()
})
</script>

<template>
    <div v-for="(item, index) in lists" :key="index" :class="['right-box', item.box_class]">
        {{ item.label }}
        <span :class="item.value_class">
            <img v-for="(image, imageIndex) in item.value_images" :key="imageIndex" :src="image.src" :alt="image.alt" :width="image.width">
            {{ item.value_text }}
        </span>
        <template v-if="item.suffix">
            {{ item.suffix }}
        </template>
        <div v-if="item.link_text" class="right-box-link">
            {{ item.link_prefix }}
            <a :href="item.link_href">{{ item.link_text }}</a>
        </div>
    </div>
</template>

<style scoped src="@/assets/css/style_post_job.css"></style>