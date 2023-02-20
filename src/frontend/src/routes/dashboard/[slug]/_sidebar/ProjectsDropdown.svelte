<script lang="ts">
	import type { SoftwareProject } from "../../models";

	export let projects: Array<SoftwareProject> = [];
</script>

<div class="my-2 mb-8">
	<select class="bg-slate-800 border border-slate-600 rounded-md p-2 text-sm font-medium w-full">
		{#each projects as project}
			<option value="{project.name}">
				{project.name}
			</option>
		{/each}
	</select>
</div>
<!-- 
<ul class="mx-4  border border-slate-900 bg-slate-800 overflow-hidden rounded-md h-[60vh]">
	{#each projects as project}
		<li class="px-2 py-1 text-md text-slate-200 hover:text-slate-50 hover:bg-slate-9+00">
			<a href="/dashboard/{project.name}">
				{project.name}
			</a>
		</li>
	{/each}
</ul> -->