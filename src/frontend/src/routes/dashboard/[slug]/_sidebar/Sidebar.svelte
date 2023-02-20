<!--
	@component
	The sidebar component for the Pitstop dashboard. Contains the sidebar header and a list of projects to select from, along with user controls.
-->

<script lang="ts">
	import ProjectsDropdown from "./ProjectsDropdown.svelte";
	import SidebarHeader from "./SidebarHeader.svelte";
	import type { SoftwareProject } from "../../models";
	import ProjectViewSelect from "./ProjectViewSelect.svelte";
	import ProjectManageSelect from "./ProjectManageSelect.svelte";
	import SidebarProfileWidget from "./SidebarProfileWidget.svelte";

	export let projects: Array<SoftwareProject> = [];
</script>

<nav class="w-64 bg-slate-700 text-slate-100 h-full flex-shrink-0 flex flex-col justify-between">
	<div>
		<SidebarHeader/>

		<div class="px-6">
			<h3 class="text-xs font-semibold mb-2">Go to Project</h3>
			<ProjectsDropdown {projects}/>
			
			<h3 class="text-xs font-semibold mb-2">View</h3>
			<ProjectViewSelect/>
		
			<h3 class="text-xs font-semibold mb-2">Manage</h3>
			<ProjectManageSelect/>
		</div>
	</div>

	<div class="px-6">
		<h3 class="text-xs font-semibold mb-2">Profile</h3>
		<SidebarProfileWidget/>
	</div>
</nav>