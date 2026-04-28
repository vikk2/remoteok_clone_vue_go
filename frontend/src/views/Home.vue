<script setup>
import { ref } from 'vue'
import BenefitBox from '@/components/BenefitBox.vue';
import FooterInputBtn from '@/components/FooterInputBtn.vue';
import JobCard from '@/components/JobCard.vue';
import LabelTab from '@/components/LabelTab.vue';
import Location from '@/components/Location.vue';
import LogoDropdown from '@/components/LogoDropdown.vue';
import Navigation from '@/components/Navigation.vue';
import NoteBox from '@/components/NoteBox.vue';
import PriceRange from '@/components/PriceRange.vue';
import SearchBox from '@/components/SearchBox.vue';
import SortByBox from '@/components/SortByBox.vue';

const hide = ref(true);

function hideFooter(){
    hide.value = !hide.value;
}


</script>

<template>
    <!-- nav -->
    <div class="nav-bar">
        <Navigation/>
    </div>
    <!-- //nav -->

    <div class="main">
        <!-- notification -->
        <NoteBox/>
        <!-- //notification -->

        <!-- filter -->
        <div class="filter-bar">
            <div class="filter-row">
                <LogoDropdown/>

                <SearchBox/>
                
                <Location/>
                
                <PriceRange/>

                <BenefitBox/>

                <SortByBox/>
                
            </div>

            <LabelTab/>
        </div>
        <!-- //filter -->

        <!-- job card -->
        <JobCard/>
        <!-- job card -->

    </div>
    
    <!-- footer -->
    <div v-show="hide" class="footer" id="footer">
        <p class="footer-text">
            Get new remote jobs sent to
        </p>
        <FooterInputBtn/>
        <div class="footer-x" id="x-footer" @click="hideFooter">
            ×
        </div>
    </div>
</template>
