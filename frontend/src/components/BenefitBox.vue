<script setup>
import { ref, onMounted } from 'vue'
const lists = ref([])
onMounted(async () => {
    const res = await fetch('http://127.0.0.1:3000/api/benefit-data')
    lists.value = await res.json()
})
</script>

<template>
    <div class="search-big-box2">
        <input type="text" placeholder="🎪 Benefits" class="search-box3">
        <img src="/images/chevron-down.svg" alt="chevron-down" width="15" class="chevron-down1">
        <ul class="search-menu" id="benefit-menu">
            <li v-for="(list, index) in lists" :key="index">
                <a href="#">{{ list }}</a>
            </li>
        </ul>
    </div>
</template>