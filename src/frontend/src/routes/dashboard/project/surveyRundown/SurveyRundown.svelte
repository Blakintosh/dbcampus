<script lang="ts">
	import { SurveyFactor } from "./surveyModels";
	import SurveyItem from "./SurveyItem.svelte";

	export let lastReviewed: Date = new Date();
	export let factors: Array<SurveyFactor> = [
		new SurveyFactor("Factor 1", "Description 1", 0.87, "normal"),
		new SurveyFactor("Factor 2", "Description 2", 0.62, "warning"),
		new SurveyFactor("Factor 3", "Description 3", 0.23, "danger"),
	];
</script>

<div class="mb-4">
	Last Survey: <span class="ml-1 font-medium">{
		lastReviewed.toLocaleDateString("en-GB", {
			year: "numeric",
			month: "long",
			day: "numeric"
		})
		}</span>
</div>
<ul>
	{#each factors as factor}
		<SurveyItem {factor} />
	{/each}
</ul>

