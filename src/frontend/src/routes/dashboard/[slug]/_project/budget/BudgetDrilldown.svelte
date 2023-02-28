<script lang="ts">
	import BudgetValue from "../budgetRundown/BudgetValue.svelte";

	export let budget: number;
	export let spend: number;
</script>

<BudgetValue heading="Total Spend" value={spend} />
<BudgetValue heading="Total Budget" value={budget} />