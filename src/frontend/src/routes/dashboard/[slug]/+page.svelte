<script lang="ts">
	import { page } from "$app/stores";
	import type { HealthSummary, SoftwareProject, SurveyFactor, SurveySummary } from "../models";
	import BudgetRundown from "./_project/budgetRundown/BudgetRundown.svelte";
	import HealthRag from "./_project/health/HealthRag.svelte";
	import Participants from "./_project/participants/Participants.svelte";
	import ProjectViewLayout from "./_project/ProjectViewLayout.svelte";
	import Button from "./_project/shared/Button.svelte";
	import HeaderTile from "./_project/shared/HeaderTile.svelte";
	import Tile from "./_project/shared/Tile.svelte";
	import SurveyRundown from "./_project/surveyRundown/SurveyRundown.svelte";

	$: project = $page.data.project;
</script>

<ProjectViewLayout gridClass="grid-cols-6">
	<HeaderTile projectName={project.name} categoryName="Overview">
		<span slot="categoryIcon">
			<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 576 512"><!--! Font Awesome Pro 6.3.0 by @fontawesome - https://fontawesome.com License - https://fontawesome.com/license (Commercial License) Copyright 2023 Fonticons, Inc. --><path d="M302 240V16.6c0-9 7-16.6 16-16.6C441.7 0 542 100.3 542 224c0 9-7.6 16-16.6 16H302zM30 272C30 150.7 120.1 50.3 237 34.3c9.2-1.3 17 6.1 17 15.4V288L410.5 444.5c6.7 6.7 6.2 17.7-1.5 23.1C369.8 495.6 321.8 512 270 512C137.5 512 30 404.6 30 272zm526.4 16c9.3 0 16.6 7.8 15.4 17c-7.7 55.9-34.6 105.6-73.9 142.3c-6 5.6-15.4 5.2-21.2-.7L318 288H556.4z"/></svg>
		</span>
		<span slot="actionButtons">
			<Button label="Update Project Information"/>
			<Button label="Issue a Survey"/>
			<Button label="Project Settings" type="warning"/>
		</span>
	</HeaderTile>
	<Tile heading="Project Timeline" tileClass="col-span-6">

	</Tile>
	<Tile heading="Client Survey Snapshot" tileClass="col-span-3">
		<SurveyRundown {...project.surveys.client}/>

		<p class="flex justify-center mt-4">
			<Button label="Issue a Survey"/>
		</p>
	</Tile>
	<Tile heading="Project Budget" tileClass="col-span-3">
		<BudgetRundown />

		<p class="flex justify-center">
			<Button label="Update Spend or Budget"/>
		</p>
	</Tile>
</ProjectViewLayout>