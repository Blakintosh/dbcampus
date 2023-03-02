<script lang="ts">
	import SurveyRag from "./SurveyRag.svelte";

	export let name: string;
	export let question: string;
	export let satisfaction: number;

	let status = "r";
	$: {
		if(satisfaction >= 0.8) {
			status = "g";
		} else if(satisfaction >= 0.6) {
			status = "a";
		} else if(satisfaction >= 0.4) {
			status = "A";
		} else if(satisfaction >= 0.2) {
			status = "r";
		} else {
			status = "R";
		}
	}
</script>

<tr class="border-t border-slate-200" 
	class:bg-red-200="{status === 'R'}"
	class:bg-red-100="{status === 'r'}"
	class:bg-amber-100="{status === 'A'}"
	class:bg-amber-50="{status === 'a'}"
	class:bg-slate-100="{status === 'g'}">
	<td class="font-medium text-lg p-2">
		{name}
	</td>
	<td class="text-md p-2 border-l border-r border-slate-200">
		{question}
	</td>
	<td class="font-medium p-2">
		<SurveyRag satisfaction={satisfaction} />
	</td>
</tr>