<script lang="ts">
    import "/src/app.css";
	import SideNav from "../../components/dashboard/sidebar/SideNav.svelte";
	import ProjectViewLayout from "../../components/dashboard/project/ProjectViewLayout.svelte";

</script>

<svelte:head>
	<title>Select a project - Pitstop Dashboard</title>
	<meta name="description" content="Project overview dashboard.">
</svelte:head>

<main class="flex max-h-full h-[100vh] flex-col lg:flex-row">
	<SideNav projects={undefined}/>

	<!-- Manage Modal -->

    <ProjectViewLayout gridClass="grid-cols-3 md:grid-cols-4 lg:grid-cols-6">
        <h1 class="col-span-6 text-3xl p-2">No project selected. That's to come...</h1>
    </ProjectViewLayout>

    <!-- svelte-ignore a11y-click-events-have-key-events -->
    <div class="fixed w-full h-full bg-black/60 grid z-50 backdrop-blur-sm overflow-hidden">
    </div>
    <div class="fixed w-[30%] h-[60%] translate-x-[-50%] left-[50%] top-[50%] translate-y-[-50%] z-50 bg-slate-200 border border-slate-300 text-black shadow-xl rounded-lg flex items-stretch justify-center overflow-hidden p-8">
        <h1 class="text-md font-semibold text-center">Select a Project</h1>
    </div>
</main>