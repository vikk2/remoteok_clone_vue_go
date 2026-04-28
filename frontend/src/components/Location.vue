<script setup>
import { ref, onMounted } from 'vue'

const lists = ref([])

onMounted(async () => {
    const res = await fetch('http://127.0.0.1:3000/api/location-data')
    lists.value = await res.json()
})
</script>

<template>
    <div class="search-big-box2">
        <input type="text" placeholder="🌏 Location" class="search-box2">
        <img src="/images/chevron-down.svg" alt="chevron-down" width="15" class="chevron-down1">
        <ul class="location-menu" id="locationMenu">
            <template v-for="(list, index) in lists" :key="index">
                <li v-if="list.type === 'group'" class="group">{{list.location}}</li>
                <li v-else>
                    <a href="#">{{ list.location }}</a>
                </li>
            </template>
        </ul>
    </div>
</template>