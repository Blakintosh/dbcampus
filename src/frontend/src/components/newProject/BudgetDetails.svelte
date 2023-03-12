<script lang="ts">
	import FormField from "../form/FormField.svelte";
	import FormTextField from "../form/FormTextField.svelte";
</script>

<h1 class="text-xl">Budget Details</h1>

<p class="text-sm">
	Please fill out all details relating to the budget of the project and its expected costs.
</p>
<p class="text-sm">
	<span class="text-red-600 font-semibold">*</span> = Mandatory
</p>

<FormField label="What currency are you using?" isFor="currency" required>
	<select name="currency" class="w-full border border-slate-300 rounded-lg p-2 mt-2">
		<option value="GBP">GBP</option>
		<option value="USD">USD</option>
		<option value="EUR">EUR</option>
	</select>
</FormField>

<FormField label="What is the project budget?" isFor="budget" required>
	<input type="number" name="budget" placeholder="Budget" min="0" class="w-full border border-slate-300 rounded-lg p-2 mt-2"/>
</FormField>

<FormField label="What's your expected regular monthly expenditure?" isFor="monthlyExpenditure" required>
	<input type="number" name="monthlyExpenditure" placeholder="Monthly Expenditure" min="0" class="w-full border border-slate-300 rounded-lg p-2 mt-2"/>
</FormField>

<FormField label="Are there any costs up to now that you would like to add?" isFor="initialCost">
	<input type="number" name="initialCost" placeholder="Initial Expenses" min="0" class="w-full border border-slate-300 rounded-lg p-2 mt-2"/>
</FormField>