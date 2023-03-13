<script lang="ts">
	import { page } from "$app/stores";
	import BudgetRundown from "../../../components/dashboard/project/budgetRundown/BudgetRundown.svelte";
	import ProjectViewLayout from "../../../components/dashboard/project/ProjectViewLayout.svelte";
	import Button from "../../../components/common/SmallButton.svelte";
	import HeaderTile from "../../../components/dashboard/project/shared/HeaderTile.svelte";
	import Tile from "../../../components/dashboard/project/shared/Tile.svelte";
	import SurveyRundown from "../../../components/dashboard/project/surveyRundown/SurveyRundown.svelte";

    import { modalCategory, modalVisible } from "../../../util/stores";
	import { ProjectManageCategory, type HealthInformation, type SoftwareProject } from '../../../util/models';
	import HealthTrafficLight from "../../../components/dashboard/project/health/HealthTrafficLight.svelte";
	import RagCircle from "../../../components/dashboard/project/common/RagCircle.svelte";

    const openModal = (category: ProjectManageCategory) => {
        modalVisible.set(true);
        modalCategory.set(category);
    }

	$: project = $page.data.project;

    let projectHealth = "unknown";
    let healthColourClass = "bg-grey-600";
    let healthHeading = "No data."
    $: {
        const health: HealthInformation = project.health;
        const percentageHealth = health.percentageHealth;

        if(percentageHealth <= 0.4) {
            projectHealth = "danger";
            healthColourClass = "bg-red-600";
            healthHeading = "Significant issues present.";
        } else if(percentageHealth <= 0.7) {
            projectHealth = "warning";
            healthColourClass = "bg-amber-600";
            healthHeading = "Issues present.";
        } else {
            projectHealth = "normal";
            healthColourClass = "bg-green-600";
            healthHeading = "Project is OK.";
        }
    }
</script>

<ProjectViewLayout gridClass="grid-cols-3 md:grid-cols-4 lg:grid-cols-6">
	<HeaderTile projectName={project.name} categoryName="Overview" health={null}>
		<span slot="categoryIcon">
			<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 576 512"><!--! Font Awesome Pro 6.3.0 by @fontawesome - https://fontawesome.com License - https://fontawesome.com/license (Commercial License) Copyright 2023 Fonticons, Inc. --><path d="M302 240V16.6c0-9 7-16.6 16-16.6C441.7 0 542 100.3 542 224c0 9-7.6 16-16.6 16H302zM30 272C30 150.7 120.1 50.3 237 34.3c9.2-1.3 17 6.1 17 15.4V288L410.5 444.5c6.7 6.7 6.2 17.7-1.5 23.1C369.8 495.6 321.8 512 270 512C137.5 512 30 404.6 30 272zm526.4 16c9.3 0 16.6 7.8 15.4 17c-7.7 55.9-34.6 105.6-73.9 142.3c-6 5.6-15.4 5.2-21.2-.7L318 288H556.4z"/></svg>
		</span>
		<span slot="actionButtons" class="flex gap-2 md:gap-8 flex-wrap justify-center">
			<Button label="Update Project Information" on:click={() => openModal(ProjectManageCategory.General)}/>
			<Button label="Issue a Survey" on:click={() => openModal(ProjectManageCategory.ClientSurvey)}/>
			<Button label="Project Settings" type="warning" on:click={() => openModal(ProjectManageCategory.Settings)}/>
		</span>
	</HeaderTile>
    <Tile heading="Project Health" tileClass="col-span-3 md:col-span-4 lg:col-span-2" isDark>
        <div class="flex gap-4 xl:gap-8 items-center">
            <HealthTrafficLight status={projectHealth} sizeClass="w-10 h-10 m-1 xl:w-14 xl:h-14 xl:m-1.5"/>
            <div class="w-full">
                <div class="inline-block">
                    <h1 class="text-5xl xl:text-7xl">{Math.round(project.health.percentageHealth * 1000) / 10}<span class="text-2xl xl:text-4xl font-medium text-slate-200 mx-2">%</span></h1>
                    <div class="{healthColourClass} w-full h-1 my-2 rounded-sm"></div>
                    <h3 class="font-medium text-md xl:text-lg">{healthHeading}</h3>
                </div>
                <p class="text-xs xl:text-md">{project.health.message}</p>
            </div>
        </div>
    </Tile>
    <Tile heading="What can be improved?" tileClass="col-span-3 md:col-span-4" isDark>
		<h3 class="text-sm xl:text-md font-medium my-1 xl:my-2">5 issues and suggestions.</h3>
		<table class="w-full border-collapse rounded-md overflow-hidden my-4">
			<thead>
				<tr class="text-slate-50 text-left">
					<th class="p-1 px-2 bg-slate-500 font-semibold w-8"></th>
					<th class="p-1 px-2 bg-slate-500 font-semibold">Category</th>
					<th class="p-1 px-2 bg-slate-500 font-semibold">Suggestion</th>
				</tr>
			</thead>
			<tbody>
				<tr class="bg-slate-700">
					<td class="p-2 border border-slate-600">
						<RagCircle status="r"/>
					</td>
					<td class="p-2 border border-slate-600">
						Surveys
					</td>
					<td class="p-2 border text-xs md:text-sm border-slate-600">
						Client survey feedback has shown critically low satisfaction for multiple areas. Try to address this by engaging in dialog with the customer.
					</td>
				</tr>
			</tbody>
		</table>
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
		<BudgetRundown {...project.budget} />

		<p class="flex justify-center">
			<Button label="Update Spend or Budget"/>
		</p>
	</Tile>
</ProjectViewLayout>