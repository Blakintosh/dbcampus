<script lang="ts">
    import "/src/app.css";
	import SideNav from "../../../components/dashboard/sidebar/SideNav.svelte";
	import type { SoftwareProject } from "../../../util/models";
	import type { LayoutData } from "./$types";
	import ManageProjectModal from "../../../components/dashboard/project/manage/ManageProjectModal.svelte";

	export let data: LayoutData;

    console.log("test" + data.availableProjects);
</script>

<svelte:head>
	<title>{data.project.name} - Pitstop Dashboard</title>
	<meta name="description" content="Project overview dashboard.">
</svelte:head>

<ManageProjectModal/>

<main class="flex max-h-full h-[100vh] flex-col lg:flex-row">
	<SideNav projects={data.availableProjects}/>

	<!-- Manage Modal -->

    <slot/>
</main>