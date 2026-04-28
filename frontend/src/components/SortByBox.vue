<script setup>
import { ref, onMounted } from 'vue'

const lists =ref([])

onMounted(async () => {
    const res = await fetch('http://127.0.0.1:3000/api/sortby-data')
    lists.value = await res.json()
})


</script>

<template>
    <div class="sort-box">
        <select name="sort_by" class="sort-select" id="sort-by-box">
            <option v-for="(list, index) in lists" :key="index" v-bind:value="list.value">{{ list.items }}</option>
        </select>
    </div>
</template>