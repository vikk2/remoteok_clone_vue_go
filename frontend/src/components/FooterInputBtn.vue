<script setup>
import { ref, onMounted } from 'vue'

const lists = ref({ footer_btn: [] })
onMounted(async () => {
    const res = await fetch('http://127.0.0.1:3000/api/btn-data')
    lists.value = await res.json()
    lists.value = lists.value[0]
})

</script>

<template>
    <div v-for="(list, index) in lists.footer_btn" :key="index" class="wrapper" id="input-btn-wrapper">
        <input v-if="list.footerInput" type="text" :placeholder="list.footerInput.placeHolder" :class="list.footerInput.className">
        <div v-if="list.subscribeBtn" :class="['btn', list.subscribeBtn.className]">{{ list.subscribeBtn.text }}</div>
    </div>
</template>