<script lang="ts">
    import "/src/app.css";
	import SideNav from "../../components/dashboard/sidebar/SideNav.svelte";
	import ProjectViewLayout from "../../components/dashboard/project/ProjectViewLayout.svelte";
	import { error, redirect } from "@sveltejs/kit";
	import type { SoftwareProjectSnippet } from "../../util/models";
	import PageHeading from "../../components/common/PageHeading.svelte";
	import { goto } from "$app/navigation";
	import ActionButton from "../../components/common/ActionButton.svelte";

    let projects: Array<SoftwareProjectSnippet> = [];
    const getProjects = async () => {
        const availableProjectsResponse = await fetch("/api/dashboard/getProjects", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
        });

        if(availableProjectsResponse.status === 401) {
            goto(301, "/auth/login");
        } else if(availableProjectsResponse.status === 500) {
            throw error(500, "Unable to contact the backend. Please try again later.");
        }

        const value = await availableProjectsResponse.json();

        projects = value.data;

        console.log(projects.length);
        if(projects.length === 0) {
            goto("/dashboard/new");
        }
    }
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
    <div class="fixed w-[30%] h-[60%] translate-x-[-50%] left-[50%] top-[50%] translate-y-[-50%] z-50 bg-slate-200 border border-slate-300 text-black shadow-xl rounded-lg flex flex-col items-center overflow-hidden">
        <div class="px-8 py-3 border-b border-slate-300 bg-slate-100 shrink-0 w-full">
            <PageHeading section="Dashboard" title="Select a Project"/>
        </div>

        <div class="p-4">
            <h3>Select a project from the list below to continue to the <span class="font-medium">Dashboard</span>.</h3>
    
            <table class="w-[100%] border-t border-slate-300 my-8 overflow-hidden rounded-md">
                <thead>
                    <tr class="bg-slate-600 text-slate-100 text-xs text-left rounded-t-lg">
                        <th class="font-semibold px-2 py-1 w-16">Code</th>
                        <th class="font-semibold px-2 py-1">Name</th>
                    </tr>
                </thead>
                <tbody>
                    {#await getProjects()}
                        <tr class="bg-slate-300/50 border-b border-slate-300">
                            <td class="text-sm p-2">
                                <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 512 512" class="fill-red-500 w-4 h-4 animate-spin">
                                    <!--! Font Awesome Pro 6.3.0 by @fontawesome - https://fontawesome.com License - https://fontawesome.com/license (Commercial License) Copyright 2023 Fonticons, Inc. -->
                                    <path
                                        d="M304 48c0-26.5-21.5-48-48-48s-48 21.5-48 48s21.5 48 48 48s48-21.5 48-48zm0 416c0-26.5-21.5-48-48-48s-48 21.5-48 48s21.5 48 48 48s48-21.5 48-48zM48 304c26.5 0 48-21.5 48-48s-21.5-48-48-48s-48 21.5-48 48s21.5 48 48 48zm464-48c0-26.5-21.5-48-48-48s-48 21.5-48 48s21.5 48 48 48s48-21.5 48-48zM142.9 437c18.7-18.7 18.7-49.1 0-67.9s-49.1-18.7-67.9 0s-18.7 49.1 0 67.9s49.1 18.7 67.9 0zm0-294.2c18.7-18.7 18.7-49.1 0-67.9S93.7 56.2 75 75s-18.7 49.1 0 67.9s49.1 18.7 67.9 0zM369.1 437c18.7 18.7 49.1 18.7 67.9 0s18.7-49.1 0-67.9s-49.1-18.7-67.9 0s-18.7 49.1 0 67.9z"
                                    />
                                </svg>
                            </td>
                            <td class="font-serif text-md p-2">Please wait...</td>
                        </tr>
                    {:then}
                        {#each projects as project, index}
                            <tr class="{(index % 2) == 0 ? "bg-slate-300/50 hover:bg-amber-100/20" : "bg-slate-300 hover:bg-amber-100/40"} duration-75 hover:text-red-700 cursor-pointer" on:click={() => goto(`/dashboard/${project.code}`)}>
                                <td class="text-sm p-2">{project.code}</td>
                                <td class="font-serif p-2 text-md">{project.name}</td>
                            </tr>
                        {/each}
                    {/await}
                </tbody>
            </table>

            <div class="flex">
                <ActionButton label="Create new project..." on:click={() => goto("/dashboard/new")} additionalClasses="grow" loading={false}/>
            </div>
        </div>
    </div>
</main>