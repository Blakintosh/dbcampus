<script lang="ts">
	import BudgetValue from "./BudgetValue.svelte";

	export let date: string;
    $: jsDate = new Date(date);
</script>

<div class="mb-4">
	Last Budget Report: <span class="ml-1 font-medium">{
		jsDate.toLocaleDateString("en-GB", {
			year: "numeric",
			month: "long",
			day: "numeric"
		})
		}</span>
</div>

<BudgetValue heading="Current Spend" value={13452.50} currency="£"/>

<BudgetValue heading="Total Budget" value={23000} currency="£"/>