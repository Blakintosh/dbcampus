<script lang="ts">
    import { page } from '$app/stores';

	export let label: string = "New";
	export let href: string = "#";
</script>

<svelte:element this={$page.url.pathname === href ? "div" : "a"} {href}>
	<li class="flex gap-4 items-center px-2 py-2 my-0 rounded-sm text-slate-50 hover:text-amber-100 hover:bg-slate-800/50 duration-75">
		<span class="fill-amber-300 w-4">
			<slot/>
		</span>
		<span class="font-medium text-sm">{label}</span>
	</li>
</svelte:element>