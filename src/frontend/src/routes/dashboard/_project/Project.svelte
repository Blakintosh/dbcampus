<script lang="ts">
	import type { HealthSummary, SoftwareProject, SurveyFactor, SurveySummary } from "../models";
	import HealthRag from "./health/HealthRag.svelte";
	import Button from "./shared/Button.svelte";
	import Tile from "./shared/Tile.svelte";
	import SurveyRundown from "./surveyRundown/SurveyRundown.svelte";

	export let name: string;
	export let health: HealthSummary;
	export let survey: SurveySummary;
</script>

<article class="p-4 shadow-inner flex-grow bg-slate-100 flex">
	<section class="flex-grow">
		<Tile>
			<h1 class="text-3xl font-serif mb-4">{name}</h1>
			
			<div class="flex justify-between items-center">
				<p class="text-md mx-4">
					Actions
				</p>
				<p class="flex-grow flex justify-center">
					<Button label="Update Project Information"/>
					<Button label="Issue a Survey"/>
					<Button label="Project Settings" type="warning"/>
				</p>
			</div>
		</Tile>
		<Tile heading="Project Timeline">

		</Tile>
		<div class="flex">
			<div class="w-[55%] flex-shrink-0">
				<Tile heading="Surveys Rundown">
					<SurveyRundown {...survey}/>

					<p class="flex justify-center mt-4">
						<Button label="Issue a Survey"/>
					</p>
				</Tile>
			</div>
			<div class="flex-grow">
				<Tile heading="Project Budget">
	
				</Tile>
			</div>
		</div>
	</section>

	<section class="w-96 flex-shrink-0">
		<Tile heading="Project Health" isDark>
			<HealthRag status={health.status}/>

			<p class="my-2 text-center text-sm">
				{health.issues} issues and suggestions.
			</p>

			<p class="flex justify-center">
				<Button label="View Issues"/>
			</p>

		</Tile>
		<Tile heading="Participants" isDark>

		</Tile>
	</section>
</article>