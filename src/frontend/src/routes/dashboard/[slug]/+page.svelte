<script lang="ts">
	import { page } from "$app/stores";
	import BudgetRundown from "../../../components/dashboard/project/budgetRundown/BudgetRundown.svelte";
	import ProjectViewLayout from "../../../components/dashboard/project/ProjectViewLayout.svelte";
	import Button from "../../../components/common/SmallButton.svelte";
	import HeaderTile from "../../../components/dashboard/project/shared/HeaderTile.svelte";
	import Tile from "../../../components/dashboard/project/shared/Tile.svelte";
	import SurveyRundown from "../../../components/dashboard/project/surveyRundown/SurveyRundown.svelte";

    import { modalCategory, modalVisible } from "../../../util/stores";
	import { ProjectManageCategory } from '../../../util/models';
	import HealthTrafficLight from "../../../components/dashboard/project/health/HealthTrafficLight.svelte";

    const openModal = (category: ProjectManageCategory) => {
        modalVisible.set(true);
        modalCategory.set(category);
    }

	$: project = $page.data.project;
</script>

<ProjectViewLayout gridClass="grid-cols-3 md:grid-cols-4 lg:grid-cols-6">
	<HeaderTile projectName={project.name} categoryName="Overview" health={null}>
		<span slot="categoryIcon">
			<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 576 512"><!--! Font Awesome Pro 6.3.0 by @fontawesome - https://fontawesome.com License - https://fontawesome.com/license (Commercial License) Copyright 2023 Fonticons, Inc. --><path d="M302 240V16.6c0-9 7-16.6 16-16.6C441.7 0 542 100.3 542 224c0 9-7.6 16-16.6 16H302zM30 272C30 150.7 120.1 50.3 237 34.3c9.2-1.3 17 6.1 17 15.4V288L410.5 444.5c6.7 6.7 6.2 17.7-1.5 23.1C369.8 495.6 321.8 512 270 512C137.5 512 30 404.6 30 272zm526.4 16c9.3 0 16.6 7.8 15.4 17c-7.7 55.9-34.6 105.6-73.9 142.3c-6 5.6-15.4 5.2-21.2-.7L318 288H556.4z"/></svg>
		</span>
		<span slot="actionButtons" class="flex gap-8">
			<Button label="Update Project Information" on:click={() => openModal(ProjectManageCategory.General)}/>
			<Button label="Issue a Survey" on:click={() => openModal(ProjectManageCategory.ClientSurvey)}/>
			<Button label="Project Settings" type="warning" on:click={() => openModal(ProjectManageCategory.Settings)}/>
		</span>
	</HeaderTile>
    <Tile heading="Project Health" tileClass="col-span-2 md:col-span-2 lg:col-span-2" isDark>
        <div class="flex gap-8">
            <HealthTrafficLight status="danger" sizeClass="w-14 h-14 m-1.5"/>
            <div class="w-full">
                <div class="inline-block">
                    <h1 class="text-7xl">38.4<span class="text-4xl font-medium text-slate-200 mx-2">%</span></h1>
                    <h2 class="text-xl text-slate-200">&pm; <span class="text-2xl text-slate-100">5.2</span>%</h2>
                    <div class="bg-red-600 w-full h-1 my-2 rounded-sm"></div>
                    <h3 class="font-medium text-lg">Significant issues present.</h3>
                </div>
                <p>The latest assessment has identified multiple areas of issue which need to be addressed urgently in order to get this project back on track.</p>
            </div>
        </div>
    </Tile>
    <Tile heading="What can be improved?" tileClass="col-span-4" isDark>

    </Tile>
	<Tile heading="Project Timeline" tileClass="col-span-3 md:col-span-4 lg:col-span-6">

	</Tile>
	<Tile heading="Client Survey Snapshot" tileClass="col-span-3 md:col-span-4 lg:col-span-6 xl:col-span-3">
		<SurveyRundown {...project.surveys.client}/>

		<p class="flex justify-center mt-4">
			<Button label="Issue a Survey"/>
		</p>
	</Tile>
	<Tile heading="Project Budget" tileClass="col-span-3 md:col-span-4 lg:col-span-6 xl:col-span-3">
		<BudgetRundown />

		<p class="flex justify-center">
			<Button label="Update Spend or Budget"/>
		</p>
	</Tile>
</ProjectViewLayout>