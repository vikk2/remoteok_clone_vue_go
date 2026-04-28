<script setup>
import { ref } from "vue";
import PostHideBtn from '@/components/PostHideBtn.vue';

const showNoteBox = ref(true);
</script>

<template>
    <div v-if="showNoteBox" class="note_box" id="note-box">
        <a href="pages/post_job.html" class="note_box_link">
            <div class="note_text">
                👉 Hiring remotely? Reach 
                <span class="note-text-span">
                    2,500,000+ remote workers</span>
                on the 🏆 #1 Remote Job Platform
            </div>
        </a>

        <PostHideBtn @hide="showNoteBox = false"/>
    </div>
</template>