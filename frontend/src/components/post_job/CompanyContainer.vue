<script setup></script>

<template>
    <div class="left-company">
        <!-- compamy twitter -->
        <div class="header-label">
            Company
        </div>
        <div class="left-info-position" style="padding-top: 0;">
            <div class="position-title">
                COMPANY TWITTER*
            </div>
            <input type="text" placeholder="username" class="position-input focus-red-outline">
            <p class="position-text">
                Twitter username without @. Not required, but used to tag your company when we tweet out your job post.
            </p>
        </div>
        <!-- //compamy twitter -->

        <!-- company email -->
        <div class="left-info-position" style="padding-top: 0;">
            <div class="position-title">
                COMPANY EMAIL* (STAYS PRIVATE, FOR INVOICE + EDIT LINK)
            </div>
            <input type="text" class="position-input focus-red-outline">
            <p class="position-text">
                Make sure this email is accessible by you! We use this to send the invoice and edit link. We can not and do not manually resend it! If you use your company domain (same as company name), we will show a [ Verified ] tag on your job post.
            </p>
        </div>
        <!-- //company email -->

        <!-- invoice email -->
        <div class="left-info-position" style="padding-top: 0;">
            <div class="position-title">
                Invoice email (stays private)
            </div>
            <input type="text" class="position-input focus-red-outline">
            <p class="position-text">
                We send a copy of the invoice and edit link to here too. You can write your finance department or accountant expenses email here so they get a copy of the invoice for your bookkeeping.
            </p>   
        </div>
        <!-- //invoice email -->

        <!-- invoice address -->
        <div class="left-info-position" style="padding-top: 0;">
            <div class="position-title">
                Invoice address*
            </div>
            <textarea name="text-area" class="text-area focus-red-outline" placeholder="e.g. your company's full name and full invoice address, including building, street, city and country; also things like your VAT number, this is shown on the invoice."></textarea>
            <p class="position-text">
                Specify your company address here and we'll put it on your invoice for your bookkeeping. Be sure to [ Save changes ] in bottom right after editing your invoice address. Then it'll be instantly updated on the invoice.
            </p>        
        </div>
        <!-- //invoice address -->

        <!-- invoice notes -->
        <div class="left-info-position" style="padding-top: 0;">
            <div class="position-title">
                Invoice notes / PO number
            </div>
            <input type="text" class="position-input focus-red-outline">
            <p class="position-text">
                Not required. Add notes here that you'd like to see on the invoice/receipt such as a Purchase Order number or any other internal notes you need for reference. You can add or edit this later.
            </p>
        </div>
        <!-- //invoice notes -->

        <!-- pay later -->
        <div class="left-info-position" style="padding-top: 0;">
            <div class="position-title">
                Pay later
            </div>
            <div class="detail-job-checkbox pay-later">
                <input type="checkbox" class="checkbox focus-red-outline">
                I'd like to pay later
            </div>
            <p class="position-text">
                Need to get approval for this payment? Or send the invoice to your finance department first? No problem, we'll save your job post and send you (and your finance department below) a payment link. Once it's paid we immediately publish it!
            </p>        
        </div>
        <!-- //pay later -->
    </div>
</template>

<style scoped src="@/assets/css/style_post_job.css"></style>