<script setup>
import { ref, onMounted } from 'vue'

const lists = ref({ post_btn: [], hide_btn: [] })
onMounted(async () => {
    const res = await fetch('http://127.0.0.1:3000/api/btn-data')
    lists.value = await res.json()
    lists.value = lists.value[0]
})

const emit = defineEmits(['hide'])
</script>

<template>
    <div id="post_and_hideBTN">
        <router-link v-for="list in lists.post_btn" :key="list.text" :to="list.href" :class="['btn', list.className]">
            {{ list.text }}
        </router-link>
        <button
            v-for="list in lists.hide_btn" :key="list.id" :class="['btn', list.className]" @click="emit('hide')">
            {{ list.text }}
        </button>
    </div>
</template>
