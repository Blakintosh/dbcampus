<script lang="ts">
    import "/src/app.css";
	import SideNav from "../../../components/dashboard/sidebar/SideNav.svelte";
	import type { SoftwareProject } from "../../../util/models";
	import type { LayoutData } from "./$types";

	export let data: LayoutData;
</script>

<svelte:head>
	<title>{data.project.name} - Pitstop Dashboard</title>
	<meta name="description" content="Project overview dashboard.">
</svelte:head>

<main class="flex max-h-full h-[100vh] flex-col lg:flex-row">
	<SideNav projects={data.availableProjects}/>

	<!-- Manage Modal -->

	<slot/>
</main>