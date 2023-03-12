<script lang="ts">
	import { newProjectData, newProjectSectionDataValid, newProjectShowIssues } from "../../util/stores";
	import FormField from "../form/FormField.svelte";
	import FormTextField from "../form/FormTextField.svelte";

    $: {
        $newProjectShowIssues = false;
        $newProjectSectionDataValid = $newProjectData.budget > 0 && 
            $newProjectData.budget <= 1000000000000 &&
            $newProjectData.monthlyExpenses > 0 && 
            $newProjectData.monthlyExpenses <= 1000000000000 &&
            $newProjectData.customSpendings >= 0 &&
            $newProjectData.customSpendings <= 1000000000000;
    };

    let budgetError: string = "";
    let monthlyExpenditureError: string = "";
    let customSpendingsError: string = "";
    $: {
        if($newProjectShowIssues) {
            if($newProjectData.budget <= 0) {
                budgetError = "Please enter a value.";
            } else if($newProjectData.budget > 1000000000000) {
                budgetError = "Value is too large.";
            } else {
                budgetError = "";
            }

            if($newProjectData.monthlyExpenses <= 0) {
                monthlyExpenditureError = "Please enter a value.";
            } else if($newProjectData.monthlyExpenses > 1000000000000) {
                monthlyExpenditureError = "Value is too large.";
            } else {
                monthlyExpenditureError = "";
            }

            if($newProjectData.customSpendings > 1000000000000) {
                customSpendingsError = "Value is too large.";
            }
        }
    };
</script>

<h1 class="text-xl">Budget Details</h1>

<p class="text-sm">
	Please fill out all details relating to the budget of the project and its expected costs.
</p>
<p class="text-sm">
	<span class="text-red-600 font-semibold">*</span> = Mandatory
</p>

<FormField label="What currency are you using?" isFor="currency" required>
	<select name="currency" class="w-full border border-slate-300 rounded-lg p-2 mt-2 opacity-25" disabled aria-disabled>
		<option value="GBP">GBP</option>
		<option value="USD" selected>USD</option>
		<option value="EUR">EUR</option>
	</select>
</FormField>

<FormField label="What is the project budget?" isFor="budget" required error={budgetError}>
	<input type="number" name="budget" placeholder="Budget" min="0" class="w-full border {budgetError ? "border-red-400 bg-red-100/80" : "border-slate-300"} rounded-lg p-2 mt-2" bind:value={$newProjectData.budget}/>
</FormField>

<FormField label="What's your expected regular monthly expenditure?" isFor="monthlyExpenditure" required error={monthlyExpenditureError}>
	<input type="number" name="monthlyExpenditure" placeholder="Monthly Expenditure" min="0" class="w-full border {monthlyExpenditureError ? "border-red-400 bg-red-100/80" : "border-slate-300"} rounded-lg p-2 mt-2" bind:value={$newProjectData.monthlyExpenses}/>
</FormField>

<FormField label="Are there any costs up to now that you would like to add?" isFor="initialCost">
	<input type="number" name="initialCost" placeholder="Initial Expenses" min="0" class="w-full border {customSpendingsError ? "border-red-400 bg-red-100/80" : "border-slate-300"} rounded-lg p-2 mt-2" bind:value={$newProjectData.customSpendings}/>
</FormField>