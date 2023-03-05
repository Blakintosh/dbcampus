<script lang="ts">
	import Tile from "./Tile.svelte";
	import HealthRag from "../health/HealthRag.svelte";
	import type { HealthSummary } from "../../../../util/models";

    export let projectName: string = "";
    export let categoryName: string = "";
	export let health: HealthSummary | null = {
		status: "unknown",
		message: "",
		issues: 0
	};
</script>

<div class="flex col-span-3 md:col-span-4 lg:col-span-6 h-min">
	<Tile tileClass="flex-grow">
		<div class="flex justify-between flex-col h-full">
			<div>
				<div class="flex items-center gap-4 w-full mb-2">
					<span class="w-6 fill-amber-400">
						<slot name="categoryIcon"/>
					</span>
					<span class="text-lg font-medium text-red-600">
						{categoryName}
					</span>
				</div>
				<h1 class="text-3xl font-serif mb-4">{projectName}</h1>
			</div>
				
			<div class="flex justify-between items-center">
				<p class="text-md font-semibold mx-2 uppercase">
					Actions
				</p>
				<p class="flex-grow flex justify-center gap-8">
					<slot name="actionButtons" class="flex gap-8"/>
				</p>
			</div>
		</div>
	</Tile>
	{#if health}
		<div class="w-96 flex-shrink-0">
			<Tile isDark>
				<HealthRag {...health}/>
			</Tile>
		</div>
	{/if}
</div>