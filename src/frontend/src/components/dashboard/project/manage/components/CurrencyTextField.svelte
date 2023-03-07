<script lang="ts">
	export let currency: string = "$";
	export let placeholder: string = "";
	export let value: string = "";
</script>

<div class="relative">
	<div class="absolute left-2 top-[50%] translate-y-[-50%] text-amber-400 text-lg leading-normal">
		{currency}
	</div>
	<input type="text" class="bg-slate-800 text-slate-100 py-2 px-4 pl-8 rounded-sm" {placeholder} bind:value={value}/>
</div>