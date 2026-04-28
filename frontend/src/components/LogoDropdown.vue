<script setup>
import { onBeforeUnmount, onMounted, ref } from 'vue'

const menu = ref([])
const show = ref(false)
const dropdownRef = ref(null)

onMounted(async () => {
    document.addEventListener('click', handleOutsideClick)

    try {
        const res = await fetch('http://127.0.0.1:3000/api/logo-dropdown-data')

        if (!res.ok) {
            throw new Error(`Request failed with status ${res.status}`)
        }

        menu.value = await res.json()
    } catch (error) {
        console.error('Failed to load logo dropdown data:', error)
        menu.value = []
    }
})

onBeforeUnmount(() => {
    document.removeEventListener('click', handleOutsideClick)
})

function handleOutsideClick(event) {
    if (!dropdownRef.value?.contains(event.target)) {
        show.value = false
    }
}

function toggleDropdown() {
    show.value = !show.value
}

function groupItem(items = []) {
    const groups = []

    for (let i = 0; i < items.length; i += 2) {
        groups.push(items.slice(i, i + 2))
    }

    return groups
}

function isRouterLink(href) {
    return href?.startsWith('/') ?? false
}
</script>

<template>
    <div ref="dropdownRef" class="logo-dropdown-wrapper">
        <button class="logo-remote" id="logo-click-to-dropmenu" type="button" @click="toggleDropdown">
            <img src="/images/logo.webp" alt="logo" width="40" class="logo-ok">
            <img src="/images/chevron-down.svg" alt="chevron-down" width="15" class="chevron-down">
        </button>
        <div v-show="show" class="logo-dropdown" id="drop-menu">
            <div v-for="(list, listIndex) in menu" :key="listIndex">
                <div v-if="list.title" class="logo-dropdown-group-line">
                    <div class="logo-dropdown-group">{{ list.title }}</div>
                </div>

                <div v-for="(group, index) in groupItem(list.items)" :key="index" class="group">
                    <component
                        :is="isRouterLink(item.href) ? 'router-link' : 'a'"
                        v-for="(item, itemIndex) in group"
                        :key="`${item.href}-${itemIndex}`"
                        :to="isRouterLink(item.href) ? item.href : undefined"
                        :href="isRouterLink(item.href) ? undefined : item.href"
                    >
                        <img v-if="item.img" :src="item.img" width="13" alt="logo" class="group-logo">
                        {{ item.label }}
                    </component>
                </div>
            </div>
        </div>
    </div>
</template>
