<script lang="ts">
	import { page } from "$app/stores";
	import { goto } from "$app/navigation";
	import type { SoftwareProjectSnippet } from "../../../util/models";

	export let projects: Array<SoftwareProjectSnippet> = [];

	let projectTarget: string = $page.data.project?.id ?? "none";

	$: urlEnd = $page.url.pathname.split("/").at(-1);


	const changeProject = () => {
		console.log(urlEnd);
		const destination = urlEnd !== undefined && ["health", "jira", "budget", "surveys"].includes(urlEnd)
			? `/dashboard/${projectTarget}/${urlEnd}`
			: `/dashboard/${projectTarget}`;
		console.log(destination);
		goto(destination);
	};
</script>

<div class="my-2 mb-2">
	<select bind:value={projectTarget} on:change={(e) => changeProject()}
		class="bg-slate-800 border border-slate-600 rounded-md p-2 text-sm font-medium w-full">
		{#each projects as project}
			<option value="{project.id}">
				{project.name}
			</option>
        {:else}
            <option value="none">
                No project selected
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