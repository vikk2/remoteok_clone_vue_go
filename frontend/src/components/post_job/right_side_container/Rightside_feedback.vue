<script setup>
import { ref, onMounted } from 'vue'

const lists = ref([])
onMounted(async () => {
    const res = await fetch('http://127.0.0.1:3000/api/feedback-data')
    lists.value = await res.json()
})
</script>

<template>
    <div v-for="(list, index) in lists" :key="index" class="right-box pf">
        <div class="img-n-icon">
            <img :src="list.profile_img" alt="pf-1" width="200">
            <img v-if="list.icon" :src="list.icon" alt="green-circle" width="50" class="icon-on-pf">
        </div>
        <p v-html="list.feedback">
        </p>
    </div>
</template>

<style scoped src="@/assets/css/style_post_job.css"></style>

<style scoped>
.pf :deep(a) {
    color: #0072b1;
    text-decoration: none;
    border-bottom: 2px solid #0072b1;
    font-weight: 800;
}

.pf :deep(.a-gray) {
    color: #888;
    text-decoration: none;
    border-bottom: none;
    font-weight: 800;
}
</style>
