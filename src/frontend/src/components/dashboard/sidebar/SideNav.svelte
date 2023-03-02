<!--
	@component
	The sidebar component for the Pitstop dashboard. Contains the sidebar header and a list of projects to select from, along with user controls.
-->

<script lang="ts">
	import ProjectsDropdown from "./ProjectsDropdown.svelte";
	import SidebarHeader from "./SidebarHeader.svelte";
	import type { SoftwareProjectSnippet } from "../../../util/models";
	import ProjectViewSelect from "./ProjectViewSelect.svelte";
	import ProjectManageSelect from "./ProjectManageSelect.svelte";
	import SidebarProfileWidget from "./SidebarProfileWidget.svelte";
	import { onMount } from "svelte";
	import { fly } from "svelte/transition";

	export let projects: Array<SoftwareProjectSnippet> = [];
	
	let sidebarOpen = true;
	let responsiveView = false;

	const updateResponsiveView = () => {
		if(!sidebarOpen && window.innerWidth > 1024) {
			sidebarOpen = true;
		}
		responsiveView = window.innerWidth <= 1024;
	}

	onMount(() => {
		sidebarOpen = window.innerWidth > 1024;

		updateResponsiveView();
		window.addEventListener('resize', updateResponsiveView);
	});

	const toggleSidebar = () => {
		sidebarOpen = !sidebarOpen;
	}

	$: sidebarVisibilityClass = sidebarOpen ? (responsiveView ? "transition-transform duration-200 translate-x-0" : "") : "transition-transform duration-200 -translate-x-full";
	$: overlayVisibilityClass = sidebarOpen && responsiveView ? "block" : "hidden";
</script>

<div class="flex items-center gap-8 lg:hidden border-b border-slate-200">
	<button class="fill-slate-50 bg-slate-700 p-3 border-slate-200 border-r" on:click={toggleSidebar}>
		<svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" viewBox="0 0 448 512"><!--! Font Awesome Pro 6.3.0 by @fontawesome - https://fontawesome.com License - https://fontawesome.com/license (Commercial License) Copyright 2023 Fonticons, Inc. --><path d="M0 96C0 78.3 14.3 64 32 64H416c17.7 0 32 14.3 32 32s-14.3 32-32 32H32C14.3 128 0 113.7 0 96zM0 256c0-17.7 14.3-32 32-32H416c17.7 0 32 14.3 32 32s-14.3 32-32 32H32c-17.7 0-32-14.3-32-32zM448 416c0 17.7-14.3 32-32 32H32c-17.7 0-32-14.3-32-32s14.3-32 32-32H416c17.7 0 32 14.3 32 32z"/></svg>
	</button>
	<span class="font-medium text-red-600 text-lg">
		Project Dashboard
	</span>
</div>

<div class="fixed h-full w-full bg-black/30 z-10 backdrop-blur-sm {overlayVisibilityClass}" on:click={toggleSidebar}>
</div>

<nav class="w-64 flex-shrink-0 bg-slate-50 text-slate-100 h-full min-h-0 max-h-full fixed z-40 lg:relative overflow-y-auto {sidebarVisibilityClass}" on:resize={updateResponsiveView}>
	<div class="sticky top-0 right-0">
		<button class="text-slate-800 absolute top-1 right-1 lg:hidden" on:click={toggleSidebar}>
			<svg class="h-6 w-6" fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" viewBox="0 0 24 24" stroke="currentColor">
				<path d="M6 18L18 6M6 6l12 12"></path>
			</svg>
		</button>
	</div>
	<div class="w-full flex-shrink-0 flex flex-col justify-between h-full bg-slate-700">
		<div class="bg-slate-700">
			<SidebarHeader/>
	
			<div class="px-6 overflow-y-auto bg-slate-700">
				<h3 class="text-xs font-semibold mb-2">Go to Project</h3>
				<ProjectsDropdown {projects}/>
				
				<h3 class="text-xs font-semibold mb-2">View</h3>
				<ProjectViewSelect/>
			
				<h3 class="text-xs font-semibold mb-2">Manage</h3>
				<ProjectManageSelect/>
			</div>
		</div>
	
		<div class="px-6 bg-slate-700">
			<h3 class="text-xs font-semibold mb-2">Profile</h3>
			<SidebarProfileWidget/>
		</div>
	</div>
</nav>