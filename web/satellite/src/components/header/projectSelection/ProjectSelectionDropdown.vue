// Copyright (C) 2019 Storj Labs, Inc.
// See LICENSE for copying information.

<template>
    <div class="project-selection-choice-container" id="projectDropdown">
        <div class="project-selection-overflow-container">
            <!-- loop for rendering projects -->
            <!-- TODO: add selection logic onclick -->
            <div class="project-selection-overflow-container__project-choice" @click="onProjectSelected(project.id)" v-for="project in projects" :key="project.id" >
                <div class="project-selection-overflow-container__project-choice__mark-container">
                    <svg class="project-selection-overflow-container__project-choice__mark-container__image" v-if="project.isSelected" width="15" height="13" viewBox="0 0 15 13" fill="none" xmlns="http://www.w3.org/2000/svg">
                        <path d="M14.0928 3.02746C14.6603 2.4239 14.631 1.4746 14.0275 0.907152C13.4239 0.339699 12.4746 0.368972 11.9072 0.972536L14.0928 3.02746ZM4.53846 11L3.44613 12.028C3.72968 12.3293 4.12509 12.5001 4.53884 12.5C4.95258 12.4999 5.34791 12.3289 5.63131 12.0275L4.53846 11ZM3.09234 7.27469C2.52458 6.67141 1.57527 6.64261 0.971991 7.21036C0.36871 7.77812 0.339911 8.72743 0.907664 9.33071L3.09234 7.27469ZM11.9072 0.972536L3.44561 9.97254L5.63131 12.0275L14.0928 3.02746L11.9072 0.972536ZM5.6308 9.97199L3.09234 7.27469L0.907664 9.33071L3.44613 12.028L5.6308 9.97199Z" fill="#2683FF"/>
                    </svg>
                </div>
                <h2 class="project-selection-overflow-container__project-choice__unselected" :class="{'selected': project.isSelected}">{{project.name}}</h2>
            </div>
        </div>
    </div>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator';

import { BUCKET_ACTIONS } from '@/store/modules/buckets';
import { PROJECTS_ACTIONS } from '@/store/modules/projects';
import { PROJECT_USAGE_ACTIONS } from '@/store/modules/usage';
import { Project } from '@/types/projects';
import {
    API_KEYS_ACTIONS,
    APP_STATE_ACTIONS,
    NOTIFICATION_ACTIONS,
    PM_ACTIONS,
} from '@/utils/constants/actionNames';

@Component
export default class ProjectSelectionDropdown extends Vue {
    private FIRST_PAGE = 1;

    public async onProjectSelected(projectID: string): Promise<void> {
        this.$store.dispatch(PROJECTS_ACTIONS.SELECT, projectID);
        this.$store.dispatch(APP_STATE_ACTIONS.TOGGLE_PROJECTS);
        this.$store.dispatch(PM_ACTIONS.SET_SEARCH_QUERY, '');

        try {
            await this.$store.dispatch(PROJECT_USAGE_ACTIONS.FETCH_CURRENT_ROLLUP);
        } catch (error) {
            await this.$store.dispatch(NOTIFICATION_ACTIONS.ERROR, `Unable to fetch project usage. ${error.message}`);
        }

        try {
            await this.$store.dispatch(PM_ACTIONS.FETCH, this.FIRST_PAGE);
        } catch (error) {
            this.$store.dispatch(NOTIFICATION_ACTIONS.ERROR, `Unable to fetch project members. ${error.message}`);
        }

        try {
            await this.$store.dispatch(API_KEYS_ACTIONS.FETCH, this.FIRST_PAGE);
        } catch (error) {
            this.$store.dispatch(NOTIFICATION_ACTIONS.ERROR, `Unable to fetch api keys. ${error.message}`);
        }

        try {
            await this.$store.dispatch(BUCKET_ACTIONS.FETCH, this.FIRST_PAGE);
        } catch (error) {
            this.$store.dispatch(NOTIFICATION_ACTIONS.ERROR, 'Unable to fetch buckets: ' + error.message);
        }
    }

    public get projects(): Project[] {
        return this.$store.getters.projects;
    }
}
</script>

<style scoped lang="scss">
    .project-selection-choice-container {
        position: absolute;
        top: 9vh;
        left: -5px;
        border-radius: 4px;
        padding: 10px 0 10px 0;
        box-shadow: 0 4px rgba(231, 232, 238, 0.6);
        background-color: #FFFFFF;
        z-index: 1120;
    }

    .project-selection-overflow-container {
        position: relative;
        width: 226px;
        overflow-y: auto;
        overflow-x: hidden;
        height: auto;
        max-height: 240px;
        background-color: #FFFFFF;
        font-family: 'font_regular';

        &__project-choice {
            display: flex;
            flex-direction: row;
            align-items: center;
            justify-content: flex-start;
            padding-left: 20px;
            padding-right: 20px;

            &__unselected {
                margin-left: 20px;
                font-size: 14px;
                line-height: 20px;
                color: #354049;
            }

            &:hover {
                background-color: #F2F2F6;
            }

            &__mark-container {
                width: 10px;

                &__image {
                    object-fit: cover;
                }
            }
        }
    }

    .selected {
        font-family: 'font_bold';
    }

    /* width */
    ::-webkit-scrollbar {
        width: 4px;
    }

    /* Track */
    ::-webkit-scrollbar-track {
        box-shadow: inset 0 0 5px #fff;
    }

    /* Handle */
    ::-webkit-scrollbar-thumb {
        background: #AFB7C1;
        border-radius: 6px;
        height: 5px;
    }

    @media screen and (max-width: 1024px) {
        .project-selection-choice-container {
            top: 50px;
        }
    }
</style>
