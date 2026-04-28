<script setup>
</script>

<template>
    <div class="left-info">
        <div class="left-info-position">
            <div class="header-label">
                let's start
            </div>
            <div class="position-title">
                POSITION*
            </div>
            <input type="text" placeholder="Position" class="position-input focus-red-outline">
            <p class="position-text">
                Please specify as single job position like "Marketing Manager" or "Node JS Developer", not a sentence like "Looking for PM / Biz Dev / Manager". We know your job is important but please DO NOT WRITE IN FULL CAPS. If posting multiple roles, please create multiple job posts. A job post is limited to a single job. We only allow real jobs, absolutely no MLM-type courses "learn how to work online" please.
            </p>
        </div>
        
        <div class="left-info-company">
            <div class="company-title">
                COMPANY NAME*
            </div>
            <input type="text" placeholder="Company name" class="company-input focus-red-outline">
            <p class="company-text">
                Your company's brand/trade name: without Inc., Ltd., B.V., Pte., etc.
            </p>
        </div>

        <div class="left-info-description">
            <div class="title-btn-wrapper">
                <div class="description-title">
                    JOB DESCRIPTION
                </div>
                <div class="btn description-btn b-blue-hover">
                    ✨ Write with AI
                </div>
            </div>
            <div class="edit-container">
                <div class="upper-container">
                    <div class="div-tools">
                        <img src="/images/undo.svg" alt="undo">
                        <img src="/images/redo.svg" alt="redo">
                    </div>
                    <div class="div-tools">
                        <img src="/images/bold.svg" alt="bold">
                        <img src="/images/italic.svg" alt="italic">
                    </div>
                    <div class="div-tools">
                        <img src="/images/h1.svg" alt="h1">
                        <img src="/images/h2.svg" alt="h2">
                    </div>
                    <div class="div-tools">
                        <img src="/images/bullet-point.svg" alt="list-bullet">
                        <img src="/images/bullet-num.svg" alt="list-number">
                    </div>
                    <div class="div-tools">
                        <img src="/images/decrease.svg" alt="decrease-indent">
                        <img src="/images/increase.svg" alt="increase-indent">
                    </div>
                    <div class="div-tools" style="border-right: none;">
                        <img src="/images/clip2.svg" alt="clip">
                    </div>
                </div>

                <div class="lower-container">

                </div>
            </div>
        </div>

        <div class="left-info-company">
            <div class="company-title">
                EMPLOYMENT TYPE*
            </div>
            <select name="employment-type" class="select-employment focus-red-outline">
                <option selected value="default">Full-time</option>
                <option value="part-time">Part-time</option>
                <option value="contractor">Contractor</option>
                <option value="temporary">Temporary</option>
                <option value="internship">Internship</option>
                <option value="perdiem">Per diem</option>
                <option value="volunteer">Volunteer</option>
                <option value="onsite">Onsite</option>
            </select>
        </div>

        <div class="left-info-company">
            <div class="company-title">
                PRIMARY TAG*
            </div>
            <select name="employment-type" class="select-employment focus-red-outline">
                <option selected value="default">Select a primary tag</option>
                <option value="software">Software Development</option>
                <option value="customer-spp">Customer Support</option>
                <option value="sale">Sales</option>
                <option value="marketing">Marketing</option>
                <option value="design">Design</option>
                <option value="front">Front End</option>
                <option value="back">Back End</option>
                <option value="legal">Legal</option>
                <option value="testing">Testing</option>
                <option value="assurance">Quality Assurance</option>
                <option value="non-tech">Non-Tech</option>
                <option value="other">Other</option>
            </select>
            <p class="company-text">
                This primary tag shows first and increases visibility in the main sections. Your job is shown on every page that is tagged with though. E.g. if you tag it as PHP, it shows for Remote PHP Jobs etc.
            </p>
        </div>

        <div class="search-big-box">
            <div class="company-title">
                TAGS, KEYWORDS OR STACK*
            </div>
            <input type="text" placeholder="Type a tag or keyword to search and add it" class="search-box focus-red-outline">
            <ul class="search-menu">
                <li><a href="#">🎧 Support</a></li>
                <li><a href="#">🤓 Engineer</a></li>
                <li><a href="#">🤓 Software</a></li>
                <li><a href="#">👵 Senior</a></li>
                <li><a href="#">🛠 Technical</a></li>
                <li><a href="#">💼 Management</a></li>
                <li><a href="#">🚀 Growth</a></li>
                <li><a href="#">👩‍✈️ Lead</a></li>
                <li><a href="#">🎨 Design</a></li>
                <li><a href="#">💼 Manager</a></li>
                <li><a href="#">💼 Sales</a></li>
                <li><a href="#">💰 Financial</a></li>
                <li><a href="#">🚥 Marketing</a></li>
                <li><a href="#">🔑 Security</a></li>
                <li><a href="#">♾️ Operations</a></li>
                <li><a href="#">☁️ Cloud</a></li>
                <li><a href="#">♾️ System</a></li>
                <li><a href="#">♾️ Operational</a></li>
                <li><a href="#">💼 Strategy</a></li>
                <li><a href="#">💼 Leader</a></li>
                <li><a href="#">🤓 Code</a></li>
                <li><a href="#">🔌 Non Tech</a></li>
                <li><a href="#">👨‍🏫 Training</a></li>
                <li><a href="#">✍️ Content</a></li>
            </ul>
            <p class="company-text">
                Short tags are preferred. Use tags like industry and tech stack. The first 3 or 4 tags are shown on the site, the other tags aren't but the job will be shown on each tag specific page (like /remote-react-jobs). We also sometimes generate tags automatically after you post/edit to supplement.
            </p>
        </div>

        <div class="search-big-box">
            <div class="company-title">
                TAGS, KEYWORDS OR STACK*
            </div>
            <input type="text" placeholder="Type a tag or keyword to search and add it" class="search-box focus-red-outline">
            <ul class="search-menu">
                <li class="group">Regions</li>
                <li><a href="#">🌏 Worldwide</a></li>
                <li><a href="#">⛰️ North America</a></li>
                <li><a href="#">💃 Latin America</a></li>
                <li class="group">🇪🇺 Europe</li>
                <li><a href="#">🦁 Africa</a></li>
                <li><a href="#">🕌 Middle East</a></li>
                <li><a href="#">⛩ Asia</a></li>
                <li><a href="#">🌊 Oceania</a></li>
                <li class="group">Countries</li>
                <li><a href="#">🇰🇭 Cambodia</a></li>
                <li><a href="#">🇺🇸 United States</a></li>
                <li><a href="#">🇨🇦 Canada</a></li>
                <li><a href="#">🇬🇧 United Kingdom</a></li>
            </ul>
            <p class="company-text">
                If you'd only like to hire people from a specific location or timezone this remote job is restricted to (e.g. Europe, United States or Japan). If not restricted, please leave it as "Worldwide". The less restricted this is, the more applicants you will get. Keeping it "Worldwide" is highly recommended as you'll have access to a worldwide pool of talent. <strong>To promote fairness in remote work positions, <a href="#">worldwide jobs are ranked higher.</a></strong>
            </p>
            </div>

    </div>
</template>

<style scoped src="@/assets/css/style_post_job.css"></style>