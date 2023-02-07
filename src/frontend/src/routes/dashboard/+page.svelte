<script lang="ts">
	import type { SoftwareProject } from "./models";
	import Project from "./_project/Project.svelte";
	import Sidebar from "./_sidebar/Sidebar.svelte";

	// Fetch data from BE
	let project: SoftwareProject = {
		name: "Placeholder Project",
		health: {
			status: "warning",
			issues: 2
		},
		survey: {
			date: new Date(),
			factors: [
				{
					name: "Factor 1",
					description: "Description 1",
					satisfaction: 0.87,
					status: "normal"
				},
				{
					name: "Factor 2",
					description: "Description 3",
					satisfaction: 0.62,
					status: "warning"
				}
			]
		}
	} satisfies SoftwareProject;
</script>
<main class="flex max-h-full h-[100vh]">
	<Sidebar />

	<Project {...project}/>
</main>