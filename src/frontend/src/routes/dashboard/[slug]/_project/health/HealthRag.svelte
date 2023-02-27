<script lang="ts">
	import HealthTrafficLight from "./HealthTrafficLight.svelte";

	export let status: string = "unknown";
	export let message: string = "They didn't tell me anything!";
</script>

<!--
	{#if status === "danger"}
				This project has critical issues that must be dealt with as soon as possible to rescue it.
			{:else if status === "warning"}
				This project has issues that may threaten the project's future success unless dealt with.
			{:else if status === "normal"}
				This project has no issues present threatening its future success.
			{:else}
				We don't have any information on the health of this project.
			{/if}
-->

<div class="flex">
	<HealthTrafficLight {status}/>

	<div class="px-4">
		<h1 class="text-lg">
			{#if status === "danger"}
				<span class="text-red-500">Significant issues present.</span>
			{:else if status === "warning"}
				<span class="text-amber-500">Issues present.</span>
			{:else if status === "normal"}
				<span class="text-green-500">Project is OK.</span>
			{:else}
				<span class="text-gray-300">No health rating.</span>
			{/if}
		</h1>
		<p class="text-md">
			{message}
		</p>
	</div>
</div>