<script lang="ts">
    import "/src/app.css";
	import Sidebar from "./_sidebar/Sidebar.svelte";
	import type { SoftwareProject } from "../models";
	import type { LayoutData } from "./$types";

	export let data: LayoutData;
</script>

<svelte:head>
	<title>{data.project.name} - Pitstop Dashboard</title>
	<meta name="description" content="Project overview dashboard.">
</svelte:head>

<main class="flex max-h-full h-[100vh]">
	<Sidebar projects={data.availableProjects}/>

	<slot/>
</main>