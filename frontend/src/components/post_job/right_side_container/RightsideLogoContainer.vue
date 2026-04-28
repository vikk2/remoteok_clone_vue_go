<script setup>
import { ref, onMounted } from 'vue'

const logos = ref([])

onMounted(async () => {
    const res = await fetch('http://127.0.0.1:3000/api/post-job-logo')
    logos.value = await res.json()
})
</script>

<template>
    <div class="right-logo">
        <img v-for="(logo, index) in logos" :key="index" :src="logo.img">

        <div class="show-more">
            show more
        </div>
    </div>
</template>

<style scoped src="@/assets/css/style_post_job.css"></style>