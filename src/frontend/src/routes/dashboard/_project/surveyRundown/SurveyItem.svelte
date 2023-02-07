<script lang="ts">
	import RagCircle from "../common/RagCircle.svelte";
import type { SurveyFactor } from "./surveyModels";
	
	export let factor: SurveyFactor;
</script>

<li class="flex justify-between w-full items-center border-b bg-slate-100 px-4">
	{#if factor}
		<div class="text-lg font-medium py-2">{factor.factor}</div>
		<div class="flex flex-row items-center">
			<div class="text-xl font-medium">{
				factor.satisfactionPercent.toLocaleString(undefined,{
					style: 'percent', 
					minimumFractionDigits: 1
				})
			}
			</div>
			<div class="text-sm font-normal mx-2">satisfied</div>
			<div class="text-sm">
				<RagCircle status={factor.ragStatus}/>
			</div>
		</div>
	{/if}
</li>