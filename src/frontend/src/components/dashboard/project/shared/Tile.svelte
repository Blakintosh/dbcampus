<!--
	@component
	A reusable tile component that can be used as a container to display information on a specific topic in the dashboard.
	
	@prop heading - (Default: `""`) The heading of the tile. Not rendered if falsy.
	@prop isDark - (Default: `false`) When set, renders the tile with a dark background.
-->

<script lang="ts">
	export let heading: string = "";
	export let isDark: boolean = false;
	export let tileClass: string = "";
</script>

<div class="p-4 rounded-md {isDark ? 'bg-slate-600' : 'bg-slate-50'} {isDark ? 'text-slate-50' : 'text-slate-800'} shadow-md m-2 {tileClass}">
	{#if heading}
		<h1 class="font-medium text-xl mb-2">{heading}</h1>
	{/if}
	<span class="text-sm">
		<slot>
			Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.
		</slot>
	</span>
</div>