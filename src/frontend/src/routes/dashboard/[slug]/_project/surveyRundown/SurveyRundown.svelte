<script lang="ts">
	import type { SurveyFactor } from "../../models";
	import SurveyItem from "./SurveyItem.svelte";

	export let date: Date;
	export let factors: Array<SurveyFactor>;
</script>

<div class="mb-4">
	Last Survey: <span class="ml-1 font-medium">{
		date.toLocaleDateString("en-GB", {
			year: "numeric",
			month: "long",
			day: "numeric"
		})
		}</span>
</div>
<ul>
	{#each factors as factor}
		<SurveyItem {...factor} />
	{/each}
</ul>

