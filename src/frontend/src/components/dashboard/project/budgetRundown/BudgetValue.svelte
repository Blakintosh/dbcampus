<script lang="ts">
	export let heading: string;
	export let value: number;
	export let currency: string = "$";
</script>

<div class="font-medium">
	{heading}
</div>
<div class="flex items-center mb-8">
	<div class="text-xl xl:text-2xl font-medium mr-1 xl:mr-2">{currency}</div>
	<div class="text-4xl xl:text-5xl">{value.toLocaleString("en", { minimumFractionDigits: 2, maximumFractionDigits: 2 })}</div>
</div>