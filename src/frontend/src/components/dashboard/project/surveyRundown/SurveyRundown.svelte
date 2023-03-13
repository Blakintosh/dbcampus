<script lang="ts">
	import type { SurveyFactor } from "../../../../util/models";
	import SurveyItem from "./SurveyItem.svelte";

	export let date: string;
    console.log("Oi "+date);
    $: jsDate = new Date(date);
	export let factors: Array<SurveyFactor>;
</script>

<div class="mb-4">
	Last Client Survey: <span class="ml-1 font-medium">{
		jsDate.toLocaleDateString("en-GB", {
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

