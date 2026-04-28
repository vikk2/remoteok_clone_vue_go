<script setup>
import { ref, onMounted } from 'vue'

const cards = ref([]);

onMounted(async() => {
    const res = await fetch('http://127.0.0.1:3000/api/card-info');
    cards.value = await res.json();
});


const open = ref(false);

function cardContainer (){
    open.value = !open.value;
}
</script>

<template>
    <div id="jobCard">
        <template v-for="(card, index) in cards" :key="index">
            <table v-if="card.type === 'job' && card.jobClass === 'adRow'" class="job-row-0">
                <tr>
                    <td class="logo">
                        <img :src="card.logo" alt="logo-biz-0" width="56" style="background-color: #283C49;">
                    </td>
                    <td class="info">
                        <table style="width: 100%;">
                            <tr>
                                <td>
                                    <a href="#" class="title">{{ card.jobTitle }}</a>
                                </td>
                            </tr>
                            <tr>
                                <td style="font-weight: 400; color: #fff; padding: 0;" class="sub-title">{{ card.company }}</td>
                            </tr>
                            <tr>
                                <td style="padding: 0;">
                                    <table>
                                        <tr>
                                            <td style="padding: 0;">
                                            </td>
                                        </tr>
                                    </table>
                                </td>
                            </tr>
                        </table>
                    </td>
                    <td class="label-col">
                    </td>
                    <td class="date date-ad">
                        Ad
                    </td>
                    <td class="apply-col0">
                        <div class="btn apply-btn-0">
                            Sign up today
                        </div>
                    </td>
                </tr>
            </table>

            <table v-if="card.cardVariant === 'normal'" :class="['job-row', card.jobClass]">
                <tr class="special-tr-1">
                    <td v-if="card.logoType != 'text'" class="logo">
                        <img :src="card.logo" alt="logo-biz" width="56">
                    </td>
                    <td class="info">
                        <table style="width: 100%;">
                            <tr>
                                <td>
                                    <a href="#" class="title">{{card.jobTitle}}</a>
                                    <span v-if="card.balloonIcon">🎈</span>
                                    <span v-if="card.verify" class="verify-badge">VERIFIED</span>
                                </td>
                            </tr>
                            <tr>
                                <td :class="card.jobClass === 'job2' || card.jobClass === 'job3' ? 'sub-title2' : 'sub-title'">
                                    {{ card.company }}
                                    <img v-if="card.new" src="/images/new-gif.webp" alt="new-gif" width="30">
                                </td>
                            </tr>
                            <tr>
                                <td style="padding: 0;">
                                    <table>
                                        <tr>
                                            <td style="padding: 0;">
                                                <div class="label">
                                                    <span v-if="card.location" class="label-box">{{ card.location }}</span>
                                                    <span v-if="card.salary" class="label-box">{{ card.salary }}</span>
                                                    <span v-if="card.schedule" class="label-box">{{ card.schedule }}</span>
                                                </div>
                                            </td>
                                        </tr>
                                    </table>
                                </td>
                            </tr>
                        </table>
                    </td>
                    <td class="label-col">
                        <div class="label2">
                            <span v-for="(label, index) in card.label" :key="index" :class="card.jobClass === 'job2' || card.jobClass === 'job3' ? 'label-box2-2' : 'label-box2'">
                                {{ label }}
                                <span class="label-box-tooltip">Add to filters</span>
                            </span>
                        </div>
                    </td>
                    <td :class="card.jobClass === 'job2' || card.jobClass === 'job3' ? 'date2' : 'date'">
                        <img src="/images/paper-clip.svg" alt="clip"> {{ card.date }}
                    </td>
                    <td class="apply-col">
                        <div :class="card.jobClass === 'job2' || card.jobClass === 'job3' ? 'apply-btn-1' : 'btn apply-btn'">
                            Apply
                        </div>
                    </td>
                </tr>
            </table>

            <table v-if="card.jobClass === 'job1' && card.cardVariant === 'expand'" class="job-row job1">
                <tr class="special-tr-1" id="tr-1" @click="cardContainer">
                    <td v-if="card.logoType === 'text'" class="logo" id="border-radius-logo" :style="{borderBottomLeftRadius: open ? '0px' : '12px'}">
                        <div class="GP-logo">
                            <p>{{ card.logoText }}</p>
                        </div>
                    </td>
                    <td class="info">
                        <table style="width: 100%;">
                            <tr>
                                <td>
                                    <a href="#" class="title">{{ card.jobTitle }}</a>
                                    <span v-if="card.balloonIcon">🎈</span>
                                    <span v-if="card.verify" class="verify-badge">VERIFIED</span>
                                </td>
                            </tr>
                            <tr>
                                <td :class="card.jobClass === 'job2' || card.jobClass === 'job3' ? 'sub-title2' : 'sub-title'">
                                    {{ card.company }}
                                    <img v-if="card.new" src="/images/new-gif.webp" alt="new-gif" width="30">
                                </td>
                            </tr>
                            <tr>
                                <td style="padding: 0;">
                                    <table>
                                        <tr>
                                            <td style="padding: 0;">
                                                <div class="label">
                                                    <span v-if="card.location" class="label-box">{{ card.location }}</span>
                                                    <span v-if="card.salary" class="label-box">{{ card.salary }}</span>
                                                    <span v-if="card.schedule" class="label-box">{{ card.schedule }}</span>    
                                                </div>
                                            </td>
                                        </tr>
                                    </table>
                                </td>
                            </tr>
                        </table>
                    </td>
                    <td class="label-col">
                        <div class="label2">
                            <span v-for="(label, index) in card.label" :key="index" :class="card.jobClass === 'job2' || card.jobClass === 'job3' ? 'label-box2-2' : 'label-box2'">
                                {{ label }}
                                <span class="label-box-tooltip">Add to filters</span>
                            </span>
                        </div>
                    </td>
                    <td class="date">
                        {{ card.date }}
                    </td>
                    <td class="apply-col" id="border-radius-apply-col" :style="{borderBottomRightRadius: open ? '0px' : '12px'}">
                        <div class="btn apply-btn">
                            Apply
                        </div>
                    </td>
                </tr>
                <tr v-show="open" v-if="card.jobClass === 'job1' && card.cardVariant === 'expand'" class="special-tr-2" id="tr-2">
                    <td colspan="5" style="padding: 0;">
                        <div class="job-detail-whole">
                            <div class="job-detail">
                                <div class="detail-upper">
                                    <div class="detail-logo">
                                        <h2 class="gno-h2">
                                            {{ card.company }}
                                        </h2>
                                        <p><a href="#" class="gno-a">{{ card.details.websiteURL }}</a></p>
                                        <div class="btn-apply-now">
                                            Apply now
                                        </div>
                                        <p>
                                            Share this job:
                                        </p>
                                        <img :src="card.details.qr_code" alt="qrcode" width="80" class="qrcode">
                                        <input type="text" value="rok.co/ks">
                                    </div>
                                    <div class="detail-title">
                                        <span>{{ card.company }} is hiring a</span>
                                        <h1>Remote {{ card.jobTitle }}</h1>
                                    </div>
                                    <div class="detail">
                                        <div v-html="card.details.job_description"></div>
                                        <h2>
                                            Salary and compensation
                                        </h2>
                                        <br>
                                        <p>$220,000 — $300,000/year</p>
                                    </div>
                                </div>
                            </div>
                            <div class="apply-box">
                                <div>
                                    <h2>How do you apply?</h2>
                                    <p><strong>Please submit yout application through Breezy HR:</strong> https://gno-partners.breezy.hr/p/154fdc8fa066-senior-amazon-brand-manager?state=published</p>
                                </div>
                                <div class="apply-box-btn">
                                    Apply for this job
                                </div>
                            </div>
                        </div>
                    </td>
                </tr>
            </table>
            <a v-if="card.type === 'ad' && card.jobClass === 'adReview'" href="#" class="ad-link">
                <div class="ad">
                    <div class="ad-image">
                        <img :src="card.adImage" alt="ad-pic">
                    </div>
                    <div class="ad-text">
                        <h2>{{ card.adTitle }}</h2>
                        <p v-html="card.adDescription">
                        </p>
                        <div class="ad-rating">
                            <div class="ad-rating-star">★★★★★</div>
                            <div class="ad-rating-text">4.8/5 from 10,000+ reviews</div>
                        </div>
                        <div class="ad-btn">
                            Get instant access
                        </div>
                    </div>
                </div>
            </a>
        </template>
    </div>
    
</template>
