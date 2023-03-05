<script lang="ts">
	import Checkbox from "./Checkbox.svelte";

	export let category: string;
	export let question: string;
	export let enabled: boolean = true;
	// type (is it a 1-5 scale, or a yes/no question)

	function toggle() {
		enabled = !enabled;
	}
</script>

<li>
	<button on:click={toggle} class="flex w-full text-left items-center cursor-pointer p-2 gap-8 {enabled ? "bg-slate-700 text-slate-100 hover:text-slate-50 hover:bg-slate-600/50" : "bg-slate-700/20 text-slate-300 hover:bg-slate-700/40 hover:text-slate-200"} border-b border-slate-600 duration-75">
		<!-- checkbox -->
		<Checkbox {enabled} focusable={false}/>
		<!-- question -->
		<div class="flex flex-col">
			<h4 class="font-medium text-lg">{category}</h4>
			<p>{question}</p>
		</div>
	</button>
</li>