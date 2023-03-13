<script lang="ts">
	import type { CreateProjectData } from "../../util/models";
	import FormCheckValue from "./FormCheckValue.svelte";

    export let values: CreateProjectData;
</script>

<table class="w-full text-left my-4 border border-slate-300">
    <tbody>
        <FormCheckValue key="Project Name" value={values.projectName}/>
        <FormCheckValue key="Project Code" value={values.projectCode}/>
        <FormCheckValue key="Deadline" value={values.deadline.toISOString().split("T")[0]} divide/>
        <FormCheckValue key="Budget" value={values.budget.toLocaleString("en-us", { currency: "usd", style: "currency" })}/>
        <FormCheckValue key="Monthly Expenses" value={values.monthlyExpenses.toLocaleString("en-us", { currency: "usd", style: "currency" })}/>
        <FormCheckValue key="Initial Spend" value={values.customSpendings.toLocaleString("en-us", { currency: "usd", style: "currency" })} divide/>
        {#if values.jiraURL && values.jiraApiToken}
            <FormCheckValue key="Jira URL" value={values.jiraURL}/>
            <FormCheckValue key="Jira Project Code" value={values.jiraProjectId} divide/>
        {:else}
            <FormCheckValue key="Jira Integration" value="None" divide/>
        {/if}

    </tbody>
</table>