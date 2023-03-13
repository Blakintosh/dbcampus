<script lang="ts">
	import type { SvelteComponent } from "svelte";
	import { ProjectManageCategory } from "../../../../util/models";
	import { modalCategory, modalVisible } from "../../../../util/stores";
	import SmallButton from "../../../common/SmallButton.svelte";
	import ProjectManageSelect from "../../sidebar/ProjectManageSelect.svelte";
	import UnsavedChanges from "./components/UnsavedChanges.svelte";
	import ManageModalTab from "./ManageModalTab.svelte";
	import BudgetProjectSettings from "./sections/BudgetProjectSettings.svelte";
	import ClientSurveyProjectAction from "./sections/ClientSurveyProjectAction.svelte";
	import GeneralProjectSettings from "./sections/GeneralProjectSettings.svelte";
	import ManageProjectSettings from "./sections/ManageProjectSettings.svelte";
	import TeamSurveyProjectAction from "./sections/TeamSurveyProjectAction.svelte";

    type ProjectManageTab = {
        category: ProjectManageCategory;
        name: string;
        component: any;
    };

	export let categories: Map<ProjectManageCategory, ProjectManageTab> = new Map<ProjectManageCategory, ProjectManageTab>([
            [
                ProjectManageCategory.General,
                {
                    category: ProjectManageCategory.General,
                    name: "General",
                    component: GeneralProjectSettings
                }
            ],
            [
                ProjectManageCategory.Budget,
                {
                    category: ProjectManageCategory.Budget,
                    name: "Budget",
                    component: BudgetProjectSettings
                }
            ],
			[
				ProjectManageCategory.ClientSurvey,
				{
					category: ProjectManageCategory.ClientSurvey,
					name: "Issue a Client Survey",
					component: ClientSurveyProjectAction
				}
			],
			[
				ProjectManageCategory.TeamSurvey,
				{
					category: ProjectManageCategory.TeamSurvey,
					name: "Issue a Developer Survey",
					component: TeamSurveyProjectAction
				}
			],
            [
                ProjectManageCategory.Settings,
                {
                    category: ProjectManageCategory.Settings,
                    name: "Project Settings",
                    component: ManageProjectSettings
                }
            ]
        ]);

    const closeModal = () => {
        modalVisible.set(false);
    }

    const goToCategory = (category: ProjectManageCategory) => {
        modalCategory.set(category);
    }

    $: categoryData = categories.get($modalCategory);
</script>

<!-- svelte-ignore a11y-click-events-have-key-events -->
<div class="fixed w-full h-full bg-black/60 grid z-50 backdrop-blur-sm overflow-hidden" class:block={$modalVisible} class:hidden={!$modalVisible} on:click={closeModal}>
</div>
<div class="fixed w-[80%] h-[75%] translate-x-[-50%] left-[50%] top-[50%] translate-y-[-50%] z-50 bg-slate-800 border border-slate-700 text-slate-50 shadow-xl rounded-lg flex items-stretch overflow-hidden"
    class:block={$modalVisible} class:hidden={!$modalVisible}>
    <button class="absolute z-50 w-4 h-4 right-4 top-4 fill-slate-300/80 hover:fill-red-500" on:click={closeModal}>
        <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 320 512"><!--! Font Awesome Pro 6.3.0 by @fontawesome - https://fontawesome.com License - https://fontawesome.com/license (Commercial License) Copyright 2023 Fonticons, Inc. --><path d="M310.6 150.6c12.5-12.5 12.5-32.8 0-45.3s-32.8-12.5-45.3 0L160 210.7 54.6 105.4c-12.5-12.5-32.8-12.5-45.3 0s-12.5 32.8 0 45.3L114.7 256 9.4 361.4c-12.5 12.5-12.5 32.8 0 45.3s32.8 12.5 45.3 0L160 301.3 265.4 406.6c12.5 12.5 32.8 12.5 45.3 0s12.5-32.8 0-45.3L205.3 256 310.6 150.6z"/></svg>
    </button>
    <div class="w-72 bg-slate-800 p-4 flex flex-col gap-4">
        <div class="border-b border-slate-600 py-4">
            <h1 class="font-medium text-xs">Manage</h1>
            <h2 class="text-xl font-serif">Foobar</h2>
        </div>
        <div class="flex flex-col justify-between grow">
            <div>
                <div class="border-b border-slate-600">
                    <h1 class="font-medium text-xs mb-2">Settings</h1>
                    <ul class="w-full">
                        <ManageModalTab label="General" selected={$modalCategory === ProjectManageCategory.General} on:click={() => goToCategory(ProjectManageCategory.General)}>
                            <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 576 512" class="w-3 h-3"><!--! Font Awesome Pro 6.3.0 by @fontawesome - https://fontawesome.com License - https://fontawesome.com/license (Commercial License) Copyright 2023 Fonticons, Inc. --><path d="M575.8 255.5c0 18-15 32.1-32 32.1h-32l.7 160.2c0 2.7-.2 5.4-.5 8.1V472c0 22.1-17.9 40-40 40H456c-1.1 0-2.2 0-3.3-.1c-1.4 .1-2.8 .1-4.2 .1H416 392c-22.1 0-40-17.9-40-40V448 384c0-17.7-14.3-32-32-32H256c-17.7 0-32 14.3-32 32v64 24c0 22.1-17.9 40-40 40H160 128.1c-1.5 0-3-.1-4.5-.2c-1.2 .1-2.4 .2-3.6 .2H104c-22.1 0-40-17.9-40-40V360c0-.9 0-1.9 .1-2.8V287.6H32c-18 0-32-14-32-32.1c0-9 3-17 10-24L266.4 8c7-7 15-8 22-8s15 2 21 7L564.8 231.5c8 7 12 15 11 24z"/></svg>
                        </ManageModalTab>
                        <ManageModalTab label="Budget" selected={$modalCategory === ProjectManageCategory.Budget} on:click={() => goToCategory(ProjectManageCategory.Budget)}>
                            <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 320 512" class="w-3 h-3"><!--! Font Awesome Pro 6.3.0 by @fontawesome - https://fontawesome.com License - https://fontawesome.com/license (Commercial License) Copyright 2023 Fonticons, Inc. --><path d="M146 0c17.7 0 32 14.3 32 32V67.7c1.6 .2 3.1 .4 4.7 .7c.4 .1 .7 .1 1.1 .2l48 8.8c17.4 3.2 28.9 19.9 25.7 37.2s-19.9 28.9-37.2 25.7l-47.5-8.7c-31.3-4.6-58.9-1.5-78.3 6.2s-27.2 18.3-29 28.1c-2 10.7-.5 16.7 1.2 20.4c1.8 3.9 5.5 8.3 12.8 13.2c16.3 10.7 41.3 17.7 73.7 26.3l2.9 .8c28.6 7.6 63.6 16.8 89.6 33.8c14.2 9.3 27.6 21.9 35.9 39.5c8.5 17.9 10.3 37.9 6.4 59.2c-6.9 38-33.1 63.4-65.6 76.7c-13.7 5.6-28.6 9.2-44.4 11V480c0 17.7-14.3 32-32 32s-32-14.3-32-32V445.1c-.4-.1-.9-.1-1.3-.2l-.2 0 0 0c-24.4-3.8-64.5-14.3-91.5-26.3C4.9 411.4-2.4 392.5 4.8 376.3s26.1-23.4 42.2-16.2c20.9 9.3 55.3 18.5 75.2 21.6c31.9 4.7 58.2 2 76-5.3c16.9-6.9 24.6-16.9 26.8-28.9c1.9-10.6 .4-16.7-1.3-20.4c-1.9-4-5.6-8.4-13-13.3c-16.4-10.7-41.5-17.7-74-26.3l-2.8-.7 0 0C105.4 279.3 70.4 270 44.4 253c-14.2-9.3-27.5-22-35.8-39.6C.3 195.4-1.4 175.4 2.5 154.1C9.7 116 38.3 91.2 70.8 78.3c13.3-5.3 27.9-8.9 43.2-11V32c0-17.7 14.3-32 32-32z"/></svg>
                        </ManageModalTab>
                    </ul>
                </div>
                <div class="border-b border-slate-600 mt-4">
                    <h1 class="font-medium text-xs mb-2">Actions</h1>
                    <ul class="w-full">
                        <ManageModalTab label="Issue Client Survey" selected={$modalCategory === ProjectManageCategory.ClientSurvey} on:click={() => goToCategory(ProjectManageCategory.ClientSurvey)}>
                            <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 448 512" class="w-3 h-3"><!--! Font Awesome Pro 6.3.0 by @fontawesome - https://fontawesome.com License - https://fontawesome.com/license (Commercial License) Copyright 2023 Fonticons, Inc. --><path d="M224 256A128 128 0 1 0 224 0a128 128 0 1 0 0 256zm-45.7 48C79.8 304 0 383.8 0 482.3C0 498.7 13.3 512 29.7 512H418.3c16.4 0 29.7-13.3 29.7-29.7C448 383.8 368.2 304 269.7 304H178.3z"/></svg>
                        </ManageModalTab>
                        <ManageModalTab label="Issue Developer Survey" selected={$modalCategory === ProjectManageCategory.TeamSurvey} on:click={() => goToCategory(ProjectManageCategory.TeamSurvey)}>
                            <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 640 512" class="w-3 h-3"><!--! Font Awesome Pro 6.3.0 by @fontawesome - https://fontawesome.com License - https://fontawesome.com/license (Commercial License) Copyright 2023 Fonticons, Inc. --><path d="M392.8 1.2c-17-4.9-34.7 5-39.6 22l-128 448c-4.9 17 5 34.7 22 39.6s34.7-5 39.6-22l128-448c4.9-17-5-34.7-22-39.6zm80.6 120.1c-12.5 12.5-12.5 32.8 0 45.3L562.7 256l-89.4 89.4c-12.5 12.5-12.5 32.8 0 45.3s32.8 12.5 45.3 0l112-112c12.5-12.5 12.5-32.8 0-45.3l-112-112c-12.5-12.5-32.8-12.5-45.3 0zm-306.7 0c-12.5-12.5-32.8-12.5-45.3 0l-112 112c-12.5 12.5-12.5 32.8 0 45.3l112 112c12.5 12.5 32.8 12.5 45.3 0s12.5-32.8 0-45.3L77.3 256l89.4-89.4c12.5-12.5 12.5-32.8 0-45.3z"/></svg>
                        </ManageModalTab>
                    </ul>
                </div>
            </div>
            <div class="border-b border-slate-600">
                <h1 class="font-medium text-xs mb-2">Danger Zone</h1>
                <ul class="w-full">
                    <ManageModalTab label="Project Settings" selected={$modalCategory === ProjectManageCategory.Settings} on:click={() => goToCategory(ProjectManageCategory.Settings)}>
                        <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 512 512" class="w-3 h-3"><!--! Font Awesome Pro 6.3.0 by @fontawesome - https://fontawesome.com License - https://fontawesome.com/license (Commercial License) Copyright 2023 Fonticons, Inc. --><path d="M481.9 166.6c3.2 8.7 .5 18.4-6.4 24.6l-30.9 28.1c-7.7 7.1-11.4 17.5-10.9 27.9c.1 2.9 .2 5.8 .2 8.8s-.1 5.9-.2 8.8c-.5 10.5 3.1 20.9 10.9 27.9l30.9 28.1c6.9 6.2 9.6 15.9 6.4 24.6c-4.4 11.9-9.7 23.3-15.8 34.3l-4.7 8.1c-6.6 11-14 21.4-22.1 31.2c-5.9 7.2-15.7 9.6-24.5 6.8l-39.7-12.6c-10-3.2-20.8-1.1-29.7 4.6c-4.9 3.1-9.9 6.1-15.1 8.7c-9.3 4.8-16.5 13.2-18.8 23.4l-8.9 40.7c-2 9.1-9 16.3-18.2 17.8c-13.8 2.3-28 3.5-42.5 3.5s-28.7-1.2-42.5-3.5c-9.2-1.5-16.2-8.7-18.2-17.8l-8.9-40.7c-2.2-10.2-9.5-18.6-18.8-23.4c-5.2-2.7-10.2-5.6-15.1-8.7c-8.8-5.7-19.7-7.8-29.7-4.6L69.1 425.9c-8.8 2.8-18.6 .3-24.5-6.8c-8.1-9.8-15.5-20.2-22.1-31.2l-4.7-8.1c-6.1-11-11.4-22.4-15.8-34.3c-3.2-8.7-.5-18.4 6.4-24.6l30.9-28.1c7.7-7.1 11.4-17.5 10.9-27.9c-.1-2.9-.2-5.8-.2-8.8s.1-5.9 .2-8.8c.5-10.5-3.1-20.9-10.9-27.9L8.4 191.2c-6.9-6.2-9.6-15.9-6.4-24.6c4.4-11.9 9.7-23.3 15.8-34.3l4.7-8.1c6.6-11 14-21.4 22.1-31.2c5.9-7.2 15.7-9.6 24.5-6.8l39.7 12.6c10 3.2 20.8 1.1 29.7-4.6c4.9-3.1 9.9-6.1 15.1-8.7c9.3-4.8 16.5-13.2 18.8-23.4l8.9-40.7c2-9.1 9-16.3 18.2-17.8C213.3 1.2 227.5 0 242 0s28.7 1.2 42.5 3.5c9.2 1.5 16.2 8.7 18.2 17.8l8.9 40.7c2.2 10.2 9.4 18.6 18.8 23.4c5.2 2.7 10.2 5.6 15.1 8.7c8.8 5.7 19.7 7.7 29.7 4.6l39.7-12.6c8.8-2.8 18.6-.3 24.5 6.8c8.1 9.8 15.5 20.2 22.1 31.2l4.7 8.1c6.1 11 11.4 22.4 15.8 34.3zM242 336a80 80 0 1 0 0-160 80 80 0 1 0 0 160z"/></svg>
                    </ManageModalTab>
                </ul>
            </div>
        </div>
    </div>
    <div class="p-6 bg-slate-700/50 grow relative flex flex-col">
        <h1 class="font-normal text-3xl mb-4">{categoryData?.name}</h1>
        <div class="text-sm overflow-auto grow">
            <svelte:component this={categoryData?.component} />
        </div>
		<UnsavedChanges enabled={false}/>
    </div>
</div>