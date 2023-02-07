<!--
	@component
	The sidebar component for the Pitstop dashboard. Contains the sidebar header and a list of projects to select from, along with user controls.
-->

<script>
	import SidebarHeader from "./SidebarHeader.svelte";

</script>

<nav class="w-64 bg-slate-700 text-slate-100 h-full flex-shrink-0">
	<SidebarHeader/>
	<h3 class="text-xs text-center font-semibold mb-8">Your Projects</h3>
</nav>