// Copyright (C) 2019 Storj Labs, Inc.
// See LICENSE for copying information.

<template>
    <div class="dialog">
        <p class="label dialog__make-default">Make Default</p>
        <p class="label dialog__delete">Delete</p>
    </div>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator';

@Component
export default class CardDialog extends Vue {}
</script>

<style scoped lang="scss">
    .dialog {
        position: absolute;
        top: 22px;
        right: -30px;
        display: flex;
        flex-direction: column;
        align-items: center;
        justify-content: flex-end;
        z-index: 100;
        background-image: url("../../../../static/images/payments/Dialog.svg");
        background-size: contain;
        width: 167px;
        height: 122px;
        cursor: initial;

        &__make-default {
            color: #61666B;
        }

        &__delete {
            color: #EB5757;
        }
    }

    .label {
        font-family: 'font_medium';
        font-size: 16px;
        margin: 0;
        height: 35%;
        text-align: center;
        cursor: pointer;

        &:hover {
            text-decoration: underline;
        }
    }
</style>
