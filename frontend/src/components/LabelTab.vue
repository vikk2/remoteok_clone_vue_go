<script setup>
import { ref, onMounted } from 'vue'

const lists = ref([])

onMounted(async () => {
    const res = await fetch('http://127.0.0.1:3000/api/tab-label-data')
    lists.value = await res.json()
    lists.value = lists.value[0]
})
</script>

<template>
    <div class="lower-row" id="tab-box">
        <template v-for="(list, index) in lists" :key="index">
            <div v-if="index === 0" class="tag-box first-box">
                {{ list }}
                <span class="tag-box-tooltip">Suggested filter, tap to add</span>
            </div>
            <div v-else class="tag-box">
                {{ list }}
            <span class="tag-box-tooltip">Suggested filter, tap to add</span>
            </div>
        </template>
    </div>
</template>