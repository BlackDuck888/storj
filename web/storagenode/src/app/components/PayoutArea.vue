// Copyright (C) 2019 Storj Labs, Inc.
// See LICENSE for copying information.

<template>
    <div class="payout-container">
        <svg class="payout-container__image" width="40" height="35" viewBox="0 0 40 35" fill="none" xmlns="http://www.w3.org/2000/svg" alt="wallet image">
            <path fill-rule="evenodd" clip-rule="evenodd" d="M2 7C2.57391 7.33667 3.23903 7.53563 3.94595 7.53563C14.9031 7.52607 27.1226 7.53563 36.5405 7.53563C37.3653 7.53563 38 8.17454 38 9.00473V13.9017H23.8919C22.5616 13.9017 21.4595 15.0112 21.4595 16.3502V24.1854C21.4595 25.5244 22.5616 26.6339 23.8919 26.6339H38V31.5309C38 32.3611 37.3653 33 36.5405 33H3.45946C2.63472 33 2 32.3611 2 31.5309V7Z" fill="#A5C7EF"/>
            <path fill-rule="evenodd" clip-rule="evenodd" d="M3.90244 0C1.75688 0 0 1.75078 0 3.88889V31.5972C0 33.4657 1.53961 35 3.41463 35H36.5854C38.4604 35 40 33.4657 40 31.5972V9.23611C40 7.3676 38.4604 5.83333 36.5854 5.83333H34.1463V0.972222C34.1463 0.463313 33.6814 0 33.1707 0H3.90244ZM3.90244 1.94444H32.1951V5.83333H3.90244C2.80488 5.83333 1.95122 4.98264 1.95122 3.88889C1.95122 2.79514 2.80488 1.94444 3.90244 1.94444ZM3.90244 7.77768C3.19361 7.77768 2.52668 7.58017 1.95122 7.24597V31.5971C1.95122 32.4212 2.58766 33.0555 3.41463 33.0555H36.5854C37.4123 33.0555 38.0488 32.4212 38.0488 31.5971V26.736H23.9024C22.5686 26.736 21.4634 25.6347 21.4634 24.3055V16.5277C21.4634 15.1985 22.5686 14.0971 23.9024 14.0971H38.0488V9.23601C38.0488 8.41191 37.4123 7.77768 36.5854 7.77768C33.5841 7.77768 30.2991 7.77672 26.8612 7.77572C19.4816 7.77356 11.3975 7.7712 3.90244 7.77768ZM23.9024 16.0417H38.0488V24.7917H23.9024C23.6147 24.7917 23.4146 24.5923 23.4146 24.3056V16.5278C23.4146 16.241 23.6147 16.0417 23.9024 16.0417ZM26.8293 20.4167C26.8293 19.3419 27.702 18.4722 28.7805 18.4722C29.859 18.4722 30.7317 19.3419 30.7317 20.4167C30.7317 21.4914 29.859 22.3611 28.7805 22.3611C27.702 22.3611 26.8293 21.4914 26.8293 20.4167Z" fill="#224CA5"/>
        </svg>
        <div class="payout-container__wallet-address-section">
            <p class="payout-container__wallet-address-section__label">{{label}}</p>
            <p><b class="payout-container__wallet-address-section__bold-text">{{walletAddress}}</b></p>
        </div>
        <a :href="'https://etherscan.io/address/' + walletAddress + '#tokentxns'" target="_blank" rel="noopener">
            <div class="payout-container__button"><b class="payout-container-button-label">View on Etherscan</b></div>
        </a>
    </div>
</template>

<script lang="ts">
import { Component, Prop, Vue } from 'vue-property-decorator';

@Component
export default class PayoutArea extends Vue {
    @Prop({default: ''})
    private readonly label: string;
    @Prop({default: ''})
    private readonly walletAddress: string;
}
</script>

<style scoped lang="scss">
    .payout-container {
        background-color: #FFFFFF;
        padding: 40px;
        display: flex;
        justify-content: flex-start;
        align-items: center;
        border: 1px solid #EAEAEA;
        border-radius: 12px;
        position: relative;

        &__image {
            margin-right: 40px;
        }

        &__wallet-address-section {
            height: 50px;
            display: flex;
            flex-direction: column;
            justify-content: space-between;

            &__label {
                font-size: 14px;
                color: #586C86;
            }

            &__bold-text {
                font-size: 18px;
                color: #535F77;
            }
        }

        &__button {
            font-size: 14px;
            width: 168px;
            height: 44px;
            display: flex;
            align-items: center;
            justify-content: center;
            background-color: #F4F6F9;
            border: 1px solid #E8E8E8;
            border-radius: 12px;
            position: absolute;
            top: 42px;
            right: 40px;
            color: #535F77;

            &:hover {
                background-color: #4D72B7;
                cursor: pointer;

                .payout-container-button-label {
                    color: #FFFFFF;
                }
            }
        }
    }
</style>
