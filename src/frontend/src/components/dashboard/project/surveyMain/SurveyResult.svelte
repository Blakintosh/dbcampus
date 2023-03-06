<script lang="ts">
	import type { SurveyFactor } from "../../../../util/models";
	import SurveyRag from "./SurveyRag.svelte";
	import SurveyResultRow from "./SurveyResultRow.svelte";

    export let date: Date;
	export let factors: Array<SurveyFactor>;
	export let suggestions: Array<string>;
</script>

<div>
    <h1 class="text-lg my-2">On {date.toLocaleDateString("en-gb")}, you asked, and they answered:</h1>

    <table class="w-full border-collapse h-full rounded-md overflow-hidden my-4">
        <thead>
            <tr class="text-slate-50 text-left">
                <th class="sticky top-0 p-1 px-2 bg-slate-500 font-semibold"></th>
                <th class="sticky top-0 p-1 px-2 bg-slate-500 font-semibold">Question</th>
                <th class="sticky top-0 p-1 px-2 bg-slate-500 font-semibold">Feedback</th>
            </tr>
        </thead>
        <tbody>
			{#each factors as factor}
				<SurveyResultRow {...factor} />
			{/each}
        </tbody>
    </table>

    <h1 class="text-lg my-2">Based on these survey results, here are some suggestions to try:</h1>
    <ul class="list-disc pl-6 mb-2">
		{#each suggestions as suggestion}
			<li class="px-2">{suggestion}</li>
		{:else}
			<li class="px-2">No suggestions to give.</li>
		{/each}
    </ul>
</div>