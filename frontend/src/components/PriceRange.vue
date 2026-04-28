<script setup>
import { onMounted, ref } from 'vue'

const price = ref({})
const currentPrice = ref("0")

onMounted(async () => {
    const res = await fetch('http://127.0.0.1:3000/api/price-range-data')
    price.value = await res.json()
    price.value = price.value[0]
    currentPrice.value = price.value.defaultPrice
})

</script>

<template>
    <div class="search-big-box2">
        <input type="text" placeholder="💵 Salary" class="search-box3">
        <img src="/images/chevron-down.svg" alt="chevron-down" width="15" class="chevron-down1">
        <div class="salary" id="salaryPrice">
            <div class="mini">Minimum</div>
            <div class="price">${{ currentPrice }}k/year</div>
            <input v-model="currentPrice" type="range" class="range-price" :min="price.minimumPrice" :max="price.maximumPrice" :step="price.step">
        </div>
    </div>
</template>
