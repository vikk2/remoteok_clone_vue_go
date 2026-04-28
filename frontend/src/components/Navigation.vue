<script setup>
import { ref, onMounted } from 'vue'

const lists = ref({ nav_btns: [] })
onMounted(async () => {
    const res = await fetch('http://127.0.0.1:3000/api/btn-data')
    lists.value = await res.json()
    lists.value = lists.value[0]
})

</script>


<template>
    <ul class="nav-menu" id="navBtn">
        <li v-for="list in lists.nav_btns" :key="list.id">
            <router-link :to="list.href" v-bind:class="['btn', list.className]">
                {{ list.text }}
            </router-link>
        </li>
    </ul>
</template>


