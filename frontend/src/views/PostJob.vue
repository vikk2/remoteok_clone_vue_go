<script setup>
import CompanyContainer from '@/components/post_job/CompanyContainer.vue';
import CompanyInfo from '@/components/post_job/CompanyInfo.vue';
import DesignJobContainer from '@/components/post_job/DesignJobContainer.vue';
import FeedbackContainer from '@/components/post_job/FeedbackContainer.vue';
import LeftFooter from '@/components/post_job/footer/LeftFooter.vue';
import RightFooter from '@/components/post_job/footer/RightFooter.vue';
import JobDetailContainer from '@/components/post_job/JobDetailContainer.vue';
import PreviewContainer from '@/components/post_job/PreviewContainer.vue';
import PreviewPartnerContainer from '@/components/post_job/PreviewPartnerContainer.vue';
import PromotionHeader from '@/components/post_job/PromotionHeader.vue';
import Rightside_feedback from '@/components/post_job/right_side_container/Rightside_feedback.vue';
import Rightside_info from '@/components/post_job/right_side_container/Rightside_info.vue';
import RightsideLogoContainer from '@/components/post_job/right_side_container/RightsideLogoContainer.vue';
import { ref } from 'vue'

const trigger = ref(true);

function closeSidebar(){
    trigger.value = !trigger.value;
}
</script>

<template>
    <div class="post-job-page">
        <div class="nav-bar">
            <router-link to="/">
                <div class="logo">
                    <span class=remote>
                        remote
                    </span>
                    <span class="logo-ok-post">
                        OK
                        <span class="circle-r">®</span>
                    </span>
                </div>
            </router-link>
            <h1 class="nav-bar-h1">
                <router-link to="/">Hire Remotely</router-link>
            </h1>
            <div class="btn buy b-blue-hover">
                Buy a bundle →
            </div>
        </div>
        
        <!-- //nav -->

        <PromotionHeader/>


        <!-- main content -->
        <div class="big-container">

            <!-- left-container -->
            <div class="container-left" id="left-container" :style="{width:  trigger ? '' : '100%'}">

                <CompanyInfo/>

                <DesignJobContainer/>

                <JobDetailContainer/>

                <CompanyContainer/>

                <FeedbackContainer/>

                <PreviewContainer/>

                <PreviewPartnerContainer/>
            </div>

            <!-- //left container -->

            <!-- chevron to open the right container -->
            <div class="chevron to-open" id="ch-to-open" @click="closeSidebar">
                <img src="/images/ch-down.webp" alt="ch-open">
            </div>
            <!-- //chevron to open the right container -->


            <!-- right container -->
            <div v-show="trigger" class="container-right" id="right-container">

                <!-- chevron to close the right container -->
                <div class="chevron to-close" id="ch-to-close" @click="closeSidebar">
                    <img src="/images/ch-down.webp" alt="ch-down">
                </div>
                <!-- //chevron to closes the right container -->


                <div class="right-text">
                    Remote OK is <a href="#">the most popular remote jobs board in the world</a> trusted by millions of remote workers and leading remote companies like
                </div>
                
                <RightsideLogoContainer/>

                <Rightside_info/>
                
                <Rightside_feedback/>

            </div>
            <!-- //right container -->
        </div>

        <!-- footer -->
        <div class="footer">
            <LeftFooter/>
            <RightFooter/>
        </div>
        <!-- //footer -->
    </div>
</template>

<style scoped src="@/assets/css/style_post_job.css"></style>

<style scoped>
.post-job-page {
    min-height: 100vh;
    font-family: system-ui, sans-serif, "Nunito", "Helvetica", Arial, sans-serif;
    font-weight: 800;
    background-color: rgb(251, 250, 248);
}

.post-job-page .nav-bar {
    display: flex;
    justify-content: space-between;
    align-items: center;
    border-bottom: 1px solid #ddd;
    padding: 14px;
    padding-left: 20px;
    background-color: #fff;
}

.post-job-page .nav-bar a {
    text-decoration: none;
    color: #0f1115;
}

.post-job-page .logo {
    width: auto;
    padding: 0;
    text-align: left;
    display: flex;
    align-items: baseline;
    gap: 4px;
    white-space: nowrap;
}

.post-job-page .btn {
    background-color: #0072b1;
    color: #fff;
    padding: 8px;
    border-radius: 12px;
    border: 2px solid #0072b1;
    cursor: pointer;
}

.post-job-page .promo-code {
    text-align: center;
    padding: 14px;
    background-color: #fff;
    font-weight: bold;
}

.post-job-page .discount {
    background-color: #0072b1;
    color: #fff;
    text-align: center;
    font-weight: bold;
    padding: 14px;
}

.post-job-page .discount a {
    color: #fff;
}

.post-job-page .search-big-box {
    position: relative;
    padding: 14px;
    padding-top: 4px;
}

.post-job-page .search-box {
    outline: none;
    border: none;
    border-radius: 6px;
    padding: 7px;
    font-size: 16px;
    width: 100%;
    box-shadow: 0 0 0 1px #dddddd;
    background-color: #fbfaf8;
    margin-top: 7px;
    margin-bottom: 7px;
}

.post-job-page .footer {
    display: flex;
    background-color: #fff;
    position: fixed;
    bottom: 0;
    width: 100%;
    justify-content: flex-start;
    align-items: stretch;
    min-height: 0;
    padding: 0;
    border-top: none;
}

.post-job-page .footer-left {
    width: 67%;
    padding: 14px;
    border: 1px solid #ddd;
}

.post-job-page .footer-right {
    width: 33%;
    padding: 28px;
    padding-bottom: 21px;
    border: 1px solid #ddd;
    border-left: none;
    border-right: none;
}

.post-job-page .footer-right .big-btn {
    display: block;
    font-size: 24px;
    padding: 19px;
    border-radius: 6px;
    text-align: center;
    background-color: #0072b1;
    color: #fff;
    border: 2px solid #0072b1;
}

.post-job-page .footer-right .big-btn span {
    text-decoration: line-through;
}

.post-job-page .pf a {
    color: #0072b1;
    text-decoration: none;
    border-bottom: 2px solid #0072b1;
    font-weight: 800;
}

.post-job-page .pf .a-gray {
    color: #888;
    text-decoration: none;
    border-bottom: none;
    font-weight: 800;
}

.post-job-page .b-blue-hover:hover{
    background-color: #fff;
    border-color: #0072b1;
    color: #0072b1;
}

@media (max-width: 1144px) {
    .post-job-page .footer-left {
        display: none;
    }

    .post-job-page .footer-right {
        width: 100%;
    }
}
.nav-bar-h1 a{
    font-family: "system-ui";
}

</style>
