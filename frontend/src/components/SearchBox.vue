<script setup>
import { ref, onMounted } from 'vue'

const lists = ref([])

onMounted(async () => {
    const res = await fetch('http://127.0.0.1:3000/api/search-tab-data')
    lists.value = await res.json()
    lists.value = lists.value[0]
})
</script>

<template>
    <div class="search-big-box">
        <input type="text" placeholder="🔍 Search" class="search-box">
        <img src="/images/chevron-down.svg" alt="chevron-down" width="15" class="chevron-down1">
        <ul class="search-menu" id="search-menu">
            <li v-for="(list, index) in lists" :key="index">
                <a href="#">{{ list }}</a>
            </li>
        </ul>
    </div>
</template>