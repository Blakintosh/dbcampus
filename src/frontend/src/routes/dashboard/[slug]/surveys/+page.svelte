<script lang="ts">
	import { page } from '$app/stores';
	import ProjectViewLayout from "../_project/ProjectViewLayout.svelte";
	import HeaderTile from "../_project/shared/HeaderTile.svelte";
	import Tile from "../_project/shared/Tile.svelte";
	import Button from "../../../../components/common/SmallButton.svelte";
	import SurveyResult from "../_project/surveyMain/SurveyResult.svelte";

	$: surveys = $page.data.project.surveys;
</script>

<ProjectViewLayout gridClass="grid-cols-6">
	<HeaderTile projectName={$page.data.project.name} categoryName="Surveys" health={surveys.health}>
		<span slot="categoryIcon">
			<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 448 512"><!--! Font Awesome Pro 6.3.0 by @fontawesome - https://fontawesome.com License - https://fontawesome.com/license (Commercial License) Copyright 2023 Fonticons, Inc. --><path d="M448 96c0-35.3-28.7-64-64-64L64 32C28.7 32 0 60.7 0 96L0 416c0 35.3 28.7 64 64 64l320 0c35.3 0 64-28.7 64-64l0-320zM256 160c0 17.7-14.3 32-32 32l-96 0c-17.7 0-32-14.3-32-32s14.3-32 32-32l96 0c17.7 0 32 14.3 32 32zm64 64c17.7 0 32 14.3 32 32s-14.3 32-32 32l-192 0c-17.7 0-32-14.3-32-32s14.3-32 32-32l192 0zM192 352c0 17.7-14.3 32-32 32l-32 0c-17.7 0-32-14.3-32-32s14.3-32 32-32l32 0c17.7 0 32 14.3 32 32z"/></svg>
		</span>
		<span slot="actionButtons">
			<Button label="Issue a Survey"/>
			<Button label="Project Settings" type="warning"/>
		</span>
	</HeaderTile>

	<Tile heading="Client Satisfaction" tileClass="col-span-6">
		<SurveyResult {...surveys.client}/>

		<p class="flex justify-center">
			<Button label="Issue a Client Survey"/>
		</p>
	</Tile>
	<Tile heading="Developer Satisfaction" tileClass="col-span-6">
		<SurveyResult {...surveys.team}/>

		<p class="flex justify-center">
			<Button label="Issue a Developer Survey"/>
		</p>
	</Tile>
</ProjectViewLayout>