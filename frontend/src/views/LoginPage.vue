<script setup>
import FormLogin from '@/components/login_page/FormLogin.vue';
import LogoDropdown from '@/components/LogoDropdown.vue';
import Navigation from '@/components/Navigation.vue';


</script>

<template>
    <div class="login-page">
        <div class="nav-bar">
            <LogoDropdown/>
            <Navigation/>
        </div>
        <!-- //nav -->

        <FormLogin/>
    </div>
</template>

<style src="@/assets/css/style_login.css"></style>
